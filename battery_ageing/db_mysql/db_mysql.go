package db_mysql

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	//"time"
)

const DATABASE_NAME = "battery_ageing"
const DATABASE_USER = "phpmyadmin"
const DATABASE_PASSWORD = "PHP@password123"
const DATABASE_ADDR = "127.0.0.1"
const DATABASE_PORT = 3306

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
	Id                  int32
	Timestamp           string
	Batterysn           string
	Slotnum             int32
	Voltage             int32
	Current             int32
	Temprature          int32
	Elapsed             int32
	Discharging         int32
	XTemprature         int32
	XVoltage            int32
	XFullChargecapacity int32
	XRemainingcapacity  int32
	XAveragecurrent     int32
	XCyclecount         int32
	XBmssafetyStatus    int32
	XBmsflags           int32
	XBatterystatus      int32
	XChargestatus       int32
	XEnablestatus       int32
	XSlotstatus         int32
	XDestroyed          int32
	XHasbms             int32
	XRadio              int32
}

type OptionRecord struct {
	Id                  int32
	Timestamp           string
	User                string
	Batterysn           string
	Slotnum             int32
	Cmd                 int32
    Level               int32
	Rcc                 int32
}

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
func (config MySQLConfig) ensureUpinfoTableExists(table_name string) error {
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


	if _, err := conn.Exec("USE " + DATABASE_NAME); err != nil {
		// MySQL error 1049 is "database does not exist"
		if mErr, ok := err.(*mysql.MySQLError); ok && mErr.Number == 1049 {
			return createUpinfoTable(conn, DATABASE_NAME, table_name)
		}
	}

	if _, err := conn.Exec("DESCRIBE " + table_name); err != nil {
		// MySQL error 1146 is "table does not exist"
		if mErr, ok := err.(*mysql.MySQLError); ok && mErr.Number == 1146 {
			return createUpinfoTable(conn, DATABASE_NAME, table_name)
		}
		// Unknown error.
		return fmt.Errorf("mysql: could not connect to the database: %v", err)
	}

	return nil
}


// ensureTableExists checks the table exists. If not, it creates it.
func (config MySQLConfig) ensureOptTableExists(table_name string) error {
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


	if _, err := conn.Exec("USE " + DATABASE_NAME); err != nil {
		// MySQL error 1049 is "database does not exist"
		if mErr, ok := err.(*mysql.MySQLError); ok && mErr.Number == 1049 {
			return createOptTable(conn, DATABASE_NAME, table_name)
		}
	}

	if _, err := conn.Exec("DESCRIBE " + table_name); err != nil {
		// MySQL error 1146 is "table does not exist"
		if mErr, ok := err.(*mysql.MySQLError); ok && mErr.Number == 1146 {
			return createOptTable(conn, DATABASE_NAME, table_name)
		}
		// Unknown error.
		return fmt.Errorf("mysql: could not connect to the database: %v", err)
	}

	return nil
}

// createTable creates the table, and if necessary, the database.
func createUpinfoTable(conn *sql.DB, database_name, table_name string) error {

	createTableStatements := make([]string, 3)
	createTableStatements[0] = fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v %v", database_name, "DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';")
	createTableStatements[1] = fmt.Sprintf("USE %v", database_name)
	createTableStatements[2] = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v %v", table_name,
		"(id INT UNSIGNED NOT NULL AUTO_INCREMENT, " +
		"timestamp           TEXT NULL, " +
		"batterysn           TEXT NULL, " +
		"slotnum             INT  NULL, " +
		"voltage             INT  NULL, " +
		"current             INT  NULL, " +
		"temprature          INT  NULL, " +
		"elapsed             INT  NULL, " +
		"discharging         INT  NULL, " +
		"xtemprature         INT  NULL, " +
		"xvoltage            INT  NULL, " +
		"xfullchargecapacity INT  NULL, " +
		"xremainingcapacity  INT  NULL, " +
		"xaveragecurrent     INT  NULL, " +
		"xcyclecount         INT  NULL, " +
		"xbmssafetystatus    INT  NULL, " +
		"xbmsflags           INT  NULL, " +
		"xbatterystatus      INT  NULL, " +
		"xchargestatus       INT  NULL, " +
		"xenablestatus       INT  NULL, " +
		"xslotstatus         INT  NULL, " +
		"xdestroyed          INT  NULL, " +
		"xhasbms             INT  NULL, " +
		"xradio              INT  NULL, " +
		"PRIMARY KEY (id))")

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

