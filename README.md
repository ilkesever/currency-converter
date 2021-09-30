1. Install Golang at https://golang.org/doc/install
2. Unzip project 
3. cd into project and install dependencies with `go get ./...`

### CLI
1. A built application and 2 example json files have been added at `./main/cli`
2. To rebuild application cd into `./main/cli` and run `go build`
3. To run the application cd into `./main/cli` and run `./cli(.exe) -file='<json/file>' -target-currency='<Target-Currancy>'` for example: `./cli(.exe) -file='test.json' -target-currency='JPY'`

## API
1. To run locally cd into `./main/api` and run `go run main.go`
2. You can hit the endpoint with curl request: 
    ```curl --location --request POST 'http://localhost:4000/conversion' \
       --header 'Content-Type: application/json' \
       --data-raw '{
       	"conversion" : "USD",
       	"data" : [
       		{ "value" : 1.00, "currency" : "EUR" },
       { "value" : 1.00, "currency" : "USD" },
       { "value" : 1.00, "currency" : "JPY" },
       { "value" : 1.00, "currency" : "EUR" }
       		]
       }'

## OPEN DISCUSSION QUESTIONS
1. It will work. Because it does not tries to keep all of te input file in memory. Rather, takes line by line and print as it reads.
2. I have returned {"value":0,"currency":"Error"} if file has malformed JSON line.
3. Currently, I am using very simple cache. So it would work if the currency is requested before. However, considering the production; I would improve my httpclient so that it would try gracefully to obtain higher resiliency. If still no response is coming back; I would have two opsions: 1st return latest currency but add a warning each line and mention the currency date, 2nd return http error mentioning service is down temporarily.

