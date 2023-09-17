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
    description TEXT        DEFAULT 'empty',
    parent_id   INT
);


CREATE TABLE Products
(
    id          SERIAL      NOT NULL PRIMARY KEY,
    name        TEXT        NOT NULL,
    description TEXT        ,
    price       INT         NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE ProductsCategories
(
    id              SERIAL      NOT NULL PRIMARY KEY,
    category_id     INT         NOT NULL,
    product_id      INT         NOT NULL,
    FOREIGN KEY (category_id) REFERENCES Categories (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES Products (id) ON DELETE CASCADE
);

CREATE TABLE ProductPhotos
(
    id              SERIAL      NOT NULL PRIMARY KEY,
    product_id      INT         NOT NULL,
    photo_obj_name  Text         NOT NULL,
    FOREIGN KEY (product_id) REFERENCES Products (id) ON DELETE CASCADE
);

CREATE TABLE ProductDocuments
(
    id              SERIAL      NOT NULL PRIMARY KEY,
    product_id      INT         NOT NULL,
    document_obj_name  Text         NOT NULL,
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
    addedBy      INT         NOT NULL,
    amount          INT         NOT NULL,
    pay_way         TEXT        DEFAULT 'empty',
    delivery_way    TEXT        DEFAULT 'empty',
    adress          TEXT        DEFAULT 'empty',
    FOREIGN KEY (addedBy) REFERENCES Users (id) ON DELETE CASCADE,
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

CREATE TABLE Chats 
(
    id                        SERIAL                    NOT NULL PRIMARY KEY,
    name                      TEXT                      DEFAULT 'New Chat',
    creator_id                INT                       NOT NULL,
    product_id                INT                       NOT NULL,
    status                    TEXT                      DEFAULT 'Обсуждение',
    FOREIGN KEY (creator_id)   REFERENCES Users (id)    ON DELETE CASCADE,
    FOREIGN KEY (product_id)  REFERENCES Products (id)  ON DELETE CASCADE,
    created_at                TIMESTAMPTZ               NOT NULL DEFAULT NOW(),
    updated_at                TIMESTAMPTZ               NOT NULL DEFAULT NOW()
);

CREATE TABLE Msgs 
(
    id                      SERIAL                  NOT NULL PRIMARY KEY,
    chat_id                 INT                     NOT NULL,
    sender_id               INT                       NOT NULL,
    receiver_id             INT                       NOT NULL,
    checked                 BOOLEAN                 NOT NULL DEFAULT FALSE,
    text                    TEXT                    NOT NULL,
    type                    TEXT                    NOT NULL DEFAULT 'regular msg',
    created_at              TIMESTAMPTZ             NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ             NOT NULL DEFAULT NOW(),
    FOREIGN KEY (chat_id)   REFERENCES Chats (id)   ON DELETE CASCADE,
    FOREIGN KEY (sender_id)   REFERENCES Users (id)   ON DELETE CASCADE,
    FOREIGN KEY (receiver_id)   REFERENCES Users (id)   ON DELETE CASCADE
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


-- INSERT INTO products(name, description, price) VALUES('1_product', 'test_desc1', 1001);
-- INSERT INTO products(name, description, price) VALUES('2_product', 'test_desc2', 1002);


INSERT INTO groupaccessrights(add, edit, del) VALUES(true, true, true);
INSERT INTO groups(name, access_rights) VALUES('Владелец', 1);


