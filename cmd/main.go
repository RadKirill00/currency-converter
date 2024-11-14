package main

import (
	"currency-converter/internal/config"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)

}




