package room

type Room struct {
	Name     string
	RoomType string
}

func NewRoom(name, roomType string) *Room {
	return &Room{
		Name:     name,
		RoomType: roomType,
	}
}
