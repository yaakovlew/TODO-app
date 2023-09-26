CREATE TABLE IF NOT EXISTS task(
    id bigserial not null Primary key,
    header varchar(300) not null,
    description varchar(500),
    date DATE not null,
    is_done bool not null 
)