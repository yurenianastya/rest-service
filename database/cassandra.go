package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
)

var Session *gocql.Session

func init() {
	var err error
	cluster := gocql.NewCluster(goDotEnvVariable("CASSANDRA_URL"))
	cluster.ProtoVersion = 4
	Session, err := cluster.CreateSession()
	if err != nil {
		log.Println(err)
	}
	err = Session.Query("CREATE KEYSPACE IF NOT EXISTS songs WITH REPLICATION " +
		"= {'class' : 'SimpleStrategy', 'replication_factor' : 1};").Exec()
	if err != nil {
		log.Println(err)
		return
	}
	err = Session.Query("CREATE TABLE IF NOT EXISTS songs.song " +
		"(id int, title text, artist text, release_year int, PRIMARY KEY (id));").Exec()
	if err != nil {
		log.Println(err)
		return
	}
	err = Session.Query("INSERT INTO songs.song (id, title, artist, release_year)" +
		" VALUES (1, 'Lovesick Girls', 'Blackpink', 2020);").Exec()
	err = Session.Query("INSERT INTO songs.song (id, title, artist, release_year)" +
		" VALUES (2, 'Psycho', 'Red Velvet', 2019);").Exec()
	err = Session.Query("INSERT INTO songs.song (id, title, artist, release_year)" +
		" VALUES (3, 'Feel Special', 'Twice', 2019);").Exec()
	err = Session.Query("INSERT INTO songs.song (id, title, artist, release_year)" +
		" VALUES (4, 'Not Shy', 'Itzy', 2020);").Exec()
	if err != nil {
		log.Println(err)
		return
	}
	cluster.Keyspace = "songs"
	Session.Close()
	fmt.Println("cassandra well initialized")
}
