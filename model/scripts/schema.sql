begin;

create table if not exists product (
	id serial primary key,
	name varchar(255) not null,
	description varchar(255),
	price decimal,
	quantity integer
);

commit;