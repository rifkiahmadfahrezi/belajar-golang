package routing

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

const port string = ":8080"

type Student struct {
	Nama  string
	Kelas string
}

func handleStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var filePath = path.Join("views/home", "index.html")
	var tmpl, err = template.ParseFiles(filePath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	nama := vars["nama"]
	kelas := vars["kelas"]

	var data = map[string]interface{}{
		"title": "Belajar golang web",
		"name":  "Nama saya : " + nama + ", Kelas : " + kelas,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func StartRouting() {
	r := mux.NewRouter()

	// handling static assets
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ini Halaman Home"))
	})
	r.HandleFunc("/student/{nama}/{kelas}", handleStudent)

	http.Handle("/", r)

	fmt.Println("Routing server started at http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
