package main

import (
	"http"
	"conf"
)

func main(){
	conf.GetVhosts("config/server.json")
	http.Do()
}




