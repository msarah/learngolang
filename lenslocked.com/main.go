package main //every go file must have a package name specified

import ( //these are packages imported from the golang standard library
	"fmt"      //formatting strings
	"net/http" //tools for web development
)

/*
handlerFunc takes a response writer and pointer to a request.
These two objects allow us to transmit data over the web.
When a user makes a request this function takes that request
and uses the response writer object to send back a reply.

We can do this because the ResponseWriter has the Write() method
which is what we use to write to the response body. We can also use
this object to write headers when we need to.

The pointer to request object allows us to access any information
the user may have sent with the request e.g. email for sign in.
*/

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email "+
			"to <a href=\"mailto:support@lenslocked.com\">"+
			"support@lenslocked.com</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find the page you "+
			"were looking for :(</h1>"+
			"<p>Please email us if you keep being sent to an "+
			"invalid page.</p>")
	}
}

/*
  Fprint takes
  (1) and io.Writer to write to and
  (2)Any number of interface{}s to print out. Typically strings, but could be any data type.

  An io.Writer is an interface that requires a struct to have implemented the Write([]byte) method
  so fmt.Fprint helps handle converting all of the provided interfaces to a byte array,
  and then calls the Write method on the io.Writer.

  Since we are writing a string, and strings can be treated as byte arrays,
  you could replace the line

  fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>") with
  w.Write([]byte("<h1>Welcome to my awesome site!</h1>"))

  and you would end up getting the same end result.
*/

func main() {
	http.HandleFunc("/", handlerFunc) // "/" handles all path requests
	http.ListenAndServe(":3000", nil) // starts up a local web server using default handlers
}
