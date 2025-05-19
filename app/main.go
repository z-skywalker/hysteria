package main

import (
	"log"
	"os"
)

func main() {
	f, _ := os.OpenFile("D:\\tools\\hysteria\\debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(f)

	run()
}
