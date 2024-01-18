package server

func (s *Server) findRoomById(Id string) (*Room, error) {
	var foundRoom *Room

	for room := range s.rooms {
		if room.GetId() == Id {
			foundRoom = room
			return foundRoom, nil
		}
	}

	foundRoom, err = s.runRoomFromRepository(Id)
	if err != nil {
		return nil, err
	}

	return foundRoom, nil

}
