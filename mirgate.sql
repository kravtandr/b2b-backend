-- drop table subcategories ;
-- drop table productssubcategories ;
-- drop table orderproducts;
-- drop table categories ;
-- drop table orders ;
-- drop table companiesproducts ;
-- drop table products ;
-- drop table companiesUsers;
-- drop table companies  ;
-- drop table users ;
-- drop table groups ;
-- drop table groupaccessrights ;
-- drop function trigger_set_timestamp();
-- drop table OrderForm;


CREATE
    OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at
        = NOW();
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;


CREATE TABLE Categories
(
    id          SERIAL      NOT NULL PRIMARY KEY,
    name        TEXT        NOT NULL ,
    description TEXT
);

CREATE TABLE SubCategories
(
    id          SERIAL      NOT NULL PRIMARY KEY,
    name        TEXT        NOT NULL,
    description TEXT        ,
    category_id INT         NOT NULL,
    FOREIGN KEY (category_id) REFERENCES Categories (id) ON DELETE CASCADE
);

CREATE TABLE Products
(
    id          SERIAL      NOT NULL PRIMARY KEY,
    name        TEXT        NOT NULL,
    description TEXT        ,
    price       INT         ,
    photo       TEXT      
);

CREATE TABLE ProductsSubCategories
(
    id              SERIAL      NOT NULL PRIMARY KEY,
    category_id     INT         NOT NULL,
    product_id      INT         NOT NULL,
    FOREIGN KEY (category_id) REFERENCES Categories (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES Products (id) ON DELETE CASCADE
);

CREATE TABLE GroupAccessRights
(
    id          SERIAL      NOT NULL PRIMARY KEY,
    edit        BOOLEAN     ,
    del         BOOLEAN     ,
    add         BOOLEAN
);

CREATE TABLE Groups
(
    id              SERIAL      NOT NULL PRIMARY KEY,
    name            TEXT        NOT NULL,
    access_rights   INT,
    FOREIGN KEY (access_rights) REFERENCES GroupAccessRights (id) ON DELETE CASCADE
);


CREATE TABLE Users
(
    id              SERIAL      NOT NULL PRIMARY KEY,
    name            TEXT        NOT NULL,
    surname         TEXT        NOT NULL,
    patronymic      TEXT        NOT NULL,
    country         TEXT        NOT NULL,
    email           TEXT        NOT NULL UNIQUE,
    password        TEXT        NOT NULL,
    group_id        INT         DEFAULT 1,
    FOREIGN KEY (group_id) REFERENCES Groups (id) ON DELETE CASCADE
);



CREATE TABLE Companies
(
    id              SERIAL      NOT NULL PRIMARY KEY,
    name            TEXT        NOT NULL,
    description     TEXT        DEFAULT '',
    legal_name      TEXT        NOT NULL,
    itn             TEXT        NOT NULL,
    psrn            TEXT        DEFAULT '',
    address         TEXT        DEFAULT '',
    legal_address   TEXT        DEFAULT '',
    email           TEXT        NOT NULL,
    phone           TEXT        DEFAULT '',
    link            TEXT        DEFAULT '',
    activity        TEXT        DEFAULT '',
    owner_id        INT         NOT NULL,
    rating          INT         DEFAULT 0,
    verified        INT         DEFAULT 0,
    docs            TEXT[]      ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (owner_id) REFERENCES Users (id) ON DELETE CASCADE
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON Companies
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


CREATE TABLE CompaniesUsers
(
    id              SERIAL      NOT NULL PRIMARY KEY,
    post            TEXT        ,
    company_id      INT         NOT NULL,
    user_id         INT         NOT NULL,
    itn             TEXT         NOT NULL,
    FOREIGN KEY (company_id) REFERENCES Companies (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE
);

CREATE TABLE Orders
(
    id              SERIAL          NOT NULL PRIMARY KEY,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    total_price     INT NOT NULL,
    provider_id     INT NOT NULL,
    purchaser_id    INT NOT NULL,
    FOREIGN KEY (provider_id) REFERENCES Companies (id) ON DELETE CASCADE,
    FOREIGN KEY (purchaser_id) REFERENCES Companies (id) ON DELETE CASCADE
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON Orders
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE OrderProducts
(
    id              SERIAL      NOT NULL PRIMARY KEY,
    order_id        INT         NOT NULL,
    product_id      INT         NOT NULL,
    amount          INT         NOT NULL,
    FOREIGN KEY (order_id) REFERENCES Orders (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES Products (id) ON DELETE CASCADE
);

CREATE TABLE CompaniesProducts
(
    id              SERIAL      NOT NULL PRIMARY KEY,
    company_id      INT         NOT NULL,
    product_id      INT         NOT NULL,
    amount          INT         NOT NULL,
    FOREIGN KEY (company_id) REFERENCES Companies (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES Products (id) ON DELETE CASCADE
);

CREATE TABLE OrderForm
(
    id                SERIAL      NOT NULL PRIMARY KEY,
    role              boolean     NOT NULL,
    product_category  TEXT        NOT NULL,
    product_name      TEXT        DEFAULT 'empty',
    order_text        TEXT        DEFAULT 'empty',
    order_comments    TEXT        DEFAULT 'empty',
    fio               TEXT        NOT NULL,
    email             TEXT        NOT NULL,
    phone             TEXT        DEFAULT 'empty',
    company_name      TEXT        NOT NULL,
    itn               TEXT        NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE landing_request
(
    id                SERIAL      NOT NULL PRIMARY KEY,
    product_category  TEXT        NOT NULL,
    delivery_address  TEXT        NOT NULL,
    delivery_date     TEXT        NOT NULL,
    order_text        TEXT        DEFAULT 'empty',
    email             TEXT        NOT NULL,
    itn               TEXT        NOT NULL,
    phone             TEXT        DEFAULT 'empty',
    company_name      TEXT        DEFAULT 'empty',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON OrderForm
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE Cookies (
                         id SERIAL NOT NULL PRIMARY KEY,
                         hash TEXT NOT NULL,
                         user_id INT NOT NULL,
                         created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                         FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);


INSERT INTO groupaccessrights(add, edit, del) VALUES(true, true, true);
INSERT INTO groups(name, access_rights) VALUES('Владелец', 1);
INSERT INTO products(name, description, price, photo) VALUES('1_product', 'aaa', 1000, '/photo1');
INSERT INTO products(name, description, price, photo) VALUES('2_product', 'bbb', 1000, '/photo1');
COPY categories(name) FROM './export_base_categories.csv' DELIMITER ',' CSV HEADER;
-- INSERT INTO Users ("name", "surname", "patronymic", "email", "password", "country") VALUES ('Иван', 'Иванович','Иванов','ivan@mail.ru','password123','Россия');
-- INSERT INTO Users ("name", "surname", "patronymic", "email", "password", "country") VALUES ('Петр', 'Петрович','Петров','petr@mail.ru','password123','Россия');
-- INSERT INTO Users ("name", "surname", "patronymic", "email", "password", "country") VALUES ('Алексндр', 'Александров','Александрович','alex@mail.ru','password123','Россия');
-- INSERT INTO Companies ("name", "description", "legal_name", "itn", "psrn", "address","legal_address","email", "phone", "link", "activity", "owner_id", "rating"  ) VALUES ('Весна',
--                                                                                                                                                                            'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
--                                                                                                                                                                            'ООО Весна',
--                                                                                                                                                                            '7727563379', '1162225076155',
--                                                                                                                                                                            '656922, Алтайский край, г. Барнаул, ул. Попова, д. 185г, офис 21а',
--                                                                                                                                                                            '656922, Алтайский край, г. Барнаул, ул. Попова, д. 185г, офис 21а',
--                                                                                                                                                                            'tsvetiu@mail.ru',
--                                                                                                                                                                            '89031223451',
--                                                                                                                                                                            'tsvetiu.ru',
--                                                                                                                                                                            'Розничная торговля',
--                                                                                                                                                                            1,5);
-- INSERT INTO companiesusers ("post", "company_id", "user_id") VALUES ('CEO',1,1);
-- INSERT INTO companiesusers ("post", "company_id", "user_id") VALUES ('Manager',1,2);
-- INSERT INTO companiesusers ("post", "company_id", "user_id") VALUES ('Driver',1,3);
--
-- INSERT INTO Industry(title) VALUES('Машиностроение');
-- INSERT INTO Industry(title) VALUES('MetalProd');
-- INSERT INTO Industry(title) VALUES('OilProd');
-- INSERT INTO Category(title, industry_id) VALUES('category1', 1);
-- INSERT INTO Category(title, industry_id) VALUES('category2', 1);
-- INSERT INTO Category(title, industry_id) VALUES('category3', 2);
-- INSERT INTO Category(title, industry_id) VALUES('category4', 3);
--
-- INSERT INTO Company(email, password, name, legal_name, itn, psrn, adress, phone, link, category_id) VALUES('test2@mail.ru', 'password123','Ромашка','ООО Ромашка','7727563779','1056749631995','Москва','8(915)9999998','yandex.ru',1);
-- INSERT INTO Company(email, password, name, legal_name, itn, psrn, adress, phone, link, category_id)  VALUES('test3@mail.ru', 'password123','Пчелка','ООО Пчелка','7727563719','1057749331995','Москва','8(915)9999918','yandex.ru',1);
-- INSERT INTO Company(email, password, name, legal_name, itn, psrn, adress, phone, link, category_id)  VALUES('test4@mail.ru', 'password123','Пример','ООО Пример','7727563729','1057249631995','Москва','8(915)9999928','yandex.ru',1);
-- INSERT INTO Company(email, password, name, legal_name, itn, psrn, adress, phone, link, category_id)  VALUES('test5@mail.ru', 'password123','Тест','ООО Тест','7727533779','1057741631995','Москва','8(915)9999938','yandex.ru',2);
-- INSERT INTO Company(email, password, name, legal_name, itn, psrn, adress, phone, link, category_id)  VALUES('test6@mail.ru', 'password123','Последний','ООО Последний','7727463779','1057749631995','Москва','8(915)9399998','yandex.ru',3);


-- SELECT   cu.post, cu.user_id, u.name, u.surname, u.patronymic, u.email, u.country, u.group_id FROM companiesusers AS cu  JOIN Users u on u.id = cu.user_id WHERE cu.company_id = 1;

-- SELECT
--
-- COPY categories(name)
-- FROM '/var/lib/postgresql/backend/b2b-backend/export_base_categories.csv'
-- DELIMITER ','
-- CSV HEADER;
--
--
-- SELECT name, description
-- 	FROM categories
-- 	WHERE name ~ $1;
--
--
--
--
-- UPDATE Companies SET name = $2, description = $3, legal_name = $4, address = $5, legal_address= $6, phone = $7, link = $8, activity = $9 WHERE id = $1
--
-- UPDATE Users SET name = 'asd', surname = 'asd', patronymic = 'as', email = 'asd' ,password = 'asd' WHERE id = 28;
--


--COPY categories(name) FROM '/var/lib/postgresql/data/export_base_categories.csv' DELIMITER ',' CSV HEADER;
--COPY products(name, description, price, photo) FROM '/var/lib/postgresql/data/test_products.csv' DELIMITER ';' CSV HEADER;

