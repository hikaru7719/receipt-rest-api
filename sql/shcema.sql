create table receipt (
  id int primary key auto_increment not null,
  name varchar(255) not null ,
  type varchar(255) not null ,
  date date not null,
  memo varchar(255) not null
);