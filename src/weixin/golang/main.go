package main

import (
	"os"

	"github.com/bitwormhole/passwordbox"
	"github.com/starter-go/starter"
)

func main() {

	a := os.Args
	m := passwordbox.ModuleForWeixin()
	i := starter.Init(a)

	i.MainModule(m)

	i.WithPanic(true).Run()
}
