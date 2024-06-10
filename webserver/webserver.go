package webserver

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const PORT string = ":8080"

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var siswa = map[string]string{
			"Name":    "Rifki ahmad fahrezi",
			"Message": "Hello world!",
		}

		var t, err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, siswa)
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This is about page")
	})

	fmt.Println("Server start on http://localhost" + PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
