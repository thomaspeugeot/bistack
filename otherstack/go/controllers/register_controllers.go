// generated code - do not edit
package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/thomaspeugeot/bistack/otherstack/go/orm"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// registerControllers register controllers
func registerControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/thomaspeugeot/bistack/otherstack/go")
	{ // insertion point for registrations
		v1.GET("/v1/bars", GetController().GetBars)
		v1.GET("/v1/bars/:id", GetController().GetBar)
		v1.POST("/v1/bars", GetController().PostBar)
		v1.PATCH("/v1/bars/:id", GetController().UpdateBar)
		v1.PUT("/v1/bars/:id", GetController().UpdateBar)
		v1.DELETE("/v1/bars/:id", GetController().DeleteBar)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)

		v1.GET("/v1/ws/stage", GetController().onWebSocketRequestForBackRepoContent)

		v1.GET("/v1/stacks", GetController().stacks)
	}
}

func (controller *Controller) stacks(c *gin.Context) {

	var res []string

	for k := range controller.Map_BackRepos {
		res = append(res, k)
	}

	c.JSON(http.StatusOK, res)
}

// onWebSocketRequestForBackRepoContent is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequestForBackRepoContent(c *gin.Context) {

	// log.Println("Stack github.com/thomaspeugeot/bistack/otherstack/go, onWebSocketRequestForBackRepoContent")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://localhost:8080" || origin == "http://localhost:4200"
		},
	}

	wsConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wsConnection.Close()

	// Create a context that is canceled when the connection is closed
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}

	log.Printf("Stack github.com/thomaspeugeot/bistack/otherstack/go: stack path: '%s', new ws index %d",
		stackPath, controller.listenerIndex,
	)
	controller.listenerIndex++

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/bistack/otherstack/go, Unkown stack", stackPath)
	}
	updateCommitBackRepoNbChannel := backRepo.SubscribeToCommitNb(ctx)

	// Start a goroutine to read from the WebSocket to detect disconnection
	go func() {
		for {
			// ReadMessage is used to detect client disconnection
			_, _, err := wsConnection.ReadMessage()
			if err != nil {
				log.Println("github.com/thomaspeugeot/bistack/otherstack/go", stackPath, "WS client disconnected:", err)
				cancel() // Cancel the context
				return
			}
		}
	}()

	backRepoData := new(orm.BackRepoData)
	orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)

	err = wsConnection.WriteJSON(backRepoData)
	if err != nil {
		log.Println("github.com/thomaspeugeot/bistack/otherstack/go:\n",
			"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	} else {
		log.Println(time.Now().Format(time.RFC3339Nano), "github.com/thomaspeugeot/bistack/otherstack/go: 1st sent backRepoData of stack:", stackPath)
	}
	for {
		select {
		case <-ctx.Done():
			// Context canceled, exit the loop
			return
		default:
			for nbCommitBackRepo := range updateCommitBackRepoNbChannel {
				_ = nbCommitBackRepo

				backRepoData := new(orm.BackRepoData)
				orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)

				// Set write deadline to prevent blocking indefinitely
				wsConnection.SetWriteDeadline(time.Now().Add(10 * time.Second))

				// Send backRepo data
				err = wsConnection.WriteJSON(backRepoData)
				if err != nil {
					log.Println("github.com/thomaspeugeot/bistack/otherstack/go:\n", stackPath,
						"client no longer receiver web socket message,closing websocket handler")
					fmt.Println(err)
					cancel() // Cancel the context
					return
				} else {
					log.Println(time.Now().Format(time.RFC3339Nano), "github.com/thomaspeugeot/bistack/otherstack/go: sent backRepoData of stack:", stackPath)
				}
			}
		}
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func (controller *Controller) GetLastCommitFromBackNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/bistack/otherstack/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func (controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastPushFromFrontNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/bistack/otherstack/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}