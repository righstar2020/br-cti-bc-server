
## 说明
br-cti fabric接口服务

## 编译
```
go build -o bc-server.exe
```
## 运行
```
go run main.go
```

## 调试接口
# 区块链相关接口
curl -X ANY http://localhost:7777/blockchain/queryBlock/{blockID} 
curl -X ANY http://localhost:7777/blockchain/queryChain

# 合约相关接口
curl -X POST http://localhost:7777/contract/queryContract
curl -X POST http://localhost:7777/contract/invokeContract

# 交易相关接口
curl -X POST http://localhost:7777/tx/getTransactionNonce

# 用户相关接口
curl -X POST http://localhost:7777/user/registerUserAccount
curl -X POST http://localhost:7777/user/queryUserInfo
curl -X POST http://localhost:7777/user/getUserStatistics
curl -X POST http://localhost:7777/user/queryPointTransactions

# CTI相关接口
curl -X POST http://localhost:7777/cti/registerCtiInfo
curl -X POST http://localhost:7777/cti/queryCtiInfo
curl -X POST http://localhost:7777/cti/queryCtiInfoByTypeWithPagination
curl -X POST http://localhost:7777/cti/queryCtiInfoByType
curl -X POST http://localhost:7777/cti/queryAllCtiInfoWithPagination
curl -X POST http://localhost:7777/cti/queryCtiInfoByCTIHash
curl -X POST http://localhost:7777/cti/queryCtiInfoByCreatorUserID

# 模型相关接口
curl -X POST http://localhost:7777/model/registerModelInfo
curl -X POST http://localhost:7777/model/queryModelInfo
curl -X POST http://localhost:7777/model/queryModelInfoByIDWithPagination
curl -X POST http://localhost:7777/model/queryModelsByTrafficType
curl -X POST http://localhost:7777/model/queryModelsByRefCTIId
curl -X POST http://localhost:7777/model/queryModelInfoByCreatorUserID

# 数据统计相关接口
curl -X POST http://localhost:7777/dataStat/queryCTISummaryInfo
curl -X POST http://localhost:7777/dataStat/getDataStatistics
curl -X POST http://localhost:7777/dataStat/getCTITrafficTrend
curl -X POST http://localhost:7777/dataStat/getAttackTypeRanking
curl -X POST http://localhost:7777/dataStat/getIOCsDistribution
curl -X POST http://localhost:7777/dataStat/getGlobalIOCsDistribution
curl -X POST http://localhost:7777/dataStat/getSystemOverview
