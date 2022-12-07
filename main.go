package main

import (
	"log"
	"web-shortlink/app/interfaces"
	"web-shortlink/web"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	s := web.NewServer()
	service := interfaces.InitializeService()
	s.Register(service)
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
