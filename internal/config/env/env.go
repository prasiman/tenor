package env

import (
	"fmt"

	"github.com/joho/godotenv"
)

// Loads environment variables from a .env file
func init() {
	fmt.Println("Load env")
	godotenv.Load()
}
