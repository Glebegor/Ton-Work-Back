package TonWork

type Post struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" binding:"required" db:"description"`
	Text        string `json:"text" binding:"required" db:"text"`
	Tags        string `json:"tags" db:"tags"`
	Rating      int    `json:"rating" binding:"required" db:"rating"`
}

type PostUpdate struct {
	Title       *string `json:"title" binding:"required"`
	Description *string `json:"description" binding:"required"`
	Text        *string `json:"text" binding:"required"`
	Tags        *string `json:"tags"`
}
