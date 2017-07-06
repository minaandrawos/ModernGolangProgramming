package main

import (
	"log"

	"github.com/influxdata/influxdb/client/v2"
)

func main() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: "",
		Password: "",
	})
	if err != nil {
		log.Fatal(err)
	}
	res, err := queryDB(c, "dino", `select * from weightmeasures where "animal_type" = 'Velociraptor'`)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range res {
		log.Println("messages: ", v.Messages)
		for _, s := range v.Series {
			log.Println("series name: ", s.Name)
			log.Println("series columns: ", s.Columns)
			for _, r := range s.Values {
				log.Println("row values: ", r)
			}
		}
	}
}

func queryDB(c client.Client, database, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: database,
	}
	if response, err := c.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}
