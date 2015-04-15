package main

import (
	"flag"
	"fmt"	
	"log"
    "net"
    "encoding/gob"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

var connectStr = "ars:ARSePassW0rd@/ARSdb?parseTime=true"
var dbType = "mysql"

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

type Packet struct {
	Flights []Flight
	Airports []Airport
}


//
var tickets = make([]ticket, 0)
var selected_ticket_id = 0
var dbmap *gorp.DbMap

//
func checkErr(err error, msg string) bool {

	if err != nil {
		log.Fatalln(msg, err)
	}

	return err != nil
}

func initDB() error {
	db, err := sql.Open(dbType, connectStr)

	if checkErr(err, "Database connection failed, sql.Open") {
		log.Printf("ERROR?  %+v", connectStr);	
		return err
	}

	if dbmap == nil {
		log.Printf("Started Dialing");	
		dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
		log.Printf("Finshed Dialing");	
	}

	log.Printf("Finished Init DB  \n");	

	return nil
}

func handleConnection(conn net.Conn) {

    // Recived connection request
    fmt.Printf("Received connection");
    encoder := gob.NewEncoder(conn)

    // Grab DB data
    var flights []Flight
    var airports []Airport
    listFlightsDB(&flights)
    listAirportsDB(&airports)

   	// Prepare packet to send	
	p := &Packet{}

    for _, element := range flights {
    	p.Flights = append(p.Flights, element)
	}

	for _, element := range airports {
    	p.Airports = append(p.Airports, element)
	}

    //fmt.Printf("Sending : %+v", p);
    encoder.Encode(p)
    fmt.Println("Done sending packet to client");
}


func startDBInterface() {
   fmt.Println("starting DB Service Proxy...");
   ln, err := net.Listen("tcp", ":8080")

   if err != nil {
   		// handle error
   }
   for {
        conn, err := ln.Accept() // this blocks until connection or error
        if err != nil {
            // handle error
            continue
        }
        go handleConnection(conn) // a goroutine handles conn so that the loop can accept other connections
    }
}


func listFlightsDB(flights *[]Flight){
	_, err := dbmap.Select(flights, "select f.id, f.id_str, da.short_name as depart_airport, f.depart_time, aa.short_name as arrive_airport, f.arrive_time from flights f, airports da, airports aa where f.depart_airport=da.id and f.arrive_airport=aa.id")
	
	if checkErr(err, "select from flights") {
		return;
	}

    log.Printf("listFlights called\n");
    return;
}

func listAirportsDB(airports *[]Airport) {
	_, err := dbmap.Select(airports, "select id, short_name, long_name from airports")
	if checkErr(err, "select from airports") {
		return;
	}

	return;
}


func main() {

	// command line flags
	port := flag.Int("port", 81, "port to serve on")
	dir  := flag.String("directory", "web/", "directory of web files")
	flag.Parse()

	log.Printf("", port, dir);

	// Initialize table mappings
	log.Printf(" Init DB\n")
	err := initDB()
	if err != nil {
		return
	}

	// Start the proxy server to DB
	startDBInterface();
}