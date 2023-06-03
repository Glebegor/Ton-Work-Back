package TonWork

type User struct {
	Username      string `json:"username" binding:"required"`
	Password_hash string `json:"password_hash" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Telefon       string `json:"telefon"`
	Position      string `json:"position"`
	Description   string `json:"description"`
	Subscribe     string `json:"subscribe"`
	Companies     string `json:"companies"`
	Id            int    `json:"id" db:"id"`
}
