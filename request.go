package CustomerTrafficSimGo

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func makeRequest(method string, url string, userFunction func(*http.Request, func() *http.Response), client *http.Client) error {
	req, err := http.NewRequestWithContext(
		context.Background(),
		method,
		url,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Error creating request (%s-%s): %w", method, url, err)
	}
	req.Header.Set("User-Agent", "Go-Simulator")

	var response *http.Response = nil
	userFunction(req, func() *http.Response {
		resp, err := client.Do(req)
		if err != nil {
			panic(fmt.Errorf("Error run request (%s-%s): %w", method, url, err))
		}
		resp.Body.Close()
		response = resp
		fmt.Printf("done %s %s -> Status: %d\n", method, url, resp.StatusCode)
		return resp
	})
	if response == nil {
		return fmt.Errorf("Error request not called (%s-%s)", method, url)
	}
	return nil
}

func (s *Simulator) Run() {
	client := &http.Client{Timeout: 5 * time.Second}
	for _, route := range s.routes {
		err := makeRequest(route.method, s.target+route.path, route.userFunction, client)
		if err != nil {
			fmt.Println(err)
		}
	}
}
