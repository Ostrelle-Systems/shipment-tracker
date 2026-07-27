package database

import "fmt"

// Ping pings the database to verify the connection.
func (d *GormDatabase) Ping() error {
	sqldb, err := d.DB.DB()
	if err != nil {
		return fmt.Errorf("cannot access underlying sql.DB: %w", err)
	}
	if err := sqldb.Ping(); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}
	return nil
}
