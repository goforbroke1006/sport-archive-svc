package handler

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"

	"github.com/goforbroke1006/sport-archive-svc/pkg/endpoint"
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

func HandleClientsRequests(handleAddr string, svc *endpoint.SportArchiveServiceEndpoint) {
	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	err := server.RegisterService(svc, "SportArchiveService")
	if err != nil {
		log.Fatalf("Format of service SportArchiveServiceEndpoint isn't correct. %s", err)
	}

	r := mux.NewRouter()
	r.Handle("/rpc", server)
	log.Fatal(http.ListenAndServe(handleAddr, r))
}
