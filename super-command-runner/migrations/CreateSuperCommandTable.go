package migrations

import "github.com/jmoiron/sqlx"

// CreateSuperCommandTable Creates the super_commands table in the database
func CreateSuperCommandTable(tx *sqlx.Tx) error {
	_, err := tx.Exec("CREATE TABLE `super_commands` (`name` varchar(255) NOT NULL, `proc` varchar(255) NOT NULL, PRIMARY KEY(`name`))")
	return err
}
