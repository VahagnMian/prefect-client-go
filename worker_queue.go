package prefect_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (wq WorkQueue) CreateWorkQueue(queue WorkQueue) (*WorkQueue, error) {
	queue.ID = "" // set ID to empty string to prevent it from being sent in the request
	data, err := json.Marshal(queue)
	if err != nil {
		return nil, fmt.Errorf("marshaling error: %v", err)
	}

	resp, err := http.Post("http://localhost:4200/api/work_queues/", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("response body: %s\n", string(body))
		return nil, fmt.Errorf("bad status: %s\n%s", resp.Status, body)
	}

	body, err := ioutil.ReadAll(resp.Body)

	work_queue := WorkQueue{}
	err = json.Unmarshal(body, &work_queue)
	if err != nil {
		return nil, err
	}

	return &work_queue, nil
}

func (wq WorkQueue) GetWorkQueue(wqID string) (*WorkQueue, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/work_queues/%s", "http://localhost:4200", wqID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	w_queue := WorkQueue{}
	err = json.Unmarshal(body, &w_queue)
	if err != nil {
		return nil, err
	}

	return &w_queue, nil
}

func (wq WorkQueue) UpdateWorkQueue(wqID string, queue WorkQueue) (*WorkQueue, error) {
	client := &http.Client{}
	data, err := json.Marshal(queue)
	if err != nil {
		return nil, fmt.Errorf("marshaling error: %v", err)
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/api/work_queues/%s", "http://localhost:4200", wqID), bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("response body: %s\n", string(body))
		return nil, fmt.Errorf("bad status: %s\n%s", resp.Status, body)
	}

	body, err := ioutil.ReadAll(resp.Body)

	work_queue := WorkQueue{}
	err = json.Unmarshal(body, &work_queue)
	if err != nil {
		return nil, err
	}

	return &work_queue, nil
}
