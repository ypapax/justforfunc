package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.M) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	r := t.Run()
	os.Exit(r)
}

func TestHandle(t *testing.T) {
	as := assert.New(t)
	pl, err := ioutil.ReadFile("./payload.json")
	if !as.NoError(err) {
		return
	}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:"+port, bytes.NewReader(pl))
	if !as.NoError(err) {
		return
	}
	rec := httptest.NewRecorder()
	handle(rec, req)

	res := rec.Result()
	if !as.Equal(http.StatusOK, res.StatusCode) {
		return
	}

	b, err := ioutil.ReadAll(res.Body)
	if !as.NoError(err) {
		return
	}
	defer res.Body.Close()
	t.Log("response", string(b))
	if exp := "pull request id: 191568743"; !as.Equal(exp, string(b)) {
		return
	}
}

func BenchmarkHandle(b *testing.B) {
	b.StopTimer()
	as := assert.New(b)

	pl, err := ioutil.ReadFile("./payload.json")
	if !as.NoError(err) {
		return
	}

	for i := 0; i <= b.N; i++ {
		req, err := http.NewRequest(http.MethodPost, "http://localhost:"+port, bytes.NewReader(pl))
		if !as.NoError(err) {
			return
		}
		res := httptest.NewRecorder()
		b.StartTimer()
		handle(res, req)
		b.StopTimer()
	}
}
