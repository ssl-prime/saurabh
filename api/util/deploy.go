package util

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	//"solar/abi/bind"

	"saurabh/api/model"
	file "saurabh/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-ethereum/crypto"
	"github.com/go-ethereum/ethclient"
	"github.com/rightjoin/aqua"
)

//TODO ...
//make a connection file
// get nonce of an account
//add a file on ipfs and then on blockchain
//get hash and read data from ipfs and display on front end
//every time you have to take care of nonce ,nonce is very importnat for tran

//DeploySmartContract ...
func DeploySmartContract(j aqua.Aide) (string, error) {
	j.LoadVars()
	var (
		client *ethclient.Client
		err    error
		auth   *bind.TransactOpts
		msg    string
		// 	address  common.Address
		// tx       *types.Transaction
		// instance *file.FileStorage
	)
	msg = "connection failed"
	if client, err = ConnectToBlockChain(); err == nil {
		if auth, err = GetKeyAndNonce(client, j); err == nil && auth != nil {
			// to check then returned instance is golobal or for specific address
			address, tx, instance, errDeploy := file.DeployFileStorage(auth, client)
			err = errDeploy
			if err == nil {
				fmt.Println("contract address", address.Hex())
				fmt.Println("transaction id", tx.Hash().Hex())
				_ = instance
				msg = "successFully deployed"
			} else {

				fmt.Println(err, "srt")
				//return "", instance
			}
		} else if auth == nil {
			msg = "auth is nil"
			fmt.Println("correct you given info", auth)
		}
	}
	return msg, err
}

//AddHashToBlockchain  ...
func AddHashToBlockchain(cid string, j aqua.Aide) {
	j.LoadVars()
	if client, err := ConnectToBlockChain(); err == nil {
		if auth, err := GetKeyAndNonce(client, j); err == nil && auth != nil {
			//address of smart contract
			var info model.PreDeployInfo
			if err := json.Unmarshal([]byte(j.Body), &info); err == nil {
			}
			address := common.HexToAddress(info.ContractAddress)
			fmt.Println(address)
			if instance, err := file.NewFileStorage(address, client); err == nil {
				_, err := instance.AddFile(auth, cid, "saurabh")
				fmt.Println(err)
			} else {
				fmt.Println("instance", err)
			}
		} else {
			fmt.Println("auth ----", err)
		}
	} else {
		fmt.Println(err, "connection")
	}
	return
}

//get hash from blockchain and retrive data from ipfs

//GetHashed ...
func GetHashed(j aqua.Aide) (string, error) {
	j.LoadVars()
	if client, err := ConnectToBlockChain(); err == nil {
		if auth, err := GetKeyAndNonce(client, j); err == nil && auth != nil {
			//address of smart contract
			var info model.PreDeployInfo
			if err := json.Unmarshal([]byte(j.Body), &info); err == nil {
			}
			address := common.HexToAddress(info.ContractAddress)
			//fmt.Println(address)

			if instance, err := file.NewFileStorage(address, client); err == nil {
				var opts *bind.CallOpts
				opts.From = common.HexToAddress(info.AccountNo)
				opts.Pending = true
				opts.Context = context.Background()
				//instance.GetHash()
				hash, err := instance.GetHash(opts)
				//fmt.Println(hash, err, "--------------------------")
				fmt.Println(hash, err)
			} else {
				fmt.Println("instance", err)
			}
		} else {
			fmt.Println("auth ----", err)
		}
	} else {
		fmt.Println(err, "connection")
	}
	return "hashed testing", nil
}

//ConnectToBlockChain ...
func ConnectToBlockChain() (*(ethclient.Client), error) {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	// network := "dev"
	// switch network {
	// case "development":
	// 	fmt.Println("dev connected")
	// }
	if err != nil {
		log.Fatal(err)
	}
	return client, err
}

//GetKeyAndNonce ...
func GetKeyAndNonce(client *ethclient.Client, j aqua.Aide) (auth *bind.TransactOpts, err error) {
	//fmt.Println("called")
	var info model.PreDeployInfo
	if err := json.Unmarshal([]byte(j.Body), &info); err == nil {
		//private key will be given by user during the transaction
		//fmt.Print(info.PrivateKey)
		if privateKey, err := crypto.HexToECDSA(info.PrivateKey); err == nil {
			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if ok {
				fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
				nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
				if err == nil {
					if gasPrice, err := client.SuggestGasPrice(context.Background()); err == nil {
						auth = bind.NewKeyedTransactor(privateKey)
						auth.Nonce = big.NewInt(int64(nonce))
						auth.Value = big.NewInt(0)
						auth.GasLimit = uint64(0)
						auth.GasPrice = gasPrice
					}
				}
			}
		}
	}
	return
}
