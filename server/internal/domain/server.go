package domain

type Server struct {
	users      map[*User]bool
	register   chan *User
	unregister chan *User
	broadcast  chan []byte
	rooms      map[*Room]bool
	// roomRepository RoomRepository
	// userRepository UserRepository
}

type WsServer interface {
	RegisterUser(u *User) (*User, error)
	UnregisterUser(u *User) (*User, error)
	BroadcastServiceMessage(m *ServiceMessage) error
	// FindUserById(id string) (*User, error)
	// CreateRoom(r *Room) (*Room, error)
	// FindRoomById(id string) (*Room, error)
	// ListOnlineUsers() ([]*User, error)
}
