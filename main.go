package main

import (
	"fmt"

	fabric "github.com/righstar2020/br-cti-bc-server/fabric"
	global "github.com/righstar2020/br-cti-bc-server/global"
	server "github.com/righstar2020/br-cti-bc-server/server"
)

const fabric_config_path = "./config/server/config.yaml"

func main() {
	// 启动Fabric SDK
	var err error
	//该程序只能在当前工作目录下进行(子进程调用需要切换工作目录pwd)
	global.FabricSDK, err = fabric.NewSDK(fabric_config_path) // 确保路径正确
	if err != nil {
		fmt.Printf("Failed to initialize SDK: %v", err)
	}
	defer global.FabricSDK.Close()
	//初始化fabric client
	fmt.Println("Fabric client initializing ...")
	InitFabricClient()

	r := server.NewRouter(global.FabricSDK)
	// 启动服务
	if err := r.Run(":7777"); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}

func InitFabricClient() {
	var err error
	global.ChannelClient, err = fabric.CreateChannelClient(global.FabricSDK)
	if err != nil {
		fmt.Printf("Failed to connect fabric chain: %v", err)
	}
	global.LedgerClient, err = fabric.CreateLedgerClient(global.FabricSDK)
	if err != nil {
		fmt.Printf("Failed to connect fabric chain: %v", err)
	}
}
