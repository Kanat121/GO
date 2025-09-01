package bas

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello , go standart library")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Http respon status:", resp.Status)
}
