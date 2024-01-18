package domain

type ServiceMessage struct {
	Action  string `json:"action"`
	Message string `json:"message"`
	Target  *Room  `json:"target"`
	Sender  User   `json:"sender"`
}
