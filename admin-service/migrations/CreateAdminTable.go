package migrations

import (
	"github.com/jmoiron/sqlx"
)

func CreateAdminTable(tx *sqlx.Tx) error {
	tx.MustExec("CREATE TABLE `admin` (`username` varchar(255) NOT NULL,`notes` varchar(255), PRIMARY KEY(`username`))")
	return nil
}
