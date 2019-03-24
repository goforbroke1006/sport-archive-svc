package endpoint

type GetSportArgs struct {
	Name string `json:"name"`
}

type GetSportResult struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type GetParticipantArgs struct {
	Name      string `json:"name"`
	SportName string `json:"sport_name"`
}

type GetParticipantResult struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	SportID uint64 `json:"sport_id"`
}
