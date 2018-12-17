package main

import (
	"fmt"

	router "ymo.me/sbum-end/routers"
	driver "ymo.me/sbum-end/utils"
)

func main() {
	fmt.Println("Hello World")
	router.InitConnection()
	driver.InitDriver()
}
