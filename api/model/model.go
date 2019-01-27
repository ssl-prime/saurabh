package model

//Block ...
type Block struct {
	Index      int
	Timestamp  string
	BPM        int
	Hash       string
	PrevHash   string
	Difficulty int
	Nonce      string
}

//Blockchain ...
var Blockchain []Block

//Message  ...
type Message struct {
	BPM int
}

//UserInfo ...
type UserInfo struct {
	BPM  int    `json:"bpm" valid:"required"`
	Name string `json:"name"`
}

//IPFSData ...
type IPFSData struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

//PreDeployInfo ...
type PreDeployInfo struct {
	PrivateKey      string `json:"private_key"`
	AccountNo       string `json:"account_no"`
	ContractAddress string `json:"contract_address"`
}

//PostDeploy ...
type PostDeploy struct {
	PrivateKey      string `json:"private_key"`
	ContractAddress string `json:"contract_address"`
}
