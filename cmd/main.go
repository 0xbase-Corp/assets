package main

import (
	"github.com/0xbase-Corp/assets/internal/manager"
)

func main() {
	manager.InitCommands()
	manager.Execute()
}
