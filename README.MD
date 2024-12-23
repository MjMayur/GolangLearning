 type handler interface
 handler : serve http method response writer pointing to request 

 ListenAndServe takes handler

handleFunc its a function which takes response writer and request in parameters 
handler func have its own type and request attached to it


example
package main
import (
"10"
"net/http"
)
    func d(res http. ResponseWriter, req *http.Request) { 
        io.WriteString(res, "dog dog dog")
    }
    func c(res http. ResponseWriter, req *http.Request) { 
        io.WriteString(res, "cat cat cat")
    }
    func main() {
        http.Handle("/dog", http.HandlerFunc(d)) http.Handle("/cat", http.HandlerFunc(c))
        http. ListenAndServe(":8080", nil)
    }

HII
