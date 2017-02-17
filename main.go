package main

import (
	"github.com/astaxie/beego"
	"github.com/jepsonwu/inchat_go/controllers/api"
)

func main() {
	beego.Router("/", &api.PeelController{})

	beego.Run()
}
