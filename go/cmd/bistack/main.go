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

	// setup stack1Instance1
	stack1Instance1 := bistack_stack.NewStack(r, bistack_models.Bistack_Stack1_Instance1.ToString(),
		"stage_stack1_instance1.go", "", "", *embeddedDiagrams, false)

	// A routine that, every 5 seconds,
	// - flips the name of the "Foo 1" instance between "Foo 1" and "Foo 1*"
	// - commit the stage
	//
	// This to demonstrate the websocket function of the front
	go func() {

		time.Sleep(1 * time.Second)

		// get first element
		set := (*models.GetGongstructInstancesSet[models.Foo](stack1Instance1.Stage))
		var foo *models.Foo

		for key, _ := range set {
			foo = key
			break
		}

		index := 0
		if foo != nil {
			for {
				time.Sleep(2 * time.Second)

				index++
				if index%2 == 1 {
					foo.Name = "Stack 1 Instance 1" + "*"
				} else {
					foo.Name = "Stack 1 Instance 1"
				}
				stack1Instance1.Stage.Commit()
			}
		}
	}()

	// setup stack1Instance2
	stack1Instance2 := bistack_stack.NewStack(r, bistack_models.Bistack_Stack1_Instance2.ToString(),
		"stage_stack1_instance2.go", "", "", *embeddedDiagrams, false)

	// A routine that, every 5 seconds,
	// - flips the name of the "Foo 1" instance between "Foo 1" and "Foo 1*"
	// - commit the stage
	//
	// This to demonstrate the websocket function of the front
	go func() {

		time.Sleep(2 * time.Second)

		// get first element
		set := (*models.GetGongstructInstancesSet[models.Foo](stack1Instance2.Stage))
		var foo *models.Foo

		for key, _ := range set {
			foo = key
			break
		}

		index := 0
		if foo != nil {
			for {
				time.Sleep(2 * time.Second)
				index++
				if index%2 == 1 {
					foo.Name = "Stack 1 Instance 2" + "*"
				} else {
					foo.Name = "Stack 1 Instance 2"
				}
				stack1Instance2.Stage.Commit()
			}
		}
	}()

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
