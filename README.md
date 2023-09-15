# About

The app allows you to get the weather forecast and IP address info.

# How to use it

_Prerequisite:_ obtain [the Weather API](https://www.weatherapi.com/api.aspx) key

## Usage:

`MyWeather [command]`

### Available Commands:

| Command      | Description                                                |  
|--------------|------------------------------------------------------------|  
| `completion` | Generate the autocompletion script for the specified shell |  
| `help`       | Help about any command                                     |  
| `mylocation` | Show the info about your location                          |  
| `show`       | Show current weather information                           |  
| `version`    | Print the version number of MyWeather                      |

### Global flags:

* `--config` _string_ config file (default is $APP/.env)
* `-h`, `--help` help for `MyWeather`

## Usage:

`MyWeather show [flags]`

Show a user up-to-date current weather information in the desired location.

### Flags:

* `-c`, `--city` _string_ City name e.g.: Paris
* `-l`, `--lang` _string_ Returns 'condition:text' in the desired language
* `-h`, `--help` help for `show`

## Usage:

`MyWeather mylocation [flags]`

Show the geo info about your location by your IP or any.

### Flags:

* `-i`, `--ip` _string_ IP address. If not set the app uses your current IP.
* `-h`, `--help` help for `mylocation`

Use `MyWeather [command] --help` for more information about a command.

Powered by [WeatherAPI.com](https://www.weatherapi.com/)   
![Powered by WeatherAPI.com](https://cdn.weatherapi.com/v4/images/weatherapi_logo.png)