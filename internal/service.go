package internal

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Data map[int]int

type StressTestReport struct {
	TotalReqs int
	Results   Data
}

func ExecuteStressTest(url string, requests int, concurrency int, startTime time.Time) {
	report := &StressTestReport{
		TotalReqs: 0,
		Results:   make(Data),
	}
	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	ch := make(chan *http.Request, 100)

	for w := 0; w < concurrency; w++ {
		go makeRequests(ch, report, &wg, &m)
	}

	produceRequests(ch, &wg, url, requests)
	wg.Wait()

	printReport(report, url, concurrency, startTime)
}

func produceRequests(ch chan *http.Request, wg *sync.WaitGroup, url string, requests int) {
	for i := 0; i < requests; i++ {
		wg.Add(1)
		request, _ := http.NewRequest("GET", url, nil)
		ch <- request
	}

	close(ch)
}

func makeRequests(ch chan *http.Request, report *StressTestReport, wg *sync.WaitGroup, m *sync.Mutex) {
	client := &http.Client{}

	for request := range ch {
		response, err := client.Do(request)
		if err != nil {
			m.Lock()
			setReportInfo(report, 0)
			m.Unlock()
			wg.Done()
			continue
		}

		m.Lock()
		setReportInfo(report, response.StatusCode)
		m.Unlock()
		wg.Done()
	}
}

func setReportInfo(report *StressTestReport, key int) {
	report.TotalReqs++
	report.Results[key]++
}

func printReport(report *StressTestReport, url string, concurrency int, startTime time.Time) {
	fmt.Println("##### STRESS TEST REPORT #####")
	fmt.Printf("Executed in: %v \n", time.Since(startTime))
	fmt.Printf("Url requested: %v \n", url)
	fmt.Printf("Total requests: %v \n", report.TotalReqs)
	fmt.Printf("Workers used: %v \n", concurrency)
	fmt.Printf("Requests with error: %v requests \n", report.Results[0])
	delete(report.Results, 0)

	for key, value := range report.Results {
		fmt.Printf("Response Code [%v]: %v requests \n", key, value)
	}
}
