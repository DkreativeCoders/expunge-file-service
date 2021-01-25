// De-messenger
//
// De-messenger
//
//     Schemes: http,https,127.0.0.1
//     Host: localhost:8900
//     Version: 0.0.1
//	   BasePath: /api
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"github.com/DkreativeCoders/expunge-file-service/pkg"
	"github.com/kpango/glg"
)

func main() {

	srv :=pkg.NewServer()
	//defer func() {
	//	fmt.Print("Closing Db")
	//	err := db.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}()

	glg.Info("Server listening on", srv.Addr)
	glg.Error(srv.ListenAndServe())
}
