package main

//go:generate oapi-codegen -generate types -package main -o types.gen.go api.v1.yaml
//go:generate oapi-codegen -generate server -package main -o server.gen.go api.v1.yaml
