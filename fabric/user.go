package fabric

import (
	"encoding/json"

	"github.com/righstar2020/br-cti-bc-server/global"
)

func RegisterUserAccount(userName string, publicKey string) (string, error) {
	// 调用链码注册用户账户
	client, err := CreateChannelClient(global.FabricSDK)
	if err != nil {
		return "", err
	}

	// 构造注册消息数据
	msgData := UserRegisterMsgData{
		UserName:  userName,
		PublicKey: publicKey,
	}
	msgJsonData, err := json.Marshal(msgData)
	if err != nil {
		return "", err
	}

	resp, err := InvokeChaincode(client, global.MainChaincodeName, "RegisterUserInfo", [][]byte{msgJsonData})
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func QueryUserInfo(userID string) (string, error) {
	// 创建通道客户端
	client, err := CreateChannelClient(global.FabricSDK)
	if err != nil {
		return "", err
	}

	// 构造查询参数
	args := [][]byte{[]byte(userID)}

	// 调用链码查询用户信息
	resp, err := InvokeChaincode(client, global.MainChaincodeName, "QueryUserInfo", args)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

// 查询用户积分信息
func QueryUserPointInfo(userID string) (string, error) {
	// 创建通道客户端
	client, err := CreateChannelClient(global.FabricSDK)
	if err != nil {
		return "", err
	}

	// 构造查询参数
	args := [][]byte{[]byte(userID)}

	// 调用链码查询用户积分信息
	resp, err := InvokeChaincode(client, global.MainChaincodeName, "QueryUserPointInfo", args)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

//查询用户详细信息
func QueryUserDetailInfo(userID string) (string, error) {
	// 创建通道客户端
	client, err := CreateChannelClient(global.FabricSDK)
	if err != nil {
		return "", err
	}

	// 构造查询参数
	args := [][]byte{[]byte(userID)}

	// 调用链码查询用户详细信息
	resp, err := InvokeChaincode(client, global.MainChaincodeName, "QueryUserDetailInfo", args)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}
//查询所有用户列表
func QueryAllUserList() (string, error) {
	// 创建通道客户端
	client, err := CreateChannelClient(global.FabricSDK)
	if err != nil {
		return "", err
	}

	// 调用链码查询所有用户列表
	resp, err := InvokeChaincode(client, global.MainChaincodeName, "QueryUserAccountList", [][]byte{})
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func GetUserStatistics(userID string) (string, error) {
	// 创建通道客户端
	client, err := CreateChannelClient(global.FabricSDK)
	if err != nil {
		return "", err
	}

	// 构造查询参数
	args := [][]byte{[]byte(userID)}

	// 调用链码获取用户统计数据
	resp, err := InvokeChaincode(client, global.MainChaincodeName, "GetUserStatistics", args)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func QueryPointTransactions(userID string) (string, error) {
	// 创建通道客户端
	client, err := CreateChannelClient(global.FabricSDK)
	if err != nil {
		return "", err
	}

	// 构造查询参数
	args := [][]byte{[]byte(userID)}

	// 调用链码查询用户积分交易记录
	resp, err := InvokeChaincode(client, global.MainChaincodeName, "QueryPointTransactions", args)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}
// 查询用户拥有的情报(上传+购买的)
func QueryUserOwnCTIInfos(userID string) (string, error) {
	// 创建通道客户端
	client, err := CreateChannelClient(global.FabricSDK)
	if err != nil {
		return "", err
	}

	// 构造查询参数
	args := [][]byte{[]byte(userID)}

	// 调用链码查询用户拥有的情报
	resp, err := InvokeChaincode(client, global.MainChaincodeName, "QueryUserOwnCTIInfos", args)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

