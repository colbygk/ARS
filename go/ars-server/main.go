package main

import (
	"os"
	"time"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"log"
	"encoding/gob"	
	"strconv"
	"github.com/gorilla/mux"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

var connectStr = "ars:ARSePassW0rd@/ARSdb?parseTime=true"
var dbType = "mysql"

//
var flights []Flight
var airports []Airport

func checkErrDB(err error, msg string) bool {

	if err != nil {
		log.Fatalln(msg, err)
	}

	return err != nil
}

// error response contains everything we need to use http.Error
type handlerError struct {
	Error   error
	Message string
	Code    int
}

// ticket model
type ticket struct {
	FirstName        string `json:"firstname"`
	LastName         string `json:"lastname"`
	SourceCity       string `json:"arrivecity"`
	DepartCity       string `json:"departcity"`
	FlightID         string `json:"flightid"`
	FlightDate       string `json:"flightdate"`
	Id               int    `json:"id"`
	BackgroundArrive string `json:"backgroundarrive"`
	BackgroundDepart string `json:"backgrounddepart"`
	SelectedSource   string `json:"selectedsource"`
	SelectedDest     string `json:"selecteddest"`
}

// list of all of the tickets
var tickets = make([]ticket, 0)
var selected_ticket_id = 0

type Airport struct {
	Id        int    `json:"id"`
	ShortName string `json:"shortname" db:"short_name"`
	LongName  string `json:"longname" db:"long_name"`
}

type Flight struct {
	Id            int       `json:"id" db:"id"`
	IdStr         string    `json:"idstr" db:"id_str"`
	DepartAirport string    `json:"departairport" db:"depart_airport"`
	DepartTime    time.Time `json:"departtime" db:"depart_time"`
	ArriveAirport string    `json:"arriveairport" db:"arrive_airport"`
	ArriveTime    time.Time `json:"arrivetime" db:"arrive_time"`
	TicketPrice   string    `json:"ticketprice" db:"price"`
}

var dbmap *gorp.DbMap

func initDB() error {
	db, err := sql.Open(dbType, connectStr)

	if checkErr(err, "Database connection failed, sql.Open") {
		return err
	}

	if dbmap == nil {
		dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	}

	return nil
}

// a custom type that we can use for handling errors and formatting responses
type handler func(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError)


func checkErr(err error, msg string) bool {

	if err != nil {
		log.Fatalln(msg, err)
	}

	return err != nil
}

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

func listFlights(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {	
	// DB interface call
	//listFlightsDB(&flights, w, r)
	testing(&flights, &airports)
	fmt.Printf("\n--List Flights--\n");

	// Return DB json info
	return flights, nil
}

func listAirports(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	fmt.Printf("\n--List Airports--\n");

	/*
	_, err := dbmap.Select(&airports, "select id, short_name, long_name from airports")
	if checkErr(err, "select from airports") {
		return nil, nil
	}*/

	return airports, nil
}

func listTickets(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	fmt.Printf("List Tickets");
	return tickets, nil
}

func getTicketID(w http.ResponseWriter, r *http.Request) (interface{}, *handlerError) {
	return selected_ticket_id, nil
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

type P struct {
    M, N int64
    A string
}

type Packet struct {
	Flights []Flight
	Airports []Airport
}

func handleConnection(conn net.Conn, flights *[]Flight, airports *[]Airport) {
    dec := gob.NewDecoder(conn)
    p := &Packet{}
 
    dec.Decode(p)
    flights = &p.Flights
    airports = &p.Airports
    //tickets = &p.Tickets

    fmt.Printf("Received : %+v", airports);
}


func startDBInterface2() {

	// Dial out to send request
    fmt.Println("I want to send some data to DB interface");
    conn, err := net.Dial("tcp", "localhost:8080")

    if err != nil {
        log.Fatal("Connection error", err)
    }
    
    encoder := gob.NewEncoder(conn)
    p := Packet{}

    // Send a request asking for flights[]
    fmt.Printf("Sending empty packet as request");
    encoder.Encode(p)

    // Recive flights[] from proxy
    dec := gob.NewDecoder(conn)
    p2 := Packet{}
    dec.Decode(p2)
    fmt.Printf("Received : %+v", p);

    // Close connection and listen for reply
    conn.Close()
    fmt.Println("done");
}

func testing(flights *[]Flight, airports *[]Airport) {
    strEcho := "Asking DB interface"
    servAddr := "localhost:8080"

    tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
    if err != nil {
        println("ResolveTCPAddr failed:", err.Error())
        os.Exit(1)
    }
 
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        println("Dial failed:", err.Error())
        os.Exit(1)
    }
 
 	// Send a small packet to let the DB proxy know to send the flight information
 	// This is where we can add in the date information
 	// And other searching info.
    _, err = conn.Write([]byte(strEcho))
    if err != nil {
        println("Write to server failed:", err.Error())
        os.Exit(1)
    }
 
    println("write to server = ", strEcho)
 
    // Handle the reply
    handleConnection(conn, flights, airports)
    conn.Close()
}


func main() {
	// command line flags
	port := flag.Int("port", 80, "port to serve on")
	dir := flag.String("directory", "web/", "directory of web files")
	flag.Parse()

	// handle all requests by serving a file of the same name
	fs := http.Dir(*dir)
	fileHandler := http.FileServer(fs)

	// Initialize table mappings
	err := initDB()
	if err != nil {
		return
	}

	// setup routes
	router := mux.NewRouter()
	router.Handle("/", http.RedirectHandler("/static/", 302))
	router.Handle("/flights", handler(listFlights)).Methods("GET")
	router.Handle("/airports", handler(listAirports)).Methods("GET")
	router.Handle("/tickets", handler(listTickets)).Methods("GET")
	router.Handle("/tickets", handler(addTicket)).Methods("POST")
	router.Handle("/selected_ticket_id", handler(getTicketID)).Methods("GET")
	router.Handle("/tickets/{id}", handler(getTickets)).Methods("GET")
	router.Handle("/tickets/{id}", handler(updateTicket)).Methods("POST")
	router.Handle("/tickets/{id}", handler(removeTicket)).Methods("DELETE")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileHandler))
	http.Handle("/", router)

	// bootstrap some data
	log.Printf("Running on port %d\n", *port)
	addr := fmt.Sprintf(":%d", *port)

	// this call blocks -- the progam runs here forever
	err = http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
