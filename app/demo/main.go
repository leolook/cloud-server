package main

import (
	"cloud-server/lib/log"
	"time"
)

func main() {

	i := 1
	for {
		log.Infof("current:%d", i)
		time.Sleep(1 * time.Second)
		i++
	}

}
