package main

import (
	"flag"
	"log"
	"strconv"
	"time"

	bm "github.com/thomaspeugeot/bistack/go/models"
	bistack_stack "github.com/thomaspeugeot/bistack/go/stack"
	bistack_static "github.com/thomaspeugeot/bistack/go/static"

	os "github.com/thomaspeugeot/bistack/otherstack/go/models"
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

	{
		stack1Instance1 := bistack_stack.NewStack(r, bm.Bistack_Stack1_Instance1.ToString(),
			"", "", "", *embeddedDiagrams, false).Stage

		foo := new(bm.Foo).Stage(stack1Instance1)
		updateNameBeat(1*time.Second, foo, bm.Bistack_Stack1_Instance1.ToString(), stack1Instance1)
	}

	{
		stack1Instance2 := bistack_stack.NewStack(r, bm.Bistack_Stack1_Instance2.ToString(),
			"", "", "", *embeddedDiagrams, false).Stage
		foo := new(bm.Foo).Stage(stack1Instance2)

		updateNameBeat(2*time.Second, foo, bm.Bistack_Stack1_Instance2.ToString(), stack1Instance2)
	}

	{
		otherstack_instance1 := otherstack_stack.NewStack(r, os.Otherstack_Instance1.ToString(),
			"", "", "", *embeddedDiagrams, true).Stage
		bar := new(os.Bar).Stage(otherstack_instance1)

		updateNameBeat(3*time.Second, bar, os.Otherstack_Instance1.ToString(), otherstack_instance1)
	}

	{
		otherstack_instance2 := otherstack_stack.NewStack(r, os.Otherstack_Instance2.ToString(),
			"", "", "", *embeddedDiagrams, true).Stage
		// get first element
		bar := new(os.Bar).Stage(otherstack_instance2)

		updateNameBeat(4*time.Second, bar, os.Otherstack_Instance2.ToString(), otherstack_instance2)
	}

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

type NamedStructInterface interface {
	SetName(name string)
}

type StageInterface interface {
	Commit()
}

// updateNameBeat toggles the bar's name between displayName and displayName+"*"
// every four seconds after an initial delay, committing the stage each time.
func updateNameBeat(initialDelay time.Duration, bar NamedStructInterface, displayName string,
	stage StageInterface) {
	go func() {
		time.Sleep(initialDelay)

		index := 0
		if bar != nil {
			for {
				time.Sleep(4 * time.Second)
				index++
				if index%2 == 1 {
					bar.SetName(displayName + "*********")
				} else {
					bar.SetName(displayName)
				}
				stage.Commit()
			}
		}
	}()
}
