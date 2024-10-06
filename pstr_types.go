package pstr

type Video struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string // ?
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Ratings    []Rating
	Metascore  string
	imdbRating string
	imdbVotes  string
	imdbID     string
	Type       string
	DVD        string
	BoxOffice  string
	Production string
	Website    string
	Response   string
}
type Rating struct {
	Source string
	Value  string
}

// SearchItems
// Response example:
//
//	{
//	 "Search": [
//
//	{
//	 "Title": "The Mask",
//	 "Year": "1994",
//	 "imdbID": "tt0110475",
//	 "Type": "movie",
//	 "Poster": "https://m.media-amazon.com/images/M/MV5BNGNmNjI0ZmMtMzI5MC00ZjUyLWFlZDEtYjUyMGZlN2E3N2E2XkEyXkFqcGc@._V1_SX300.jpg"
//	},
//
//	{
//	 "Title": "The Mask of Zorro",
//	 "Year": "1998",
//	 "imdbID": "tt0120746",
//	 "Type": "movie",
//	 "Poster": "https://m.media-amazon.com/images/M/MV5BMzg4ZjQ4OGUtZjkxMi00Y2I2LWEzNTAtODI2ZjkxMGVjNTQwXkEyXkFqcGdeQXVyNjgxNTAwNjQ@._V1_SX300.jpg"
//	 ...
//	},
type SearchItems struct {
	Search []SearchItem
}

type SearchItem struct {
	Title  string
	Year   string
	imdbID string
	Type   string
	Poster string
}
