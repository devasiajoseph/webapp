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