package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//Movie from DB
type Movie struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	RealeaseDate string `json:"releaseDate"`
	Director     string `json:"director"`
}

var movies []Movie

// Connect to database
func conncectDB() (db *sql.DB) {
	//dbHost := "localhost"
	dbHost := "mysql-db"

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "secret123"
	dbName := "testdb"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return db
}

//GET all Movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	setMovies()

	json.NewEncoder(w).Encode(movies)
}

//GET single Movie
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r) // get route params

	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})

}

//POST single Movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var movie Movie

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		fmt.Println(err)
	}
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

	testDB := conncectDB()

	// INSERT INTO DB
	insert, err := testDB.Exec("INSERT INTO movies(title, release_date, director) VALUES(?, ?, ?)", movie.Title, movie.RealeaseDate, movie.Director)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(insert)

	fmt.Println(movies)

}

//PUT single Movie
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

//DELETE single Movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r) // get route params

	for i, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			db := conncectDB()
			_, err := db.Exec("DELETE FROM movies WHERE id=?", movie.ID)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("delete", movie)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// set movies Slice to databse Content
func setMovies() {
	testDB := conncectDB()

	results, err := testDB.Query("SELECT * FROM movies")
	if err != nil {
		panic(err.Error())
	}

	movies = nil

	for results.Next() {
		var movie Movie

		err = results.Scan(&movie.ID, &movie.Title, &movie.RealeaseDate, &movie.Director)
		if err != nil {
			panic(err.Error())
		}

		movies = append(movies, movie)

	}
	fmt.Println(movies)
}

func addMovie(title string, releaseDate string, director string) {
	movie := Movie{
		ID:           "0",
		Title:        title,
		RealeaseDate: releaseDate,
		Director:     director,
	}
	movies = append(movies, movie)
}

func main() {

	fmt.Println("Hello Docker")

	//init router
	router := mux.NewRouter()

	// Route Handlers / Endpoints
	router.HandleFunc("/api/movies", getMovies).Methods("GET")
	router.HandleFunc("/api/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/api/movies", createMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", deleteMovie).Methods("DELETE")

	setMovies()

	log.Fatal(http.ListenAndServe(":8090", router))

}
