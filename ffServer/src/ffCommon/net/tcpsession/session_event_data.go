package tcpsession

import (
	"ffCommon/log/log"
	"ffCommon/net/base"
	"ffProto"
	"fmt"
)

type sessionNetEventData struct {
	eventType base.NetEventType

	manualClose bool

	session *tcpSession    // 事件对应的Session引用, 如果是NetEventOff事件, 则在回收事件的同时, 执行Session的back方法
	proto   *ffProto.Proto // 只有引用, 回收操作, 由事件处理者负责
}

// Back 回收
func (s *sessionNetEventData) Back() {
	log.RunLogger.Printf("sessionNetEventData.Back: %v", s)

	if s.eventType == base.NetEventOff { // 回收tcpsession
		s.session.back()
	}

	s.eventType = base.NetEventInvalid
	s.proto, s.session = nil, nil

	// 回收自身
	eventDataPool.back(s)
}

// Session Session
func (s *sessionNetEventData) Session() base.Session {
	return s.session
}

// NetEventType 获取事件类型
func (s *sessionNetEventData) NetEventType() base.NetEventType {
	return s.eventType
}

// ManualClose 当NetEvent为NetEventOff时有效, 返回是不是主动关闭引发的Session断开
func (s *sessionNetEventData) ManualClose() bool {
	return s.manualClose
}

// Proto 当NetEvent为NetEventProto时有效, 返回事件携带的协议
func (s *sessionNetEventData) Proto() *ffProto.Proto {
	return s.proto
}

func (s *sessionNetEventData) String() string {
	return fmt.Sprintf(`%p eventType[%v] uuidSession[%v]`, s, s.eventType, s.session.uuid)
}

func newSessionNetEvent() *sessionNetEventData {
	return &sessionNetEventData{}
}

func newSessionNetEventOn(session *tcpSession) base.NetEventData {
	data := eventDataPool.apply()
	data.session, data.eventType = session, base.NetEventOn
	return data
}

func newSessionNetEventOff(session *tcpSession, manualClose bool) base.NetEventData {
	data := eventDataPool.apply()
	data.session, data.eventType, data.manualClose = session, base.NetEventOff, manualClose
	return data
}

func newSessionNetEventProto(session *tcpSession, proto *ffProto.Proto) base.NetEventData {
	data := eventDataPool.apply()
	data.session, data.eventType, data.proto = session, base.NetEventProto, proto
	return data
}