// createTable creates the table, and if necessary, the database.
func createOptTable(conn *sql.DB, database_name, table_name string) error {

	createTableStatements := make([]string, 3)
	createTableStatements[0] = fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v %v", database_name, "DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';")
	createTableStatements[1] = fmt.Sprintf("USE %v", database_name)
	createTableStatements[2] = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v %v", table_name,
		"(id INT UNSIGNED NOT NULL AUTO_INCREMENT, " +
		"timestamp           TEXT NULL, " +
		"user                TEXT NULL, " +
		"batterysn           TEXT NULL, " +
		"slotnum             INT  NULL, " +
		"cmd                 INT  NULL, " +
		"level               INT  NULL, " +
		"rcc                 INT  NULL, " +
		"PRIMARY KEY (id))")

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
		id                   int32
		timestamp            sql.NullString
		batterysn            sql.NullString
		slotnum              int32
		voltage              int32
		current              int32
		temprature           int32
		elapsed              int32
		discharging          int32
		xtemprature          int32
		xvoltage             int32
		xfullchargecapacity  int32
		xremainingcapacity   int32
		xaveragecurrent      int32
		xcyclecount          int32
		xbmssafetystatus     int32
		xbmsflags            int32
		xbatterystatus       int32
		xchargestatus        int32
		xenablestatus        int32
		xslotstatus          int32
		xdestroyed           int32
		xhasbms              int32
		xradio               int32
    )

	if err := s.Scan(
		&id                  ,
		&timestamp           ,
		&batterysn           ,
		&slotnum             ,
		&voltage             ,
		&current             ,
		&temprature          ,
		&elapsed             ,
		&discharging         ,
		&xtemprature         ,
		&xvoltage            ,
		&xfullchargecapacity ,
		&xremainingcapacity  ,
		&xaveragecurrent     ,
		&xcyclecount         ,
		&xbmssafetystatus    ,
		&xbmsflags           ,
		&xbatterystatus      ,
		&xchargestatus       ,
		&xenablestatus       ,
		&xslotstatus         ,
		&xdestroyed          ,
		&xhasbms             ,
		&xradio              ,
		); err != nil {
		return nil, err
	}

	msg := &MqttMsg {
		Id                  : id                 ,
		Timestamp           : timestamp.String   ,
		Batterysn           : batterysn.String   ,
		Slotnum             : slotnum            ,
		Voltage             : voltage            ,
		Current             : current            ,
		Temprature          : temprature         ,
		Elapsed             : elapsed            ,
		Discharging         : discharging        ,
		XTemprature         : xtemprature        ,
		XVoltage            : xvoltage           ,
		XFullChargecapacity : xfullchargecapacity,
		XRemainingcapacity  : xremainingcapacity ,
		XAveragecurrent     : xaveragecurrent    ,
		XCyclecount         : xcyclecount        ,
		XBmssafetyStatus    : xbmssafetystatus   ,
		XBmsflags           : xbmsflags          ,
		XBatterystatus      : xbatterystatus     ,
		XChargestatus       : xchargestatus      ,
		XEnablestatus       : xenablestatus      ,
		XSlotstatus         : xslotstatus        ,
		XDestroyed          : xdestroyed         ,
		XHasbms             : xhasbms            ,
		XRadio              : xradio             ,
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


func (db *MysqlDB) InsertMqttMsg(config MySQLConfig, table_name string, msg *MqttMsg) (id int64, err error) {

	// Check database and table exists. If not, create it.
	if err := config.ensureUpinfoTableExists(table_name); err != nil {
		return 0, fmt.Errorf("table %v not exist %v", table_name, err)
	}

	stmt, err := db.conn.Prepare("INSERT INTO " + table_name + " (" +
	"timestamp           , " +
	"batterysn           , " +
	"slotnum             , " +
	"voltage             , " +
	"current             , " +
	"temprature          , " +
	"elapsed             , " +
	"discharging         , " +
	"xtemprature         , " +
	"xvoltage            , " +
	"xfullchargecapacity , " +
	"xremainingcapacity  , " +
	"xaveragecurrent     , " +
	"xcyclecount         , " +
	"xbmssafetystatus    , " +
	"xbmsflags           , " +
	"xbatterystatus      , " +
	"xchargestatus       , " +
	"xenablestatus       , " +
	"xslotstatus         , " +
	"xdestroyed          , " +
	"xhasbms             , " +
	"xradio                " +
	") VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("mysql: prepare list by cmd: %v", err)
	}

	r, err := execAffectingOneRow(stmt,
		msg.Timestamp           ,
		msg.Batterysn           ,
		msg.Slotnum             ,
		msg.Voltage             ,
		msg.Current             ,
		msg.Temprature          ,
		msg.Elapsed             ,
		msg.Discharging         ,
		msg.XTemprature         ,
		msg.XVoltage            ,
		msg.XFullChargecapacity ,
		msg.XRemainingcapacity  ,
		msg.XAveragecurrent     ,
		msg.XCyclecount         ,
		msg.XBmssafetyStatus    ,
		msg.XBmsflags           ,
		msg.XBatterystatus      ,
		msg.XChargestatus       ,
		msg.XEnablestatus       ,
		msg.XSlotstatus         ,
		msg.XDestroyed          ,
		msg.XHasbms             ,
		msg.XRadio              )

	if err != nil {
		return 0, err
	}

	lastInsertID, err := r.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("mysql: could not get last insert ID: %v", err)
	}
	return lastInsertID, nil
}

