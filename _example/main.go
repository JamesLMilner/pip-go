package main

import (
	"encoding/json"
	"flag"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/JamesMilnerUK/pip-go"
)

var (
	npoints = flag.Int("n", 30, "npoints")
	width   = flag.Int("w", 500, "width")
	height  = flag.Int("h", 500, "height")
)

func main() {
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	var polygon pip.Polygon
	for n := 0; n < *npoints; n++ {
		polygon.Points = append(polygon.Points, pip.Point{rand.Float64() * float64(*width), rand.Float64() * float64(*height)})
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("index.html")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
		defer f.Close()
		io.Copy(w, f)
	})
	http.HandleFunc("/polygon", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(&polygon)
	})
	http.HandleFunc("/hit", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query()
		x, _ := strconv.ParseFloat(param.Get("x"), 64)
		y, _ := strconv.ParseFloat(param.Get("y"), 64)
		res := struct {
			Result bool `json:"result"`
		}{pip.PointInPolygon(pip.Point{x, y}, polygon)}
		json.NewEncoder(w).Encode(&res)
	})
	http.ListenAndServe(":8080", nil)
}
