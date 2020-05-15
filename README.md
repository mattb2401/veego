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

Create a server 
```
import (
    "github.com/mattb2401/veego"
)

func main() {
    //Requires a config file path and config type  
    server := veego.NewServer()
}
```