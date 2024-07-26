package main

import (
	"github.com/hyst10/spinx/http"
	"github.com/hyst10/spinx/conf"
)

func main(){
	conf.GetVhosts("config/server.json")
	http.Do()
}




