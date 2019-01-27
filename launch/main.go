package main

import (
	block "saurabh/api/service"

	"github.com/rightjoin/aqua"
)

func main() {
	server := aqua.NewRestServer()
	server.AddService(&block.BlockAPI{})
	//fmt.Println("smt")
	server.Run()
}
