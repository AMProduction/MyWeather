/*
Copyright © 2023 Andrii Malchyk <snooki17@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const IPLookupUrl string = "/ip.json"

var IP string

type IPInfo struct {
	Ip            string  `json:"ip"`
	Type          string  `json:"type"`
	ContinentCode string  `json:"continent_code"`
	ContinentName string  `json:"continent_name"`
	CountryCode   string  `json:"country_code"`
	CountryName   string  `json:"country_name"`
	IsEu          string  `json:"is_eu"`
	GeonameId     uint32  `json:"geoname_id"`
	City          string  `json:"city"`
	Region        string  `json:"region"`
	Lat           float32 `json:"lat"`
	Lon           float32 `json:"lon"`
	TzId          string  `json:"tz_id"`
}

func getIPInfo(ip string) {
	if ip == "" {
		ip = "auto:ip"
	}

	if net.ParseIP(ip) == nil {
		fmt.Printf("IP Address: %s - Invalid\n", ip)
		os.Exit(1)
	}
	var ipResponse IPInfo

	ApiKey := viper.GetString("API_KEY")
	url := BaseUrl + IPLookupUrl + "?key=" + ApiKey + "&q=" + ip
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Could not create request: %s\n", err)
		os.Exit(1)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making http request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Could not read response body: %s\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal(body, &ipResponse)
	if err != nil {
		fmt.Printf("Could not read unmarshall JSON: %s\n", err)
	}
	fmt.Println("IP address: " + ipResponse.Ip)
	fmt.Println("Continent name: " + ipResponse.ContinentName)
	fmt.Println("Name of country: " + ipResponse.CountryName)
	fmt.Println("City name: " + ipResponse.City)
	fmt.Println("Region name: " + ipResponse.Region)
	fmt.Println("Time zone: " + ipResponse.TzId)
	fmt.Println("Position: lat " + fmt.Sprintf("%f", ipResponse.Lat) + " lan " + fmt.Sprintf("%f", ipResponse.Lon))
}

// mylocationCmd represents the mylocation command
var mylocationCmd = &cobra.Command{
	Use:   "mylocation",
	Short: "Show the info about your location",
	Long:  `Show the geo info about your location by your IP or any.`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		getIPInfo(IP)
	},
}

func init() {
	rootCmd.AddCommand(mylocationCmd)
	mylocationCmd.Flags().StringVarP(&IP, "ip", "i", "", "IP address")
}
