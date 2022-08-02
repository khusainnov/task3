CREATE TABLE tax_rate (
    zip_code int not null primary key unique,
    rate numeric(10, 8) not null
);