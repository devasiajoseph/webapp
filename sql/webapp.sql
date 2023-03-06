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
(14,'Wayanad',1),
(15,'Lakshadweep',1),
(16,'Coimbatore',1),
(17,'Tenkasi',1),
(18,'Kanyakumari',1),
(19,'Tirunelveli',1);


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


CREATE TABLE user_account(
       user_account_id serial PRIMARY KEY,
       phone VARCHAR(20) UNIQUE NOT NULL,
       email VARCHAR (355) NOT NULL default '',
       password VARCHAR (1024) NOT NULL,       
       full_name VARCHAR (250) NOT NULL default '',
       active BOOLEAN NOT NULL default false,
       created_on TIMESTAMP NOT NULL default NOW(),
       last_login TIMESTAMP,
       role VARCHAR (20) not null default 'user'
       );

CREATE INDEX index_user_account_email ON user_account(email);
CREATE INDEX index_user_account_phone ON user_account(phone);

CREATE TABLE user_address(
       user_address_id serial PRIMARY KEY,
       address_line_1 VARCHAR (255) NOT NULL default '',
       address_line_2 VARCHAR (255) NOT NULL default '',
       city_id integer NULL,
       CONSTRAINT user_account_city_id_fkey FOREIGN KEY (city_id)
       REFERENCES city (city_id) MATCH SIMPLE 
       ON DELETE CASCADE
);



CREATE TABLE user_keys(
       user_keys_id serial PRIMARY KEY,
       user_account_id integer NOT NULL,
       key_name VARCHAR (255),
       key_value VARCHAR (255),
       otp varchar(6),
       CONSTRAINT user_registration_user_account_id_fkey FOREIGN KEY (user_account_id)
       REFERENCES user_account (user_account_id) MATCH SIMPLE 
       ON DELETE CASCADE
       );

