package pstr

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var search = flag.String("search", "", "--search cat")
var page = flag.String("page", "1", "--page 2")
var apikey = flag.String("apikey", "", "--apikey <API_KEY>")

func DownloadPoster(searchItems *SearchItems) ([]string, error) {
	si := make([]string, 24)

	h := strconv.Itoa(time.Now().Hour())
	m := strconv.Itoa(time.Now().Minute())
	s := strconv.Itoa(time.Now().Second())
	now := time.Now().Format("02012006") + "_" + h + m + s
	for _, searchItem := range searchItems.Search {

		link := searchItem.Poster
		si = append(si, link)

		resp, err := http.Get(link)
		if err != nil {
			return nil, fmt.Errorf("http.Get: %s\n", err)
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("resp.StatusCode: %s\n", err)
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("io.ReadAll: %s\n", err)
		}

		title := strings.Replace(searchItem.Title, " ", "_", -1) // replace all spaces to underline symbol

		sep := string(os.PathSeparator)

		tmpDir := os.TempDir() + sep + *search + now
		err = os.MkdirAll(tmpDir, 0755)
		if err != nil {
			return nil, fmt.Errorf("os.MkdirAll: %s\n", err)
		}

		f, err := os.Create(tmpDir + string(os.PathSeparator) + title + filepath.Ext(link))
		if err != nil {
			return nil, fmt.Errorf("os.Create: %s\n", err)
		}

		_, err = f.Write(b)
		if err != nil {
			return nil, fmt.Errorf("f.Write: %s\n", err)
		}
	}
	return si, nil
}

// Search Example Variant 1:
// $ go test -v --apikey <API_KEY> --search mask --page 2
//
// Search Example Variant 2:
// $ go test -v --apikey=<API_KEY> --search=doctor --page=2
//
// https://omdbapi.com/?s=mask&apikey=<API_KEY>&page=2
func Search() (*SearchItems, error) {
	flag.Parse()
	url := "https://omdbapi.com?s=" + *search + "&page=" + *page + "&apikey=" + *apikey
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("сбой запроса: %s", resp.Status)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var searchItems *SearchItems
	err = json.Unmarshal(b, &searchItems)
	return searchItems, err
}
