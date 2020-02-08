package migrations

import "github.com/jmoiron/sqlx"

func CreateCustomCommandTables(tx *sqlx.Tx) error {
	_, err := tx.Exec("CREATE TABLE `custom_commands` (`name` varchar(255) NOT NULL, `proc` varchar(255) NOT NULL, `description` varchar(255) NOT NULL, PRIMARY KEY(`name`,`proc`))")
	if err != nil {
		return err
	}
	_, err = tx.Exec("CREATE TABLE `custom_command_channels` (`name` varchar(255) NOT NULL, `channel` varchar(255) NOT NULL, `enabled` TINYINT NOT NULL DEFAULT 1, `restricted` TINYINT NOT NULL DEFAULT 0, PRIMARY KEY(`name`,`channel`))")
	return err
}
