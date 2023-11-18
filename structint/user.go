package structint

type User struct {
	Username      string `json:"username" binding:"required"db:"username"`
	Password_hash string `json:"password_hash" binding:"required"db:"password_hash"`
	Email         string `json:"email" binding:"required"db:"email"`
	Telefon       string `json:"telefon"db:"telefon"`
	Position      string `json:"position"db:"position"`
	Description   string `json:"description"db:"description"`
	Subscribe     string `json:"subscribe"db:"subscribe"`
	Companies     string `json:"companies"db:"companies"`
	Name          string `json:"name"db:"name"`
	Surname       string `json:"surname"db:"surname"`
	Id            int    `json:"id" db:"id"`
}
