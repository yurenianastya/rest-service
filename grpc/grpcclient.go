package grpc

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	sg "rest-service/song"
)

func Client() {
	serverAddress := "localhost:9000"
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Can't dial server")
	}
	defer conn.Close()

	client := sg.NewSongServiceClient(conn)
	create, err := client.CreateSong(context.Background(), &sg.Song{Id: 0, Artist: "Red Velvet",
		Title: "Psycho", ReleaseYear: 2019})
	if err != nil {
		log.Fatal("Can't create song on server")
	}
	log.Println(create)

	get, err := client.GetSong(context.Background(), &sg.Id{Value: 0})
	if err != nil {
		log.Fatal("Can't get a song from server")
	}
	log.Println(get)

	update, err := client.UpdateSong(context.Background(), &sg.Song{Id: 0, Title: "How you like that",
		Artist: "Blackpink", ReleaseYear: 2020})
	log.Println(update)

	deleting, err := client.DeleteSong(context.Background(), &sg.Id{Value: 0})
	if err != nil {
		log.Fatal("Can't delete a song from server")
	}
	log.Println(deleting)
}
