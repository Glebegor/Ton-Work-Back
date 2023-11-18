package structint

type Work struct {
	Id              int    `json:"id" db:"id"`
	Title           string `json:"title" binding:"required" db:"title"`
	Description     string `json:"description" binding:"required" db:"description"`
	Text            string `json:"text" binding:"required" db:"text"`
	Tags            string `json:"tags" db:"tags"`
	Technologies    string `json:"technologies" db:"technologies"`
	Company         string `json:"company" db:"company"`
	Price           int    `json:"price" binding:"required" db:"price"`
	ExperienceLevel string `json:"experienceLevel" binding:"required" db:"experienceLevel"`
	Type_of_job     string `json:"type_of_job" db:"type_of_job"`
	Invites         int    `json:"invites" db:"invites"`
	Rating          int    `json:"rating" binding:"required" db:"rating"`
}
type WorkUpdate struct {
	Title           *string `json:"title"`
	Description     *string `json:"description"`
	Text            *string `json:"text"`
	Tags            *string `json:"tags"`
	Technologies    *string `json:"technologies" `
	Company         *string `json:"company" `
	Price           *int    `json:"price"`
	ExperienceLevel *string `json:"experienceLevel"`
	Type_of_job     *string `json:"type_of_job" `
}
