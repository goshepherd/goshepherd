package main

import (
	"github.com/kataras/iris"
	"github.com/youkyll/goshepherd/app/endpoint"
)

func main() {
	endpoint.Routes()

	iris.Listen(":8000")
}
