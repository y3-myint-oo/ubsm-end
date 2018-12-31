package main

import (
	router "ymo.me/sbum-end/routers"
	driver "ymo.me/sbum-end/utils"
)

func main() {
	//fmt.Println("Hello World")
	driver.InitDriver()
	router.InitConnection()

}
