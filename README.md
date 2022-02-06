# project2
Service that calls uzma24/project1 service, takes input from .txt file and prints JSON output returned from the service. Program can take large input files. 

## How to run the project:
1. Clone the repo in your local along with repo : ```uzma24/project1```   
2. Run both main.go file(use command from terminal: go run main.go)  
3. Use this curl request in postman:``` curl --location --request GET 'localhost:8081/getFrequency' \
--data-raw '' ```  

4. The API will return 10 most frequent words with frequency in JSON format
