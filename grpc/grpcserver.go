package grpc

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"rest-service/song"
	sg "rest-service/song"
	_ "strings"
)


const (
	port = ":9000"
)

type Song struct {
	sg.UnimplementedSongServiceServer
	savedSongs []*song.Song
}

func (p *Song) CreateSong(c context.Context, input *song.Song) (*song.Song, error) {
	p.savedSongs = append(p.savedSongs, input)
	return &song.Song{Id: input.Id, Title: input.Title,
		Artist: input.Artist, ReleaseYear: input.ReleaseYear}, nil
}

func (p *Song) UpdateSong(c context.Context, input *song.Song) (*song.Song, error) {
	for _, elem := range p.savedSongs {
		if elem.Id == input.Id {
			p.savedSongs[elem.Id].Title = input.Title
			p.savedSongs[elem.Id].Artist = input.Artist
			p.savedSongs[elem.Id].ReleaseYear = input.ReleaseYear
			return &song.Song{Id: input.Id, Title: input.Title,
				Artist: input.Artist, ReleaseYear: input.ReleaseYear}, nil
		}
	}
	return &song.Song{Id: 0, Title: "", Artist: "", ReleaseYear: 0}, errors.New("no song found")
}

func (p *Song) GetAllSongs(n *emptypb.Empty, stream song.SongService_GetAllSongsServer) error {
	for _, elem := range p.savedSongs {
		err := stream.Send(elem)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Song) GetSong(c context.Context, Id *song.Id) (*song.Song, error) {
	for _, elem := range p.savedSongs {
		if Id.Value == elem.Id {
			return &song.Song{Id: elem.Id, Title: elem.Title,
				Artist: elem.Artist, ReleaseYear: elem.ReleaseYear}, nil
		}
	}
	return &song.Song{Id: 0, Title: "", Artist: "", ReleaseYear: 0}, errors.New("no song found")
}

func (p *Song) DeleteSong(c context.Context, Id *song.Id) (*song.Song, error) {
	for _, elem := range p.savedSongs {
		if Id.Value == elem.Id {
			p.savedSongs = append(p.savedSongs[:elem.Id], p.savedSongs[elem.Id+1:]...)
			return &song.Song{Id: elem.Id, Title: elem.Title, Artist: elem.Artist,
				ReleaseYear: elem.ReleaseYear}, nil
		}
	}
	return &song.Song{Id: 0, Title: "", Artist: "", ReleaseYear: 0}, errors.New("no song found")
}

func Server() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to listen: ", err)
		return
	}
	// Creates a new gRPC Song
	s := grpc.NewServer()
	song.RegisterSongServiceServer(s, &Song{})
	reflection.Register(s)
	fmt.Println("Starting rpc server")
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}