package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	BaseURL string
	Token   string
}

func NewClient(baseURL, username string) (*Client, error) {
	client := &Client{BaseURL: baseURL}

	// Request token
	body, _ := json.Marshal(map[string]string{"username": username})
	resp, err := http.Post(fmt.Sprintf("%s/token", baseURL), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	token, ok := result["token"]
	if !ok {
		return nil, fmt.Errorf("failed to retrieve token")
	}

	client.Token = token
	return client, nil
}

func (c *Client) Set(key string, value interface{}, ttl int) error {
	url := fmt.Sprintf("%s/set", c.BaseURL)
	body, _ := json.Marshal(map[string]interface{}{
		"key":   key,
		"value": value,
		"ttl":   ttl,
	})

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to set value: %s", resp.Status)
	}

	return nil
}

func (c *Client) Get(key string) (interface{}, error) {
	url := fmt.Sprintf("%s/get/%s", c.BaseURL, key)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result interface{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &result)

	return result, nil
}

func (c *Client) Delete(key string) error {
	url := fmt.Sprintf("%s/delete/%s", c.BaseURL, key)
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to delete key: %s", resp.Status)
	}

	return nil
}

func (c *Client) Push(key string, value interface{}) error {
	url := fmt.Sprintf("%s/list/push", c.BaseURL)
	body, _ := json.Marshal(map[string]interface{}{
		"key":   key,
		"value": value,
	})

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to push to list: %s", resp.Status)
	}

	return nil
}

func (c *Client) Pop(key string) (interface{}, error) {
	url := fmt.Sprintf("%s/list/pop/%s", c.BaseURL, key)
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result interface{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &result)

	return result, nil
}
