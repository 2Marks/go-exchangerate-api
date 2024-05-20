<!-- ABOUT THE PROJECT-->
## About The Project
Exchange rate API built with [Go](https://go.dev/) leveraging a third-party API

### Prerequisites
Make sure you have go installed on your machine
* go
```sh
download from the [Go official site](https://go.dev/dl/)
```

### Installation
1. Get a Free API key at [https://exchangesrateapi.com/](https://exchangesrateapi.com/)
2. Clone the repository
3. Install go packages
```sh
go mod install
```
4. Make a copy of .example.env file and rename to .env like so:
```sh
cp .example.env .env
```
5. Update the env file with the right values as applicable


## Usage
To start the app, please run: (tested on mac)
```sh
make run
```
or:
```sh
go run cmd/main.go
```

### Example usage: Accessing the API endpoint:
to specify multiple to currencies: send them as a comma seperated value like so: (USD,GBP) e.g
```sh
http://localhost:8060/api/v1/exchange-rates?fromCurrency=EUR&toCurrency=USD,GBP
```

<!-- Feedback-->
### Feedback on the project stucture, quality are always welcome. thanks as you check it out 