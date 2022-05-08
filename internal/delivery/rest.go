package delivery

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/garixx/workshop-app/internal/authtoken/repository"
	"github.com/garixx/workshop-app/internal/domain"
	"github.com/garixx/workshop-app/internal/inventory"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type RestFrontEnd struct {
	i inventory.Inventory
}

var io inventory.Inventory

func (f RestFrontEnd) Start(ii inventory.Inventory) error {
	log.Println("it's me, rest frontend")
	io = ii
	r := mux.NewRouter()

	r.HandleFunc("/register", RegisterHandler).Methods("POST")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/version", VersionHandler).Methods("GET")

	srv := &http.Server{
		Addr:         viper.GetString("server.address") + ":" + viper.GetString("server.port"),
		WriteTimeout: time.Second * viper.GetDuration("server.timeout.write"),
		ReadTimeout:  time.Second * viper.GetDuration("server.timeout.read"),
		IdleTimeout:  time.Second * viper.GetDuration("server.timeout.idle"),
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c
	gracefulShutdown(srv)
	return nil
}

func gracefulShutdown(srv *http.Server) {
	//TODO: add DB closing and other necessary actions
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down ...")
	os.Exit(0)
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	if _, err := io.AuthToken.FetchToken(token); err != nil {
		if errors.Is(err, repository.InvalidTokenError) {
			http.Error(w, "invalid token", http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "{ \"version\": \"3.1\" }")
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var auth domain.User

	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := io.User.CreateUser(auth)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	neww, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(neww)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	getUser, err := io.User.GetUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := io.AuthToken.StoreToken(domain.AuthTokenParams{User: getUser})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	neww, err := json.Marshal(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(neww)
}
