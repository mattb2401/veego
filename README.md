# **Veego** Microservice framework

 This framework helps you build a web microservice in minutes

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See how to use it in your project.

### Prerequisites

* Go 1.11 and above 


### Installing

A step by step series of examples that tell you how to get a development env running

Say what the step will be

```
go get github.com/mattb2401/veego
```

### Examples

Create a simple application server 
```Go
import (
    "github.com/mattb2401/veego"
)

func main() {
    config := veego.NewAppConfig()
	conf, err := config.LoadEnv(".env")
	if err != nil {
		panic(err)
	}
	baseRouter := veego.NewRouter(mux.NewRouter())
	baseRouter.Get("/", func(w *http.ResponseWriter, r *http.Request){
        w.Write([]byte(`{"code": 200}`))
        return
    })
    server := veego.NewServer()
    if err := server.Run(); err != nil {
        panic(err)
    }
}
```
## Built With

* [GorillaMux](https://www.gorillatoolkit.org/pkg/mux) - The http router and url matcher
* [Gorm](https://gorm.io/) - The fantastic ORM library for Golang, aims to be developer friendly.
* [GodotEnv](https://github.com/joho/godotenv) - A Go port of Ruby's dotenv library (Loads environment variables from `.env`.) 
* [GorillaHandlers](https://github.com/gorilla/handlers) - A collection of useful middleware for Go HTTP services & web applications
* [AnyXML](https://github.com/clbanning/anyxml) - Marshal XML from map[string]interface{}, arrays, slices, and alpha/numeric values.
* [Parsrus](https://github.com/mattb2401/parsrus) - Return JSON and XML for Golang web apis

## Contributing

Give me a hand and please be nice

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/mattb2401/veego/tags). 

## Authors

See also the list of [contributors](https://github.com/mattb2401/veego/contributors) who participated in this project.


## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc