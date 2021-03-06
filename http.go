package main

import "net/http"
import "fmt"
import "os"
import "io"

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.ca")

	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
