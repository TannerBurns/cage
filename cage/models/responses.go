package models

type Exception struct {
	Message string `json:"message"`
}

type JwtToken struct {
	Token      string `json:"token"`
	EmployeeID int    `json:"employee_id"`
}

type RespError struct {
	Error string `json:"error"`
}

type ManagerResp struct {
	PlayerID        int    `json:"player_id"`
	First           string `json:"first"`
	Last            string `json:"last"`
	Status          string `json:"status"`
	CheckedInTime   int    `json:"checkedin_time"`
	TotalTimePlayed int    `json:"total_time_played"`
	AmountOwed      int    `json:"amount_owed"`
}

type ManagerRoster struct {
	Responses []ManagerResp `json:"roster"`
}
