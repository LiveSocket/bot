package migrations

import "github.com/jmoiron/sqlx"

func AlterCustomCommands(tx *sqlx.Tx) error {
	_, err := tx.Exec("ALTER TABLE `custom_commands` ADD `channel` varchar(255) NOT NULL DEFAULT '' FIRST")
	if err != nil {
		return err
	}
	_, err = tx.Exec("ALTER TABLE `custom_commands` ADD `enabled` TINYINT NOT NULL DEFAULT 1")
	if err != nil {
		return err
	}
	_, err = tx.Exec("ALTER TABLE `custom_commands` ADD `restricted` TINYINT NOT NULL DEFAULT 0")
	if err != nil {
		return err
	}
	_, err = tx.Exec("ALTER TABLE `custom_commands` DROP PRIMARY KEY")
	if err != nil {
		return err
	}
	_, err = tx.Exec("ALTER TABLE `custom_commands` ADD CONSTRAINT `custom_commands_pk` PRIMARY KEY (`channel`,`name`)")
	if err != nil {
		return err
	}
	_, err = tx.Exec("DROP TABLE `custom_command_channels`")
	if err != nil {
		return err
	}
	return nil
}
