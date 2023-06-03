package TonWork

type UserProfile struct {
	Telefon     string `json:"telefon"`
	Position    string `json:"position"`
	Description string `json:"description"`
	Subscribe   string `json:"subscribe"`
	Companies   string `json:"companies"`
}

type UserPreson struct {
	Username      string `json:"username" binding:"required"`
	Password_hash string `json:"password_hash" binding:"required"`
}

type User struct {
	Person  UserPreson
	Email   string `json:"email" binding:"required"`
	Profile UserProfile
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Id      int    `json:"id" db:"id"`
}
