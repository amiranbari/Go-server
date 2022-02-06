package main

import (
	// "errors"
	"fmt"
	"log"
	"server/pkg/config"
	"server/pkg/handlers"
	"server/pkg/renders"

	// "math/rand"
	"net/http"
	// "time"
)

// var randomSource = rand.NewSource(time.Now().Unix())
// var random = rand.New(randomSource)

const portNumber string = ":8000"

// func makeRandomNumber() int {
// 	return random.Intn(100)
// }

// func devide(rw http.ResponseWriter, r *http.Request) {
// 	var x float32 = float32(makeRandomNumber())
// 	var y float32 = float32(makeRandomNumber())
// 	f, err := devideValues(x, y)
// 	if err != nil {
// 		fmt.Fprintf(rw, "Cannot devide by zero")
// 		return
// 	}
// 	fmt.Fprintf(rw, fmt.Sprintf("%f devided by %f is %f", x, y, f))
// }

// func devideValues(x, y float32) (float32, error) {
// 	if y <= 0 {
// 		return 0, errors.New("cannot devide zero")
// 	}

// 	return x / y, nil
// }

// func addValues(x, y int) int {
// 	return x + y
// }

func main() {

	var app config.AppConfig

	tc, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	renders.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)

	http.HandleFunc("/about", handlers.Repo.About)

	// http.HandleFunc("/devide", devide)

	fmt.Println(fmt.Sprintf("starting application on port number %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}