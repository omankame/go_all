/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
        "strconv"
	"github.com/spf13/cobra"
)

var tempTo, tempFrom string

// temperatureCmd represents the temperature command
var temperatureCmd = &cobra.Command{
	Use:   "temperature",
	Short: "Convert temperature for different units",
	Long: `temperature is a sub-command for uconv (unit converter).

    It helps to convert temperature from one unit to another, Kelvin - k,
    Fahrenheit - f and Celsius - c.
 
    If no flags are specified, it converts Celsius - c to Fahrenheit - f by default.
    
    Usage:
        uconv temperature 45
        uconv temperature 45 --from f --to c
        uconv temperature 45 -f f -t c

        `,
        Args: cobra.ExactArgs(1),
        Run: executeTemperatureCmd,

	
}

func init() {
        rootCmd.AddCommand(temperatureCmd)


	temperatureCmd.Flags().StringVarP(&tempFrom, "from", "f", "c", "the unit to convert from")
	temperatureCmd.Flags().StringVarP(&tempTo, "to", "t", "f", "the unit to convert to")

//        rootCmd.AddCommand(temperatureCmd)
}

func executeTemperatureCmd(cmd *cobra.Command, args []string) {
     arg, err := strconv.ParseInt(args[0], 10, 64)
     if err != nil {
        fmt.Println(err)
        return
     }

      switch {
    case tempFrom == "c" && tempTo == "f":
        fmt.Printf("temperature: %d°C ==> %d°F\n", arg, celsiusToFahrenheit(arg))
    case tempFrom == "f" && tempTo == "c":
        fmt.Printf("temperature: %d°F ==> %d°C\n", arg, fahrenheitToCelsius(arg))
    case tempFrom == "c" && tempTo == "k":
        fmt.Printf("temperature: %d°C ==> %dK\n", arg, celsiusToKelvin(arg))
    case tempFrom == "f" && tempTo == "k":
        fmt.Printf("temperature: %d°F ==> %dK\n", arg, fahrenheitToKelvin(arg))
    case tempFrom == "k" && tempTo == "c":
        fmt.Printf("temperature: %dK ==> %d°C\n", arg, kelvinToCelsius(arg))
    case tempFrom == "k" && tempTo == "f":
        fmt.Printf("temperature: %dK ==> %d°F\n", arg, kelvinToFahrenheit(arg))
    case tempFrom == tempTo:
        fmt.Printf("temperature: %d%s\n", arg, tempTo)
    }
}

// celsiusToFahrenheit - converts Celsius to Fahrenheit
func celsiusToFahrenheit(data int64) int64 {
    return (data * 9 / 5) + 32
}

// celsiusToKelvin - converts Celsius to Kelvin
func celsiusToKelvin(data int64) int64 {
    return data + 273
}

// fahrenheitToCelsius - converts Fahrenheit to Celsius
func fahrenheitToCelsius(data int64) int64 {
    return (data - 32) * 5 / 9
}

// fahrenheitToKelvin - converts Fahrenheit to Kelvin
func fahrenheitToKelvin(data int64) int64 {
    return (data-32)*5/9 + 273
}

// kelvinToCelsius - converts Kelvin to Celsius
func kelvinToCelsius(data int64) int64 {
    return data - 273
}

// kelvinToFahrenheit - converts Kelvin to Fahrenheit
func kelvinToFahrenheit(data int64) int64 {
    return (data-273)*9/5 + 32
}
