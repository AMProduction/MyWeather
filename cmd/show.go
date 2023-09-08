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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"strconv"
)

const BaseUrl string = "https://api.weatherapi.com/v1"
const CurrentWeatherUrl string = "/current.json"

var City string
var Lang string

type CurrentResponse struct {
	Location Location
	Current  Current
}

// Test string
type Location struct {
	Name            string
	Region          string
	Country         string
	Lat             float32
	Lon             float32
	Tz_id           string
	Localtime_epoch int32
	Localtime       string
}

type Current struct {
	Last_updated       string
	Last_updated_epoch int32
	Temp_c             float32
	Temp_f             float32
	Feelslike_c        float32
	Feelslike_f        float32
	Condition          Condition
	Wind_mph           float32
	Wind_kph           float32
	Wind_degree        int16
	Wind_dir           string
	Pressure_mb        float32
	Pressure_in        float32
	Precip_mm          float32
	Precip_in          float32
	Humidity           int
	Cloud              int
	Is_day             int
	Uv                 float32
	Gust_mph           float32
	Gust_kph           float32
}

type Condition struct {
	Text string
	Icon string
	Code int16
}

func GetCurrentWeather(cityName string, lang string) {
	if cityName == "" {
		cityName = "Kyiv"
	}

	if lang == "" {
		lang = "uk"
	}

	var currentResponse CurrentResponse
	ApiKey := viper.GetString("API_KEY")
	url := BaseUrl + CurrentWeatherUrl + "?key=" + ApiKey + "&q=" + cityName + "&lang=" + lang
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

	err = json.Unmarshal(body, &currentResponse)
	if err != nil {
		fmt.Printf("Could not read unmarshall JSON: %s\n", err)
	}
	fmt.Println("The current weather in " + currentResponse.Location.Name)
	fmt.Println("Region name: " + currentResponse.Location.Region)
	fmt.Println("Name of country: " + currentResponse.Location.Country)
	fmt.Println("Local date and time: " + currentResponse.Location.Localtime)
	fmt.Println("Time zone: " + currentResponse.Location.Tz_id)
	fmt.Println("Position: lat " + fmt.Sprintf("%f", currentResponse.Location.Lat) + " lan " + fmt.Sprintf("%f", currentResponse.Location.Lon))
	fmt.Println(`_____________________________________________`)
	fmt.Println("Temperature in celsius: " + fmt.Sprintf("%.1f", currentResponse.Current.Temp_c))
	fmt.Println("Feels like: " + fmt.Sprintf("%.1f", currentResponse.Current.Feelslike_c))
	fmt.Println("Weather condition: " + currentResponse.Current.Condition.Text)
	fmt.Println("Wind speed: " + fmt.Sprintf("%.2f", currentResponse.Current.Wind_kph) + "km/h")
	fmt.Println("Wind direction: " + currentResponse.Current.Wind_dir)
	fmt.Println("Precipitation amount: " + fmt.Sprintf("%.2f", currentResponse.Current.Precip_mm) + "mm")
	fmt.Println("Humidity: " + strconv.Itoa(currentResponse.Current.Humidity) + "%")
	fmt.Println("Cloud cover: " + strconv.Itoa(currentResponse.Current.Cloud) + "%")
	fmt.Println("UV Index: " + fmt.Sprintf("%.2f", currentResponse.Current.Uv))
	fmt.Println(`_____________________________________________`)
	fmt.Println("The weather updated at: " + currentResponse.Current.Last_updated)
}

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current weather information",
	Long:  `Show a user up-to-date current weather information in the desired location.`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		GetCurrentWeather(City, Lang)
	},
}

func init() {
	showCmd.Flags().StringVarP(&City, "city", "c", "", "City name e.g.: Paris")
	showCmd.Flags().StringVarP(&Lang, "lang", "l", "", "Returns 'condition:text' in the desired language")
	rootCmd.AddCommand(showCmd)
}
