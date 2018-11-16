package db_mysql

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"time"
)

type MysqlDB struct {
	conn *sql.DB
}

type MySQLConfig struct {
	// Optional.
	Username, Password string
	// Host of the MySQL instance.  If set, UnixSocket should be unset.
	Host string
	// Port of the MySQL instance.  If set, UnixSocket should be unset.
	Port int
	// UnixSocket is the filepath to a unix socket. If set, Host and Port should be unset.
	UnixSocket string
}

type MqttMsg struct {
	Id                                       int32
	Timestamp, Batterysn   string 
    Slotnum, Temprature, Voltage, Current, Elapsed int32
}

const DATABASE_NAME = "battery_ageing"

// dataStoreName returns a connection string suitable for sql.Open.
func (c MySQLConfig) dataStoreName(databaseName string) string {
	var cred string
	// [username[:password]@]
	if c.Username != "" {
		cred = c.Username
		if c.Password != "" {
			cred = cred + ":" + c.Password
		}
		cred = cred + "@"
	}

	if c.UnixSocket != "" {
		return fmt.Sprintf("%sunix(%s)/%s", cred, c.UnixSocket, databaseName)
	}

	return fmt.Sprintf("%stcp(%s:%d)/%s", cred, c.Host, c.Port, databaseName)
}

// ensureTableExists checks the table exists. If not, it creates it.
func (config MySQLConfig) ensureTableExists() error {
	conn, err := sql.Open("mysql", config.dataStoreName(""))
	if err != nil {
		fmt.Println("could not get connection")
		return fmt.Errorf("mysql: could not get a connection: %v", err)
	}

	defer conn.Close()

	// Check the connection.
	if conn.Ping() == driver.ErrBadConn {
		fmt.Println("could not connect to database")
		return fmt.Errorf("mysql: could not connect to the database. " +
			"could be bad address, or this address is not whitelisted for access.")
	}

	table_name := "battery_ageing" + time.Now().Local().Format("20060102")

	if _, err := conn.Exec("USE " + DATABASE_NAME); err != nil {
		// MySQL error 1049 is "database does not exist"
		if mErr, ok := err.(*mysql.MySQLError); ok && mErr.Number == 1049 {
			return createTable(conn, DATABASE_NAME, table_name)
		}
	}

	if _, err := conn.Exec("DESCRIBE " + table_name); err != nil {
		// MySQL error 1146 is "table does not exist"
		if mErr, ok := err.(*mysql.MySQLError); ok && mErr.Number == 1146 {
			return createTable(conn, DATABASE_NAME, table_name)
		}
		// Unknown error.
		return fmt.Errorf("mysql: could not connect to the database: %v", err)
	}

	return nil
}

// createTable creates the table, and if necessary, the database.
func createTable(conn *sql.DB, database_name, table_name string) error {

	createTableStatements := make([]string, 3)
	createTableStatements[0] = fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v %v", database_name, "DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';")
	createTableStatements[1] = fmt.Sprintf("USE %v", database_name)
	createTableStatements[2] = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v %v", table_name,
		"(id INT UNSIGNED NOT NULL AUTO_INCREMENT, timestamp TEXT NULL, batterysn TEXT NULL, slotnum INT NULL, temprature INT NULL, voltage INT NULL, current INT NULL, elapsed INT NULL, PRIMARY KEY (id))")

	for _, stmt := range createTableStatements {
		fmt.Println("sql cmd : ", stmt)
		_, err := conn.Exec(stmt)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

// execAffectingOneRow executes a given statement, expecting one row to be affected.
func execAffectingOneRow(stmt *sql.Stmt, args ...interface{}) (sql.Result, error) {
	r, err := stmt.Exec(args...)
	if err != nil {
		return r, fmt.Errorf("mysql: could not execute statement: %v", err)
	}
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return r, fmt.Errorf("mysql: could not get rows affected: %v", err)
	} else if rowsAffected != 1 {
		return r, fmt.Errorf("mysql: expected 1 row affected, got %d", rowsAffected)
	}
	return r, nil
}

// NewMySQLDB creates a new BookDatabase backed by a given MySQL server.
func NewMySQLDB(config MySQLConfig) (*MysqlDB, error) {
	// Check database and table exists. If not, create it.
	if err := config.ensureTableExists(); err != nil {
		return nil, err
	}

	conn, err := sql.Open("mysql", config.dataStoreName(DATABASE_NAME))
	if err != nil {
		return nil, fmt.Errorf("mysql: could not get a connection: %v", err)
	}
	if err := conn.Ping(); err != nil {
		conn.Close()
		return nil, fmt.Errorf("mysql: could not establish a good connection: %v", err)
	}

	db := &MysqlDB{
		conn: conn,
	}

	return db, nil
}

