package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type job struct {
	jobName string
	cost    float64
}

type invoice struct {
	name  string
	jobs  []string
	costs []float64
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Hey there!")
	fmt.Println("So you'd like to make an invoice?")
	fmt.Println("Who should I make it out to?")
	employer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Okay great! What should I invoice ", strings.TrimSpace(employer), " for?")
	jobName, err := reader.ReadString('\n')

	fmt.Println("Okay great! How much should I invoice for ", strings.TrimSpace(jobName))
	cost, err := reader.ReadString('\n')
	convertCost, err := strconv.ParseFloat(strings.TrimSpace(cost), 64)
	job := job{jobName: strings.TrimSpace(jobName), cost: convertCost}
	invoice := invoice{name: employer, jobs: []string{job.jobName}, costs: []float64{job.cost}}

	done := checkDone()
	for !done {
		jobName, cost := otherItems(invoice)
		invoice.jobs = append(invoice.jobs, jobName)
		invoice.costs = append(invoice.costs, cost)
		done = checkDone()
	}

	fmt.Println(invoice)
}

func otherItems(invoice) (string, float64) {
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Println("No probs, what is the next item on the invoice?")
	jobName, err := reader2.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Nice! How much should I invoice for ", strings.TrimSpace(jobName))
	cost, err := reader2.ReadString('\n')
	convertCost, err := strconv.ParseFloat(strings.TrimSpace(cost), 64)

	return jobName, convertCost
}

func checkDone() bool {
	doneReader := bufio.NewReader(os.Stdin)
	fmt.Println("Okay, too easy! Would you like me to invoice for anything else? (type: yes or no)")
	input, err := doneReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	if strings.TrimSpace(input) == "yes" {
		return false
	}
	return true
}
