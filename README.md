# first-go-api
My first go API

on run docker-compose up access `localhost:16543/` and log in with PGADMIN_DEFAULT_EMAIL and PGADMIN_DEFAULT_PASSWORD used on docker compose.

to register a server: 

- Rigth click on `Servers` > `Register` > `Servers` (again xD)
- On section `Connection` set host name to `postgres-contaner`
    - If you don't change this values:
        - Username: postgres
        - password: postgres


script:
create table if not exists  product (
	id serial primary key,
	name varchar(255) not null,
	description varchar(255),
	price decimal,
	quantity integer
);

insert into product(name, description, price, quantity)
	values('Shoe', 'Blue', 10, 2);
	
insert into product(name, description, price, quantity)
	values('Tshirt', 'White', 8, 1);
	
select * from product;