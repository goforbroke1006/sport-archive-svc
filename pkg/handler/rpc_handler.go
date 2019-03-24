package handler

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"io"
	"log"
	"net/http"

	"github.com/goforbroke1006/sport-archive-svc/pkg/service"
)

type HttpConn struct {
	in  io.Reader
	out io.Writer
}

func (c *HttpConn) Read(p []byte) (n int, err error) {
	return c.in.Read(p)
}

func (c *HttpConn) Write(d []byte) (n int, err error) {
	return c.out.Write(d)
}
func (c *HttpConn) Close() error {
	return nil
}

func HandleClientsRequests(handleAddr string, svc *service.SportArchiveService) {
	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	err := server.RegisterService(svc, "SportArchiveService")
	if err != nil {
		log.Fatalf("Format of service SportArchiveService isn't correct. %s", err)
	}

	r := mux.NewRouter()
	r.Handle("/rpc", server)
	log.Fatal(http.ListenAndServe(handleAddr, r))

	//rpc.HandleHTTP()
	//
	//l, e := net.Listen("tcp", handleAddr)
	//if e != nil {
	//	log.Fatalf("Couldn't start listening on port '%s'. Error %s", handleAddr, e)
	//}
	//log.Println("Serving RPC handler on " + handleAddr)
	//err = http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//
	//	if r.URL.Path == "/rpc" {
	//		serverCodec := jsonrpc.NewServerCodec(&HttpConn{in: r.Body, out: w})
	//		w.Header().Set("Content-type", "application/json")
	//		w.WriteHeader(200)
	//		err := server.ServeRequest(serverCodec)
	//		if err != nil {
	//			log.Printf("Error while serving JSON request: %v", err)
	//			http.Error(w, "Error while serving JSON request, details have been logged.", 500)
	//			return
	//		}
	//	} else {
	//		http.Error(w, "Unknown request", 404)
	//	}
	//
	//}))
	//if err != nil {
	//	log.Fatalf("Error serving: %s", err)
	//}
}
