package main

const databaseDSN string = "db.sqlite?__journal_mode=WAL&_busy_timeout=5000"

const appTitle string = "Base Go SQLite App"

const databaseSchema string = `
	CREATE TABLE IF NOT EXISTS Users 
	(
		ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		FullName TEXT, 
		Email TEXT, 
		Password TEXT,   
		Active TEXT
	);

	CREATE TABLE IF NOT EXISTS Services 
	(
		ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		Title TEXT, 
		Description TEXT,  
		Duration INTEGER,  
		Price NUMERIC, 
		Active TEXT
	);

	CREATE TABLE IF NOT EXISTS Customers 
	(
		ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		FullName TEXT,
		Phone TEXT, 
		Email TEXT, 
		Password TEXT, 
		Address TEXT,
		Created TEXT
	);

	CREATE TABLE IF NOT EXISTS Orders 
	(
		ID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		DateTime TEXT,
		Details TEXT
	);
`

const databaseData string = `
	INSERT INTO Users (FullName, Email, Password, Active) VALUES ("Super Admin", "admin@example.com", "21232f297a57a5a743894a0e4a801fc3", "Y");

	INSERT INTO Services (Title, Description, Duration, Price, Active)
		VALUES("Tire Rotation", "Take off current tires and put on alternative tires.", 30, 59.99, "Y");

	INSERT INTO Services (Title, Description, Duration, Price, Active)
		VALUES("Tire Rotation with replacement", "Take off current tires and put on alternative tires.", 60, 99.99, "Y");
`
