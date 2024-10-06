package pstr

import (
	"flag"
	"testing"
)

func IsFlagExist(name string) bool {
	exist := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			exist = true
		}
	})
	return exist
}

func TestSearch(t *testing.T) {

	if IsFlagExist("apikey") == false {
		t.Fatalf("flag apikey is not set (--apikey 123456)")
	}
	if IsFlagExist("search") == false {
		t.Fatalf("flag search is not set (--search text)")
	}

	searchItems, err := Search()
	if err != nil {
		t.Errorf("Search: %s", err)
	}

	_, err = DownloadPoster(searchItems)
	if err != nil {
		t.Errorf("DownloadPoster: %s", err)
	}

	for _, si := range searchItems.Search {
		t.Log("---------------------------------------------------------------------------------------------------------------------------------")
		t.Logf("poster: %s\n", si.Poster)
		t.Logf("title: %s\n", si.Title)
		t.Logf("type: %s\n", si.Type)
		t.Logf("year: %s\n", si.Year)
		t.Logf("imdbID: %s\n", si.imdbID)
	}
	t.Logf("all posters successfully uploaded")
	//for _, poster := range posters {
	//	if len(poster) > 0 {
	//		t.Logf("%s\n", poster)
	//	}
	//}
}
