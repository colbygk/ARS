
drop table if exists flights;
drop table if exists airports;


create table airports (
    id int not null auto_increment primary key,
    short_name varchar(25) not null,
    long_name  varchar(100) not null
    ) engine=InnoDB;

create table flights (
    id int not null auto_increment primary key,
    id_str varchar(25) not null,
    depart_airport int not null,
    depart_time datetime not null,
    arrive_airport int not null,
    arrive_time datetime not null,
    foreign key fk_depart_airport(depart_airport) references airports(id),
    foreign key fk_arrive_airport(arrive_airport) references airports(id)
    ) engine=InnoDB;


/* The following insert statements were autogenerated by the
   convert_json_to_sql.js script in ../web */
insert into airports (short_name,long_name) values ('ABQ','ABQ%20%28Albuquerque%20NM%2C%20Sunport%29');
insert into airports (short_name,long_name) values ('BOS','BOS%20%28Boston%20MA%2C%20Logan%29');
insert into airports (short_name,long_name) values ('DFW','DFW%20%28Dallas%20TX%2C%20Dallas-Fort%20Worth%29');
insert into airports (short_name,long_name) values ('OAK','OAK%20%28Oakland%20CA%2C%20Oakland%20International%29');
insert into airports (short_name,long_name) values ('ORD','ORD%20%28Chicago%20IL%2C%20O%27Hare%29');
insert into airports (short_name,long_name) values ('LAX','LAX%20%28Los%20Angeles%20CA%2C%20International%29');
insert into airports (short_name,long_name) values ('NIP','NIP%20%28Pandora%2C%20Na%27vi%20Interstellar%20Port%29');
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('420',(select id from airports where short_name='ABQ'),STR_TO_DATE('04/27/2015 09:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='OAK'),STR_TO_DATE('04/27/2015 10:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('602',(select id from airports where short_name='ABQ'),STR_TO_DATE('04/27/2015 14:05:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='BOS'),STR_TO_DATE('04/27/2015 17:25:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('1577',(select id from airports where short_name='ABQ'),STR_TO_DATE('04/27/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='DFW'),STR_TO_DATE('04/27/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('507',(select id from airports where short_name='ABQ'),STR_TO_DATE('04/27/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='DFW'),STR_TO_DATE('04/27/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('632',(select id from airports where short_name='ABQ'),STR_TO_DATE('04/27/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='NIP'),STR_TO_DATE('04/27/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('1822',(select id from airports where short_name='ABQ'),STR_TO_DATE('04/27/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='LAX'),STR_TO_DATE('04/27/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('705',(select id from airports where short_name='ABQ'),STR_TO_DATE('04/27/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='LAX'),STR_TO_DATE('04/27/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('427',(select id from airports where short_name='ABQ'),STR_TO_DATE('04/27/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='LAX'),STR_TO_DATE('04/27/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('421',(select id from airports where short_name='OAK'),STR_TO_DATE('04/28/2015 09:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='ABQ'),STR_TO_DATE('04/28/2015 10:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('603',(select id from airports where short_name='BOS'),STR_TO_DATE('04/27/2015 14:05:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='ABQ'),STR_TO_DATE('04/27/2015 17:25:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('1578',(select id from airports where short_name='LAX'),STR_TO_DATE('04/27/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='ABQ'),STR_TO_DATE('04/27/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('507',(select id from airports where short_name='ORD'),STR_TO_DATE('04/27/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='ABQ'),STR_TO_DATE('04/27/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('632',(select id from airports where short_name='NIP'),STR_TO_DATE('04/27/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='ABQ'),STR_TO_DATE('04/27/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('1822',(select id from airports where short_name='ORD'),STR_TO_DATE('04/27/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='ABQ'),STR_TO_DATE('04/27/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('705',(select id from airports where short_name='BOS'),STR_TO_DATE('04/27/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='ABQ'),STR_TO_DATE('04/27/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('427',(select id from airports where short_name='DFW'),STR_TO_DATE('04/28/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='ABQ'),STR_TO_DATE('04/27/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('420',(select id from airports where short_name='ABQ'),STR_TO_DATE('04/28/2015 09:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='OAK'),STR_TO_DATE('04/28/2015 10:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('602',(select id from airports where short_name='ABQ'),STR_TO_DATE('04/28/2015 14:05:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='BOS'),STR_TO_DATE('04/28/2015 17:25:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('1577',(select id from airports where short_name='ABQ'),STR_TO_DATE('04/28/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='LAX'),STR_TO_DATE('04/28/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('507',(select id from airports where short_name='ABQ'),STR_TO_DATE('04/28/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='ORD'),STR_TO_DATE('04/28/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('632',(select id from airports where short_name='ABQ'),STR_TO_DATE('04/28/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='NIP'),STR_TO_DATE('04/28/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('1822',(select id from airports where short_name='ABQ'),STR_TO_DATE('04/28/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='LAX'),STR_TO_DATE('04/28/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('705',(select id from airports where short_name='ABQ'),STR_TO_DATE('04/28/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='LAX'),STR_TO_DATE('04/28/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('428',(select id from airports where short_name='ABQ'),STR_TO_DATE('04/28/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='LAX'),STR_TO_DATE('04/28/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('421',(select id from airports where short_name='OAK'),STR_TO_DATE('04/28/2015 09:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='ABQ'),STR_TO_DATE('04/28/2015 10:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('603',(select id from airports where short_name='BOS'),STR_TO_DATE('04/28/2015 14:05:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='ABQ'),STR_TO_DATE('04/28/2015 17:25:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('1578',(select id from airports where short_name='LAX'),STR_TO_DATE('04/28/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='ABQ'),STR_TO_DATE('04/28/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('507',(select id from airports where short_name='ORD'),STR_TO_DATE('04/28/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='ABQ'),STR_TO_DATE('04/28/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('632',(select id from airports where short_name='NIP'),STR_TO_DATE('04/28/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='ABQ'),STR_TO_DATE('04/28/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('1822',(select id from airports where short_name='ORD'),STR_TO_DATE('04/28/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='ABQ'),STR_TO_DATE('04/28/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('705',(select id from airports where short_name='BOS'),STR_TO_DATE('04/28/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='ABQ'),STR_TO_DATE('04/28/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('428',(select id from airports where short_name='LAX'),STR_TO_DATE('04/28/2015 10:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='ABQ'),STR_TO_DATE('04/28/2015 11:45:00','%m/%d/%Y %h:%i'));
insert into flights (id_str,depart_airport,depart_time,arrive_airport,arrive_time) values ('429',(select id from airports where short_name='LAX'),STR_TO_DATE('04/28/2015 11:45:00','%m/%d/%Y %h:%i'),(select id from airports where short_name='ABQ'),STR_TO_DATE('04/28/2015 12:45:00','%m/%d/%Y %h:%i'));
