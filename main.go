package main

import "web-shortlink/web"

func main() {
	app := &web.App{}
	app.Initalize()
	app.Run("127.0.0.1:8080")
}
