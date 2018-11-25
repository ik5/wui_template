package db

import (
	"fmt"

	"github.com/ik5/wui_template/logging"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// DB is the database connection
var DB *sqlx.DB

// Init initialize the DB variable
func Init() error {
	dbName := viper.GetString("dbname")
	dbPort := viper.GetInt("dbport")
	dbUser := viper.GetString("dbuser")
	dbPassword := viper.GetString("dbpassword")
	dbHost := viper.GetString("dbhost")
	connStr := fmt.Sprintf("user='%s' password='%s' dbname='%s' host='%s' port=%d sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	var err error
	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		return err
	}
	return nil
}

// LogError override the string
func LogError(e *pq.Error) {
	logging.Logger.WithFields(logrus.Fields{
		"Severity":         e.Severity,
		"Code":             e.Code,
		"Message":          e.Message,
		"Detail":           e.Detail,
		"Hint":             e.Hint,
		"Position":         e.Position,
		"InternalPosition": e.InternalPosition,
		"InternalQuery":    e.InternalQuery,
		"Where":            e.Where,
		"Schema":           e.Schema,
		"Table":            e.Table,
		"Column":           e.Column,
		"DataTypeName":     e.DataTypeName,
		"Constraint":       e.Constraint,
		"File":             e.File,
		"Line":             e.Line,
		"Routine":          e.Routine,
	}).Errorf("DB Error: %s", e)
}
