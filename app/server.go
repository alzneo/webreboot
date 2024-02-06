package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed embedWww
var embedFS embed.FS

func reboot(w http.ResponseWriter, r *http.Request) {
	if err := execute(argKeys.reboot); err != nil {
		fmt.Fprintln(w, err.Error())
	} else {
		http.Redirect(w, r, "/cancel.html", http.StatusSeeOther)
	}
}

func shutdown(w http.ResponseWriter, r *http.Request) {

	if err := execute(argKeys.poweroff); err != nil {
		fmt.Fprintln(w, err.Error())
	} else {
		http.Redirect(w, r, "/cancel.html", http.StatusSeeOther)
	}
}

func cancel(w http.ResponseWriter, r *http.Request) {
	if err := execute(argKeys.abort); err != nil {
		fmt.Fprintln(w, err.Error())
	} else {
		http.Redirect(w, r, "/done.html", http.StatusSeeOther)
	}
}

func serve() {
	http.HandleFunc("/reboot", reboot)
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/cancel", cancel)

	embedRoot, _ := fs.Sub(embedFS, "embedWww")
	http.Handle("/", http.FileServer(http.FS(embedRoot)))

	err := http.ListenAndServe(bind, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
