package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func get_internal() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
	os.Exit(0)
}
func main() {
	fmt.Printf("%%")
	get_internal()

	//insertData()
	//selectData()
	var t time.Time
	fmt.Println(t, t.IsZero())
	k := time.Now()
	hour, _ := time.ParseDuration("-48h")
	minute, _ := time.ParseDuration("-10m")
	fmt.Println(k.Add(hour).Format("2006-01-02 15:04:05"))
	fmt.Println(k.Add(hour).Add(minute).Format("2006-01-02 15:04:05"))
}

func insertData() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8")
	checkErr(err)
	stmt, err := db.Prepare(`INSERT user (user_name,user_age,user_sex) values (?,?,?)`)
	checkErr(err)
	res, err := stmt.Exec("tony", 20, 1)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func selectData() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8")
	checkErr(err)
	rows, err := db.Query("SELECT *FROM USER")
	checkErr(err)
	/*
		for rows.Next() {
			var userId int
			var userName string
			var userAge int
			var userSex int
			rows.Columns()
			err = rows.Scan(&userId, &userName, &userAge, &userSex)
			checkErr(err)
			fmt.Println(userId)
			fmt.Println(userName)
			fmt.Println(userAge)
			fmt.Println(userSex)
		}
	*/
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}

func update() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare(`UPDATE user SET user_age=?, user_sex=? WHERE user_id=?`)
	checkErr(err)
	res, err := stmt.Exec(21, 2, 1)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

func remove() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare(`DELETE FROM user where user_id=?`)
	checkErr(err)
	res, err := stmt.Exec(1)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
