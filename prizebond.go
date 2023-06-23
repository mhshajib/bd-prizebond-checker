package prizebond

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
)

// Real implementation of Prizebond interface
type PrizebondConnection struct {
	BaseUrl     string
	Url         string
	BondNumbers []string
}

/*
Init prizebond with url
*/
func (p *PrizebondConnection) Init(bondNumbers []string) {
	p.BaseUrl = "https://www.bb.org.bd/"
	p.Url = "https://www.bb.org.bd/en/index.php/investfacility/prizebond"
	p.BondNumbers = bondNumbers
}

// Check Bond Number Length
func bondNumberValidation(bondNumbers []string) error {
	if len(bondNumbers) < 1 {
		err := errors.New("You've to provide at least 1 prizebond number")
		return err
	}
	return nil
}

// Getting raw html data for data scraping
func getHtml(url string, bondNumbers []string) (string, error) {
	var responseString string

	bondNumbersString := strings.Join(bondNumbers[:], ",")
	//Orchastrating Payload
	payload := strings.NewReader(fmt.Sprintf("gsearch=%s", bondNumbersString))

	//Creating new request with URL and Payload
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return responseString, err
	}

	//Setting essential request headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	//Executing Request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return responseString, err
	}
	defer resp.Body.Close()

	//Reading data from response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return responseString, err
	}

	//returning html response data
	return string(body), nil
}

// Getting Absolute Image URL from Relative URL
func getAbsoluteImage(baseUrl string, relativeImageUrl string) string {
	imagePathWithoutBaseUrl := strings.ReplaceAll(relativeImageUrl, "../", "")
	return fmt.Sprintf("%s/%s", baseUrl, imagePathWithoutBaseUrl)
}

// Scraping Html data to JSON
func parseHtml(baseUrl string, htmlData string) ([]map[string]string, error) {
	var results []map[string]string

	// Parse the HTML response body
	doc, err := htmlquery.Parse(strings.NewReader(htmlData))
	if err != nil {
		return results, err
	}

	// Find the table with the class "tableData"
	table := htmlquery.FindOne(doc, "//table[@class='tableData']")
	if table == nil {
		return results, errors.New("No Record Available")
	}

	// Find all table rows within the table
	rows := htmlquery.Find(table, "//tr")

	// Iterate over the rows, skipping the header row
	for i := 1; i < len(rows); i++ {
		row := rows[i]

		// Find all table data cells within the row
		cells := htmlquery.Find(row, "//td")

		// Create a map to hold the row data
		rowData := make(map[string]string)

		// Extract the text from each cell and store it in the map
		for j := 0; j < len(cells); j++ {
			cell := cells[j]
			cellText := htmlquery.InnerText(cell)
			headerCell := htmlquery.FindOne(table, "//tr[@class='tableSubHeader']/td["+fmt.Sprint(j+1)+"]")
			headerText := htmlquery.InnerText(headerCell)
			if headerText == "Eligible Series" {
				imgNode := htmlquery.FindOne(cell, "//img")
				relativeImageUrl := htmlquery.SelectAttr(imgNode, "src")
				rowData[headerText] = getAbsoluteImage(baseUrl, relativeImageUrl)
			} else {
				rowData[headerText] = cellText
			}
		}

		// Append the row data to the slice
		results = append(results, rowData)
	}
	return results, nil
}

// Send sms
func (p *PrizebondConnection) Fetch() ([]map[string]string, error) {
	var results []map[string]string

	//Checking is there any mobile number
	err := bondNumberValidation(p.BondNumbers)
	if err != nil {
		return results, err
	}

	//Getting html string from request
	htmlString, err := getHtml(p.Url, p.BondNumbers)
	if err != nil {
		return results, err
	}

	results, err = parseHtml(p.BaseUrl, htmlString)
	if err != nil {
		return results, err
	}

	return results, nil
}
