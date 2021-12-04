package controllers

import "github.com/khafido/ronda-app-api/middlewares"

func (s *Server) initializeRoutes() {
	//---Users routes
	// Login Route
	s.Router.HandleFunc("/api/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	// Register
	s.Router.HandleFunc("/api/register", middlewares.SetMiddlewareJSON(s.Register)).Methods("POST")
	
	// //Select All
	// s.Router.HandleFunc("/api/users", middlewares.SetMiddlewareAuthentication(s.GetUsers)).Methods("GET")
	// //Select by ID
	// s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareAuthentication(s.GetUser)).Methods("GET")
	// //Update
	// s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	// //Delete
	// s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.DeleteUser))).Methods("DELETE")

	//Products routes
	//Insert
	s.Router.HandleFunc("/api/products", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreateProduct))).Methods("POST")
	//Select All
	s.Router.HandleFunc("/api/products", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetProducts))).Methods("GET")
	//Select by ID
	s.Router.HandleFunc("/api/products/{id_produk}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetProduct))).Methods("GET")
	//Update
	s.Router.HandleFunc("/api/products/{id_produk}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateProduct))).Methods("PUT")
	//Delete
	s.Router.HandleFunc("/api/products/{id_produk}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.DeleteProduct))).Methods("DELETE")



}
