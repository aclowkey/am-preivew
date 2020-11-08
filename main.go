package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/prometheus/alertmanager/template"
)

type RenderRequest struct {
	Template string `json:"template"`
	Data     string `json:"data"`
}

func main() {
	t, err := template.FromGlobs()
	if err != nil {
		panic(err)
	}

	for {
		data := new(template.Data)
		d, err := ioutil.ReadFile("test.gotempl")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		b, err := ioutil.ReadFile("data.json")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		err = json.Unmarshal(b, data)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		res, err := t.ExecuteTextString(string(d), data)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		err = ioutil.WriteFile("/tmp/output.html", []byte(res), 0644)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Done")
		time.Sleep(5 * time.Second)
	}

}
