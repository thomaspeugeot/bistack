// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/thomaspeugeot/bistack/otherstack/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	barDBs map[uint]*BarDB

	nextIDBarDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		barDBs: make(map[uint]*BarDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/thomaspeugeot/bistack/otherstack/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *BarDB:
		db.nextIDBarDB++
		v.ID = db.nextIDBarDB
		db.barDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/thomaspeugeot/bistack/otherstack/go, unsupported type in Create")
	}
	return db, nil
}

// Unscoped sets the unscoped flag for soft-deletes (not used in this implementation)
func (db *DBLite) Unscoped() (db.DBInterface, error) {
	return db, nil
}

// Model is a placeholder in this implementation
func (db *DBLite) Model(instanceDB any) (db.DBInterface, error) {
	// Not implemented as types are handled directly
	return db, nil
}

// Delete removes a record from the database
func (db *DBLite) Delete(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/thomaspeugeot/bistack/otherstack/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *BarDB:
		delete(db.barDBs, v.ID)
	default:
		return nil, errors.New("github.com/thomaspeugeot/bistack/otherstack/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/thomaspeugeot/bistack/otherstack/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *BarDB:
		db.barDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/thomaspeugeot/bistack/otherstack/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/thomaspeugeot/bistack/otherstack/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *BarDB:
		if existing, ok := db.barDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Bar github.com/thomaspeugeot/bistack/otherstack/go, record not found")
		}
	default:
		return nil, errors.New("github.com/thomaspeugeot/bistack/otherstack/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]BarDB:
		*ptr = make([]BarDB, 0, len(db.barDBs))
		for _, v := range db.barDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/thomaspeugeot/bistack/otherstack/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/thomaspeugeot/bistack/otherstack/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/thomaspeugeot/bistack/otherstack/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/thomaspeugeot/bistack/otherstack/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *BarDB:
		tmp, ok := db.barDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Bar Unkown entry %d", i))
		}

		barDB, _ := instanceDB.(*BarDB)
		*barDB = *tmp
		
	default:
		return nil, errors.New("github.com/thomaspeugeot/bistack/otherstack/go, Unkown type")
	}
	
	return db, nil
}

