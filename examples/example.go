package examples

import (
	"fmt"
	"net/http"
	"time"

	"github.com/osvaldoabel/rest-party/gohttp"
)

var (
	githubHttpClient = gohttp.NewBuilder().Build()
)

func getGithubClient() gohttp.Client {
	client := gohttp.NewBuilder().
		DisableTimeouts(true).
		SetMaxIdleConnections(20).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(50 * time.Millisecond).
		Build()

	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")

	client.SetHeaders(commonHeaders)
	return client
}

func main() {
	getUrls()
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func createUser(user User) {
	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status())
	fmt.Println(response.StatusCode())
	fmt.Println(response.String())

}

func getUrls() {
	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status())
	fmt.Println(response.StatusCode())
	fmt.Println(response.String())

	// var user User
	// if err := response.UnmarshalJson(&user); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(user.FirstName)

	// //second
	// bytes2, _ := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// var user2 User
	// if err := json.Unmarshal(bytes2, &user2); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(user2.FirstName)
}
