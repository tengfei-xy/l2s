package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

func getShortURL() string {
	rand.Seed(time.Now().UnixNano())
	const pool = "qazwsxedcrfvtgbyhnujmikolpQAZWSXEDCRFVTGBYHNUJMIKOLP1234567890"
	bytes := make([]byte, I.Application.Short_url_length)
	for i := 0; i < I.Application.Short_url_length; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}

func shorturl_get(l *logrus.Entry, ip string, longurl string) (string, int) {
	var query_count int
	shorturl := getShortURL()

	if err := DB.QueryRow(`select shorturl,query_count from url where longurl=?`, longurl).Scan(&shorturl, &query_count); err == sql.ErrNoRows {
		if _, err2 := DB.Exec(`INSERT INTO url(longurl,shorturl) values(?,?)`, longurl, shorturl); err2 != nil {
			return fmt.Sprintf("Failed to query:%s", err.Error()), 403
		}

	} else if err != nil {
		return fmt.Sprintf("Failed to query:%s", err.Error()), 403
	}
	shorturl = fmt.Sprintf("%s/%s", I.Application.Domain, shorturl)
	l.Infof("query_count:%d shorturl:%s", query_count, shorturl)
	return shorturl, 200
}

func longurl_get(l *logrus.Entry, ip string, shorturl string) (string, int) {
	var query_count int
	var longurl string
	if err := DB.QueryRow(`select longurl,query_count from url where shorturl=?`, shorturl).Scan(&longurl, &query_count); err == sql.ErrNoRows {
		return fmt.Sprintf("Failed to query:%s", err.Error()), 403
	} else if err != nil {
		return fmt.Sprintf("Failed to query:%s", err.Error()), 403
	}
	if _, err := DB.Exec(`UPDATE url SET query_count=? WHERE longurl=?`, query_count+1, longurl); err != nil {
		l.Error("failed to update when query_count=%s and longurl=%s, error:%s", query_count+1, longurl, err.Error())
	}
	l.Infof("query_count:%d longurl:%s", query_count, longurl)

	return longurl, 301

}
