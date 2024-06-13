package website

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

func StartPartialRender() {
	// Fungsi ParseGlob digunakan untuk memparsing semua file
	//  yang match dengan pattern/pola yang ditentukan (views/partialrender/*)
	// jike ingin menggunakan function ini, silahkan komen kode dibawah yg menggunakan ParseFiles
	// var tmpl, err = template.ParseGlob("views/partialrender/*")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// // parsing dengan template.ParseGlob
	// http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
	// 	var data = M{"name": "Batman"}
	// 	err = tmpl.ExecuteTemplate(w, "index", data)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// })

	// http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
	// 	var data = M{"name": "Batman"}
	// 	err = tmpl.ExecuteTemplate(w, "about", data)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// })

	//  Parsing dengan template.ParseFiles
	// jike ingin menggunakan function ini, silahkan komen kode diatas yg menggunakan parseGlob
	// Fungsi ini selain bisa digunakan untuk parsing satu buah file saja
	// (seperti yang sudah dicontohkan pada chapter sebelumnya), bisa digunakan untuk parsing banyak file.
	http.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		var data = M{
			"blogTitle":   "Ini contoh blog",
			"blogContent": "lorem ipsum",
		}
		var tmpl = template.Must(template.ParseFiles(
			"views/partialrender/blog.html",
			"views/partialrender/_header.html",
			"views/partialrender/_message.html",
		))

		var err = tmpl.ExecuteTemplate(w, "blog", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at http://localhost:9000")
	http.ListenAndServe(":9000", nil)

}
