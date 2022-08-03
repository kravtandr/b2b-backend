
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
    name        TEXT        NOT NULL,
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
    price       int
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
    description     TEXT,
    legal_name      TEXT        NOT NULL,
    itn             TEXT        NOT NULL,
    psrn            TEXT        DEFAULT '',
    address         TEXT        DEFAULT '',
    legal_address   TEXT        DEFAULT '',
    email           TEXT        ,
    phone           TEXT        DEFAULT '',
    link            TEXT        DEFAULT '',
    activity        TEXT        DEFAULT '',
    owner_id        INT         NOT NULL,
    rating          INT         DEFAULT 0,
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

INSERT INTO groupaccessrights(add, edit, del) VALUES(true, true, true);
INSERT INTO groups(name, access_rights) VALUES('Владелец', 1);

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


