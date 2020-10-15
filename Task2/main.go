
// main.go

package main

// importing libraries

import (

	"fmt"
	"log"
	"net/http"
	"time"
	"encoding/json"
	"io/ioutil"
	"strings"
	"strconv"
	
)

// Defining article structure

type Article struct {
    Id          string `json:"id"`
    Title       string `json:"title"`
    SubTitle    string `json:"subtitle"`
    Content     string `json:"content"`
    Timestamp   time.Time `json:"ts"`
}

// declared a global Articles array
// that we can then populate in our main function
// to simulate a database

var Articles []Article

// assign value for offset "default value taken as 0"

func  findOffset(Query1 string) int{

	if(Query1==""){
		return 0
	}else{
		ret1,_  := strconv.Atoi(Query1) 
		return ret1
	}

}

// assign value for limit "default value taken as 5"

func  findLimit(Query2 string) int{

	if(Query2==""){
		return 5
	}else{
		ret2,_  := strconv.Atoi(Query2) 
		return ret2
	}

}


// Pagination for the struct array of Article

func paginate(x []Article, skip int, size int) []Article {

		

		limit := func() int {
		    if skip+size > len(x) {
		        return len(x)
		    } else {
		        return skip + size
		    }

		}

		start := func() int {
		    if skip > len(x) {
		        return len(x)
		    } else {
		        return skip
		    }

		}

		return x[start():limit()]


}



func articleFunction(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

		case "GET":

		   // extract path parameter 

	       ID := r.URL.Path[len("/articles/"):]

	       // extract value assigned with key named q

	       Q  := r.URL.Query().Get("q")
 

	      // if query present [for /articles/search?q=<> GET route ]

	       if(Q != ""){
	       		
	       		for _, article := range Articles {

	       		 	// Loop over all of our Articles
				    // if the article.Title or article.SubTitle or article.Content contains
				    // the key we pass in irrespective of lower or upper case
				    // return the article encoded as JSON

			        if (strings.Contains( strings.ToLower(article.Title),strings.ToLower(Q)) || strings.Contains( strings.ToLower(article.SubTitle), strings.ToLower(Q)) || strings.Contains( strings.ToLower(article.Content),strings.ToLower(Q))){

			            json.NewEncoder(w).Encode(article)

			        }
			    }

	       }else if(ID != ""){            //if path parameter present [for /articles/:id GET route]

				// Loop over all of our Articles
			    // if the article.Id equals the key we pass in
			    // return the article encoded as JSON

			    for _, article := range Articles {
			        if article.Id == ID {

			            json.NewEncoder(w).Encode(article)
			            
			        }
			    }

			}else{

			   // To return all the articles present [for /articles/  GET route]
			   // Extracting query values if present for /articles?offset=<>&limit=<>/

			   Query1 := r.URL.Query().Get("offset")
			   Query2 := r.URL.Query().Get("limit")

			   start := findOffset(Query1)  //funtion to get value for start index of the list

			   limit := findLimit(Query2)  //function to get the end index of the list

			   fmt.Println("Endpoint Hit: returnAllArticles")

			   json.NewEncoder(w).Encode(paginate(Articles, start, limit)) //show articles

			}

		// To add new article [for /articles/ POST route]

		case "POST":   

			    // get the body of our POST request

			    reqBody, _ := ioutil.ReadAll(r.Body)

			     // return the string response containing the request body

			    fmt.Fprintf(w, "%+v", string(reqBody))

			    // unmarshal this into a new Article struct
    			// append this to our Articles array.

			    var article Article 
			    json.Unmarshal(reqBody, &article)

			    // update our global Articles array to include
			    // our new Article

			    Articles = append(Articles, article)
			    json.NewEncoder(w).Encode(article)

		default:

			// In case of an error show method not found

	        w.WriteHeader(http.StatusNotFound)
	        w.Write([]byte(`{"message": "Can't find method requested"}`))
	} 
    
}


// landin page

func homePage(w http.ResponseWriter, r *http.Request){

    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")

}


func handleRequests() {
	
    http.HandleFunc("/", homePage)
   	http.HandleFunc("/articles/",  articleFunction)
    
    log.Fatal(http.ListenAndServe(":10000", nil))

}

func main() {
	
	Articles = []Article{
        Article{ Id : "1", Title: "Hello1", SubTitle: "Article Description1", Content: "Article Content1", Timestamp : time.Now()},
        Article{Id : "2", Title: "Hello2", SubTitle: "Article Description2", Content: "Article Content2", Timestamp : time.Now()},
        Article{ Id : "1", Title: "Hello1", SubTitle: "Article Description1", Content: "Article Content1", Timestamp : time.Now()},
        Article{Id : "2", Title: "Hello2", SubTitle: "Article Description2", Content: "Article Content2", Timestamp : time.Now()},
        Article{ Id : "1", Title: "Hello1", SubTitle: "Article Description1", Content: "Article Content1", Timestamp : time.Now()},
        Article{Id : "2", Title: "Hello2", SubTitle: "Article Description2", Content: "Article Content2", Timestamp : time.Now()},

    }

    handleRequests()
}

