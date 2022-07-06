// go get "github.com/gorilla/mux"
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Return all movies
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// params is an array with the parameters
	// index = name of the parameter
	// value on that position = value of the parameter
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			// Return the movie that matched the ID
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// params is an array with the parameters
	// index = name of the parameter
	// value on that position = value of the parameter
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			// Deletes the movie with append function:
			// this appends the start of the array until the index and
			// the other slice of the array after the index
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	// Return all the movies after the deletion
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	// Decode the body and save it in the variable movie direction
	_ = json.NewDecoder(r.Body).Decode(&movie)

	// Get a random number and convert it to string
	// Assign this value to the ID value of the movie
	movie.ID = strconv.Itoa(rand.Intn(100000000))

	// Append the movie to our array of movies
	movies = append(movies, movie)

	// Return the created movie
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// params is an array with the parameters
	// index = name of the parameter
	// value on that position = value of the parameter
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			// Delete the movie when you found it
			movies = append(movies[:index], movies[index+1:]...)

			// Decode the movie from the body and assign it to the variable
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)

			// Assign the same ID that it had before
			movie.ID = params["id"]

			// Append the new movie with the same ID
			movies = append(movies, movie)

			// Return the updated movie
			json.NewEncoder(w).Encode(movie)
		}
	}

}

func main() {
	// Creates a router
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "4387227",
		Title: "The Black Phone",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Wayne"}})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "4343255",
		Title: "The Irishman",
		Director: &Director{
			Firstname: "Quentin",
			Lastname:  "Tarantino"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")

	log.Fatal(http.ListenAndServe(":8000", r))
}
