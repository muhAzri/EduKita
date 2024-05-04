package entity

import (
	"EduKita/modules/core/utils"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	FirebaseID     string    `json:"firebaseId"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	ProfilePicture string    `json:"profilePicture"`
	CreatedAt      int64     `json:"createdAt"`
	UpdatedAt      int64     `json:"updatedAt"`
}

func MigrateUser(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id uuid PRIMARY KEY,
		firebase_id TEXT UNIQUE NOT NULL,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		profile_picture TEXT,
		created_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM CURRENT_TIMESTAMP) * 1000,
		updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM CURRENT_TIMESTAMP) * 1000
	);
	`

	t := reflect.TypeOf(User{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Name != "ID" && field.Name != "CreatedAt" && field.Name != "UpdatedAt" {
			query += fmt.Sprintf("ALTER TABLE IF EXISTS users ADD COLUMN IF NOT EXISTS %s TEXT;\n", utils.ToSnakeCase(field.Name))
		}
	}

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
