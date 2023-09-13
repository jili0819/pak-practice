package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
)

func main() {
	_ = city()
}

// 全国城市解析
// city
func city() (err error) {
	// Request the HTML page.
	res, err := http.Get("https://ip.bczs.net/diming")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = res.Body.Close()
	}()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var citySqls []City
	// Find the review items
	doc.Find(".container-ip .well table tbody tr").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		cityInfo := City{}
		s.Find("td").Each(func(i int, selection *goquery.Selection) {
			title := selection.Text()
			switch i {
			case 0:
				cityInfo.Code = strings.TrimSpace(title)
			case 1:
				cityInfo.Name = strings.TrimSpace(title)
			case 2:
				cityInfo.FullName = strings.TrimSpace(title)
			case 3:
				cityInfo.Level = strings.TrimSpace(title)
			}
			fmt.Printf("Review %d: %s\n", i, title)
		})
		citySqls = append(citySqls, cityInfo)
	})
	_ = insertMysql(citySqls)
	return
}

type City struct {
	Code     string
	Name     string
	FullName string
	Level    string
}

// 赋值mysql
// insertMysql
func insertMysql(cityList []City) (err error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:mysqlps@123@tcp(127.0.0.1:3306)/city?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                                  // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                 // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                 // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                 // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	return db.Table("city").Create(cityList).Error
}
