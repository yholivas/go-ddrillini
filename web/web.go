package web

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/yholivas/go-ddrillini/api"
)

var db *sql.DB

var home string
var imgRegexp = regexp.MustCompile(`^/img(/[^/]+)(/[^\.]+.png)$`)

const packDir = "/.stepmania-5.1/Songs"

func imgHandler(w http.ResponseWriter, r *http.Request) {
	img := imgRegexp.FindStringSubmatch(r.URL.Path)
	if img == nil {
		http.NotFound(w, r)
		return
	}
	imgPath := home + packDir + img[1] + img[2]
	fmt.Println(imgPath)
	http.ServeFile(w, r, imgPath)
}

// make a handler and a template to view a pack table
func viewPacks(w http.ResponseWriter, r *http.Request) {
	p, err := api.GetAllPacks(db)
	t, err := template.ParseFiles("web/packs.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, p)
}

func viewPack(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/pack/"):]
	id, _ := strconv.ParseInt(idStr, 10, 0)
	p, err := api.GetPack(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t, err := template.ParseFiles("web/pack.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, p)
}

func Serve(dbCon *sql.DB) {
	db = dbCon
	home, _ = os.UserHomeDir()
	http.HandleFunc("/style.css",
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "web/style.css")
		})
	// img handler
	http.HandleFunc("/img/", imgHandler)
	http.HandleFunc("/pack/", viewPack)
	http.HandleFunc("/", viewPacks)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
