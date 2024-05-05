package model

type UserModel struct {
	ID             string `json:"id"`
	FirebaseId     string `json:"firebaseId"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profilePicture"`
	CreatedAt      int64  `json:"createdAt"`
	UpdatedAt      int64  `json:"updatedAt"`
}

