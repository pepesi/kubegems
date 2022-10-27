package main

import (
	"kubegems.io/kubegems/pkg/apiserver/infrastructure/orm"
	"kubegems.io/kubegems/pkg/apiserver/interfaces"
)

func main() {
	db := orm.Init()
	println("start server")
	if e := interfaces.InitHTTPServer(db); e != nil {
		panic(e)
	}
}
