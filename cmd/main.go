package main

import "github.com/bingodfok/freshguard-boot/cmd/application"

func main() {
	appCtx := application.NewApplication()
	application.Run(appCtx)
}
