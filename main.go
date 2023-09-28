package main

import (
	"fmt"
	config "vid/config"
)

func main() {
	fmt.Println("Hello, world!")
	db := config.DB{}
	con, _ := db.Connect()
	fmt.Println(con)

}
