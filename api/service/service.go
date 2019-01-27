package service

import (
	"saurabh/api/util"

	"github.com/rightjoin/aqua"
)

//BlockAPI ...
type BlockAPI struct {
	aqua.RestService    `prefix:"blockchain/" root:"/" version:"1"`
	addData             aqua.POST `url:"/addData"`
	uploadFile          aqua.POST `url:"/uploadFile"`
	getFile             aqua.POST `url:"/getFile"`
	ipfs                aqua.POST `url:"/ipfs"`
	addEncFile          aqua.POST `url:"/addEncFile"`
	decEncFile          aqua.POST `url:"/decEncFile"`
	deploySmartContract aqua.POST `url:"/deploy"`
	getItemHash         aqua.POST `url:"/getItemHash"`
}

//AddData ...
func (atdn *BlockAPI) AddData(j aqua.Aide) (response interface{}, err error) {
	if err = util.ValidateAddData(j); err == nil {
		response, err = util.AddData(j)
	}
	return
}

//UploadFile ...
func (atdn *BlockAPI) UploadFile(j aqua.Aide) (response interface{}, err error) {
	if err = util.ValidateUploadFile(j); err == nil {
		response, err = util.UploadFile(j)
		//}

	}
	return
}

//Ipfs ...
func (atdn *BlockAPI) Ipfs(j aqua.Aide) (response interface{}, err error) {
	response, err = util.Ipfs(j)
	return
}

//AddEncFile ...
func (atdn *BlockAPI) AddEncFile(j aqua.Aide) (response interface{}, err error) {
	response, err = util.AddEncryptedFile()
	return
}

//DecEncFile ...
func (atdn *BlockAPI) DecEncFile(j aqua.Aide) (response interface{}, err error) {
	response, err = util.DecryptExistingFile()
	return
}

//DeploySmartContract ...
func (atdn *BlockAPI) DeploySmartContract(j aqua.Aide) (response interface{}, err error) {
	response, err = util.DeploySmartContract(j)
	return
}

//GetItemHash ...
func (atdn *BlockAPI) GetItemHash(j aqua.Aide) (response interface{}, err error) {
	response, err = util.GetHashed(j)
	return
}
