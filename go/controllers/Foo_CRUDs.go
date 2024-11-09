// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/thomaspeugeot/bistack/go/models"
	"github.com/thomaspeugeot/bistack/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Foo__dummysDeclaration__ models.Foo
var __Foo_time__dummyDeclaration time.Duration

var mutexFoo sync.Mutex

// An FooID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFoo updateFoo deleteFoo
type FooID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FooInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFoo updateFoo
type FooInput struct {
	// The Foo to submit or modify
	// in: body
	Foo *orm.FooAPI
}

// GetFoos
//
// swagger:route GET /foos foos getFoos
//
// # Get all foos
//
// Responses:
// default: genericError
//
//	200: fooDBResponse
func (controller *Controller) GetFoos(c *gin.Context) {

	// source slice
	var fooDBs []orm.FooDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFoos", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/bistack/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFoo.GetDB()

	_, err := db.Find(&fooDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	fooAPIs := make([]orm.FooAPI, 0)

	// for each foo, update fields from the database nullable fields
	for idx := range fooDBs {
		fooDB := &fooDBs[idx]
		_ = fooDB
		var fooAPI orm.FooAPI

		// insertion point for updating fields
		fooAPI.ID = fooDB.ID
		fooDB.CopyBasicFieldsToFoo_WOP(&fooAPI.Foo_WOP)
		fooAPI.FooPointersEncoding = fooDB.FooPointersEncoding
		fooAPIs = append(fooAPIs, fooAPI)
	}

	c.JSON(http.StatusOK, fooAPIs)
}

// PostFoo
//
// swagger:route POST /foos foos postFoo
//
// Creates a foo
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFoo(c *gin.Context) {

	mutexFoo.Lock()
	defer mutexFoo.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFoos", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/bistack/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFoo.GetDB()

	// Validate input
	var input orm.FooAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create foo
	fooDB := orm.FooDB{}
	fooDB.FooPointersEncoding = input.FooPointersEncoding
	fooDB.CopyBasicFieldsFromFoo_WOP(&input.Foo_WOP)

	_, err = db.Create(&fooDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFoo.CheckoutPhaseOneInstance(&fooDB)
	foo := backRepo.BackRepoFoo.Map_FooDBID_FooPtr[fooDB.ID]

	if foo != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), foo)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, fooDB)
}

// GetFoo
//
// swagger:route GET /foos/{ID} foos getFoo
//
// Gets the details for a foo.
//
// Responses:
// default: genericError
//
//	200: fooDBResponse
func (controller *Controller) GetFoo(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFoo", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/bistack/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFoo.GetDB()

	// Get fooDB in DB
	var fooDB orm.FooDB
	if _, err := db.First(&fooDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var fooAPI orm.FooAPI
	fooAPI.ID = fooDB.ID
	fooAPI.FooPointersEncoding = fooDB.FooPointersEncoding
	fooDB.CopyBasicFieldsToFoo_WOP(&fooAPI.Foo_WOP)

	c.JSON(http.StatusOK, fooAPI)
}

// UpdateFoo
//
// swagger:route PATCH /foos/{ID} foos updateFoo
//
// # Update a foo
//
// Responses:
// default: genericError
//
//	200: fooDBResponse
func (controller *Controller) UpdateFoo(c *gin.Context) {

	mutexFoo.Lock()
	defer mutexFoo.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFoo", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/bistack/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFoo.GetDB()

	// Validate input
	var input orm.FooAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var fooDB orm.FooDB

	// fetch the foo
	_, err := db.First(&fooDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	fooDB.CopyBasicFieldsFromFoo_WOP(&input.Foo_WOP)
	fooDB.FooPointersEncoding = input.FooPointersEncoding

	db, _ = db.Model(&fooDB)
	_, err = db.Updates(&fooDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	fooNew := new(models.Foo)
	fooDB.CopyBasicFieldsToFoo(fooNew)

	// redeem pointers
	fooDB.DecodePointers(backRepo, fooNew)

	// get stage instance from DB instance, and call callback function
	fooOld := backRepo.BackRepoFoo.Map_FooDBID_FooPtr[fooDB.ID]
	if fooOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), fooOld, fooNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the fooDB
	c.JSON(http.StatusOK, fooDB)
}

// DeleteFoo
//
// swagger:route DELETE /foos/{ID} foos deleteFoo
//
// # Delete a foo
//
// default: genericError
//
//	200: fooDBResponse
func (controller *Controller) DeleteFoo(c *gin.Context) {

	mutexFoo.Lock()
	defer mutexFoo.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFoo", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/thomaspeugeot/bistack/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFoo.GetDB()

	// Get model if exist
	var fooDB orm.FooDB
	if _, err := db.First(&fooDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&fooDB)

	// get an instance (not staged) from DB instance, and call callback function
	fooDeleted := new(models.Foo)
	fooDB.CopyBasicFieldsToFoo(fooDeleted)

	// get stage instance from DB instance, and call callback function
	fooStaged := backRepo.BackRepoFoo.Map_FooDBID_FooPtr[fooDB.ID]
	if fooStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), fooStaged, fooDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
