-- PostgreSQL for ERD Generated

CREATE TABLE Users (
    ID INTEGER PRIMARY KEY,
    CreatedAt TIMESTAMP NOT NULL,
    UpdatedAt TIMESTAMP NOT NULL,
    Username VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL,
    Password VARCHAR(255) NOT NULL,
    Points INTEGER NOT NULL
);

CREATE TABLE Products (
    ID INTEGER PRIMARY KEY,
    CreatedAt TIMESTAMP NOT NULL,
    UpdatedAt TIMESTAMP NOT NULL,
    Category VARCHAR(255) NOT NULL,
    Name VARCHAR(255) NOT NULL,
    Description VARCHAR(255) NOT NULL,
    Price INTEGER NOT NULL,
    Stock INTEGER NOT NULL,
    Image VARCHAR(255) NOT NULL,
    CreatedBy VARCHAR(255) NOT NULL,
    UpdatedBy VARCHAR(255) NOT NULL,
    FOREIGN KEY (CreatedBy) REFERENCES Admins(Name),
    FOREIGN KEY (UpdatedBy) REFERENCES Admins(Name)
);

CREATE TABLE SerialNumbers (
    ID INTEGER PRIMARY KEY,
    CreatedAt TIMESTAMP NOT NULL,
    UpdatedAt TIMESTAMP NOT NULL,
    SerialNumber VARCHAR(255) NOT NULL,
    ProductID INTEGER NOT NULL,
    Status VARCHAR(255) NOT NULL,
    CreatedBy VARCHAR(255) NOT NULL,
    FOREIGN KEY (ProductID) REFERENCES Products(ID),
    FOREIGN KEY (CreatedBy) REFERENCES Admins(Name)
);

CREATE TABLE Transactions (
    ID INTEGER PRIMARY KEY,
    CreatedAt TIMESTAMP NOT NULL,
    UpdatedAt TIMESTAMP NOT NULL,
    UserID INTEGER NOT NULL,
    ProductID INTEGER NOT NULL,
    SerialNumber VARCHAR(255) NOT NULL,
    IdentifierNum VARCHAR(255) NOT NULL,
    TotalPrice INTEGER NOT NULL,
    Status VARCHAR(255) NOT NULL,
    LastUpdatedBy VARCHAR(255) NOT NULL,
    FOREIGN KEY (UserID) REFERENCES Users(ID),
    FOREIGN KEY (ProductID) REFERENCES Products(ID),
    FOREIGN KEY (LastUpdatedBy) REFERENCES Admins(Name)
);

CREATE TABLE Admins (
    ID INTEGER PRIMARY KEY,
    CreatedAt TIMESTAMP NOT NULL,
    UpdatedAt TIMESTAMP NOT NULL,
    Name VARCHAR(255) NOT NULL,
    Username VARCHAR(255) NOT NULL,
    Password VARCHAR(255) NOT NULL
);

INSERT INTO users(id, created_at, updated_at, role, username, email, password)
VALUES
    ('ini-string-panjang', '2020-12-01 00:00:00', '2020-12-01 00:00:00', 'admin', 'admin01', 'admin01@gmail.com', 'adminpass'),
    ('ini-string-panjang-user01', '2020-12-01 00:00:00', '2020-12-01 00:00:00', 'user', 'user01', 'user01@gmail.com', 'userpass'),
    ('ini-string-panjang-user02', '2020-12-01 00:00:00', '2020-12-01 00:00:00', 'user', 'user02', 'user02@gmail.com', 'userpass');

INSERT INTO transactions(id, created_at, updated_at, payment_method, user_id, product_id, serial_number, identifier_num, price, status)
VALUES
    ('ini-string-panjang-transaksi01', '2020-12-01 00:00:00', '2020-12-01 00:00:00', 'redeem', 'ini-string-panjang-user01', 'ini-string-panjang-produk', 234131423213, '081239684271', 10000, 'pending'),
    ('ini-string-panjang-transaksi02', '2020-12-01 00:00:00', '2020-12-01 00:00:00', 'redeem', 'ini-string-panjang-user02', 'ini-string-panjang-produk', 234131423213, '081239684271', 10000, 'pending'),
    ('ini-string-panjang-transaksi03', '2020-12-01 00:00:00', '2020-12-01 00:00:00', 'buy', 'ini-string-panjang-user01', 'ini-string-panjang-produk', 234131423213, '081239684271', 10000, 'pending');

INSERT INTO products(id, created_at, updated_at, category, name, description, price, stock, image)
VALUES
    ('ini-string-panjang-produk', '2020-12-01 00:00:00', '2020-12-01 00:00:00', 'pulsa', 'pulsa10k', 'ini pulsa dapet 10 ribu seharga 10k', 10000, 1, 'ini link image pulsa10k');

INSERT INTO serial_numbers(id, created_at, updated_at, product_id, serial, status)
VALUES
    ('ini-string-panjang-serial', '2020-12-01 00:00:00', '2020-12-01 00:00:00', 'ini-string-panjang-produk', 234131423213, 'available');

INSERT INTO points(id, user_id, points)
VALUES
    ('ini-string-panjang-point', 'ini-string-panjang-user01', 10000);


-- eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluMDEiLCJleHAiOjE2Njk4ODUwNDMsImVtYWlsIjoiYWRtaW4ifQ.4ENFOaeG_yyZzWbiWBcN67RHOPfelifSkpyvlQ8G948
-- eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXIwMUBnbWFpbC5jb20iLCJleHAiOjE2Njk4ODUwNDMsInVzZXJuYW1lIjoidXNlcjAxIn0.w8r4Q5eB_qDNKSPj-JF0mH_y7nA9gPxXJyNIfDGxPlU

-- By User
-- {
--     "paymentMethod" : "redeem",
--     "productID" : "ini-string-panjang-produk",
--     "identifierNum" : "08765432100"
-- }

-- By Admin
-- {
--     "paymentMethod" : "buy",
--     "userID" : "dummy-from-admin",
--     "productID" : "ini-string-panjang-produk",
--     "identifierNum" : "08765432100"
-- }


-- http://localhost:8080/api/auth/transactions/c267a4d5-09a4-4f17-a7c0-75a6e0616522
-- {
--     "status" : "succes"
-- }