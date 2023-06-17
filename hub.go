package TonWork

type Room struct {
	Id      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms map[string]*Room
}

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[string]*Room),
	}
}
