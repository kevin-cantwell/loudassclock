package main

import (
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/kevin-cantwell/loudassclock/phantom"
	"github.com/timehop/golog/log"
)

var pool phantom.RenderServer

func init() {
	pool = phantom.NewRenderServerPool(4)
}

func main() {
	if err := pool.Start(); err != nil {
		log.Fatal("loudassclock/main", "Failed to start phantomjs servers", "error", err)
	}

	prepareForShutdownDown()

	r := mux.NewRouter()
	r.HandleFunc("/{tzCode}/clock.png", ClockRenderHandler).Methods("GET")
	r.HandleFunc("/images/{file}", ImageHandler).Methods("GET")
	r.HandleFunc("/{tzCode:.*}", ClockHandler).Methods("GET")
	http.Handle("/", r)
	port := os.Getenv("PORT")
	log.Info("loudassclock/main", "Starting server...", "port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("loudassclock/main", "Failed to start server", "error", err)
	}
}

func prepareForShutdownDown() {
	// Make sure to kill all spawned processes if this proc gets killed
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		log.Info("loudassclock/main", "Received os.Signal", "signal", <-sig)
		pool.Shutdown()
		os.Exit(1)
	}()
}

func ClockHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	tzCode := vars["tzCode"]
	t, err := template.ParseFiles("clock.html")
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}
	p := struct{ Timezone string }{Timezone: tzCode}
	t.Execute(response, &p)
}

func ImageHandler(response http.ResponseWriter, request *http.Request) {
	filePath := filepath.Join("images", request.URL.Path[len("/images/"):])
	http.ServeFile(response, request, filePath)
}

func ClockRenderHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	tzCode := vars["tzCode"]
	body, err := pool.RenderClock(tzCode)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("500 Internal Server Error"))
		return
	}
	response.Header().Set("Content-Type", "image/png")
	response.WriteHeader(http.StatusOK)
	response.Write(body)
}
