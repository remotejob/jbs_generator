package handlers

import (
//	"fmt"
	"net/http"
//	"log"
)

type Startpar struct {
	
	Webrootdir string
	
	
}

type Entrystartpar struct {
	
	Modstartpar Startpar
	
}

func NewEntrystartpar() *Entrystartpar {
	
	return &Entrystartpar{Startpar{}}	
	
}

func (startpar *Entrystartpar) StaticHandler(w http.ResponseWriter, r *http.Request) {

//	fmt.Println(r.URL.Path)
//	fmt.Println(startpar.Modstartpar.Webrootdir)		
		
	http.ServeFile(w, r, startpar.Modstartpar.Webrootdir)
		
}
