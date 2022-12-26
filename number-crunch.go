package main

import(
	"net/http"
	"math"
	// "fmt"
	"encoding/json"
	"strconv"
)

const (
	defaultHTTPPort = "8080"
)

func SquareRoot(x float64) float64 {
	return math.Sqrt(x)
}

type SQRoot struct{
	InputNumber float64
	SquareRoot float64
}

func squareRootHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	inputNumber, err := strconv.ParseFloat(r.URL.Path[len("/square-root/"):], 64)
	if err != nil {
		panic(err)
	}
	sq := SQRoot{inputNumber, SquareRoot(inputNumber)}
	json.NewEncoder(w).Encode(sq)
}

func main(){
	http.HandleFunc("/square-root/", squareRootHandler)
	http.ListenAndServe(":8080", nil)
	// fmt.Println(SquareRoot(2))
}



