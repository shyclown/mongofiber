package database

import (
	"embed"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"log"
)

//go:embed migrations/*.sql
var fs embed.FS

func CheckMigrationVersion(m migrate.Migrate) {
	ver, dirty, err := m.Version()
	if err != nil {
		fmt.Println("Error getting migration version:", err)
	}
	fmt.Println("Migration version:", ver)
	fmt.Println("Migration dirty:", dirty)
}

func RunMigration() {

	d, err := iofs.New(fs, "migrations")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Running migration")

	driver, err := mysql.WithInstance(DB, &mysql.Config{})

	if err != nil {
		fmt.Println("Failed to instantiate driver: ", err)
	}

	m, err := migrate.NewWithInstance("iofs",
		d,
		"fiber_local_01",
		driver,
	)
	CheckMigrationVersion(*m)

	if err != nil {
		fmt.Println("Failed to instantiate migration: ", err)
	}

	err = m.Steps(2)
	if err != nil {
		fmt.Println("Steps: ", err)
	}
	err = m.Up()
	if err != nil {
		fmt.Println("Up: ", err)
	}
}
