package main

import (
	"github.com/Math-dev264/AutoTrade/Backend/src/configs"
	"github.com/Math-dev264/AutoTrade/Backend/src/router"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	router.Initialize()
}