// Close closes the database, freeing up any resources.
func (db *MysqlDB) Close() {
	db.conn.Close()
}

// rowScanner is implemented by sql.Row and sql.Rows
type rowScanner interface {
	Scan(dest ...interface{}) error
}

// scanMqttMsg reads a mqtt message from a sql.Row or sql.Rows
func scanMqttMsg(s rowScanner) (*MqttMsg, error) {
	var (
        id        int32
        timestamp sql.NullString
        batterysn sql.NullString
        slotnum   int32
        temprature int32
        voltage   int32
        current   int32
        elapsed   int32
    )

	if err := s.Scan(&id, &batterysn, &slotnum, &temprature, &voltage, &current, &elapsed); err != nil {
		return nil, err
	}

	msg := &MqttMsg{
		Id:        id,
		Timestamp: timestamp.String,
        Batterysn: batterysn.String,
        Slotnum: slotnum,
        Temprature: temprature,
        Voltage: voltage,
        Current: current,
        Elapsed: elapsed,
	}

	return msg, nil
}

func (db *MysqlDB) ListMqttMessages(table_name string) ([]*MqttMsg, error) {

	rows, err := db.conn.Query("SELECT * FROM " + table_name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var msgs []*MqttMsg
	for rows.Next() {
		msg, err := scanMqttMsg(rows)
		if err != nil {
			return nil, fmt.Errorf("mysql: could not read row: %v", err)
		}

		msgs = append(msgs, msg)
	}

	return msgs, nil
}

func (db *MysqlDB) ListMqttmsgsOfCmd(table_name, cmd string) ([]*MqttMsg, error) {
	if cmd == "" {
		return db.ListMqttMessages(table_name)
	}

	stmt, err := db.conn.Prepare("SELECT * FROM " + table_name + " WHERE cmd= ?")
	if err != nil {
		return nil, fmt.Errorf("mysql: prepare select by cmd: %v", err)
	}

	rows, err := stmt.Query(cmd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var msgs []*MqttMsg
	for rows.Next() {
		msg, err := scanMqttMsg(rows)
		if err != nil {
			return nil, fmt.Errorf("mysql: could not read row: %v", err)
		}

		msgs = append(msgs, msg)
	}

	return msgs, nil
}

func (db *MysqlDB) GetMqttMsgById(table_name string, id int64) (*MqttMsg, error) {
	stmt, err := db.conn.Prepare("SELECT * FROM " + table_name + "  WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("mysql: prepare select by id: %v", err)
	}

	msg, err := scanMqttMsg(stmt.QueryRow(id))
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("mysql: could not find msg with id %d", id)
	}
	if err != nil {
		return nil, fmt.Errorf("mysql: could not get msg: %v", err)
	}
	return msg, nil
}

  


func (db *MysqlDB) InsertMqttMsg(table_name string, msg *MqttMsg) (id int64, err error) {
	stmt, err := db.conn.Prepare("INSERT INTO " + table_name + " (timestamp, batterysn, slotnum, temprature, voltage, current, elapsed) VALUES ( ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("mysql: prepare list by cmd: %v", err)
	}

	r, err := execAffectingOneRow(stmt, msg.Timestamp, msg.Batterysn, msg.Slotnum, msg.Temprature, msg.Voltage, msg.Current, msg.Elapsed)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := r.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("mysql: could not get last insert ID: %v", err)
	}
	return lastInsertID, nil
}

func (db *MysqlDB) DeleteMqttMsg(table_name string, id int64) error {

	if id == 0 {
		return errors.New("mysql: delete msg id == 0")
	}

	stmt, err := db.conn.Prepare("DELETE FROM " + table_name + " WHERE id = ?")
	_, err = execAffectingOneRow(stmt, id)
	return err
}

func (db *MysqlDB) UpdateMqttMsg(table_name string, msg *MqttMsg) error {
	if msg.Id == 0 {
		return errors.New("mysql: update msg id == 0")
	}

   
	stmt, err := db.conn.Prepare("UPDATE " + table_name + " SET id=?, timestamp=?, batterysn=?, slotnum=?, temprature=?, voltage=?, current=?, elapsed=?, WHERE id = ?")
	_, err = execAffectingOneRow(stmt, msg.Id, msg.Timestamp, msg.Batterysn, msg.Slotnum, msg.Temprature, msg.Voltage, msg.Current, msg.Elapsed)
	return err
}

func Init() {
    fmt.Println("mysql init")
}
