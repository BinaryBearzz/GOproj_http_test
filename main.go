package main

import (
	"fmt"
	"mul_src/Internal/httprequest"
	"mul_src/Internal/utils"
)

func main() {
	fmt.Println("main")
	utils.Test()
	httprequest.RequestGET("catfact.ninja", "/fact", 443)
}
