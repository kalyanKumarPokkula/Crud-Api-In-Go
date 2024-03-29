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

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Year int    `json:"year"`
    Rating float64 `json:"rating"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter , req *http.Request){
	w.Header().Set("Context-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovieById(w http.ResponseWriter , req *http.Request){
	w.Header().Set("Context-Type", "application/json")

	params := mux.Vars(req)

	for index , movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]... )
			break
		}
	}

	fmt.Fprintf(w, "Succesfully deleted movie")
}

func getMovieById(w http.ResponseWriter , req *http.Request){
	w.Header().Set("Context-Type", "application/json")

	params := mux.Vars(req)

	for _ , movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			break
		}
	}
}

func createMovie(w http.ResponseWriter , req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ : json.NewDecoder(req.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovieById(w http.ResponseWriter , req *http.Request){
	w.Header().Set("Context-Type", "application/json")
	params := mux.Vars(req)
	var movie Movie
	_ : json.NewDecoder(req.Body).Decode(&movie)
	
	index := -1
	for i , item := range movies {
		if params["id"] == item.ID {
			index = i
		}
	}

	if index == -1 {
		fmt.Fprintf(w, "Movie not found")
	}

	if movie.Title != "" {
		movies[index].Title = movie.Title
	}

	if movie.Rating != 0 {
		movies[index].Rating = movie.Rating
	}

	if movie.Year != 0 {
		movies[index].Year = movie.Year
	}

	if movie.Director.Firstname != "" {
		movies[index].Director.Firstname = movie.Director.Firstname
	}

	if movie.Director.Lastname != "" {
		movies[index].Director.Lastname = movie.Director.Lastname
	}

	json.NewEncoder(w).Encode(movies[index])


	
}



func main() {
	r := mux.NewRouter()
	movies = append(movies , Movie{ID : "1"  , Isbn : "438227" , Title : "Movie one" , Year: 2022 , Rating: 4, Director : &Director{Firstname : "avinash" , Lastname : "Sura"}})
	movies = append(movies , Movie{ID : "2"  , Isbn : "438228" , Title : "Movie two" ,Year: 2022 , Rating: 4, Director : &Director{Firstname : "kalyan" , Lastname : "pokkula"}})
	r.HandleFunc("/movies" , getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovieById).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovieById).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovieById).Methods("DELETE")
	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
	

}
