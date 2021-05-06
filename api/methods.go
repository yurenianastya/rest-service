package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to homepage!")
	fmt.Println("Endpoint Hit: homePage")
}

// createNewSong godoc
// @Summary create a new song
// @Description POST method which creates new song object
// @Success 200 {object} Song
// @Router /song [post]
func createNewSong(w http.ResponseWriter, r *http.Request) {
	cluster := gocql.NewCluster(goDotEnvVariable("CASSANDRA_URL"))
	cluster.ProtoVersion = 4
	Session, err := cluster.CreateSession()
	if err != nil {
		log.Println(err)
	}
	var newSong Song
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Wrong data")
	}
	json.Unmarshal(requestBody, &newSong)
	if err := Session.Query("INSERT INTO songs.song(id, title, artist, release_year) VALUES(?, ?, ?, ?)",
		newSong.Id, newSong.Title, newSong.Artist, newSong.ReleaseYear).Exec(); err != nil {
		fmt.Println("Error while inserting")
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
		Conv, _ := json.MarshalIndent(newSong, "", " ")
		fmt.Fprintf(w, "%s", string(Conv))
	fmt.Println("Endpoint Hit: createSong")
	Session.Close()
}

// getAllSongs godoc
// @Summary Retrieve all songs
// @Description GET method which retrieves all songs
// @Produce json
// @Success 200 {object} Song
// @Router /songs [get]
func getAllSongs(w http.ResponseWriter, r *http.Request) {
	cluster := gocql.NewCluster(goDotEnvVariable("CASSANDRA_URL"))
	cluster.ProtoVersion = 4
	Session, err := cluster.CreateSession()
	if err != nil {
		log.Println(err)
	}
	var songs []Song
	m := map[string]interface{}{}

	iter := Session.Query("SELECT * FROM songs.song").Iter()

	for iter.MapScan(m) {
		songs = append(songs, Song{
			Id:         m["id"].(int),
			Title:      m["title"].(string),
			Artist:     m["artist"].(string),
			ReleaseYear:m["release_year"].(int),
		})
		m = map[string]interface{}{}
	}
	Conv, _ := json.MarshalIndent(songs, "", " ")
	fmt.Fprintf(w, "%s", Conv)
	fmt.Println("Endpoint Hit: getAllSongs")
	Session.Close()
}

// getSong godoc
// @Summary Retrieve a song
// @Description GET method which retrieves a song based on given ID
// @Produce json
// @Success 200 {object} Song
// @Router /song/{id} [get]
func getSong(w http.ResponseWriter, r *http.Request) {
	cluster := gocql.NewCluster(goDotEnvVariable("CASSANDRA_URL"))
	cluster.ProtoVersion = 4
	Session, err := cluster.CreateSession()
	if err != nil {
		log.Println(err)
	}
	SongId := mux.Vars(r)["id"]
	var songs []Song
	m := map[string]interface{}{}

	iter := Session.Query("SELECT * FROM songs.song WHERE id=?", SongId).Iter()
	for iter.MapScan(m) {
		songs = append(songs, Song{
			Id:         m["id"].(int),
			Title:      m["title"].(string),
			Artist:     m["artist"].(string),
			ReleaseYear:m["release_year"].(int),
		})
		m = map[string]interface{}{}
	}
	Conv, _ := json.MarshalIndent(songs, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
	fmt.Println("Endpoint Hit: getSong")
	Session.Close()
}

// deleteSong godoc
// @Summary delete a song
// @Description DELETE method which destroys song object
// @Success 200 {object} Song
// @Router /song/{id} [delete]
func deleteSong(w http.ResponseWriter, r *http.Request) {
	cluster := gocql.NewCluster(goDotEnvVariable("CASSANDRA_URL"))
	cluster.ProtoVersion = 4
	Session, err := cluster.CreateSession()
	if err != nil {
		log.Println(err)
	}
	SongId := mux.Vars(r)["id"]
	if err := Session.Query("DELETE FROM songs.song WHERE id = ?", SongId).Exec(); err != nil {
		fmt.Println("Error while deleting")
		fmt.Println(err)
	}
	fmt.Fprintf(w, "deleted successfully the song num %s ", SongId)
	Session.Close()
	fmt.Println("Endpoint Hit: deleteSong")
}

// updateSong godoc
// @Summary update a song
// @Description PUT method which modifies song object
// @Success 200 {object} Song
// @Router /song/{id} [put]
func updateSong(w http.ResponseWriter, r *http.Request) {
	cluster := gocql.NewCluster(goDotEnvVariable("CASSANDRA_URL"))
	cluster.ProtoVersion = 4
	Session, err := cluster.CreateSession()
	if err != nil {
		log.Println(err)
	}
	SongId := mux.Vars(r)["id"]
	var updateSong Song
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please enter data correctly")
	}
	json.Unmarshal(reqBody, &updateSong)
	if err := Session.Query("UPDATE songs.song SET title = ?, artist = ?, release_year = ? WHERE id = ?",
		updateSong.Title, updateSong.Artist, updateSong.ReleaseYear, SongId).Exec(); err != nil {
		fmt.Println("Error while updating")
		fmt.Println(err)
	}
	fmt.Fprintf(w, "updated successfully")
	Session.Close()
	fmt.Println("Endpoint Hit: updateSong")
}

