package db

import (
	// _ "app/routers"
	"app/models"
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	// "time"
	// "github.com/microcosm-cc/bluemonday"
	// "github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
)

const location = "Asia/Tokyo"

func Scraping() {
	doc, err := goquery.NewDocument("https://ja.wikipedia.org/wiki/日本のテレビドラマ一覧_(2010年代)")
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	var year []string
	doc.Find("span.mw-headline").Each(func(_ int, s *goquery.Selection) {
		text, _ := s.Attr("id")
		text = text
		if len(year) < 10 {
			year = append(year, text[:4])
			// fmt.Println(year)
		}
	})
	n := 0
	doc.Find("table.wikitable").Each(func(_ int, s *goquery.Selection) {
		season := string(s.Find("caption").Text())
		var seasonName string
		if season[:1] == "4" {
			seasonName = "春"
		} else if season[:1] == "7" {
			seasonName = "夏"
		} else if season[:2] == "10" {
			seasonName = "秋"
		} else if season[:1] == "1" {
			seasonName = "冬"
		} else {
			seasonName = ""
		}
		s.Find("tbody > tr").Each(func(_ int, t *goquery.Selection) {
			var tvProgram models.TvProgram
			seasonStruct := *new(models.Season)
			seasonStruct.Name = seasonName
			tvProgram.Season = &seasonStruct
			y, _ := strconv.Atoi(year[n])
			tvProgram.Year = y
			var data []string
			t.Find("td").Each(func(_ int, u *goquery.Selection) {
				data = append(data, u.Text())
			})
			if len(data) == 5 {
				tvProgram.Title = data[0]
				tvProgram.Category = strings.Replace(data[1], "ドラマ", "", -1)
				tvProgram.Production = data[2]
				tvProgram.Cast = data[4]
				tvProgram.ImageUrl = "http://hankodeasobu.com/wp-content/uploads/animals_02.png"

				weekStruct := *new(models.Week)
				weekStruct.Name = data[3][:3]
				tvProgram.Week = &weekStruct
				hourBlock := strings.Split(data[3][6:], " - ")
				if len(hourBlock) == 2 {
					hourStart := strings.Split(hourBlock[0], ":")
					hours, _ := strconv.Atoi(hourStart[0])
					mins, _ := strconv.Atoi(hourStart[1])
					var floatHour float32
					if 15 > mins && mins >= 0 {
						floatHour = float32(hours) + 0.0
					} else if 45 > mins && mins >= 15 {
						floatHour = float32(hours) + 0.5
					} else if 60 > mins && mins >= 45 {
						floatHour = float32(hours) + 1.0
						if floatHour == 24.0 {
							floatHour = 23.30

						}
					}
					tvProgram.Hour = floatHour
				}
				if _, err := models.AddTvProgram(&tvProgram); err != nil {
					fmt.Println(err)
				}
			}
		})
		if season[:2] == "10" {
			n++
		}
	})
}

// func init(){
//     loc, e := time.LoadLocation(location)
//     if e != nil {
//         loc = time.FixedZone(location, 9*60*60)
//     }
//     time.Local = loc

//     orm.RegisterDriver(beego.AppConfig.String("driver"), orm.DRMySQL)
//     orm.RegisterDataBase("default", beego.AppConfig.String("driver"), beego.AppConfig.String("sqlconn")+"?charset=utf8&loc=Asia%2FTokyo")
//     err := orm.RunSyncdb("default", false, false)
//     // err := orm.RunSyncdb("default", false, false)
//     if err != nil {
//         fmt.Println(err)
//     }

// }
