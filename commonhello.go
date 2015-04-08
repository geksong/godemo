package main

import (
	"fmt"
	"github.com/geksong/godemo/goweb"
	"github.com/geksong/godemo/utils"
	"io/ioutil"
)

func main() {
	fmt.Println("hello")
	fmt.Println(utils.BigdecimalMulty(5))
	m1 := 10
	m2 := 5
	m1, m2 = utils.SwapTwoInt(m1, m2)
	fmt.Printf("m1: %d , m2: %d", m1, m2)
	filecontent, fileerr := ioutil.ReadFile("/home/geksong/.viminfo")
	if fileerr != nil {
		fmt.Println("Some Things going wrong" + fileerr.Error())
		return
	}
	fmt.Println(filecontent)

	goweb.HttpHandler()

	// init struct
	//um := utils.UserMes{"Name": "geksong", "Age": 3}
}
