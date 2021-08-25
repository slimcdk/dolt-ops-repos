package main

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/dolthub/dolt/go/cmd/git-dolt/doltops"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

func main() {
	fmt.Println("Fetching daily police reports")

	err := doltops.Clone("dolthub/city-populations")
	if err != nil {
		log.Fatalln(err)
	}

	v, _ := query.Values(QueryOptions{
		ItemID:        "90DEB0B1-8DF0-4A2D-823B-CFD7A5ADD85F",
		IsNewsList:    true,
		NewsType:      "døgnrapporter",
		DistrictQuery: strings.Join(Districts, ","),
		FromDate:      time.Date(1970, 1, 0, 0, 0, 0, 0, time.Local),
		ToDate:        time.Date(9999, 12, 30, 0, 0, 0, 0, time.Local),
		Page:          0,
		PageSize:      math.MaxInt32,
	})

	var reports ReportStructSuccess
	if resp, err := resty.New().SetHostURL(APIEndpoint).R().SetQueryParamsFromValues(v).SetResult(&reports).Get(APIEndpoint); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("Got %d reports for %s\n", reports.TotalNumberOfNews, resp.Request.URL)
	}
}
