package TonWork

type Post struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Text        string `json:"text" binding:"required"`
	Tags        string `json:"tags"`
	Rating      int    `json:"rating" binding:"required"`
	Id          int    `json:"id" db:"id"`
}

type PostUpdate struct {
	Title       *string `json:"title" binding:"required"`
	Description *string `json:"description" binding:"required"`
	Text        *string `json:"text" binding:"required"`
	Tags        *string `json:"tags"`
}
