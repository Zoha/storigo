package main

import (
	"fmt"

	"github.com/zoha/storigo/common/config"
)

func main() {
	// read the dot env configs
	config.ReadConfigs()

	fmt.Printf("Hello %s", config.Get(config.APP_NAME))
}
