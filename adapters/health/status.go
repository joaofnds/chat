package health

const (
	StatusUp   = "up"
	StatusDown = "down"
)

type Status struct {
	Status string `json:"status"`
}

func NewStatus(err error) Status {
	if err != nil {
		return Status{Status: StatusDown}
	}
	return Status{Status: StatusUp}
}

func (s Status) IsUp() bool {
	return s.Status == StatusUp
}
