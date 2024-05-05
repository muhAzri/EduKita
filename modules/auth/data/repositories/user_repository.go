package repositories

import (
	"EduKita/modules/auth/data/model"
	"EduKita/modules/auth/domain/entity"
	"database/sql"
)

type UserRepository interface {
	CreateUser(user entity.User) (model.UserModel, error)
	GetUserByID(id string) (model.UserModel, error)
	GetUserByFirebaseId(id string) (model.UserModel, error)
	GetUserByEmail(email string) (model.UserModel, error)
	UpdateUser(user entity.User) error
	DeleteUser(id string) error
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (r *UserRepositoryImpl) CreateUser(user entity.User) (model.UserModel, error) {

	query := `INSERT INTO users (id, firebase_id, name, email, profile_picture, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id,firebase_id, name, email, profile_picture, created_at, updated_at`

	row := r.db.QueryRow(query, user.ID, user.FirebaseId, user.Name, user.Email, user.ProfilePicture, user.CreatedAt, user.UpdatedAt)

	var userModel model.UserModel

	err := row.Scan(&userModel.ID, &userModel.Name, &userModel.Email, &userModel.ProfilePicture, &userModel.CreatedAt, &userModel.UpdatedAt)

	if err != nil {
		return model.UserModel{}, err
	}

	return userModel, nil
}

func (r *UserRepositoryImpl) GetUserByID(id string) (model.UserModel, error) {

	query := `SELECT id,firebase_id, name, email, profile_picture, created_at, updated_at FROM users WHERE id = $1`

	row := r.db.QueryRow(query, id)

	var user model.UserModel

	err := row.Scan(&user.ID, &user.FirebaseId, &user.Name, &user.Email, &user.ProfilePicture, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return model.UserModel{}, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) GetUserByFirebaseId(id string) (model.UserModel, error) {

	query := `SELECT id,firebase_id, name, email, profile_picture, created_at, updated_at FROM users WHERE firebase_id = $1`

	row := r.db.QueryRow(query, id)

	var user model.UserModel

	err := row.Scan(&user.ID, &user.FirebaseId, &user.Name, &user.Email, &user.ProfilePicture, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return model.UserModel{}, nil
	}

	if err != nil {
		return model.UserModel{}, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) GetUserByEmail(email string) (model.UserModel, error) {

	query := `SELECT id,firebase_id, name, email, profile_picture, created_at, updated_at FROM users WHERE email = $1`

	row := r.db.QueryRow(query, email)

	var user model.UserModel

	err := row.Scan(&user.ID, &user.FirebaseId, &user.Name, &user.Email, &user.ProfilePicture, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return model.UserModel{}, err
	}

	return user, nil

}

func (r *UserRepositoryImpl) UpdateUser(user entity.User) error {

	query := `UPDATE users SET name = $1, email = $2, profile_picture = $3, updated_at = $4 WHERE id = $5`

	_, err := r.db.Exec(query, user.Name, user.Email, user.ProfilePicture, user.UpdatedAt, user.ID)

	if err != nil {
		return err
	}

	return nil

}

func (r *UserRepositoryImpl) DeleteUser(id string) error {

	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil

}
