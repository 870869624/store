create table if not exists CODE (
    id int(11) not null auto_increment primary key,
    code varchar(11) not null,
    user_id int(11) not null 
)