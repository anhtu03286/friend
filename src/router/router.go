package router

import (
	"database/sql"
	"github.com/anhtu03286/friend/controller"
	"github.com/anhtu03286/friend/repository"
	"github.com/anhtu03286/friend/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func initUserController(db *sql.DB) controller.UserController {
	userRepository := repository.UserRepository{DB: db}
	userService := service.UserService{IUserRepository: userRepository}
	return controller.UserController{UserService: userService}
}
//
//func initFriendController(db *sql.DB) controller.FriendController {
//	relationshipRepository := repository.RelationshipRepository{DB: db}
//	relationshipService := service.RelationshipService{IRelationshipRepository: relationshipRepository}
//
//	userRepository := repository.UserRepository{DB: db}
//	userService := service.UserService{IUserRepository: userRepository}
//
//	friendService := service.FriendService{IRelationshipService: relationshipService, IUserService: userService}
//	return controller.FriendController{FriendService: friendService}
//}

func HandleRequest(db *sql.DB) {
	myRouter := mux.NewRouter().StrictSlash(true)

	userHandle := initUserController(db)
	//friendHandle := initFriendController(db)

	myRouter.HandleFunc("/users", userHandle.GetAllUsers).Methods("GET")
	myRouter.HandleFunc("/users", userHandle.CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}