package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
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
	fmt.Println("So you'd like to make an invoice...")
	fmt.Println("Who should I make it out to?")
	employer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Okay great! What should I invoice ", strings.TrimSpace(employer), " for?")
		jobName, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

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
		jobs := invoice.jobs
		costs := invoice.costs
		createTable(jobs, costs, employer)
	}
}

func createTable(jobs []string, costs []float64, employer string) {
	currDate := time.Now().Format("20060102150405")
	printDate := currDate[0:10]
	fN := strings.TrimSpace(employer) + "-" + printDate
	_, err := os.Create(fN + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	output, err := os.OpenFile(fN+".txt", os.O_APPEND|os.O_WRONLY, 0644)
	defer output.Close()

	writer := tabwriter.NewWriter(output, 0, 8, 1, '\t', tabwriter.AlignRight)
	length := len(jobs)

	fmt.Fprintln(writer, "Invoice to\t"+strings.TrimSpace(employer))
	fmt.Fprintln(writer, "\t"+"\t")
	fmt.Fprintln(writer, "Job\tUnit Cost")
	for i := 0; i < length; i++ {
		output := jobs[i] + "\t" + fmt.Sprintf("%.2f", costs[i])
		fmt.Fprintln(writer, output)
	}
	total := 0.00
	for i := 0; i < length; i++ {
		total += costs[i]
	}
	fmt.Fprintln(writer, "\t"+"\t")
	fmt.Fprintln(writer, "total:\t"+fmt.Sprintf("%g", total))
	fmt.Fprintln(writer, "\t"+"\t")
	fmt.Fprintln(writer, "Name: \tFelix Watson")
	fmt.Fprintln(writer, "BSB: \t033091")
	fmt.Fprintln(writer, "ACC: \t349128")
	fmt.Fprintln(writer, "ABN: \t52048795150")
	writer.Flush()
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

	return strings.TrimSpace(jobName), convertCost
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
