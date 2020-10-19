package veego

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/alexedwards/scs/v2"
	h "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type server struct {
	baseRouter *mux.Router
	session    *scs.SessionManager
}

var (
	host = "0.0.0.0"
	port = "8080"
)

func NewServer(baseRouter *mux.Router, session *scs.SessionManager) *server {
	return &server{
		baseRouter: baseRouter,
		session:    session,
	}
}

func (s *server) Run(args ...interface{}) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-c
		fmt.Printf("%v", oscall)
		cancel()
	}()
	if len(args) > 0 {
		host = args[0].(string)
		port = args[1].(string)
	}
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", host, port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: h.CORS(h.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Origin"}),
			h.AllowedMethods([]string{"GET"}),
			h.AllowedOrigins([]string{"*"}))(s.baseRouter),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("serve failed: %v \n", err.Error())
		}
	}()
	fmt.Printf("Appliction running on %s:%s...\n", host, port)
	<-ctx.Done()
	fmt.Printf("Appliction stopped \n")
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := server.Shutdown(ctxShutDown); err != nil {
		return err
	}
	return nil
}

func (s *server) RunWithSession(args ...interface{}) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-c
		fmt.Printf("%v", oscall)
		cancel()
	}()
	if len(args) > 0 {
		host = args[0].(string)
		port = args[1].(string)
	}
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", host, port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      s.session.LoadAndSave(s.baseRouter),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("serve failed: %v \n", err.Error())
		}
	}()
	fmt.Printf("Appliction running on %s:%s...\n", host, port)
	<-ctx.Done()
	fmt.Printf("Appliction stopped \n")
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := server.Shutdown(ctxShutDown); err != nil {
		return err
	}
	return nil
}
