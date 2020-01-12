package main

import (
	"com/dao"
	"com/server"
	"fmt"
)

func main() {
	for _,u := range dao.GetUsers() {
		fmt.Println(u)
	}
	port := "8090"
	fmt.Print("Server running at port: " + port)
	server.RunServer(port)
}
