package main

import (
	_ "main/internal/config/env"
	"main/internal/main/routes"
)

func main() {
	routes.SetupRouter()
}
