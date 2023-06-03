package TonWork

type Work struct {
	Title           string   `json:"title" binding:"required"`
	Description     string   `json:"description" binding:"required"`
	Text            string   `json:"text" binding:"required"`
	Tags            []string `json:"tags" `
	Technologies    []string `json:"technologies" `
	Company         string   `json:"company" `
	Price           int      `json:"price" binding:"required"`
	ExperienceLevel string   `json:"experienceLevel" binding:"required"`
	Type_of_job     string   `json:"type_of_job" `
	Invites         int      `json:"invites" `
	Rating          int      `json:"rating" binding:"required"`
	Id              int      `json:"id" db:'id'`
}