func (db *MysqlDB) InsertOptionRecord(config MySQLConfig, table_name string, opt *OptionRecord) (id int64, err error) {

	// Check database and table exists. If not, create it.
	if err := config.ensureOptTableExists(table_name); err != nil {
		return 0, fmt.Errorf("table %v not exist %v", table_name, err)
	}

	stmt, err := db.conn.Prepare("INSERT INTO " + table_name + " (" +
	"timestamp           , " +
	"user                , " +
	"batterysn           , " +
	"slotnum             , " +
	"cmd                 , " +
	"level               , " +
	"rcc                   " +
	") VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("mysql: prepare list by cmd: %v", err)
	}

	r, err := execAffectingOneRow(stmt,
		opt.Timestamp           ,
		opt.User                ,
		opt.Batterysn           ,
		opt.Slotnum             ,
		opt.Cmd                 ,
		opt.Level               ,
		opt.Rcc)

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


	stmt, err := db.conn.Prepare("UPDATE " + table_name + " SET " +
		"id                  =?, " +
		"timestamp           =?, " +
		"batterysn           =?, " +
		"slotnum             =?, " +
		"voltage             =?, " +
		"current             =?, " +
		"temprature=         =?, " +
		"elapsed             =?, " +
		"discharging         =?, " +
		"xtemprature         =?, " +
		"xvoltage            =?, " +
		"xfullchargecapacity =?, " +
		"xremainingcapacity  =?, " +
		"xaveragecurrent     =?, " +
		"xcyclecount         =?, " +
		"xbmssafetystatus    =?, " +
		"xbmsflags           =?, " +
		"xbatterystatus      =?, " +
		"xchargestatus       =?, " +
		"xenablestatus       =?, " +
		"xslotstatus         =?, " +
		"xdestroyed          =?, " +
		"xhasbms             =?, " +
		"xradio              =?, " +
		"WHERE id = ?")

		_, err = execAffectingOneRow(stmt,
			msg.Id                  ,
			msg.Timestamp           ,
			msg.Batterysn           ,
			msg.Slotnum             ,
			msg.Voltage             ,
			msg.Current             ,
			msg.Temprature          ,
			msg.Elapsed             ,
			msg.Discharging         ,
			msg.XTemprature         ,
			msg.XVoltage            ,
			msg.XFullChargecapacity ,
			msg.XRemainingcapacity  ,
			msg.XAveragecurrent     ,
			msg.XCyclecount         ,
			msg.XBmssafetyStatus    ,
			msg.XBmsflags           ,
			msg.XBatterystatus      ,
			msg.XChargestatus       ,
			msg.XEnablestatus       ,
			msg.XSlotstatus         ,
			msg.XDestroyed          ,
			msg.XHasbms             ,
			msg.XRadio              )

	return err
}

func Init() {
    fmt.Println("mysql init")
}
