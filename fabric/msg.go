package fabric

//数据传输结构
//不需要签名的消息
type UserRegisterMsgData struct {
	UserName  string `json:"user_name" default:""`  //用户名称
	PublicKey string `json:"public_key" default:""` //用户公钥(pem string)
}

//交易数据结构(需要签名的数据)
type TxMsgRawData struct {
	UserID         string `json:"user_id"`         //用户ID
	TxData         string `json:"tx_data"`         //交易数据 base64
	Nonce          string `json:"nonce"`           //随机数(base64)
	TxSignature    string `json:"tx_signature"`    //交易签名(Base64 ASN.1 DER)
	NonceSignature string `json:"nonce_signature"` //随机数签名(Base64 ASN.1 DER)
}

//情报交易数据结构
type CtiTxData struct {
	CTIID              string   `json:"cti_id"`                // 情报ID(链上生成)
	CTIHash            string   `json:"cti_hash"`              // 情报HASH(sha256链下生成)
	CTIName            string   `json:"cti_name"`              // 情报名称(可为空)
	CTIType            int      `json:"cti_type"`              // 情报类型（1:恶意流量、2:蜜罐情报、3:僵尸网络、4:应用层攻击、5:开源情报）
	CTITrafficType     int      `json:"cti_traffic_type"`      // 流量情报类型（0:非流量、1:5G、2:卫星网络、3:SDN）
	OpenSource         int      `json:"open_source"`           // 是否开源（0不开源，1开源）
	CreatorUserID      string   `json:"creator_user_id"`       // 创建者ID(公钥sha256)
	Tags               []string `json:"tags"`                  // 情报标签数组
	IOCs               []string `json:"iocs"`                  // 包含的沦陷指标（IP, Port, Payload,URL, Hash）
	StatisticInfo      string   `json:"statistic_info"`        // 统计信息(IPFS地址)
	StixData           string   `json:"stix_data"`             // STIX数据（JSON []byte）可以有多条
	StixIPFSHash       string   `json:"stix_ipfs_hash"`        // STIX数据,IPFS地址
	Description        string   `json:"description"`           // 情报描述
	DataSize           int      `json:"data_size"`             // 数据大小（B）
	DataSourceHash     string   `json:"data_source_hash"`      // 数据源HASH（sha256）
	DataSourceIPFSHash string   `json:"data_source_ipfs_hash"` // 数据源IPFS地址
	Need               int      `json:"need"`                  // 情报需求量(销售数量)
	Value              float64      `json:"value"`                 // 情报价值（积分）
	CompreValue        float64      `json:"compre_value"`          // 综合价值（积分激励算法定价）
}

type PurchaseCtiTxData struct {
	CTIID  string `json:"cti_id"`  // 情报ID
	UserID string `json:"user_id"` // 用户ID
}
type PurchaseModelTxData struct {
	ModelID string `json:"model_id"` // 模型ID
	UserID  string `json:"user_id"`  // 用户ID
}

//模型交易数据结构
type ModelTxData struct {
	ModelID             string   `json:"model_id"`              // 模型ID
	ModelHash           string   `json:"model_hash"`            // 模型hash
	ModelName           string   `json:"model_name"`            // 模型名称
	ModelCreatorUserID  string   `json:"model_creator_user_id"` // 模型创建者ID
	ModelDataType       int      `json:"model_data_type"`       // 模型数据类型(1:流量(数据集)、2:情报(文本))
	ModelType           int      `json:"model_type"`            // 模型类型(1:分类模型、2:回归模型、3:聚类模型、4:NLP模型)
	ModelAlgorithm      string   `json:"model_algorithm"`       // 模型算法
	ModelTrainFramework string   `json:"model_train_framework"` // 模型训练框架(Scikit-learn、Pytorch、TensorFlow)
	ModelOpenSource     int      `json:"model_open_source"`     // 是否开源
	ModelFeatures       []string `json:"model_features"`        // 模型特征
	ModelTags           []string `json:"model_tags"`            // 模型标签
	ModelDescription    string   `json:"model_description"`     // 模型描述
	ModelSize           int      `json:"model_size"`            // 模型大小
	ModelDataSize       int      `json:"model_data_size"`       // 模型训练数据大小
	ModelDataIPFSHash   string   `json:"model_data_ipfs_hash"`  // 模型训练数据IPFS地址
	ModelValue          float64      `json:"model_value"`           // 模型价值
	ModelIPFSHash       string   `json:"model_ipfs_hash"`       // 模型IPFS地址
	RefCTIId            string   `json:"ref_cti_id"`            // 关联情报ID(使用哪个情报训练的模型)
}
