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

// User will hold user details
type User struct {
	ID       int64    `json:"id,omitempty" datastore:"-" db:"-"`
	EMAIL    string   `json:"email"`
	PASSWORD string   `json:"password"`
	ACTIVE   int      `json:"active"`
	ROLES    []string `json:"roles"`
}

// GameDetail will hold each news data
type GameDetail struct {
	ID         int64     `json:"id" datastore:"-" db:"-"`
	USERID     string    `json:"userId"`
	GAME       string    `json:"game"`
	DIFFICULTY string    `json:"difficulty"`
	STARTTIME  time.Time `json:"startTime,omitempty"`
	ENDTIME    time.Time `json:"endTime,omitempty"`
	SCORE      int       `json:"score,omitempty"`
}

func init() {
	http.HandleFunc("/", generateReport)
	http.HandleFunc("/report", generateReport)
}
