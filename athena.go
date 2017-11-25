package athena

import (
	"net/http"
	"time"
)

// Response will be the structure that will be used of all responses
type Response struct {
	SUCCESS     bool         `json:"success"`
	MESSAGE     string       `json:"message,omitempty"`
	GAMEDETAILS []GameDetail `json:"gameDetails,omitempty"`
}

// GameDetail will hold each news data
type GameDetail struct {
	ID         int64     `json:"id" datastore:"-" db:"-"`
	USERNAME   string    `json:"username"`
	GAME       string    `json:"game"`
	DIFFICULTY string    `json:"difficulty"`
	STARTTIME  time.Time `json:"startTime,omitempty"`
	ENDTIME    time.Time `json:"endTime,omitempty"`
	SCORE      int       `json:"score,omitempty"`
}

var defaultPageSize = 30

func init() {
	// Get Game Details for Report
	http.HandleFunc("/report", generateReport)
	// Load Data
	http.HandleFunc("/load", loadData)

}
