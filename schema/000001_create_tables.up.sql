create table if not exists users(
    id UUID PRIMARY KEY,
    email VARCHAR(120),
    firstname VARCHAR(200),
    lastname VARCHAR(200)
); 


create table if not exists auction(
    id UUID PRIMARY KEY,
    title VARCHAR(400),
    owner UUID, 
);

//participants supposed to be stored in mongo-db with id of auction and array with participants
