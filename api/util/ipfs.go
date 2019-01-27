package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/rightjoin/aqua"

	shell "github.com/ipfs/go-ipfs-api"
)

//Ipfs ...
func Ipfs(j aqua.Aide) (interface{}, error) {
	//connecting to ipfs node
	sh := shell.NewShell("localhost:5001")

	//read data from local file
	//after uploading file encrypt the hash of file or encrypt the file
	dat, err := ioutil.ReadFile("/home/coinmark-003/test.txt")
	fmt.Println(err)
	//	check(err)
	fmt.Print(string(dat))
	//path := `home/coinmark-003/test.txt`

	//add file to ipfs
	cid, err := sh.Add(strings.NewReader(string(dat)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	// if err := sh.Get("QmbCMoABa9ETBKyAUMt7axvEyBsQsi4oBtMfzCm8WKQvG7", "/home/coinmark-003/go/src/saurabh/files/test.txt"); err == nil {
	// 	fmt.Println("file stored")
	// }

	//pin file to ipfs
	if err := sh.Pin(cid); err == nil {
		fmt.Println("pinned")
	} else {
		fmt.Println("errrrr", err)
	}

	fmt.Printf("added %s", cid)
	AddHashToBlockchain(cid, j)
	return "ipfs holding", nil
}

//AddEncryptedFile ...
func AddEncryptedFile() (interface{}, error) {
	path := "/home/coinmark-003/go/src/saurabh/files/testser.txt"
	fmt.Println(path)
	EncryptFile(path, []byte("saurabh 123456 hngh"), "1234567890")
	return "encrypted", nil
}

//DecryptExistingFile ...
func DecryptExistingFile() (interface{}, error) {
	path := "/home/coinmark-003/go/src/saurabh/files/testser.txt"
	data := DecryptFile(path, "1234567890")
	return string(data), nil

}
