package main

import (
	"fmt"
	"url-shortener/internal/config"
)


func main(){

	cfg := config.MustLoad("../config/config.yaml")
	fmt.Println(cfg)
}