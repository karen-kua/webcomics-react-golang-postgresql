package main

import (
	"fmt"
	"net/http"

	"github.com/azukimochi/webcomics-react-golang-postgresql/server/controllers"
	"github.com/azukimochi/webcomics-react-golang-postgresql/server/routes"
)

func main() {
	r := routes.Router()
	defer controllers.DB.Close()
	fmt.Println("Starting listening on port 5000!")
	http.ListenAndServe(":5000", r)
}
