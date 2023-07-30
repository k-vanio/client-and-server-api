server:
	go mod tidy
	go run cmd/server/*

client:
	go mod tidy
	go run cmd/client/*
