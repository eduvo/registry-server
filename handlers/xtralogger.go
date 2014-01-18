package handler

import (
	"github.com/codegangsta/martini"
	//"github.com/davecgh/go-spew/spew"
	"log"
	"net/http"
	"time"
)

func Xtralogger(res http.ResponseWriter, req *http.Request, c martini.Context, log *log.Logger) {
	start := time.Now()
	log.Printf("-- %v Started %s %s%s", req.RemoteAddr, req.Method, req.Host, req.RequestURI)
	// if len(c.Params) > 0 {
	// 	log.Printf("   Get  params: %v", c.Params)
	// }
	if len(req.Form) > 0 {
		log.Printf("   Form params: %v", req.Form)
	}
	if len(req.PostForm) > 0 {
		log.Printf("   Post params: %v", req.PostForm)
	}

	//spew.Dump(req)
	rw := res.(martini.ResponseWriter)
	c.Next()

	log.Printf(
		"-- Completed %v %s in %v\n\n",
		rw.Status(),
		http.StatusText(rw.Status()),
		time.Since(start),
	)

}
