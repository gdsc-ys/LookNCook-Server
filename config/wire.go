//go:build wireinject
// +build wireinject

package config

import (
	"github.com/gdsc-ys/lookncook-server/src/handler"
	"github.com/gdsc-ys/lookncook-server/src/service"
	"github.com/google/wire"
)

var firebaseApp = wire.NewSet(InitializeFirebaseApp)
var firebaseAuthClient = wire.NewSet(NewFirebaseAuthClient)
var firebaseStorageClient = wire.NewSet(NewFirebaseStorageClient)
var fireStoreClient = wire.NewSet(NewFireStoreClient)

var userServiceSet = wire.NewSet(service.UserServiceInit, wire.Bind(new(service.UserService), new(*service.UserServiceImpl)))
var kitchenServiceSet = wire.NewSet(service.KitchenServiceInit, wire.Bind(new(service.KitchenService), new(*service.KitchenServiceImpl)))
var recipeServiceSet = wire.NewSet(service.RecipeServiceInit, wire.Bind(new(service.RecipeService), new(*service.RecipeServiceImpl)))
var geminiServiceSet = wire.NewSet(service.GeminiServiceInit, wire.Bind(new(service.GeminiService), new(*service.GeminiServiceImpl)))

var userHandlerSet = wire.NewSet(handler.UserHandlerInit, wire.Bind(new(handler.UserHandler), new(*handler.UserHandlerImpl)))
var kitchenHandlerSet = wire.NewSet(handler.KitchenHandlerInit, wire.Bind(new(handler.KitchenHandler), new(*handler.KitchenHandlerImpl)))
var recipeHandlerSet = wire.NewSet(handler.RecipeHandlerInit, wire.Bind(new(handler.RecipeHandler), new(*handler.RecipeHandlerImpl)))

func InitializeApp() *Initialization {
	wire.Build(
		NewInitialization,
		firebaseApp,
		//firebaseAuthClient,
		fireStoreClient,
		userServiceSet,
		kitchenServiceSet,
		recipeServiceSet,
		geminiServiceSet,
		userHandlerSet,
		kitchenHandlerSet,
		recipeHandlerSet,
	)
	return nil
}
