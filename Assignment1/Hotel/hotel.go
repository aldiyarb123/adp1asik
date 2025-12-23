package Hotel

import "fmt"

type Room struct {
	RoomNumber     string
	Type           string
	PricePerNight  float64
	IsOccupied     bool
}

type Hotel struct {
	Rooms map[string]Room
}

func NewHotel() Hotel {
	return Hotel{Rooms: make(map[string]Room)}
}

func (h *Hotel) AddRoom(room Room) {
	h.Rooms[room.RoomNumber] = room
}

func (h *Hotel) CheckIn(roomNumber string) {
	room := h.Rooms[roomNumber]
	room.IsOccupied = true
	h.Rooms[roomNumber] = room
}

func (h *Hotel) CheckOut(roomNumber string) {
	room := h.Rooms[roomNumber]
	room.IsOccupied = false
	h.Rooms[roomNumber] = room
}

func (h *Hotel) ListVacantRooms() {
	for _, room := range h.Rooms {
		if !room.IsOccupied {
			fmt.Println(room)
		}
	}
}

func RunHotelMenu() {
	hotel := NewHotel()

	hotel.AddRoom(Room{"101", "Single", 100, false})
	hotel.AddRoom(Room{"102", "Double", 150, false})

	fmt.Println("Vacant rooms:")
	hotel.ListVacantRooms()

	hotel.CheckIn("101")
	fmt.Println("After check-in:")
	hotel.ListVacantRooms()
}