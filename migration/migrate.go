package migration

import (
	"app-note-go/initializer"
	"app-note-go/models"
)

func Migrate() {
	db := initializer.DB

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto";`)

	db.AutoMigrate(&models.User{}, &models.Note{})

	db.Exec(`
		ALTER TABLE users
		ALTER COLUMN id SET DEFAULT gen_random_uuid();
	`)

	db.Exec(`
		ALTER TABLE notes
		ALTER COLUMN id SET DEFAULT gen_random_uuid();
	`)
}
