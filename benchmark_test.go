package benchmark

import (
	"bytes"
	"net/http"
	"testing"
)

func BenchmarkFramework(b *testing.B) {
	b.Run("get", func(b *testing.B) {
		b.ReportAllocs()
		url := "http://localhost:8080/item?id=1"
		for i := 0; i < b.N; i++ {
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}

			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				b.Fatalf("HTTP status code is not 200: %d", resp.StatusCode)
			}
		}
	})

	b.Run("post", func(b *testing.B) {
		b.ReportAllocs()
		url := "http://localhost:8080/item"
		body := []byte(`{"id": 2, "name": "new item"}`)
		for i := 0; i < b.N; i++ {
			req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			if err != nil {
				panic(err)
			}
			req.Header.Set("header", "header")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}

			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				b.Fatalf("HTTP status code is not 200: %d", resp.StatusCode)
			}
		}
	})
}
