package crawler

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/schul-cloud/mathehappen-content-crawler/settings"
)

func downloadContent() []byte {
	source := settings.Settings[settings.ContentLocation]
	resp, err := http.PostForm(source, autheticationData())
	panicOnDownloadErr(err)
	panicOnNon200OKResp(resp, source)

	defer resp.Body.Close()
	return readContent(resp.Body)
}

func autheticationData() url.Values {
	username := settings.Settings[settings.MHappenUser]
	password := settings.Settings[settings.MHappenPassword]

	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)
	return data
}

func panicOnDownloadErr(err error) {
	if err != nil {
		panic("unable to get data from mathehappen.de: " + err.Error())
	}
}
func panicOnNon200OKResp(response *http.Response, source string) {
	if response.StatusCode != 200 {
		message := fmt.Sprintf("mathehappen.de returned status %q upon calling %q",
			response.Status, source)
		panic(message)
	}
}

func readContent(responseBody io.ReadCloser) []byte {
	content, err := ioutil.ReadAll(responseBody)
	if err != nil {
		panic("unable to read data sent from mathehappen.de")
	}
	return content
}

func postContent(dataOut [][]byte) {
	for _, data := range dataOut {
		req := assembleRequest(data)
		response, err := http.DefaultClient.Do(req)
		panicCouldntPostOnError(err)

		if response.StatusCode != http.StatusCreated {
			defer response.Body.Close()
			answer := readContent(response.Body)
			panic("content server answered " + response.Status + "\n" + string(answer))
		}
	}
}

func assembleRequest(data []byte) *http.Request {
	target := settings.Settings[settings.TargetURL]
	req, err := http.NewRequest(http.MethodPost, target, bytes.NewReader(data))
	if err != nil {
		panic("cannot post data to target: " + err.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	user := settings.Settings[settings.BasicAuthUser]
	password := settings.Settings[settings.BasicAuthPassword]
	req.SetBasicAuth(user, password)

	return req
}

func panicCouldntPostOnError(err error) {
	if err != nil {
		panic("couldn't post data to target: " + err.Error())
	}
}
