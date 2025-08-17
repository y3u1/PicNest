package main

import (
	"PicNext/router"
)

func main() {
	r := router.SetRouter()
	r.Run(":8080")
}
