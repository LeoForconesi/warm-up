package main

import "warm-up/internal/app"

func main() {
	srv := app.New()
	srv.Run()
}
