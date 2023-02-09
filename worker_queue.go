package prefect_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateWorkQueue(queue WorkQueue) error {
	data, err := json.Marshal(queue)
	if err != nil {
		return fmt.Errorf("marshaling error: %v", err)
	}

	resp, err := http.Post("http://localhost:4200/api/work_queues/", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("response body: %s\n", string(body))
		return fmt.Errorf("bad status: %s\n%s", resp.Status, body)
	}

	return nil
}
