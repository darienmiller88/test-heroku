package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	router := chi.NewRouter()

	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		render.JSON(res, req, map[string]interface{}{
			"number": 45,
		})
	})

	fmt.Println("running on port:", os.Getenv("PORT"))

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router)

}