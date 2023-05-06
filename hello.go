package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type users struct {
	name     string
	password string
	email    string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/Account", AccountManagerHandler)
	http.HandleFunc("/Uaccount", UnsafeAccountManagerHendler)
	http.HandleFunc("/LogIn", LogIn)
	http.HandleFunc("/SignIn", SignIn)
	http.HandleFunc("/UnsafeLogIn", UnsafeLogIn)
	http.HandleFunc("/UnsafeSignIn", UnsafeSignIn)

	http.HandleFunc("/Application", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.gtpl")
	t.Execute(w, nil)
}

func AccountManagerHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("Account.gtpl")
	if err != nil {
		log.Fatalf("Error parsing template file: %v", err)
	}

	if err := tpl.Execute(w, nil); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("Application.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()

		// Get the username and password from the form data
		username := r.Form.Get("Login_username")
		password := r.Form.Get("Login_password")
		email := "fahrettin_basturk@hotmail.com"

		// it is SQL ATTACK.
		if SqlAttackChecker(username, password, email) {
			log.Printf("SQL injection attack detected!")
			http.Redirect(w, r, "/error", http.StatusSeeOther)
			return
		} else {
			// No attack detected
			log.Printf("SQL injection attack not detected.")
		}

		// Query the database to check if the user exists and the password is correct
		db, err := databaseConnection()
		if err != nil {
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
			return
		}

		ok, err := loginUser(db, username, password)
		if err != nil {
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
			return
		}

		data := struct {
			Name     string
			Password string
			Email    string
		}{
			Name:     username,
			Password: password,
			Email:    "",
		}

		db.Close()

		// If the user exists and the password is correct, redirect to the home page
		if ok {
			t, _ := template.ParseFiles("Application.gtpl")
			t.Execute(w, data)
			return
		} else {
			t, _ := template.ParseFiles("Account.gtpl")
			t.Execute(w, nil)
			return
		}

		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("Application.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()

		//IP := r.RemoteAddr

		// logic part of Sign in
		username := r.Form.Get("SignIn_username")
		password := r.Form.Get("SignIn_password")
		email := r.Form.Get("SignIn_email")

		// it is SQL ATTACK.
		if SqlAttackChecker(username, password, email) {
			log.Printf("SQL injection attack detected!")
			http.Redirect(w, r, "/error", http.StatusSeeOther)
			return
		} else {
			// No attack detected
			log.Printf("No attack detected.")
		}

		data := struct {
			name     string
			password string
			email    string
		}{
			name:     username,
			password: password,
			email:    email,
		}

		db, err := databaseConnection()
		if err != nil {
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
			return
		}

		if IsItTakenDatabase(db, data) {
			log.Printf("bu kullanıcı adı alınmamış")
			err = insertUserDatabase(db, data)
			if err != nil {
				log.Printf("Insert user failed with error %s", err)
			} else {
				log.Printf("bu kullanıcı alındı")
			}
		} else {
			log.Printf("bu kullanıcı adı alınmış")
		}
		db.Close()

		t, _ := template.ParseFiles("Application.gtpl")
		t.Execute(w, data)
	}
}

func IsItTakenDatabase(db *sql.DB, data users) bool {
	var count int

	errrr := db.QueryRow("SELECT count(*) FROM users WHERE username = ?", data.name).Scan(&count)

	fmt.Println(errrr)

	if count > 0 {
		fmt.Println("This account cannot taken from u because there is an account for that username.", http.StatusBadRequest)
		return false
	} else {
		fmt.Println("This account could be taken from u")
		return true
	}

}

func insertUserDatabase(db *sql.DB, p users) error {
	query := "INSERT INTO users(username, password, email) VALUES (?,?,?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	// Generate a hash of the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error %s when hashing password", err)
		return err
	}

	res, err := stmt.ExecContext(ctx, p.name, hashedPassword, p.email)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d products created ", rows)
	return nil
}

func databaseConnection() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// MySQL bağlantısı için gerekli değişkenleri .env dosyasından okuyun
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Database bağlantısını kontrol et
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}

