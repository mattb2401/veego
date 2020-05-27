package veego

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	h "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type server struct {
	config     *AppConfig
	baseRouter *mux.Router
}

func NewServer(config *AppConfig, baseRouter *mux.Router) *server {
	server := &server{
		config:     config,
		baseRouter: baseRouter,
	}
	return server
}

func (s *server) Run() error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-c
		fmt.Printf("%v", oscall)
		cancel()
	}()
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", s.config.Host, s.config.Port),
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
	fmt.Printf("Appliction running on %s:%s...\n", s.config.Host, s.config.Port)
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
