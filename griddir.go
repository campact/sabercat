/*
Package sabercat implements net/http.FileSystem to serve
contents directly from MongoDB's GridFS.

To serve all files from a database containing a GridFS
called `fs` via http:

	session, err := mgo.Dial("mongodb://localhost/database")
	if err != nil {
		log.Fatalf("Could not connect to mongodb: %s", err)
	}
	gfs := session.DB("").GridFS("fs")
	http.Handle("/", http.FileServer(sabercat.GridDir{gfs})

Directory listing has not been implemented.
*/
package sabercat

import (
	"labix.org/v2/mgo"
	"log"
	"net/http"
)

const (
	VERSION = "1.1.2"
)

type GridDir struct {
	GridFS *mgo.GridFS
}

func (gd GridDir) Open(filename string) (http.File, error) {
	f, err := gd.GridFS.Open(filename)
	if err != nil {
		log.Printf("Error %s: %s", filename, err)
	}
	return &gridFile{f}, err
}
