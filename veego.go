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
	Config     *AppConfig
	BaseRouter *mux.Router
}

func NewServer(config *AppConfig, baseRouter *mux.Router) *server {
	server := &server{
		Config:     config,
		BaseRouter: baseRouter,
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
		Addr:         fmt.Sprintf("%s:%s", s.Config.Host, s.Config.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: h.CORS(h.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Origin"}),
			h.AllowedMethods([]string{"GET"}),
			h.AllowedOrigins([]string{"*"}))(s.BaseRouter),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("serve failed: %v \n", err.Error())
		}
	}()
	fmt.Printf("Appliction running on %s:%s...\n", s.Config.Host, s.Config.Port)
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
