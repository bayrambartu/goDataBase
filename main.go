package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1234567@/demodb") // connection strgin ve baglanti
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // baglantiyi kapatmak icin

	var (
		// ID        int
		Username  string
		Email     string
		Password  string
		FirstName string
		LastName  string
		BirthDate string
		IsActive  bool
	)

	// baglanti yap ,hata varsa yonet, baglantiyi kapat..

	/*createStatement := "`users`(`ID` int(11) NOT NULL AUTO_INCREMENT,`Username` varchar(45) NOT NULL,`Email` varchar(45) NOT NULL,`Password` varchar(45) NOT NULL,`FirstName` varchar(45) NOT NULL,`LastName`varchar(45) NOT NULL,`BirthDate varchar(45) DEFAULT NULL,`IsActive` tinyint(1) DEFAULT NULL,PRIMARY KEY (`ID`),UNIQUE KEY `ID_UNIQUE`(`ID`)) ENGINE = InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET_utf8;"
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS" + createStatement)
	if err != nil {
		log.Fatal(err)
	}

	// Veri Ekleme
	res, err := db.Exec("INSERT INTO users(Username, Email,Password,FirstName,LastName,BirthDate,IsActive) VALUES('Denem1','12345!','deneme@dene.com','Bartu','Kurnaz','2003.1.1',1)")

	if err != nil {
		log.Fatal(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Inserted %d rows", rowCount)

	lastID, err1 := res.LastInsertId()
	if err != nil {
		log.Fatal(err1)
	}
	log.Printf("Inserted ID : %d", lastID)

	var (
		ID        int
		Username  string
		Email     string
		Password  string
		FirstName string
		LastName  string
		BirthDate string
		IsActive  bool
	)

	// eklenen kayıtları getir.
	rows, err := db.Query("SELECTED * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	// rows colums
	for rows.Next() {
		err = rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Bulunan satır içeriği: %q", strconv.Itoa(ID)+" "+Username+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))
	}
	// alternatif
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()
	*/

	// ............

	/* rows, errQ := db.Query("SELECT * FROM users WHERE ID = ?", 12)
	if errQ != nil {
		log.Fatal(errQ)
	}
	defer rows.Close()

	for rows.Next() {
		errQ = rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Bulunan satır içeriği: %q", strconv.Itoa(ID)+" "+Username+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))

	}

	errQ = rows.Err()
	if errQ != nil {
		log.Fatal(errQ)
	}
	*/
	// ..........................

	/*
		// var x string
		err := db.QueryRow("SELECT * FROM users limit 1").Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
		if err != nil {
			if err == sql.ErrNoRows { // kayit satir yok hatası
				// kayıt yoksa , yapıalcak işler..

			} else {
				log.Fatal(err)
			}

		}
		// log.Println("Bir satir bulundu:", Username)
		log.Printf("Bulunan satır içeriği: %q", strconv.Itoa(ID)+" "+Username+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))
	*/
	//----------------------

	// Multiple Active Result Set : MARS
	// _, err := db.Exec("DELETE FROM xTable1;DELETE FROM xTable2")

	// Preparing Queries

	/*
		stmt, errQ := db.Prepare("SELECT * FROM users WHERE ID = ?")
		if errQ != nil {
			log.Fatal(errQ)
		}
		defer stmt.Close()
		rows, errX := stmt.Query(12)
		if errX != nil {
			log.Fatal(errX)
		}
		defer rows.Close()
		for rows.Next() {
			scanErr = rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
			if scanErr != nil {
				log.Fatal(scanErr)
			}
			log.Printf("Bulunan satır içeriği: %q", strconv.Itoa(ID)+" "+Username+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))
		}
	*/

	// Preparing Query - Single rows

	/*
		stmt, errQ := db.Prepare("SELECT * FROM users WHERE ID = ?")
		if errQ != nil {
			log.Fatal(errQ)
		}
		errX := stmt.QueryRow(2).Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
		if errX != nil {
			log.Fatal(errX)
		}
		fmt.Println(FirstName + " " + LastName)
	*/

	// ...........................

	//Modifying Data
	/*
		res, _ := db.Exec("DELETE FROM users LIMIT 1")
		rowCount, _ := res.RowsAffected()
		fmt.Println(rowCount)
		lastID, _ := res.LastInsertId()
		fmt.Println(lastID)

	*/
	// Insert İşlemi

	/*
		stmt, err := db.Prepare("INSERT INTO users(Username, Email,Password,FirstName,LastName,BirthDate,IsActive) VALUES(?,?,?,?,?,?,?)")
		Username = "CO"
		Email = "abc@xyz.co"
		Password = "1234!"
		FirstName = "ABC"
		LastName = "XYZ"
		BirthDate = "2003.1.2"
		IsActive = true
		res, errStmt := stmt.Exec(Username, Email, Password, FirstName, LastName, BirthDate, IsActive)
		if errStmt != nil {
			log.Fatal(errStmt)
		}
		fmt.Println(res.LastInsertId())
	*/

	// Transaction

	/*
		tx, errTx := db.Begin()
		// ...
		_, errTx = db.Exec("Update..", args)
		// ....
		if errTx != nil {
			log.Fatal(errTx)
			tx.Rollback() // geri alma hata anında
		}
		//...
		errTx = tx.Commit()

	*/
}
