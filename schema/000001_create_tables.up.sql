create table if not exists users(
    id UUID PRIMARY KEY,
    email VARCHAR(120),
    firstname VARCHAR(200),
    lastname VARCHAR(200),
    password VARCHAR(300);
); 


create table if not exists auctions(
    id UUID PRIMARY KEY,
    title VARCHAR(400),
    owner UUID, 
);

