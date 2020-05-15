package veego

import (
	"fmt"
	"net/http"
	"time"

	h "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	ConfigFile string
	ConfigType string
}

func NewVeegoServer(configFile string, configType string) *Server {
	server := &Server{
		ConfigFile: configFile,
		ConfigType: configType,
	}
	return server
}

func (s *Server) Run() error {
	appConfig := NewAppConfig()
	conf := &AppConfig{}
	var err error
	switch s.ConfigType {
	case "env": 
		conf, err = appConfig.LoadEnv(s.ConfigFile)
		if err != nil {
			return err
		}
	case "yml":
		conf, err = appConfig.LoadYML(s.ConfigFile)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("config file type not supported")
	}
	fmt.Printf("Appliction running on %s:%s...\n", conf.Host, conf.Port)
	router := NewRouter(mux.NewRouter())
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: h.CORS(h.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Origin"}),
			h.AllowedMethods([]string{"GET"}),
			h.AllowedOrigins([]string{"*"}))(router.Route),
	}
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil 
}
