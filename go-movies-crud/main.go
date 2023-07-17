package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/jesseokeya/go-httplogger"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` //Pointer to director struct.
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// ? slice movies to store the Movies of struct type Movie
var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies) //? send Json of all the movies in the slice
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //? Get params from the request header

	//? Loop through movies and find with id
	for _, item := range movies {
		//? if movie found with ID in the slice is equql to the id in the params
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item) //? send Json of the movie in the slice
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{}) //? if no movie found send empty movie
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //? Get params from the request header
	//? Loop through movies and find with id
	for index, item := range movies {
		//? if movie found with ID in the slice is equal to the id in the params
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) //? delete the movie from the slice with the id
			break
		}
	}
	json.NewEncoder(w).Encode(movies) //? send Json of all the movies in the slice after deleting the movie
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) //? get the body of the request and decode it to the movie struct
	movie.ID = strconv.Itoa(rand.Intn(1000))   //? Mock ID - not safe
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie) //? send Json of all the movies in the slice
}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) //? Get params from the request header

	//? Loop through movies and find with id
	for index, item := range movies {

		//* if movie found with ID in the slice is equal to the id in the params
		if item.ID == params["id"] {

			//* delete the movie from the slice with the id
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie

			//* and add the new movie to the slice with the same id
			_ = json.NewDecoder(r.Body).Decode(&movie) //? get the body of the request and decode it to the movie struct
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie) //? send Json of all the movies in the slice
			return
		}
	}

}




func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "43823324", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "23423415", Title: "Movie Two", Director: &Director{Firstname: "Sam", Lastname: "Kolder"}})
	movies = append(movies, Movie{ID: "3", Isbn: "23423423", Title: "Movie Three", Director: &Director{Firstname: "Jobin", Lastname: "Jose"}})
	movies = append(movies, Movie{ID: "4", Isbn: "22423534", Title: "Movie Four", Director: &Director{Firstname: "Cristopher", Lastname: "Nolan"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", httplogger.Golog(r)))
}
