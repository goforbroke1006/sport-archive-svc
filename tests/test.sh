#!/usr/bin/env bash

#curl -v -X CONNECT --url localhost:10001/_goRPC_
#curl -X CONNECT --url http://localhost:10001/_goRPC_ -d '{"jsonrpc":"2.0","id":"123","method":"SportArchiveService.GetSportID","params":["soccer"]}'
#curl -X POST -H "Content-Type: application/json" --url http://localhost:10001/rpc -d '{"method":"SportArchiveService.GetSport","params":{"name":"soccer"},"id":"123"}'
curl -X POST -H "Content-Type: application/json" --url http://localhost:10001/rpc -d '
{"method":"SportArchiveService.GetSport","params":[{"name":"soccer"}],"id":"123"}'
