package main

import (
    "fmt"
    "net/http"
    //"html/template"
    //"log"
    "strings"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8") 
	
	fmt.Fprintf(w, "<html><head></head><body>")

        urlValue := r.URL.Path[1:] 
        
	if (urlValue == "") {
	fmt.Fprintf(w, "Hello World! ")
        } else if (urlValue == "Hello"||urlValue == "World") {
	fmt.Fprintf(w, "%s ", urlValue)
	
        if(r.URL.Query()!=nil){
           ifUpperCase := r.URL.Query().Get("uppercase")
           ifReverse := r.URL.Query().Get("reverse") 
	   if (ifUpperCase == "true"){
 	   fmt.Fprintf(w, "Value After Capitalization: %s <br>",strings.ToUpper(urlValue))
           }
	   if (ifReverse == "true"){
           fmt.Fprintf(w, "Value After Reversing: %s <br>",strReverse(urlValue))
           }  
        }
        } else{
        fmt.Fprintf(w, "Invalid Option: %s ",urlValue)       
        }
}

func strReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
    http.HandleFunc("/", Homepage)
    http.ListenAndServe(":8090", nil)
}
