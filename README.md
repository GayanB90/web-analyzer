# web-analyzer
A simple web analyzer written in GoLang

## Project Overview
    Web analyzer is a simple HTML web page analyzing application. Users can provide the URL of the web page to be analyzed and the
    analyzer will parse the page and provide following information as the analysis result.
    01. The HTML version of the page
    02. The title text of the HTML page
    03. Whether the HTML page contains a login form
    04. No of headings and their level respectively
    05. All hyperlinks available in the page
    06. All broken links in the page

## Prerequisites
    GoLang 1.22.4

## How to build
   Run 'go build -o web-analyzer ./cmd/webanalyzer/main.go' 
   
## How to run
   Run './web-analyzer'

### FE URL for localhost
   http://localhost:8080/static/index.html
### Backend URL
   http://localhost:8080/analyze
### CURL command for backend endpoint
   curl http://localhost:8080/analyze \           
   --include \
   --header "Content-Type: application/json" \
   --request "POST" \
   --data '{"requestId":"1234", "webUrl":"http://www.google.com"}'
## Using docker
   Run 'docker build -t web-analyzer .' to build the docker image
   Run 'docker run -p 8080:8080 web-analyzer' to run the image
## Assumptions
   It is assumed that both signup and login forms will not be rendered in the same web page
   It is assumed that login form will be using a formal 'Submit' button for submitting the form
   It is assumed that signup form will use terms 'signup' or 'register' as part of the html form
## Approaches to improve the application
   We can move the lookup terms of signup forms into a config for better flexibility in the logic
   We can integrate a centralized logging middleware like DataDog/Kibana since the application publishes JSON formatted logs to the default output
   We can push the performance metrics to the logging middleware/monitoring solution and improve the observability of the application
   
   