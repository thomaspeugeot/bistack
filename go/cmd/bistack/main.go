package main

import (
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/thomaspeugeot/bistack/go/models"
	bistack_models "github.com/thomaspeugeot/bistack/go/models"
	bistack_stack "github.com/thomaspeugeot/bistack/go/stack"
	bistack_static "github.com/thomaspeugeot/bistack/go/static"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("bistack: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := bistack_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	stack := bistack_stack.NewStack(r, bistack_models.Bistack.ToString(),
		*unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, false)

	stage := stack.Stage
	// A routine that, every 5 seconds,
	// - flips the name of the "Foo 1" instance between "Foo 1" and "Foo 1*"
	// - commit the stage
	//
	// This to demonstrate the websocket function of the front
	go func() {

		time.Sleep(5 * time.Second)

		// get first element
		set_A := (*models.GetGongstructInstancesSet[models.Foo](stage))
		var foo *models.Foo

		for key, _ := range set_A {
			foo = key
			break
		}

		index := 0
		if foo != nil {
			for {
				time.Sleep(5 * time.Second)
				stage.Checkout()
				index++
				if index%2 == 1 {
					foo.Name = foo.Name + "*"
				} else {
					foo.Name = "Foo 1"
				}
				stage.Commit()
			}
		}

	}()

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
