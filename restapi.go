package main

import (
	"Shellback.nl/Restapi/router"
)

func main() {
	r := router.RequestHandler()
	r.Run("176.57.189.22:8002")
}
