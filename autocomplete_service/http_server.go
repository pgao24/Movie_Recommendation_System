package main

import (
	"fmt"
	"net/http"
	"io"
)

func root_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, World!")
}

func relay_google(w http.ResponseWriter, r *http.Request){
 	res, err := http.Get("https://google.com")
 	if err != nil {
 		fmt.Fprintf(w, "error making http request: %s\n", err)
  		return
	}
        defer res.Body.Close()
	if res.StatusCode == http.StatusOK {    	
		bodyBytes, err := io.ReadAll(res.Body)
    		if err != nil {
        		fmt.Fprintf(w, "error accessing response body: %s\n", err)
    			return
		}
    		bodyString := string(bodyBytes)
    		fmt.Fprintf(w, bodyString)
		return
	} 
	fmt.Fprintf(w, "error getting reponse status: %s\n", res.StatusCode)
}

func search_google(w http.ResponseWriter, r *http.Request){
	// Handling user query
	queryValue := r.URL.Query().Get("q")
	if queryValue != "" {
		fmt.Println("Retriving google search result with:", queryValue)
	}
	
	// Modify target URL with extracted queryValue
	targetURL := fmt.Sprintf("https://www.google.com/search?q=%s", queryValue)  
	
 	res, err := http.Get(targetURL)
 	if err != nil {
 		fmt.Fprintf(w, "error making http request: %s\n", err)
  		return
	}
        defer res.Body.Close()
	if res.StatusCode == http.StatusOK {    	
		bodyBytes, err := io.ReadAll(res.Body)
    		if err != nil {
        		fmt.Fprintf(w, "error accessing response body: %s\n", err)
    			return
		}
    		bodyString := string(bodyBytes)
    		fmt.Fprintf(w, bodyString)
		return
	} 
	fmt.Fprintf(w, "error getting reponse status: %s\n", res.StatusCode)

}

func main() {
	fmt.Println("Starting server")

	// Handler function for the :root path ("/")
	http.HandleFunc("/", root_handler)

        http.HandleFunc("/relay_google", relay_google)

	http.HandleFunc("/search", search_google)

	// Start the server listening on port 8080
	fmt.Println("Server listening on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
