package TonWork

type Post struct {
	Title       string `json:'title'binding:'required'`
	Description string `json:'description'binding:'required'`
	Text        string `json:'text'binding:'required'`
	Tags        string `json:'tags'`
	Rating      int    `json:'rating'`
	Id          int    `db:'id'`
}
