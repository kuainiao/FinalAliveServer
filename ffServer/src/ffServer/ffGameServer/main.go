package main

import (
	"ffCommon/log/log"
	"ffCommon/util"
	"ffLogic/ffGameWorld"
	"time"
)

func main() {
	defer util.PanicProtect(func(isPanic bool) {
		if isPanic {
			log.RunLogger.Println("异常退出, 以上是错误堆栈")
			<-time.After(time.Hour)
		}
	}, "ffGameServer")

	// 初始化
	err := startup()
	if err != nil {
		log.FatalLogger.Println(err)
		return
	}

	// 创建游戏世界
	world, err = ffGameWorld.NewGameWorld(worldFrame)
	if err != nil {
		log.FatalLogger.Println(err)
		return
	}

	// 启动连接
	if err = agentServerMgr.start(); err != nil {
		log.FatalLogger.Println(err)
		return
	}

	go util.SafeGo(worldFrame.mainLoop, nil)

	// 等待关闭
	select {}
}
