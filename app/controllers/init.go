package controllers

import (
	"github.com/revel/revel"
	"github.com/coopernurse/gorp"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"strings"
	"bareskin-api/app/models/db"
	"path/filepath"
	"github.com/smotes/purse"
)

func init() {
	revel.OnAppStart(InitDb)
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}

var InitDb func() = func() {
	connectionString := getConnectionString()

	if db, err := sql.Open("mysql", connectionString); err != nil {
		revel.ERROR.Fatal(err)
	} else {
		Dbm = &gorp.DbMap{
			Db:      db,
			Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	}

	defineTables(Dbm)

	if revel.DevMode {
		revel.INFO.Printf("%s", "Dropping Tables")
		Dbm.DropTables()
		if err := Dbm.CreateTablesIfNotExists(); err != nil {
			revel.ERROR.Fatal(err)
		}
		inserTestData(Dbm)

	} else {
		if err := Dbm.CreateTablesIfNotExists(); err != nil {
			revel.ERROR.Fatal(err)
		}
	}
}

func inserTestData(dbm *gorp.DbMap) {
	ps, _ := purse.New(filepath.Join("dbdata/", ""))

	contents, ok := ps.Get("insertData.sql")
	if !ok {
		fmt.Println("SQL file not loaded")
	}

	requests := strings.Split(contents, ";")

	for _, request := range requests {

		if _, err := dbm.Exec(request); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func defineTables(dbm *gorp.DbMap) {

	dbm.AddTable(db.Order{}).SetKeys(true, "order_id")
	dbm.AddTable(db.OrderedProducts{}).SetKeys(true, "order_product_id")
	dbm.AddTable(db.PaymentChannel{}).SetKeys(true, "payment_channel_id")
	dbm.AddTable(db.Product{}).SetKeys(true, "id")
	dbm.AddTable(db.ProductTags{}).SetKeys(true, "product_tag_id")
	dbm.AddTable(db.Sale{}).SetKeys(true, "sale_id")
	dbm.AddTable(db.SoldProducts{}).SetKeys(true, "sold_product_id")
	dbm.AddTable(db.Tag{}).SetKeys(true, "tag_id")

	t := dbm.AddTable(db.User{}).SetKeys(true, "user_id")
	t.ColMap("username").SetMaxSize(25)
	t.ColMap("created_date").SetNotNull(false)
}

func getParamString(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == "" {
			revel.ERROR.Fatal("Cound not find parameter: " + param)
		} else {
			return defaultValue
		}
	}
	return p
}

func getConnectionString() string {
	host := getParamString("db.host", "")
	port := getParamString("db.port", "3306")
	user := getParamString("db.user", "")
	pass := getParamString("db.password", "")
	dbname := getParamString("db.name", "bareskindb")
	protocol := getParamString("db.protocol", "tcp")
	dbargs := getParamString("db.args", " ")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		user, pass, protocol, host, port, dbname, dbargs)
}
