CREATE EXTENSION Postgis;

CREATE TABLE country(
       country_id serial PRIMARY KEY,
       country_name varchar(255) UNIQUE NOT NULL,
       country_phone_code varchar(10) NOT NULL
);

INSERT INTO country (country_name,country_phone_code) VALUES ('India','91');
CREATE TABLE state(
       state_id serial PRIMARY KEY,
       state_name varchar(255) NOT NULL,
       country_id integer,
       CONSTRAINT state_country_id_fkey FOREIGN KEY (country_id)
       REFERENCES country (country_id) MATCH SIMPLE 
       ON DELETE CASCADE,
       unique (state_name, country_id)
);
INSERT INTO state (state_name, country_id) VALUES
('Kerala', 1);

CREATE TABLE district(
       district_id serial PRIMARY KEY,
       district_name varchar(255) NOT NULL,
       state_id integer,
       CONSTRAINT dsitrict_state_id_fkey FOREIGN KEY (state_id)
       REFERENCES state (state_id) MATCH SIMPLE 
       ON DELETE CASCADE,
       unique (district_name, state_id)
);

INSERT INTO district (district_id, district_name, state_id) VALUES
(1,'Alappuzha',1),
(2,'Ernakulam',1),
(3,'Idukki',1),
(4,'Kannur',1),
(5,'Kasaragod',1),
(6,'Kollam',1),
(7,'Kottayam',1),
(8,'Kozhikode',1),
(9,'Malappuram',1),
(10,'Palakkad',1),
(11,'Pathanamthitta',1),
(12,'Thiruvananthapuram',1),
(13,'Thrissur',1),
(14,'Wayanad',1);


CREATE TABLE city_type (
       city_type_id serial PRIMARY KEY,
       city_type_name varchar(100)
);

CREATE TABLE city(
       city_id serial PRIMARY KEY,
       city_name varchar(255) NOT NULL,
       city_type varchar(32),
       location geography,
       latitude varchar(32) not null default '',
       longitude varchar(32) not null default '',
       district_id integer,
       district_centre boolean not null default false,
       CONSTRAINT city_district_id_fkey FOREIGN KEY (district_id)
       REFERENCES district (district_id) MATCH SIMPLE 
       ON DELETE CASCADE,
       unique (city_name, district_id)
);


CREATE TABLE company(
       company_id serial PRIMARY KEY,
       company_name VARCHAR (1024) NOT NULL,
       company_description text,
       company_type varchar(100) NOT NULL default 'sp',
       company_reg_no varchar(20) NOT NULL default '',
       company_pan varchar(20) NOT NULL default '',
       company_gst varchar(20) NOT NULL default '',
       company_website varchar(255) NOT NULL default '',
       company_email varchar(255) NOT NULL default '',
       individual boolean NOT NULL DEFAULT false,
       company_owner integer not null,
       active boolean,
       created_date timestamp not null default NOW()
       );

CREATE TABLE company_address(
       company_address_id serial PRIMARY KEY,
       company_address_line1 VARCHAR(255) NOT NULL default '',
       company_address_line2 VARCHAR(255) NOT NULL default '',
       company_phone1 VARCHAR(255) NOT NULL default '',
       company_phone2 VARCHAR(255) NOT NULL default '',
       company_phone3 VARCHAR(255) NOT NULL default '',
       company_email VARCHAR(255) NOT NULL default '',
       company_postal_code varchar(10) NOT NULL default '',
       city_id integer,
       latitude varchar(32) not null default '',
       longitude varchar(32) not null default '', 
       district_id integer,
       state_id integer,
       main boolean default false,
       company_id integer,
       CONSTRAINT company_city_id_fkey FOREIGN KEY (city_id)
       REFERENCES city (city_id) MATCH SIMPLE 
       ON DELETE CASCADE,
       CONSTRAINT company_address_company_id_fkey FOREIGN KEY (company_id)
       REFERENCES company (company_id) MATCH SIMPLE 
       ON DELETE CASCADE
);

CREATE TABLE company_manager(
       company_manager_id serial PRIMARY KEY,
       company_id integer,
       user_account_id integer,
       CONSTRAINT company_manger_user_account_id_fkey FOREIGN KEY (user_account_id)
       REFERENCES user_account (user_account_id) MATCH SIMPLE 
       ON DELETE CASCADE,
       CONSTRAINT company_manger_company_id_fkey FOREIGN KEY (company_id)
       REFERENCES company (company_id) MATCH SIMPLE 
       ON DELETE CASCADE
);