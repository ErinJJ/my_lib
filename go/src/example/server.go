package main

import (
	"code.byted.org/gopkg/logs"
	"code.byted.org/kite/kite"
)

func init() {

}

func main() {

	kite.Init()

	logs.Error("%v", kite.Run())
	logs.Stop()
}
