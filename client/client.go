package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web-service-gin/server"
)

func GetUserByID(host string, id string) (*server.Album, error) {
	uri := fmt.Sprintf("http://%s/albums/%s", host, id)
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var album server.Album
	err = json.NewDecoder(resp.Body).Decode(&album)
	if err != nil {
		return nil, err
	}

	return &album, nil
}