CREATE TABLE user_session(
       user_session_id serial PRIMARY KEY,
       user_account_id integer NOT NULL,
       auth_token VARCHAR(64),
       session_expiry TIMESTAMP NOT NULL,
       CONSTRAINT user_session_user_account_id_fkey FOREIGN KEY (user_account_id)
       REFERENCES user_account (user_account_id) MATCH SIMPLE 
       ON DELETE CASCADE
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

CREATE TABLE company_user(
       company_user_id serial PRIMARY KEY,
       company_id integer,
       user_account_id integer,
       role varchar(10) not null default 'manager',
       CONSTRAINT company_manger_user_account_id_fkey FOREIGN KEY (user_account_id)
       REFERENCES user_account (user_account_id) MATCH SIMPLE 
       ON DELETE CASCADE,
       CONSTRAINT company_manger_company_id_fkey FOREIGN KEY (company_id)
       REFERENCES company (company_id) MATCH SIMPLE 
       ON DELETE CASCADE
);

 ALTER TABLE company_user 
       ADD CONSTRAINT roles 
       CHECK (role IN ('admin', 'manager','owner'));

CREATE TABLE item(
       item_id serial PRIMARY KEY,
       item_name VARCHAR (1024) NOT NULL,
       item_category VARCHAR (1024) NOT NULL,
       item_type VARCHAR NOT NULL default 'service',
       item_brand varchar (255) not null default '',
       item_description text,       
       item_price numeric(10,2),
       item_price_discuss boolean not null default false,
       item_service_interval varchar(20),
       item_stock_unit varchar(10),
       item_stock decimal(10,3),
       item_available boolean NOT NULL DEFAULT FALSE,
       item_available_from TIMESTAMP NULL,
       item_available_to TIMESTAMP NULL,
       item_dim_unit varchar(10),
       item_length decimal(10,3),
       item_width decimal(10,3),
       item_height decimal(10,3),
       item_weight_unit varchar(10),
       item_weight decimal(10,3),
       item_on_offer boolean NULL,
       item_offer_price numeric(10,2) NOT NULL DEFAULT 0.0,
       item_offer_start_date TIMESTAMP NULL,
       item_offer_end_date TIMESTAMP NULL,
       allow_preorder boolean NOT NULL DEFAULT FALSE,
       allow_backorder boolean NOT NULL DEFAULT FALSE,
       shipping_cost decimal(10,3),
       shipping_instructions text,
       url_slug varchar(255),
       company_id integer,
       CONSTRAINT item_company_id_fkey FOREIGN KEY (company_id)
       REFERENCES company (company_id) MATCH SIMPLE 
       ON DELETE CASCADE

       /*ALTER TABLE item 
       ADD CONSTRAINT item_types 
       CHECK (item_type IN ('service', 'product')*/ 
       );

CREATE TABLE item_image(
       image_id serial PRIMARY KEY,
       original_image VARCHAR(255),
       resized_image VARCHAR(255),
       filename VARCHAR(255), 
       thumbnail VARCHAR(255),
       display_thumbnail BOOLEAN NOT NULL default false,
       uploaded_time TIMESTAMP NOT NULL,
       file_size INTEGER,
       priority INTEGER NOT NULL DEFAULT 0,
       item_id  INTEGER NOT NULL,
       CONSTRAINT item_image_item_id_fkey FOREIGN KEY (item_id)
       REFERENCES item (item_id) MATCH SIMPLE 
       ON DELETE CASCADE
);


CREATE TABLE item_search_index (
       index_id serial PRIMARY KEY,
       item_id integer,
       item_name VARCHAR (1024) NOT NULL,
       item_category VARCHAR (1024) NOT NULL,
       item_type VARCHAR NOT NULL default 'service',
       item_brand varchar(255),
       item_description text,
       item_available boolean NOT NULL DEFAULT FALSE,
       thumbnail VARCHAR(255),
       company_id integer,
       company_name VARCHAR (1024) NOT NULL,
       company_postal_code varchar(10) NOT NULL default '',
       company_address_id integer,
       city_id integer,
       location geography(point) ,
       district_id integer,
       state_id integer,
       company_phone1 varchar(20),
       popularity integer not null default 0,
       last_updated timestamp not null DEFAULT NOW(),
       url_slug varchar(255),
       tag_index text,
       search_index text,
       search_vector tsvector,
       CONSTRAINT index_item_id_fkey FOREIGN KEY (item_id)
       REFERENCES item (item_id) MATCH SIMPLE 
       ON DELETE CASCADE,
       CONSTRAINT index_company_id_fkey FOREIGN KEY (company_id)
       REFERENCES company (company_id) MATCH SIMPLE 
       ON DELETE CASCADE,
       CONSTRAINT index_company_address_id_fkey FOREIGN KEY (company_address_id)
       REFERENCES company_address (company_address_id) MATCH SIMPLE 
       ON DELETE CASCADE    
);




create table analytics(
       analytics_id serial primary key,
       item_id integer,
       ip varchar(64),
       phone varchar(20),
       referrer varchar(1024),
       date timestamp not null DEFAULT NOW()
);

create table query_log(
       query_log_id serial primary key,
       query varchar(1024),
       referer varchar(1024),
       user_agent varchar(1024),
       uri text,
       ip varchar(64),
       phone varchar(20),
       date timestamp not null DEFAULT NOW()
);

create table search_log(
       search_log_id serial primary key,
       query varchar(256),
       item_id integer,
       ip varchar(64),
       phone varchar(20),
       city_id integer, 
       date timestamp not null DEFAULT NOW(),
       CONSTRAINT search_log_item_id_fkey FOREIGN KEY (item_id)
       REFERENCES item (item_id) MATCH SIMPLE 
       ON DELETE CASCADE
);

create table report (
       report_id serial primary key,
       date  TIMESTAMP,
       signups_total integer,
       signups_today integer,
       companies_total integer,
       products_total integer,
       free_listing integer,
       paid_listing integer,
       payment_today integer,
       payments_total integer
);


create table tag (
       tag_id serial primary key,
       tag_name varchar(256) not null unique,
       tag_description text not null default ''
);

create table item_tag(
       item_tag_id serial primary key,
       item_id integer,
       tag_id integer,
       unique (item_id, tag_id),
       CONSTRAINT tag_id_fkey FOREIGN KEY (tag_id)
       REFERENCES tag (tag_id) MATCH SIMPLE 
       ON DELETE CASCADE,
       CONSTRAINT item_id_fkey FOREIGN KEY (item_id)
       REFERENCES item (item_id) MATCH SIMPLE 
       ON DELETE CASCADE
);

create table job (
       job_id serial primary key,
       job_title varchar(255) not null,
       job_description text,
       city_id integer,
       job_contact_number1 varchar(20),
       job_contact_number2 varchar(20),
       job_contact_number3 varchar(20),
       job_contact_email varchar(256),
       job_contact_address text,
       job_contact_pincode varchar(10),
       company_id integer,
       CONSTRAINT job_city_id_fkey FOREIGN KEY (city_id)
       REFERENCES city (city_id) MATCH SIMPLE 
       ON DELETE CASCADE,
       CONSTRAINT job_company_id_fkey FOREIGN KEY (company_id)
       REFERENCES company (company_id) MATCH SIMPLE 
       ON DELETE CASCADE
);


create table job_profile(
       job_profile_id serial primary key,
       full_name varchar(255),
       phone varchar(20),
       email varchar(255),
       pincode varchar(10),
       city_id integer,
       district_id integer,
       state_id integer,
       address text,
       user_account_id integer,
       active boolean not null default true,
       CONSTRAINT job_profile_user_account_id_fkey FOREIGN KEY (user_account_id)
       REFERENCES user_account (user_account_id) MATCH SIMPLE 
       ON DELETE CASCADE
);


create table settings (
       settings_id serial primary key,
       settings_name varchar(255),
       settings_value varchar(255)
);

insert into settings (settings_name, settings_value) values ('country','India');



create table rp_order(
       payment_id serial primary key,
       order_id varchar(255),
       company_id int,
       amount decimal(10,3),
       order_time timestamp not null DEFAULT NOW()
);


create table login_trials(
       login_trials_id serial primary key,
       ip varchar(255),
       trial_time TIMESTAMP NOT NULL DEFAULT NOW()
);




create table search_user(
       search_user_id serial primary key,
       phone varchar(20) UNIQUE NOT NULL,
       otp varchar(6) NOT NULL,
       otp_key varchar(64) NOT NULL,
       search_auth varchar(128) NOT NULL UNIQUE
      
);

create table applog (
       applog_id serial primary key,
       action varchar(20),
       ip varchar(128),
       time timestamp not null default now(),
       log text,
       raw text
);



create table page (
    page_id  serial primary key,
    page_slug varchar(1024),
    page_file varchar(255),
    base_page_file varchar(255) not null default 'base.html'
);

insert into page (page_slug,page_file,base_page_file) values ('components','components.html','base.html');
insert into page (page_slug,page_file,base_page_file) values ('login','login.html','base.html');
insert into page (page_slug,page_file,base_page_file) values ('register','register.html','base.html');
insert into page (page_slug,page_file,base_page_file) values ('feed','feed.html','base.html');

insert into page (page_slug,page_file,base_page_file) values ('dashboard','dashboard.html','base.html');
insert into page (page_slug,page_file,base_page_file) values ('profile','profile.html','base.html');