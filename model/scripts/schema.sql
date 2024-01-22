CREATE DATABASE if not exists store
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LOCALE_PROVIDER = 'libc'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

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