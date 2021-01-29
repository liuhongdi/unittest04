package main

import (
	"fmt"
	"github.com/liuhongdi/unittest04/controller"
	"github.com/liuhongdi/unittest04/global"
	"log"
)

//定义一个加法方法
func add(a, b int) int {
	return a + b
}

func init() {
	//mysql link
	err := global.SetupDBLink()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func main() {
	goods,err := controller.GoodsOne(1)
	if (err != nil){
		log.Fatalf("err:%v",err)
	} else {
		fmt.Println(goods)
	}
}


