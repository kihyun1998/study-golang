package main

import (
	"container/list"
	"log"
	"net/http"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

// ===============================구조체=================================
// 채팅 이벤트 구조체 정의
type Event struct {
	EvtType   string //이벤트 타입
	User      string // 사용자 이름
	Timestamp int    // 시간값
	Text      string // 메시지 텍스트``
}

// 구독 구조체 정의
type Subscription struct {
	Archive []Event      // 이벤트 저장 슬라이스
	New     <-chan Event // 이벤트 채널
}

//=====================================================================

func NewEvent(evtType, user, msg string) Event {
	return Event{evtType, user, int(time.Now().Unix()), msg}
}

var (
	subscribe   = make(chan (chan<- Subscription), 10) // 구독 채널(send-only)
	unsubscribe = make(chan (<-chan Event), 10)        // 구독 해지 채널(receive-only)
	publish     = make(chan Event, 10)                 // 이벤트 발행 채널
)

// ============================행동에 대한 함수===========================
// 새 사용자가 들어왔을 때 이벤트 구독 함수
func Subscribe() Subscription {
	c := make(chan Subscription)
	subscribe <- c // 구독채널에 보냄
	return <-c
}

// 사용자 나가면 구독 취소
func (s Subscription) Cancel() {
	unsubscribe <- s.New //구독 해지 채널로 보냄
	for {
		select {
		case _, ok := <-s.New: //채널에서 값 꺼냄, 저장은 안함
			if !ok { //값 모두 꺼내면 return
				return
			}
		default:
			return
		}
	}
}

// 기존 사용자 들어올 시 이벤트 발행
func Join(user string) {
	publish <- NewEvent("join", user, "")
}

// 채팅 메시지 보냈을 때 이벤트 발행
func Say(user, message string) {
	publish <- NewEvent("message", user, message)
}

// 사용자가 떠났을 때 이벤트 발행
func Leave(user string) {
	publish <- NewEvent("leave", user, "")
}

//=====================================================================

// 구독, 구독 해지, 발행된 이벤트 처리 함수
func Chatroom() {
	archive := list.New()     //이벤트 연결 리스트
	subscribers := list.New() //구독자 연결 리스트

	for {
		select {
		// 새로운 사용자 방문 시
		case c := <-subscribe:
			var events []Event
			for e := archive.Front(); e != nil; e = e.Next() { //이벤트 쌓인 것들 슬라이스에 이벤트 저장
				events = append(events, e.Value.(Event))
			}
			subscriber := make(chan Event, 10) //이벤트 채널
			subscribers.PushBack(subscriber)   //구독자 목록에 추가
			c <- Subscription{events, subscriber}

		// 새 이벤트 발행 시
		case event := <-publish:
			// 모든 사용자에게 이벤트 전달
			for e := subscribers.Front(); e != nil; e = e.Next() {
				subscriber := e.Value.(chan Event)
				subscriber <- event
			}
			if archive.Len() >= 20 {
				archive.Remove(archive.Front())
			}
			archive.PushBack(event)

		// 사용자가 나갈 시
		case c := <-unsubscribe:
			for e := subscribers.Front(); e != nil; e.Next() {
				subscriber := e.Value.(chan Event)
				// 구독자 목록에 있는 이벤트와 채널 c가 같다면
				if subscriber == c {
					subscribers.Remove(e) //구독자 목록에서 삭제
					break
				}
			}
		}
	}
}

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	go Chatroom()

	server.On("connection", func(so socketio.Socket) {
		s := Subscribe()
		Join(so.Id())
		for _, event := range s.Archive {
			so.Emit("event", event)
		}

		newMessages := make(chan string)
		so.On("message", func(msg string) {
			newMessages <- msg
		})

		so.On("disconnection", func() {
			Leave(so.Id())
			s.Cancel()
		})

		go func() {
			for {
				select {
				case event := <-s.New:
					so.Emit("event", event)
				case msg := <-newMessages:
					Say(so.Id(), msg)
				}
			}
		}()

	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":5000", nil)

}
