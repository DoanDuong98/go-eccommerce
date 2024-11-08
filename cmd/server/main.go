package main

import (
	"go-ecommerce-be/internal/routers"
)

func main() {
	r := routers.NewRouter()
	err := r.Run(":8082")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
