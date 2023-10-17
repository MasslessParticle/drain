package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/masslessparticle/drain"
)

func main() {
	logger := drain.New(drain.DefaultConfig())
	f, err := os.Open("")
	if err != nil {
		panic(err)
	}

	start := time.Now()

	var lineNum uint64
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		logger.Train(sc.Text())
		lineNum++
		if lineNum%10000 == 0 {
			fmt.Printf("Processing line: %d, %d clusters so far\n", lineNum, logger.ClusterLen())
		}

		if lineNum >= 1000000 {
			break
		}
	}

	logger.WriteToFile()
	fmt.Println(time.Since(start))

	//for _, cluster := range logger.Clusters() {
	//	println(cluster.String())
	//}
	//
	//cluster := logger.Match("user faceair logged in")
	//if cluster == nil {
	//	println("no match")
	//} else {
	//	fmt.Printf("cluster matched: %s", cluster.String())
	//}
}
