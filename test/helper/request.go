package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kichikawa/auth"
	"github.com/kichikawa/router"
)

type Result struct {
	Url    string      `json:"url"`
	Type   string      `json:"type"`
	Method string      `json:"method"`
	Body   interface{} `json:"body"`
}

func Request(t *testing.T, pass string, id ...uuid.UUID) *httptest.ResponseRecorder {
	read, err := ioutil.ReadFile(pass)
	if err != nil {
		t.Errorf("request fail to file error %s", err.Error())
		return nil
	}

	var result Result
	json.Unmarshal([]byte(read), &result)

	jsonData, _ := json.Marshal(result.Body)

	body := buffer(string(jsonData))

	r := router.SetupRouter()
	w := httptest.NewRecorder()

	req, errReq := http.NewRequest(result.Method, result.Url, body)
	if errReq != nil {
		log.Fatal(errReq)
		return nil
	}

	if len(id) != 0 {
		time := time.Now()
		token, tokenErr := auth.Generate(id[0], time)

		if tokenErr != nil {
			log.Fatal(tokenErr)
			return nil
		}

		req.Header.Add(auth.HeaderKey, fmt.Sprintf("Bearer %s", token))
	}
	r.ServeHTTP(w, req)

	return w
}

func buffer(str string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(str))
}
