package main

import "warm-up/internal/server"

func main() {
	srv := server.New()
	srv.Run()
}
