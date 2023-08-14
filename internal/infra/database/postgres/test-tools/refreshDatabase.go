package testtools

import (
	"database/sql"
	"fmt"
)

type SchemaInformation struct {
	TableName string
}

type DatabaseOperations struct {
	connection *sql.DB
}

func NewTestDatabaseOperations(connection *sql.DB) *DatabaseOperations {
	return &DatabaseOperations{
		connection: connection,
	}
}

func (t *DatabaseOperations) RefreshDatabase() {
	t.truncateTables(t.getTables())
}

func (t *DatabaseOperations) getTables() []SchemaInformation {
	var schemaInformation []SchemaInformation

	rows, err := t.connection.Query("select table_name from information_schema.tables where table_schema='public' AND table_name NOT IN ('_db_version', '_db_seeds')")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		schema := SchemaInformation{}
		_ = rows.Scan(&schema.TableName)
		schemaInformation = append(schemaInformation, schema)
	}

	_ = rows.Close()

	return schemaInformation
}

func (t *DatabaseOperations) truncateTables(SchemaInformation []SchemaInformation) {

	t.disableFk()

	for _, schema := range SchemaInformation {
		_, err := t.connection.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", schema.TableName))
		if err != nil {
			panic(err)
		}
	}

	t.enableFk()
}

func (t *DatabaseOperations) disableFk() {
	_, err := t.connection.Exec("SET session_replication_role = 'replica'")
	if err != nil {
		panic(err)
	}
}

func (t *DatabaseOperations) enableFk() {
	_, err := t.connection.Exec("SET session_replication_role = 'origin'")
	if err != nil {
		panic(err)
	}
}
