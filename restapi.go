package main

import (
	"Shellback.nl/Restapi/router"
)

func main() {
	r := router.RequestHandler()
	r.Run("localhost:8002")
}
