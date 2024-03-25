package main

import (
	"fmt"
	"github.com/The-Flash/go122-intro/middleware"
	"net/http"
)

func main() {
	stack := middleware.MiddlewareStack(
		middleware.Logger,
		middleware.IsLoggedIn,
	)
	customerRouter := http.NewServeMux()
	adminRouter := http.NewServeMux()

	// customer handlers
	customerRouter.HandleFunc("/sign-in", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Customer Sign in")
	})
	// admin handlers
	adminRouter.HandleFunc("/sign-in", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Admin Sign in")
	})

	router := http.NewServeMux()
	customerMiddlewares := middleware.MiddlewareStack(
		middleware.IsCustomer,
	)
	adminMiddlewares := middleware.MiddlewareStack(
		middleware.IsAdmin,
	)
	router.Handle("/customer/", http.StripPrefix("/customer", customerMiddlewares(customerRouter)))

	router.Handle("/admin/", http.StripPrefix("/admin", adminMiddlewares(adminRouter)))

	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", stack(router))
}
