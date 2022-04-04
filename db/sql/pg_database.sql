CREATE TABLE "customers" (
                             "customer_id" bigserial PRIMARY KEY NOT NULL,
                             "name" varchar(100) NOT NULL,
                             "date_of_birth" date NOT NULL,
                             "city" varchar(100) NOT NULL,
                             "zipcode" varchar(10) NOT NULL,
                             "status" SMALLINT NOT NULL DEFAULT 1
);

INSERT INTO customers VALUES
                          (1,'Steve','1978-12-15','Delhi','110075',1),
                          (2,'Arian','1988-05-21','Newburgh','12550',1),
                          (3,'Hadley','1988-04-30','Inglewood','07631',1),
                          (4,'Ben','1988-01-04','Manchester','03102',0),
                          (5,'Nina','1988-05-14','Blackstone','48348',1),
                          (6,'Osman','1988-11-08','Hyattsville','20782',0);
ALTER SEQUENCE customers_customer_id_seq RESTART WITH 7;