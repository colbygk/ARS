package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// error response contains everything we need to use http.Error
type handlerError struct {
	Error   error
	Message string
	Code    int
}

// ticket model
type ticket struct {
	Passenger string `json:"passenger"`
	Flight    string `json:"flight"`
	Id        int    `json:"id"`
}

// list of all of the tickets
var tickets = make([]ticket, 0)

// a custom type that we can use for handling errors and formatting responses
type handler func(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError)

// attach the standard ServeHTTP method to our handler so the http library can call it
func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// here we could do some prep work before calling the handler if we wanted to

	// call the actual handler
	response, err := fn(w, r)

	// check for errors
	if err != nil {
		log.Printf("ERROR: %v\n", err.Error)
		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Message), err.Code)
		return
	}
	if response == nil {
		log.Printf("ERROR: response from method is nil\n")
		http.Error(w, "Internal server error. Check the logs.", http.StatusInternalServerError)
		return
	}

	// turn the response into JSON
	bytes, e := json.Marshal(response)
	if e != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	// send the response and log
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
	log.Printf("%s %s %s %d", r.RemoteAddr, r.Method, r.URL, 200)
}

func listTickets(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	return tickets, nil
}

func getTickets(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	// mux.Vars grabs variables from the path
	param := mux.Vars(r)["id"]
	id, e := strconv.Atoi(param)
	if e != nil {
		return nil, &handlerError{e, "Id should be an integer", http.StatusBadRequest}
	}
	b, index := getTicketById(id)

	if index < 0 {
		return nil, &handlerError{nil, "Could not find ticket " + param, http.StatusNotFound}
	}

	return b, nil
}

func parseTicketRequest(r *http.Request) (ticket, *handlerError) {
	// the ticket payload is in the request body
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return ticket{}, &handlerError{e, "Could not read request", http.StatusBadRequest}
	}

	// turn the request body (JSON) into a ticket object
	var payload ticket
	e = json.Unmarshal(data, &payload)
	if e != nil {
		return ticket{}, &handlerError{e, "Could not parse JSON", http.StatusBadRequest}
	}

	return payload, nil
}

func addTicket(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	payload, e := parseTicketRequest(r)
	if e != nil {
		return nil, e
	}

	// it's our job to assign IDs, ignore what (if anything) the client sent
	payload.Id = getNextId()
	tickets = append(tickets, payload)

	// we return the ticket we just made so the client can see the ID if they want
	return payload, nil
}

func updateTicket(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	payload, e := parseTicketRequest(r)
	if e != nil {
		return nil, e
	}

	_, index := getTicketById(payload.Id)
	tickets[index] = payload
	return make(map[string]string), nil
}

func removeTicket(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	param := mux.Vars(r)["id"]
	id, e := strconv.Atoi(param)
	if e != nil {
		return nil, &handlerError{e, "Id should be an integer", http.StatusBadRequest}
	}
	// this is jsut to check to see if the ticket exists
	_, index := getTicketById(id)

	if index < 0 {
		return nil, &handlerError{nil, "Could not find entry " + param, http.StatusNotFound}
	}

	// remove a ticket from the list
	tickets = append(tickets[:index], tickets[index+1:]...)
	return make(map[string]string), nil
}

// searches the tickets for the ticket with `id` and returns the ticket and it's index, or -1 for 404
func getTicketById(id int) (ticket, int) {
	for i, b := range tickets {
		if b.Id == id {
			return b, i
		}
	}
	return ticket{}, -1
}

var id = 0

// increments id and returns the value
func getNextId() int {
	id += 1
	return id
}

func main() {
	// command line flags
	port := flag.Int("port", 80, "port to serve on")
	dir := flag.String("directory", "web/", "directory of web files")
	flag.Parse()

	// handle all requests by serving a file of the same name
	fs := http.Dir(*dir)
	fileHandler := http.FileServer(fs)

	// setup routes
	router := mux.NewRouter()
	router.Handle("/", http.RedirectHandler("/static/", 302))
	router.Handle("/tickets", handler(listTickets)).Methods("GET")
	router.Handle("/tickets", handler(addTicket)).Methods("POST")
	router.Handle("/tickets/{id}", handler(getTickets)).Methods("GET")
	router.Handle("/tickets/{id}", handler(updateTicket)).Methods("POST")
	router.Handle("/tickets/{id}", handler(removeTicket)).Methods("DELETE")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	http.Handle("/", router)

	// bootstrap some data
	tickets = append(tickets, ticket{"Nathan Acosta", "NMA460", getNextId()})
	tickets = append(tickets, ticket{"Colby Gutierrez-Kraybill", "NMA361", getNextId()})
	tickets = append(tickets, ticket{"Robert Herbertson", "NMA351", getNextId()})
	tickets = append(tickets, ticket{"Max Larson", "NMA362", getNextId()})
	tickets = append(tickets, ticket{"Martin Lidy", "NMA357", getNextId()})

	log.Printf("Running on port %d\n", *port)

	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
