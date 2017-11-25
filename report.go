package athena

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/net/context"

	// Imports the Google Cloud Datastore client package.
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

// generateReport sends data for generating a report
func generateReport(w http.ResponseWriter, r *http.Request) {
	var tmpResponse Response
	switch r.Method {
	case "GET": // Get the news items
		ctx := appengine.NewContext(r)
		log.Infof(ctx, "GET API CALLED")

		pageSize, err := strconv.Atoi("0" + r.URL.Query().Get("pagesize"))
		if err != nil {
			pageSize = defaultPageSize
		}
		if pageSize == 0 {
			pageSize = defaultPageSize
		}

		log.Infof(ctx, "PageSize: %v", pageSize)

		// POSTDATE is a date time field in DB (datastore)
		// Get Recent items by default
		q := datastore.NewQuery("GAMES").
			Limit(pageSize)

		// [START getall]
		gameDetails := make([]GameDetail, 0, pageSize)
		keys, err := q.GetAll(ctx, &gameDetails)
		if err != nil {
			tmpResponse.errMsg(ctx, err, "Unable to fetch data", w, r)
			return
		}
		for i := 0; i < len(keys); i++ {
			gameDetails[i].ID = keys[i].IntID()
		}
		tmpResponse.GAMEDETAILS = gameDetails
		tmpResponse.successMsg(ctx, "Fetched news from DB", w, r)
		return
	case "POST": // Add a news using JSON - Temporary
		ctx := appengine.NewContext(r)
		log.Infof(ctx, "POST API CALLED %v", "")

		var form GameDetail
		form.STARTTIME.Format(time.RFC3339)
		form.ENDTIME.Format(time.RFC3339)

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&form)
		if err != nil {
			tmpResponse.errMsg(ctx, err, "Invalid JSON Received", w, r)
			return
		}
		defer r.Body.Close()
		key := datastore.NewKey(ctx, "GAMES", "", 0, nil)
		dbresp, err := datastore.Put(ctx, key, &form)
		if err != nil {
			tmpResponse.errMsg(ctx, err, "Unable to add to DB", w, r)
			return
		}
		log.Infof(ctx, "DB response: %v", dbresp.IntID())
		form.ID = dbresp.IntID()
		var gameDetails []GameDetail
		gameDetails = append(gameDetails, form)
		tmpResponse.GAMEDETAILS = gameDetails
		tmpResponse.successMsg(ctx, "News Item added to DB", w, r)
		return
	default:
		ctx := appengine.NewContext(r)
		tmpResponse.errMsg(ctx, errors.New(" ERR INVALID REST METHOD "), "Invalid REST Request type", w, r)
	}
}

func loadData(w http.ResponseWriter, r *http.Request) {
	var tmpResponse Response

	ctx := appengine.NewContext(r)
	log.Infof(ctx, "POST API CALLED %v", "")
	for i := 0; i < 100; i++ {
		var form GameDetail
		form.STARTTIME.Format(time.RFC3339)
		form.ENDTIME.Format(time.RFC3339)

		ran := random(1, 1000)
		ranstr := strconv.Itoa(ran)
		switch ran % 3 {
		case 0:
			form.DIFFICULTY = "EASY"
		case 1:
			form.DIFFICULTY = "MEDIUM"
		case 2:
			form.DIFFICULTY = "HARD"
		default:
			form.DIFFICULTY = "TOUGH"
		}
		form.USERNAME = "TST" + ranstr
		form.GAME = "GAME" + ranstr
		form.STARTTIME = time.Now()
		d := time.Duration(ran % 30)
		form.ENDTIME = time.Now().Add(time.Minute * d)

		key := datastore.NewKey(ctx, "GAMES", "", 0, nil)
		dbresp, err := datastore.Put(ctx, key, &form)
		if err != nil {
			tmpResponse.errMsg(ctx, err, "Unable to add to DB", w, r)
			return
		}
		//log.Infof(ctx, "DB response: %v", dbresp.IntID())
		form.ID = dbresp.IntID()
		var gameDetails []GameDetail
		gameDetails = append(gameDetails, form)
		tmpResponse.GAMEDETAILS = gameDetails
		// tmpResponse.successMsg(ctx, "News Item added to DB", w, r)
	}
	log.Infof(ctx, "INSERT COMPLETE")
}
func (tmpResponse *Response) errMsg(ctx context.Context, err error, msg string, w http.ResponseWriter, r *http.Request) {
	tmpResponse.SUCCESS = false
	tmpResponse.MESSAGE = msg
	response, _ := json.Marshal(tmpResponse)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(response))
	log.Errorf(ctx, "%s - %s", msg, err.Error())
	return
}

func (tmpResponse *Response) successMsg(ctx context.Context, msg string, w http.ResponseWriter, r *http.Request) {
	tmpResponse.SUCCESS = true
	tmpResponse.MESSAGE = msg
	response, _ := json.Marshal(tmpResponse)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(response))
	log.Infof(ctx, "%s - %s", msg, string(response))
}

//GenerateRandom to generate a random number between min and max
func random(min, max int) int {
	rand.Seed(time.Now().UnixNano() / int64(time.Millisecond))
	return rand.Intn(max-min) + min
}
