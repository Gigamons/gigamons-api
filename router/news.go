package router

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Gigamons/common/helpers"
	"github.com/Gigamons/common/logger"
	"github.com/gorilla/mux"
)

// Entry is the entry of the incomming news.
type Entry struct {
	ID          int
	Timestamp   time.Time
	Title       string
	Image       string
	Description string
	Text        string
}

// News is the News router and sends the News as Json encoded.
func News(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	PageT := fmt.Sprintf("%s", vars["page"])
	SizeT := r.URL.Query().Get("s")

	Page := 0  // Default offset (page 1)
	Size := 10 // Default size (10xEntrys)

	if PageT != "" {
		var err error
		Page, err = strconv.Atoi(PageT)
		if err != nil {
			JSONAnswer(504, "Page is not an int", w, r)
			return
		}
	}

	if SizeT != "" {
		var err error
		Size, err = strconv.Atoi(SizeT)
		if err != nil {
			JSONAnswer(500, "Size is not an int", w, r)
			return
		}
	}
	rows, err := helpers.DB.Query("SELECT id, date, title, image, description, text FROM news LIMIT ? OFFSET ?", Size, Size*Page)
	if err != nil {
		fmt.Println(err)
		JSONAnswer(500, "Serverside exception!", w, r)
		return
	}
	defer rows.Close()
	var o []Entry
	for rows.Next() {
		tr := Entry{}
		tmp := ""
		err := rows.Scan(&tr.ID, &tmp, &tr.Title, &tr.Image, &tr.Description, &tr.Text)
		if err != nil {
			log.Fatal(err)
		}
		if tmp == "0000-00-00 00:00:00" || tmp == "" {
			tmp = "0001-01-01 00:00:00"
		}
		tr.Timestamp, err = time.Parse("2006-01-02 15:04:05", tmp)
		if err != nil {
			logger.Errorln(err)
		}
		o = append(o, tr)
	}
	JSONAnswer(200, o, w, r)
}
