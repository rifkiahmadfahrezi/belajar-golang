package website

// coba golang untuk web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

const port string = ":8082"

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	// ambil path ke file index.html
	// path.Join(). Fungsi ini digunakan untuk menggabungkan folder atau
	// file atau keduanya menjadi sebuah path, dengan separator relatif terhadap OS yang digunakan.
	var filepath = path.Join("views/home", "index.html") // view/index.html

	//  template.ParseFiles(), digunakan untuk parsing file template (bertipe *template.Template)
	// return 2 data -> hasil dari parsing dan err (jika terjadi error)
	var tmpl, err = template.ParseFiles(filepath)

	if err != nil {
		// fungsi http.Error() digunakan untuk menandai HTTP req jika terjadi error
		// dengan cara mengirim response berupa kode error, dan bisa juga mengirim pesan error
		// http.StatusInternalServerError = status code 500
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Belajar golang web",
		"name":  "Rifki ahmad fahrezi",
	}

	//  method Execute() digunakan untuk menyisipkan data ke template yang
	//  nantinya akan di tampilkan ke browser.
	//  Data yang bisa disipkan ke view dalam bentuk struct, map, atau interface{}.
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandlerAbout(w http.ResponseWriter, r *http.Request) {
	message := "Ini halaman About"
	w.Write([]byte(message))
}

func StartServer() {

	// handle static asset
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	// handle routing
	http.HandleFunc("/", HandlerIndex)
	http.HandleFunc("/about", HandlerAbout)

	fmt.Println("Server started on port http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
