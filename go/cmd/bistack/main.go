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

	otherstack_models "github.com/thomaspeugeot/bistack/otherstack/go/models"
	otherstack_stack "github.com/thomaspeugeot/bistack/otherstack/go/stack"
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
				time.Sleep(4 * time.Second)

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
				time.Sleep(4 * time.Second)
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

	otherstack_instance1 := otherstack_stack.NewStack(r, otherstack_models.Otherstack_Instance1.ToString(),
		"stage_stack2_instance1.go", "", "", *embeddedDiagrams, true)
	_ = otherstack_instance1

	go func() {

		time.Sleep(3 * time.Second)

		// get first element
		set := (*otherstack_models.GetGongstructInstancesSet[otherstack_models.Bar](otherstack_instance1.Stage))
		var bar *otherstack_models.Bar

		for key, _ := range set {
			bar = key
			break
		}

		index := 0
		if bar != nil {
			for {
				time.Sleep(4 * time.Second)
				index++
				if index%2 == 1 {
					bar.Name = "Stack 2 Instance 1" + "*"
				} else {
					bar.Name = "Stack 2 Instance 1"
				}
				otherstack_instance1.Stage.Commit()
			}
		}
	}()

	otherstack_instance2 := otherstack_stack.NewStack(r, otherstack_models.Otherstack_Instance2.ToString(),
		"stage_stack2_instance1.go", "", "", *embeddedDiagrams, true)
	_ = otherstack_instance2

	go func() {

		time.Sleep(4 * time.Second)

		// get first element
		set := (*otherstack_models.GetGongstructInstancesSet[otherstack_models.Bar](otherstack_instance2.Stage))
		var bar *otherstack_models.Bar

		for key, _ := range set {
			bar = key
			break
		}

		index := 0
		if bar != nil {
			for {
				time.Sleep(4 * time.Second)
				index++
				if index%2 == 1 {
					bar.Name = "Stack 2 Instance 2" + "*"
				} else {
					bar.Name = "Stack 2 Instance 2"
				}
				otherstack_instance2.Stage.Commit()
			}
		}
	}()

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
