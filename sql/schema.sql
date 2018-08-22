create table receipts (
  id int primary key auto_increment not null,
  name varchar(255) not null ,
  price int not null,
  kind varchar(255) not null ,
  date date not null,
  memo varchar(255) not null
);