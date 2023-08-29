package main

import (
	"container/list"
	"fmt"
)

// 자료구조 종류 : 연결 리스트, 힙, 링
// func New() *List : 연결리스트 생성

// func (l *List) PushBack(v interface{}) *Element : 연결리스트 맨 뒤에 데이터 추가
// func (l *List) PushFront(v interface{}) *Element : 연결리스트 맨 앞에 데이터 추가
// func (l *List) PushBackList(other List) : 연결리스트 뒤에 다른 리스트 추가
// func (l *List) PushFrontList(other List) : 연결리스트 앞에 다른 리스트 추가
// func (l *List) InsertAfter(v interface{}, mark *Element) *Element  : 특정 노드 뒤에 추가
// func (l *List) InsertBefore(v interface{}, mark *Element) *Element  : 특정 노드 앞에 추가
// func (l *List) MoveAfter(ele, mark *Element) : 노드를 특정 노드 뒤로 옮김
// func (l *List) MoveBefore(ele, mark *Element) : 노드를 특정 노드 앞으로 옮김
// func (l *List) MoveToBack(ele *Element) : 노드를 맨 뒤로 옮김
// func (l *List) MoveToFront(ele *Element) : 노드를 맨 앞으로 옮김
// func (l *List) Len() int : 연결 리스트의 길이
// func (l *List) Remove(ele *Element) any : 특정 노드를 제거

// func (l *List) Front() *Element : 연결 리스트의 맨 앞 데이터를 가져옴
// func (l *List) Back() *Element : 연결 리스트의 맨 뒤 데이터를 가져옴

func main() {
	l := list.New() // 연결리스트 생성
	l.PushBack(10)  // 데이터 추가
	l.PushBack(20)
	l.PushBack(30)

	fmt.Println("Front ", l.Front().Value) //연결 리스트의 맨 앞 데이터를 가져옴
	fmt.Println("Back ", l.Back().Value)   // 연결 리스트 맨 뒤 데이터 가져오기

	// e는 element의 약자
	for e := l.Front(); e != nil; e = e.Next() { // for in과 비슷
		fmt.Println(e.Value)
	}
}
