package entity

import (
	"EduKita/modules/core/utils"
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	FirebaseId     string    `json:"firebaseId"`
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
            )
        `

	operation := func(tx *sql.Tx) error {
		_, err := tx.Exec(query)
		return err
	}

	if err := utils.WithTransaction(db, operation); err != nil {
		panic(err)
	}
}
