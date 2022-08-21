package main

import (
	"fmt"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var DB *sql.DB

func main() {
	var err error
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, TimestampFormat: "2006-01-02 15:04:05"})

	if err := initconfig(); err != nil {
		logrus.Error(err)
		return
	}

	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", I.Mysql.User, I.Mysql.Password, I.Mysql.Ip, I.Mysql.Port, I.Mysql.Database))
	if err != nil {
		logrus.Error("open database fail")
		return
	}
	defer DB.Close()
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		logrus.Error("ping database fail")
		return
	}
	logrus.Infof("l2s is running! Listen:%s Domain:%s", I.Webserver.Listen, I.Application.Domain)
	http.HandleFunc("/", webIndex)
	logrus.Info(http.ListenAndServe(I.Webserver.Listen, nil))

}
