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
  * Using Postman to test APIs<br>
  	Above steps complete to develop a simple GET request using Go lang :smile:
	
# Making API Restful 

* Outlining the API :
	- Our simple API is going to consist just of the access and manipulation of the list of users in a fictional
          development company eg even user with beheading a name a role
	- Most of the time it comes down to what is called CRUD (Create Read Update Delete).
* REST API guidelines for each method :
	- We should use the **post** method to create a resource . The content of the documents should be properly formatted
  item in the case of success . 
  		1 . Our API returns the statuscode (2XX) and a header with the location for the created item to access
                    the collection and retrieve the list of available documents.
 	- We use **get** method on the resource User . The response should be list of users and status codes. To access an individual     item.
  		1 . We need to know its I.D. and then construct the User with it  (User/{id}) . Respose code should be 200(Success) or   		    404(Given id not found) 
  	- To update an item by placing its content use **put** method and pass the full new item content in the request.
		1 . But the response should also be two hundred and then your item content or two or four.
	- Finally **Delete** method with no body to delete an item
	
	-[Guidelines for REST API](https://hackernoon.com/restful-api-designing-guidelines-the-best-practices-60e1d954e7c9)
* Setting Up a Data Store :
	- There is a library called Storm that provides all the functions that we will require or both to perform and then some 	  	  freeing us from micromanagement.
	- Let's also install another useful utility called bson so it's a very nifty tool that was originally created
        for Mongo DB to enable storage for binary encoded decent structures but it also nicely integrates with
        both and we want it just for the ease of creating unique ideas for our records.<br>
		**go get github.com/asdine/storm**<br>
		**go get gopkg.in/mgo.v2/bson**
	- Let's not move on to creating our data record and store access. It's a best practice to create a separate package for your  	  	  data structures to avoid mixing them with your core API code. In case you want to later use it. Sometimes it's even useful to  	 create a separate package for each data type although it highly depends on how much manipulation you do with each.
	- We are only going to have a single data type a user.
	- Let's create the folder user package and then the user.go file.
	- And finally we can create the structure that will hold a single user data.
	- Let's meet user as I mentioned in previous video we are going to give it three fields.
		* The first one I.D. We'll use the Bson object idea type and the two other ones name and control will
                  be strings.
        - The field names must be capitalized to be exported and accessible both from another package and also
          during conversion to Json.

	- Technically this is not however it is customer that the field names in Jason objects are set in kind
	 of case and therefore should start with a lowercase letter.

	- We will use another powerful feature of so called tax tags are descriptors of such fields that can be
          used to add meaning intent and some additional detail about them that I feel type does not convey.
		```Go
			import "gopkg.in/mgo.v2/bson"

			//User holds data for single User
			//Language conventions : 
			type User struct {
				ID   bson.ObjectId `json:"id" storm:"id"`
				Name string        `json:"name"`
				Role string        `json:"role"`
			}
		```
* Record Manipulation : In the above user.go file add all methods which handles CRUD operation on User
* Defining custom handler : 
	- Common route - /users route
	- Algorithm :
		* Recognize the /users route
		* Parse the URL and method to call a proper handler
		* Retrieve or process data 
		* send the reposnse
  
