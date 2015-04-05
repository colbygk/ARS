
// Imports ./flights.json and ./airports.json
var flights = require('./flights');
var airports = require('./airports');

// Important to add the airports first
// to satisfy foreign key constraints
// from flights->airports
var insert_str;
airports.forEach( function(e) {
    insert_str = "insert into airports " +
    "(short_name,long_name) values ('"+e.code+"','"+escape(e.desc)+"');";
    console.log(insert_str);
    });

flights.forEach( function(e) {
    insert_str = "insert into flights " +
     "(id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('" +
     e.id+"',(select id from airports where short_name='"+e.origin+"'),'"+e.depart+
     ":00',(select id from airports where short_name='"+e.destination+"'),'"+e.arrive+
     ":00');";
     console.log(insert_str.replace(/\//g,"-"));
    });
