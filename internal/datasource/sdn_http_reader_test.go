package datasource

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewSdnHttpReaderRemoteError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
	}))

	defer server.Close()
	_, err := NewSdnHTTPReader(server.URL)
	if err == nil {
		t.Fatalf("Needs error")
	}
}

func TestNewSdnHttpReaderOk(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	_, err := NewSdnHTTPReader(server.URL)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestNewSdnHttpReaderRead(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		_, err := rw.Write([]byte("123"))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
	}))
	defer server.Close()
	reader, err := NewSdnHTTPReader(server.URL)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	r, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	err = reader.Close()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	exp := []byte("123")
	if !bytes.Equal(exp, r) {
		t.Fatalf("Expected %v to equal %v", exp, r)
	}
}
