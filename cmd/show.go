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
)

const BaseUrl string = "https://api.weatherapi.com/v1"
const CurrentWeather string = "/current.json"

var City string

type CurrentResponse struct {
	Location Location
	Current  Current
}

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

func GetCurrentWeather() {
	var current_response CurrentResponse
	ApiKey := viper.GetString("API_KEY")
	url := BaseUrl + CurrentWeather + "?key=" + ApiKey + "&q=Kiev&aqi=no&lang=uk"
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

	//fmt.Printf(string(body))
	json.Unmarshal(body, &current_response)
	fmt.Println(current_response.Current)
}

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("show called")
		fmt.Println(City)
		GetCurrentWeather()
	},
}

func init() {
	showCmd.Flags().StringVarP(&City, "city", "c", "", "City name to get weather")
	rootCmd.AddCommand(showCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
