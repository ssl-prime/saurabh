package util

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"saurabh/api/model"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/rightjoin/aqua"
)

const (
	difficulty = 1
)

var mutex = &sync.Mutex{}

//ValidateAddData ...
func ValidateAddData(j aqua.Aide) error {
	//validate all data
	if len(model.Blockchain) == 0 {
		genesisBlock()
		newblock := generateBlock(model.Blockchain[len(model.Blockchain)-1], 60)
		mutex.Lock()
		model.Blockchain = append(model.Blockchain, newblock)
		mutex.Unlock()

	} else {
		newblock := generateBlock(model.Blockchain[len(model.Blockchain)-1], 60)
		mutex.Lock()
		model.Blockchain = append(model.Blockchain, newblock)
		mutex.Unlock()
		//generateBlock(model.Blockchain[len(model.Blockchain)-1], 60)
	}
	return nil
}

//AddData ...
func AddData(j aqua.Aide) (interface{}, error) {
	j.LoadVars()
	fmt.Println(j.Body)
	var rqstBody model.UserInfo
	if err := json.Unmarshal([]byte(j.Body), &rqstBody); err == nil {
		fmt.Println("come on")
		return model.Blockchain, nil
	} else {
		return "wrong info", err
	}

}

//ValidateUploadFile ...
func ValidateUploadFile(j aqua.Aide) error {
	return nil
}

//UploadFile ...
func UploadFile(j aqua.Aide) (interface{}, error) {
	return "upload", nil
}

//ValidateGetFile ...
func ValidateGetFile(j aqua.Aide) error {
	return nil
}

//GetFile ...
func GetFile(j aqua.Aide) (interface{}, error) {
	return "upload", nil
}

//generateBlock ...
func generateBlock(oldBlock model.Block, BPM int) model.Block {
	var newBlock model.Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Difficulty = difficulty

	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)
		newBlock.Nonce = hex
		if !isHashValid(calculateHash(newBlock), newBlock.Difficulty) {
			fmt.Println(calculateHash(newBlock), " do more work!")
			time.Sleep(time.Second)
			continue
		} else {
			fmt.Println(calculateHash(newBlock), " work done!")
			newBlock.Hash = calculateHash(newBlock)
			break
		}

	}
	return newBlock
}

//isHashValid ...
func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

//isBlockValid ...
func isBlockValid(newBlock, oldBlock model.Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

//calculateHash ...
func calculateHash(block model.Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.BPM) +
		block.PrevHash + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

//genesisBlock ...
func genesisBlock() {
	t := time.Now()
	genesisBlock := model.Block{}
	genesisBlock = model.Block{0, t.String(), 0, calculateHash(genesisBlock), "", difficulty, ""}
	spew.Dump(genesisBlock)

	mutex.Lock()
	model.Blockchain = append(model.Blockchain, genesisBlock)
	mutex.Unlock()
}
