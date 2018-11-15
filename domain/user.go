package domain

// User describe user in system
type User struct {
	Model
	Name  string `json:"name"`
	Email string `json:"email"`
}
