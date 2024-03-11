package main

import (
	"FoxTok/server/route"
	"FoxTok/server/sq"
)
func main(){
	sq.Init()
	route.Init()
}