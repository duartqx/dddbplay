package main

import (
	"log"
	"net/http"
	"os/exec"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/play", playerHandler)
	mux.HandleFunc("/form", formHandler)
	if err := http.ListenAndServe(":8088", mux); err != nil {
		log.Fatalln(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func playerHandler(w http.ResponseWriter, r *http.Request) {
	video := "/home/duartqx/Media/Videos/[EMBER] Dungeon Meshi - 09.mkv"
	log.Println(video)

	cmd := exec.Command("mpv", "--fs", video)
	log.Println("start", cmd)
	if err := cmd.Start(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("past start err", cmd)

	if err := cmd.Wait(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("cmd waited, OK")

	w.Write(([]byte("OK")))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	f := r.Form["filesNames"]
	log.Println("\n", f, len(f), "\n ")
}
