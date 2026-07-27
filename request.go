package CustomerTrafficSimGo

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func makeRequest(method string, path string, client *http.Client) error {
	req, err := http.NewRequestWithContext(
		context.Background(),
		method,
		path,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Error creating request (%s-%s): %w", method, path, err)
	}
	req.Header.Set("User-Agent", "Go-Simulator")
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error run request (%s-%s): %w", method, path, err)
	}
	resp.Body.Close()
	fmt.Printf("done %s %s -> Status: %d\n", method, path, resp.StatusCode)
	return nil
}

func (s *Simulator) Run() {
	client := &http.Client{Timeout: 5 * time.Second}
	for _, route := range s.routes {
		err := makeRequest(route.method, s.target+route.path, client)
		if err != nil {
			fmt.Println(err)
		}
	}
}
