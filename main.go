package main

import (
	"log"
	"net/http"
)

func init() {
	MakeMigrations()
}

func main() {

	http.HandleFunc("/", ShowHomePage)
	http.HandleFunc("/about", ShowAboutPage)
	http.HandleFunc("/notes", GetNotes)
	http.HandleFunc("/add-note", AddNote)
	http.HandleFunc("/update-note/", CompleteNote)
	http.HandleFunc("/delete-note/", RemoveNote)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("🚀 Starting up on port 3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}

/* REFERENCES:
https://medium.com/@orlmonteverde/api-rest-con-go-golang-y-sqlite3-e378af30719c
https://medium.com/@back_to_basics/golang-template-1-bcb690165663
https://www.alexedwards.net/blog/working-with-cookies-in-go
https://lets-go.alexedwards.net/sample/02.07-html-templating-and-inheritance.html

https://htmx.org/attributes/hx-swap/
https://htmx.org/attributes/hx-target/
https://htmx.org/attributes/hx-boost/
https://hypermedia.systems/htmx-in-action/#_ajax_ifying_our_application
https://htmx.org/headers/hx-location/
https://htmx.org/essays/view-transitions/
https://hyperscript.org/
https://hypermedia.systems/book/contents/

https://github.com/moeenn/htmx-golang-demo
https://github.com/orlmonteverde/go-api-with-sqlite
https://github.com/NerdCademyDev/gophat
https://github.com/bugbytes-io/htmx-go-demo
https://github.com/awesome-club/go-htmx
https://github.com/marco-souza/marco.fly.dev
*/
