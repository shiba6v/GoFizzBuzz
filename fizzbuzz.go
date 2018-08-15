package main

//import "fmt"
import "github.com/gorilla/mux"
import "net/http"
import "strconv"

func FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num, err := strconv.Atoi(vars["id"])
	var result = vars["id"]
	if err == nil {
		if num%15 == 0 {
			result = "Fizz Buzz"
		} else if num%5 == 0 {
			result = "Buzz"
		} else if num%3 == 0 {
			result = "Fizz"
		}
	}

	w.Write([]byte(result))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fizzbuzz/{id}", FizzBuzzHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
