-- CREATE DATABASE store;
-- \connect store;


begin;

create table if not exists product (
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

commit;