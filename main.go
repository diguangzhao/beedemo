package main

import (
	_ "beedemo/routers"
	_ "beedemo/hookfunc"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

