package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/rabbanext/gosong/config"
	"github.com/rabbanext/gosong/handlers"
	"github.com/rabbanext/gosong/middlewares"
)

func main() {
	// Create a new Fiber instance

	engine := html.New("templates", ".html")
	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("login", fiber.Map{
			"Title":   "Hello, World!",
			"Content": "Lets Login",
		})
	})

	// Create a new JWT middleware
	// Note: This is just an example, please use a secure secret key
	jwt := middlewares.NewAuthMiddleware(config.Secret)

	// Create a Login route
	app.Post("/login", handlers.Login)
	// Create a protected route
	app.Get("/protected", jwt, handlers.Protected)

	// Serve static files from the "static" directory.
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server on port 8080.
	fmt.Println("Server started on http://localhost:3000")
	// http.ListenAndServe(":8080", nil)
	log.Fatal(app.Listen(":3000"))
}

func backup() {
	// http.HandleFunc("/", LoginPage)
	// http.HandleFunc("/login", LoginPage)
	// http.HandleFunc("/welcome", WelcomePage)

}

// func SignupPage(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		// Retrieve signup form data.
// 		username := r.FormValue("username")
// 		password := r.FormValue("password")

// 		// Perform signup logic here (e.g., store user data in a database).
// 		// For simplicity, we'll just print the data for demonstration.
// 		fmt.Printf("New user signup: Username - %s, Password - %s\n", username, password)

// 		// Redirect to a welcome or login page after signup.
// 		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
// 		return
// 	}

// 	// If not a POST request, serve the signup page template.
// 	tmpl, err := template.ParseFiles("templates/signup.html")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	tmpl.Execute(w, nil)
// }

// // LoginPage is the handler for the login page.
// func LoginPage(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		username := r.FormValue("username")
// 		password := r.FormValue("password")

// 		// Perform authentication logic here (e.g., check against a database).
// 		// For simplicity, we'll just check if the username and password are both "admin".
// 		if username == "admin" && password == "admin" {
// 			// Successful login, redirect to a welcome page.
// 			http.Redirect(w, r, "/welcome", http.StatusSeeOther)
// 			return
// 		}

// 		// Invalid credentials, show the login page with an error message.
// 		fmt.Fprintf(w, "Invalid credentials. Please try again.")
// 		return
// 	}

// 	// If not a POST request, serve the login page template.
// 	tmpl, err := template.ParseFiles("templates/login.html")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	tmpl.Execute(w, nil)
// }

// WelcomePage is the handler for the welcome page.
func WelcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, you have successfully logged in!")
}
