package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection(cfg *Config) *gorm.DB {
	return getPostgreSQLConn(cfg)
}

func getPostgreSQLConn(cfg *Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s database=%s port=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Database,
		cfg.Database.Port,
	)

	env := os.Getenv("GO_ENV")

	// qa env uses the same secret as dev env because the db instance is shared
	if env == "qa" {
		env = "dev"
	}

	if env != "local" {
		dbRootCert := "/secrets/server-ca/server-ca.pem"
		dbCert := "/secrets/client-cert/client-cert.pem"
		dbKey := "/secrets/client-key/client-key.pem"

		dsn += fmt.Sprintf(" sslmode=require sslrootcert=%s sslcert=%s sslkey=%s",
			dbRootCert, dbCert, dbKey)
	}

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	return conn
}
