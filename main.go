package main

import (
	"github.com/kataras/iris"
	"github.com/goshepherd/goshepherd/app/endpoint"
)

func main() {
	serve().Listen(":8000")
}

func serve() *iris.Framework {
	i := iris.New()
	return endpoint.Routes(i)
}
