package db

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Migration A Migration function to run against the database
type Migration func(*sqlx.Tx) error

type DB struct {
	*sqlx.DB
}

type model struct {
	db         *sqlx.DB
	engine     string
	table      string
	user       string
	pword      string
	address    string
	dbName     string
	connection string
	migrations []Migration
	existing   []string
}

// Init initalize and return a database connection
func Init(table string, migrations ...Migration) (*DB, func()) {
	model := model{
		address:    os.Getenv("DB_ADDRESS"),
		dbName:     os.Getenv("DB_NAME"),
		user:       os.Getenv("DB_USER"),
		pword:      os.Getenv("DB_PWORD"),
		engine:     "mysql",
		table:      table,
		migrations: migrations,
	}
	model.connection = connectionString(model)
	model.db = open(model)
	if model.table != "" {
		createMigrationTable(model)
		model.existing = getExisting(model)
		migrate(model)
	}
	return &DB{DB: model.db}, func() {
		if err := model.db.Close(); err != nil {
			panic(err)
		}
	}
}

func (db *DB) Get(dest interface{}, query string, args ...interface{}) error {
	err := db.DB.Get(dest, query, args...)
	if err != nil {
		if strings.Contains(fmt.Sprint(err), "no rows") {
			dest = nil
			return nil
		}
	}
	return err
}

func (db *DB) Select(dest interface{}, query string, args ...interface{}) error {
	err := db.DB.Select(dest, query, args...)
	if err != nil {
		if strings.Contains(fmt.Sprint(err), "no rows") {
			dest = nil
			return nil
		}
	}
	return err
}

func open(model model) *sqlx.DB {
	println("Attempting DB Connection...")
	db, err := sqlx.Connect(model.engine, model.connection)
	retries := 0
	for err != nil && retries < 30 {
		log.Print(err)
		retries++
		time.Sleep(2 * time.Second)
		println("Retrying DB Connection...")
		db, err = sqlx.Connect(model.engine, model.connection)
	}
	if err != nil {
		panic(err)
	}
	return db
}

// migrate Runs the provided migrations against the database connection
func migrate(model model) {
	println("Running migrations...")

	for _, migration := range model.migrations {
		runMigration(model, migration)
	}
}

// Creates a MySql connection string for the model
func connectionString(model model) string {
	return model.user + ":" + model.pword + "@tcp(" + model.address + ")/" + model.dbName + "?charset=utf8mb4,utf8&parseTime=True"
}

func createMigrationTable(model model) {
	model.db.MustExec("CREATE TABLE IF NOT EXISTS `" + model.table + "` (`name` varchar(255), `timestamp` datetime, PRIMARY KEY(`name`))")
}

func getExisting(model model) (existing []string) {
	if err := model.db.Select(&existing, "SELECT `name` FROM `"+model.table+"`"); err != nil {
		panic(err)
	}
	return
}

func runMigration(model model, migration Migration) {
	name := getName(migration)
	if !canMigrate(model.existing, name) {
		return
	}
	println("Running migration " + name + "...")
	doMigration(model, name, migration)
}

func doMigration(model model, name string, migration Migration) {
	tx := model.db.MustBegin()
	err := migration(tx)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	saveMigration(model, name, tx)
	tx.Commit()
}

func saveMigration(model model, name string, tx *sqlx.Tx) {
	tx.MustExec("INSERT INTO `"+model.table+"` (`name`,`timestamp`) VALUES (?,CURRENT_TIMESTAMP)", name)
}

func getName(migration Migration) string {
	return runtime.FuncForPC(reflect.ValueOf(migration).Pointer()).Name()
}

// Checks if a migration has been run before, returns false if migration entry exists in migration table
func canMigrate(collection []string, name string) bool {
	for _, item := range collection {
		if item == name {
			return false
		}
	}
	return true
}
