package objects

// Error - universal response from server
type Error struct {
	Timestamp string `json:"timestamp"` // 2020-09-18T11:00:01.211+0000
	Status    int    `json:"status"`    // 500
	Error     string `json:"error"`     // Internal Server Error
	Message   string `json:"message"`   // Error message ...
	Path      string `json:"path"`      // /api/v3/...
}
