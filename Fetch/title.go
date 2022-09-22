package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func FetchTitle(title string) ([]byte, error) {

	url := "https://imdb8.p.rapidapi.com/title/v2/find?title=" + title + "&limit=20&sortArg=moviemeter%2Casc"
	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-RapidAPI-Key", os.Getenv("X-RapidAPI-Key"))
	req.Header.Add("X-RapidAPI-Host", "imdb8.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, nil
}
