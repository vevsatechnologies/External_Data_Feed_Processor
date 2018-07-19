create table data (
id serial primary key not null
);


create table historic_data (
id serial primary key not null ,
exchangeID integer,
globaltradeid varchar (20),
tradeid varchar (20),
timestamp timestamp,
quantity  varchar (30),
price varchar (20), 
total varchar (20),
fill_type varchar (20),
order_type varchar (20)
);
