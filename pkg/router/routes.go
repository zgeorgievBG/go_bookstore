package router

import(
	"github.com/gorilla/mux"
	"github.com/zgeorgievBG/go_bookstore/pkg/controller"
)

	var RegisterBookStoreRoutes = func(router * mux.Router) {
		router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
		router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
		router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
		router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("POST")
		// router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	}
