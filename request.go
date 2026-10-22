package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"log"
)

func getRequestLocation(url string) error {
	res, err := http.Get(url)
	if err != nil{
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
    	fmt.Println("status code != 200")
     	return fmt.Errorf("non-OK HTTP status: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil{
		return err
	}

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
	}

	var loc Locations
	err = json.Unmarshal(data, &loc);
	if err != nil {
		return err
	}

	for _, area := range loc.Results {
        fmt.Println(area.Name)
    }
	return nil
}
