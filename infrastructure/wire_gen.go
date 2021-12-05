// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package infrastructure

import (
	"github.com/21hack02win/nascalay-backend/interfaces/handler"
)

// Injectors from wire.go:

func injectServer() handler.ServerInterface {
	serverInterface := handler.NewHandler()
	return serverInterface
}