package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"sync"

	"github.com/influxdata/influxdb/client/v2"
)

//collect the weights of each animal on a frequent basis => time series dataset
//this data get stored in influxdb, so that we can use it later

var animaltags = []string{"Tyrannosaurus rex;Rex", "Velociraptor;Rapto", "Velociraptor;Velo", "Carnotaurus;Carno"}

const myDB = "dino"

func main() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: "",
		Password: "",
	})
	if err != nil {
		log.Fatal(err)
	}
	queryDB(c, "", "Create DATABASE "+myDB) // Create Database dino
	//create a batch points object
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  myDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}
	wg := sync.WaitGroup{}
	detectSignal := checkStopOSSignals(&wg)
	rand.Seed(time.Now().UnixNano())
	//this loop generates random data every 1 second, then add it to the batchpoints
	for !(*detectSignal) {
		animaltag := animaltags[rand.Intn(len(animaltags))]
		split := strings.Split(animaltag, ";")
		tags := map[string]string{
			"animal_type": split[0],
			"nickname":    split[1],
		}
		fields := map[string]interface{}{
			"weight": rand.Intn(300) + 1, //1 -> 300
		}
		fmt.Println(animaltag, fields["weight"])
		pt, err := client.NewPoint("weightmeasures", tags, fields, time.Now())
		if err != nil {
			log.Println(err)
			continue
		}
		bp.AddPoint(pt)
		time.Sleep(1 * time.Second)
	}
	log.Println("Exit signal triggered, writing data... ")
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
	log.Println("Exiting program...")
}

func queryDB(c client.Client, database, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: database,
	}
	response, err := c.Query(q)
	if err != nil {
		return res, err
	}
	if response.Error() != nil {
		return res, response.Error()
	}

	return response.Results, nil
}

func checkStopOSSignals(wg *sync.WaitGroup) *bool {
	Signal := false
	go func(s *bool) {
		wg.Add(1)
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		log.Println("Exit signals received... ")
		*s = true
		wg.Done()
	}(&Signal)
	return &Signal
}
