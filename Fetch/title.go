package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchTitle(title string) ([]byte, error) {

	url := "https://imdb8.p.rapidapi.com/title/v2/find?title=" + title + "&limit=20&sortArg=moviemeter%2Casc"
	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-RapidAPI-Key", "da18f18ba0mshf2809c081a0bd98p150885jsn589bed44bafc")
	req.Header.Add("X-RapidAPI-Host", "imdb8.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, nil
}