func loginUser(db *sql.DB, username string, password string) (bool, error) {
	// Query to retrieve user with the given username
	query := "SELECT id, password FROM users WHERE username=?"

	// Query the database
	row := db.QueryRow(query, username)

	// Check if the query returns a result
	var id int
	var hashedPassword []byte
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			// If no rows were returned, then the user does not exist
			return false, nil
		}
		// If an error occurred while executing the query, return the error
		return false, err
	}

	// Compare the hashed password with the user input password
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		// If the passwords don't match, return false
		return false, nil
	}

	// If the passwords match, return true
	return true, nil
}

func SqlAttackChecker(username, password, email string) bool {
	// Regular expression to check for SQL injection attacks
	// This is just a simple example and should be improved for production use
	sqlRegex := regexp.MustCompile(`(SELECT|INSERT|UPDATE|DELETE|DROP|CREATE|ALTER|TRUNCATE|EXEC|UNION|OR|AND|\*|INTO|FROM|WHERE|LIKE|HAVING|GROUP BY|ORDER BY|--)`)

	// Additional SQL injection keywords
	additionalSqlRegex := regexp.MustCompile(`(BENCHMARK|CHAR|CHARINDEX|CAST|CONVERT|DECLARE|EXECUTE|FETCH|NCHAR|OCHAR|PRINT|PROCEDURE|REPLACE|SCRIPT|TABLE|TOP|TRANSACTION|WAITFOR|XACT_ABORT)`)

	// Tautology examples (always true)
	tautologyRegex := regexp.MustCompile(`[^\w](1=1|0=0|TRUE|FALSE)`)

	// Email regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	username = strings.ToUpper(username)
	password = strings.ToUpper(password)

	if sqlRegex.MatchString(username) || sqlRegex.MatchString(password) || additionalSqlRegex.MatchString(username) || additionalSqlRegex.MatchString(password) || tautologyRegex.MatchString(username) || tautologyRegex.MatchString(password) || !emailRegex.MatchString(email) {
		// SQL injection attack detected
		return true
	}

	// No attack detected
	return false
}

func UnsafeAccountManagerHendler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("Uaccount.gtpl")
	if err != nil {
		log.Fatalf("Error parsing template file: %v", err)
	}

	if err := tpl.Execute(w, nil); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func UnsafeLogIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("Application.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {

		r.ParseForm()

		// Get the username and password from the form data
		username := r.Form.Get("Login_username")
		password := r.Form.Get("Login_password")

		// Query the database to check if the user exists and the password is correct
		db, err := databaseConnection()
		if err != nil {
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
			return
		}

		ok, err := loginUser(db, username, password)
		if err != nil {
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
			return
		}

		data := struct {
			Name     string
			Password string
			Email    string
		}{
			Name:     username,
			Password: password,
			Email:    "",
		}

		db.Close()

		// If the user exists and the password is correct, redirect to the home page
		if ok {
			t, _ := template.ParseFiles("Application.gtpl")
			t.Execute(w, data)
			return
		} else {
			t, _ := template.ParseFiles("Uaccount.gtpl")
			t.Execute(w, nil)
			return
		}

		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}
}

func UnsafeSignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("Application.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()

		// logic part of Sign in
		username := r.Form.Get("SignIn_username")
		password := r.Form.Get("SignIn_password")
		email := r.Form.Get("SignIn_email")

		data := struct {
			name     string
			password string
			email    string
		}{
			name:     username,
			password: password,
			email:    email,
		}

		db, err := databaseConnection()
		if err != nil {
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
			return
		}

		if IsItTakenDatabase(db, data) {
			log.Printf("bu kullanıcı adı alınmamış")
			err = insertUserDatabase(db, data)
			if err != nil {
				log.Printf("Insert user failed with error %s", err)
			} else {
				log.Printf("bu kullanıcı alındı")
			}
		} else {
			log.Printf("bu kullanıcı adı alınmış")
		}
		db.Close()

		t, _ := template.ParseFiles("Application.gtpl")
		t.Execute(w, data)
	}
}

/*

	SQL:
		-- Tabloyu sil
	DROP TABLE IF EXISTS users;

	-- Tabloyu yeniden oluştur
	CREATE TABLE users (
		id INT PRIMARY KEY AUTO_INCREMENT,
		username VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	);


*/
