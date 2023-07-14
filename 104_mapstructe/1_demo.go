package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
	Age  int
	Job  string `json:"job1"`
}

func main() {

	var metadata mapstructure.Metadata
	var p Person
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &p,
		Metadata:         &metadata,
		TagName:          "json",
	})

	if err != nil {
		log.Fatal(err)
	}

	datas := []string{`
    { 
      "type": "person",
      "name":"dj",
      "age":18,
      "job1": "programmer"
    }
  `,
		`
    {
      "type": "cat",
      "name": "kitty",
      "age": 1,
      "breed": "Ragdoll"
    }
  `,
	}

	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}

		switch m["type"].(string) {
		case "person":

			//mapstructure.Decode(m, &p)
			err = decoder.Decode(m)
			fmt.Println("person", p)
		}
	}
}
