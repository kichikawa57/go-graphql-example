package main

import "github.com/kichikawa/router"

func main() {
	r := router.SetupRouter()
	r.Run("8080")
}
