package enum

type EvtType string

const (
	JOIN    EvtType = "사용자가 채팅방에 들어왔습니다."
	LEAVE   EvtType = "사용자가 채팅방을 나갔습니다."
	MESSAGE EvtType = "채팅 메시지"
)