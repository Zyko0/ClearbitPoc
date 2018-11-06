package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/clearbit/clearbit-go/clearbit"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("The company name is required as an argument")
		os.Exit(1)
	}
	companyName := os.Args[1]
	b, err := ioutil.ReadFile("api.key") // regular file containing your raw clearbit api key
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	apiKey := string(b)

	client := clearbit.NewClient(clearbit.WithAPIKey(apiKey))
	results, resp, err := client.Discovery.Search(clearbit.DiscoverySearchParams{
		Query: "name:" + companyName,
	})
	if err != nil {
		fmt.Println(results, resp, err.Error())
		os.Exit(1)
	}

	for _, company := range results.Results {
		fmt.Println("-----------------------")
		fmt.Println(company.Name)
		fmt.Println(company)
		fmt.Println("Location: ", company.Geo)
		fmt.Println("-----------------------")
	}
}
