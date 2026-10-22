package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"log"
)

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
    if cfg.nextLocationURL != nil {
        url = *cfg.nextLocationURL
    }

    data, ok := cfg.cache.Get(url)

	if ok {
		var loc Locations
		err := json.Unmarshal(data, &loc);
		if err != nil {
			return err
		}

		cfg.nextLocationURL = loc.Next
    	cfg.prevLocationURL = loc.Previous

		for _, area := range loc.Results {
        	fmt.Println(area.Name)
    	}
		return nil
	}

	res, err := http.Get(url)
	if err != nil{
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
    	fmt.Println("status code != 200")
     	return fmt.Errorf("non-OK HTTP status: %s", res.Status)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil{
		return err
	}

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
	}

	cfg.cache.Add(url, data)

	var loc Locations
	err = json.Unmarshal(data, &loc);
	if err != nil {
		return err
	}

	cfg.nextLocationURL = loc.Next
    cfg.prevLocationURL = loc.Previous

	for _, area := range loc.Results {
        fmt.Println(area.Name)
    }
	return nil
}
