package main

import "github.com/firmanmimang/go-api-trusted-news/config"

func main() {
	config.LoadConfig()
	config.ConnectDB()
}
