package models

import (
	"main/internal/config/database"
)

var db, _ = database.InitDatabase()
