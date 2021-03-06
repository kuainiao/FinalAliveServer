package main

import (
	"ffCommon/log/log"
	"ffCommon/net/netmanager"
	"ffCommon/uuid"
	"ffProto"
	"fmt"
	"sync"
	"sync/atomic"
)

type agentUserServer struct {
	netManager *netmanager.Manager

	mutexAgent sync.Mutex               // mutexAgent 锁
	mapAgent   map[uuid.UUID]*agentUser // mapAgent 所有连接
	agentPool  *agentUserPool           // agentPool 所有连接缓存
}

// Create 创建
func (server *agentUserServer) Create(netsession netmanager.INetSession) netmanager.INetSessionHandler {
	log.RunLogger.Printf("agentUserServer.Create netsession[%v]", netsession)

	server.mutexAgent.Lock()
	defer server.mutexAgent.Unlock()

	// 申请
	agent := server.agentPool.apply()

	// 初始化
	agent.Init(netsession)

	// 记录
	server.mapAgent[agent.UUID()] = agent

	return agent
}

// Back 回收
func (server *agentUserServer) Back(handler netmanager.INetSessionHandler) {
	log.RunLogger.Printf("agentUserServer.Back handler[%v]", handler)

	server.mutexAgent.Lock()
	defer server.mutexAgent.Unlock()

	agent, _ := handler.(*agentUser)

	// 清除记录
	delete(server.mapAgent, agent.UUID())

	// 回收清理
	agent.Back()

	// 缓存
	server.agentPool.back(agent)
}

// Start 开始建立服务
func (server *agentUserServer) Start() error {
	log.RunLogger.Printf("agentUserServer.Start")

	manager, err := netmanager.NewServer(server, appConfig.ServeUser, &waitApplicationQuit, chApplicationQuit)
	if err != nil {
		log.FatalLogger.Println(err)
		return err
	}

	server.netManager = manager
	server.mapAgent = make(map[uuid.UUID]*agentUser, appConfig.ServeUser.InitOnlineCount)
	server.agentPool = newAgentUserPool("agentUserServer", appConfig.ServeUser.InitOnlineCount)

	atomic.AddInt32(&waitApplicationQuit, 1)

	return err
}

// End 退出完成
func (server *agentUserServer) End() {
	log.RunLogger.Printf("agentUserServer.End")

	atomic.AddInt32(&waitApplicationQuit, -1)
}

// Status 当前状态描述
func (server *agentUserServer) Status() string {
	return fmt.Sprintf("mapAgent[%v] agentPool[%v] netManager[%v]",
		len(server.mapAgent), server.agentPool, server.netManager.Status())
}

// OnCustomLoginResult Login检查结果
func (server *agentUserServer) OnCustomLoginResult(result *httpClientCustomLoginData) {
	log.RunLogger.Printf("agentUserServer.OnCustomLoginResult result[%v]", result)

	server.mutexAgent.Lock()
	defer server.mutexAgent.Unlock()

	agent, ok := server.mapAgent[result.uuidRequester]
	log.RunLogger.Printf("agentUserServer.OnCustomLoginResult find agent[%v] result[%v]", result.uuidRequester, ok)

	if ok {
		onCustomLoginResult(agent, result)
	}
}

// OnMatchServerProto 接收到来自MatchServer的User相关协议
func (server *agentUserServer) OnMatchServerProto(proto *ffProto.Proto) bool {
	uuid := uuid.NewUUID(proto.ExtraData())
	log.RunLogger.Printf("agentUserServer.OnMatchServerProto agent[%v] proto[%v]", uuid, proto)

	server.mutexAgent.Lock()
	defer server.mutexAgent.Unlock()

	agent, ok := server.mapAgent[uuid]
	if ok {
		return ffProto.SendProtoExtraDataNormal(agent, proto, true)
	}
	log.RunLogger.Printf("agentUserServer.OnMatchServerProto agent[%v] offline", uuid)
	// todo:通知匹配服务器, 该用户已离线
	return false
}
