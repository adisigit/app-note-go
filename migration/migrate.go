package main

import (
	"app-note-go/initializer"
	"app-note-go/models"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectDB()
}

func main() {
	db := initializer.DB

	db.AutoMigrate(&models.User{}, &models.Note{})

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto";`)

	db.Exec(`
		ALTER TABLE users
		ALTER COLUMN id SET DEFAULT gen_random_uuid();
	`)

	db.Exec(`
		ALTER TABLE notes
		ALTER COLUMN id SET DEFAULT gen_random_uuid();
	`)
}
