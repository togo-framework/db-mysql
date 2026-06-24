// Package dbmysql is togo's MySQL database driver plugin.
//
// It registers the go-sql-driver/mysql database/sql driver (driver name "mysql")
// so the kernel's SQL ORM can talk to MySQL — without the app importing the raw
// driver. Use it by setting DB_DRIVER=mysql + DATABASE_URL and installing the
// plugin (`togo new --db mysql`, or `togo install togo-framework/db-mysql`). The
// kernel ships the mysql SQL dialect; this plugin only supplies the driver.
package dbmysql

import (
	_ "github.com/go-sql-driver/mysql" // registers the "mysql" database/sql driver

	"github.com/togo-framework/togo"
)

func init() {
	// Marker provider; the blank import registered the driver with database/sql at
	// import time and togo.DialectFor knows the mysql dialect, so boot is a no-op.
	togo.RegisterProviderFunc("db-mysql", togo.PriorityService, func(*togo.Kernel) error { return nil })
}
