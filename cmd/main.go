package main

import (
	"github.com/bingodfok/freshguard-boot/cmd/application"
	_ "time/tzdata"
)

func main() {
	appCtx := application.NewApplication()
	application.Run(appCtx)
}
