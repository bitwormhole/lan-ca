package main

import (
	"os"

	"github.com/bitwormhole/lan-ca/backend/modules/lanca"
	"github.com/starter-go/starter"
)

func main() {
	m := lanca.Module()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
