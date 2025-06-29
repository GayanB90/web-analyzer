# web-analyzer
A simple web analyzer written in GoLang

01. Prerequisites
    GoLang 1.22.4
2. How to build
   Run 'go build -o web-analyzer ./cmd/webanalyzer/main.go'
3. How to run
   Run './web-analyzer'
4. FE URL for localhost
   http://localhost:8080/static/index.html
5. Backend URL
   http://localhost:8080/analyze
6. CURL command for backend endpoint
   curl http://localhost:8080/analyze \           
   --include \
   --header "Content-Type: application/json" \
   --request "POST" \
   --data '{"requestId":"1234", "webUrl":"http://www.google.com"}'
7. Using docker
   Run 'docker build -t web-analyzer .' to build the docker image
   Run 'docker run -p 8080:8080 web-analyzer' to run the image
8. Assumptions
   It is assumed that both signup and login forms will not be rendered in the same web page
   It is assumed that login form will be using a formal 'Submit' button for submitting the form
   It is assumed that signup form will use terms 'signup' or 'register' as part of the html form
9. Approaches to improve the application
   We can move the lookup terms of login form and signup forms into a config for better flexibility in the logic
   