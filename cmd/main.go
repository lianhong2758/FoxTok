package main

import (
	"FoxTok/config"
	"FoxTok/server/route"
	"FoxTok/server/sq"
)
func main(){
 config.ReadConfig("config.yml")
	sq.Init( )
	route.Init( )
}