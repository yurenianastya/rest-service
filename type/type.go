package _type

//Song Json request payload is as follows,
//{
//  "id": "3",
//  "title": "Cry for me",
//  "artist":  "Twice",
//  "release_year":  "2020"
//}
type Song struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	ReleaseYear int `json:"release_year"`
}