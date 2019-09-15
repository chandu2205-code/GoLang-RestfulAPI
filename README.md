# GoLang-RestfulAPI
Sample REST API using Go lang

# First Rest API using Go
First lets start by :
  * Creating and starting a web server
       - In Go "net/http" package provides all functions required for starting web server , creating REST end points etc .
       - Check whether any port on your system is available to listen http request ,  which can be done by command : <br>
          ***curl localhost:<***any-port-number-where-u-want-to-listen-for-http-request***>***
       - Code snippet for creating a server and listening onto particular port :
          ```Go
          package main
          import ("net/http")
          func main() {
	            err := http.ListenAndServe("localhost:1111", nil)
              if err != nil {
		                panic(err)
	            }
          }
          ```

  * Creating the default route 
  	- Now lets add a default route that will send information about our API when real request is sent 
	- [Go lang net/http package documentation](https://golang.org/pkg/net/http/)
	- ***err := http.ListenandServe('string := url',handler)*** which is going to start server and listen 
	- create a root handler to receive http requests at mentioned URL and serve the response
	- ***http.HandleFunc("/", rootHandler)***
	  ```Go
	     http.HandleFunc("/", rootHandler)
	     
	     func rootHandler(w http.ResponseWriter, r *http.Request) {
	      w.WriteHeader(http.StatusOK)
	      w.Write([]byte("Go Lang Rest service \n"))
             }
	  ```  
  * Sending a message via default route
  	- In the above code snippet [Creating the default route] the message is sent via default route ("/")
  * Creating a custom error message
  	- By default above piece of code , for any incoming request sends messages via default route . In order to  avoid it
	  let's create custom error message .
	  ```Go
	  	/*
		*Method name here starts with small letter 'r'
		*which makes function 'rootHandler' private which is
		*visible only in this package
		*@param writer which is a response writer
		*@param reader which is a pointer to http.Request
 		*/
		func rootHandler(w http.ResponseWriter, r *http.Request) {

			if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Asset not found"))
			return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Go Lang Rest service \n"))
		}
	  ```
  * Using Postman to test APIs
  
