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
	}
}

func shutdown(w http.ResponseWriter, r *http.Request) {

	if err := execute(argKeys.poweroff); err != nil {
		fmt.Fprintln(w, err.Error())
	}
}

func cancel(w http.ResponseWriter, r *http.Request) {
	if err := execute(argKeys.abort); err != nil {
		fmt.Fprintln(w, err.Error())
	}
}

func serve() {
	http.HandleFunc("/power/reboot", reboot)
	http.HandleFunc("/power/shutdown", shutdown)
	http.HandleFunc("/power/cancel", cancel)

	embedRoot, _ := fs.Sub(embedFS, "embedWww")
	http.Handle("/", http.FileServer(http.FS(embedRoot)))

	err := http.ListenAndServe(bind, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
