package handler

import (
	"log"
	"net/http"

	"github.com/google/logger"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"

	"github.com/goforbroke1006/sport-archive-svc/pkg/endpoint"
)

const rpcUri = "/rpc"

func HandleClientsRequests(handleAddr string, svc *endpoint.SportArchiveServiceEndpoint) {
	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	err := server.RegisterService(svc, "SportArchiveService")
	if err != nil {
		log.Fatalf("Format of service SportArchiveServiceEndpoint isn't correct. %s", err)
	}

	r := mux.NewRouter()
	r.Handle(rpcUri, server)

	logger.Infof("Start listen address: %s", handleAddr)
	logger.Infof("Start listen path: %s", rpcUri)

	log.Fatal(http.ListenAndServe(handleAddr, r))
}
