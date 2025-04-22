package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const CREDS = "creds.txt"

func main() {
	defer measureTime()()
	creds, err := loadCreds(CREDS)
	if err != nil {
		log.Println(err)
		return
	}

	autoDeploy := AutoDeploy{
		Hostname: creds["hostname"].(string),
		Username: creds["username"].(string),
		Password: creds["password"].(string),
		Port:     uint(creds["port"].(int)),
		SiteFile: creds["siteFile"].(string),
		ListCommands: creds["commands"].([]string),
	}

	client, err := autoDeploy.NewClient()
	if err != nil {
		log.Println(err)
		return
	}
	defer client.Close()

	sites, err := autoDeploy.LoadSiteFile()
	if err != nil {
		log.Println(err)
		return
	}
	var wg sync.WaitGroup
	for _, site := range *sites {
		if len(site) == 0 {
			continue
		}
		wg.Add(1)
		go autoDeploy.Execute(client, &wg, site)
	}
	wg.Wait()
	printReport()
}

func measureTime() func() {
	start := time.Now()
	return func() {
		fmt.Println("Takes", fmt.Sprint(time.Since(start).Seconds()), "seconds")
	}
}
