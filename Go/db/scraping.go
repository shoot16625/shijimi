package db

import (
	// _ "app/routers"
	"app/models"
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	// "time"
	"github.com/microcosm-cc/bluemonday"
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

func GetTvProgramInformation(title string) (tvProgram models.TvProgram) {
	doc, err := goquery.NewDocument("https://ja.wikipedia.org/wiki/" + title)
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	s := doc.Find("table.infobox")
	tvProgram.Title = title
	p := bluemonday.NewPolicy()
	p.AllowElements("br").AllowElements("td")
	s.Each(func(_ int, u *goquery.Selection) {
		doramaFlag := false
		u.Find("tbody > tr").Each(func(_ int, t *goquery.Selection) {
			c, _ := t.Find("td").Attr("class")
			if c == "category" {
				if strings.Contains(t.Find("td").Text(), "ドラマ") {
					doramaFlag = true
				}
			}
			if doramaFlag {
				th := t.Find("th").Text()
				switch th {
				case "脚本":
					html, _ := t.Find("td").Html()
					tvProgram.Dramatist = strings.Replace(p.Sanitize(html), "<br/>", "、", -1)
				case "演出":
					html, _ := t.Find("td").Html()
					tvProgram.Director = strings.Replace(p.Sanitize(html), "<br/>", "、", -1)
				case "監督":
					html, _ := t.Find("td").Html()
					tvProgram.Supervisor = strings.Replace(p.Sanitize(html), "<br/>", "、", -1)
				case "出演者":
					html, _ := t.Find("td").Html()
					tvProgram.Cast = strings.Replace(p.Sanitize(html), "<br/>", "、", -1)
				// case "制作":
				// tvProgram.Production = t.Find("td").Text()
				case "オープニング":
					html, _ := t.Find("td").Html()
					tvProgram.Themesong = strings.Replace(p.Sanitize(html), "<br/>", "、", -1)
				case "エンディング":
					html, _ := t.Find("td").Html()
					if t.Find("td").Text() != "同上" {
						if tvProgram.Themesong == "" {
							tvProgram.Themesong = strings.Replace(p.Sanitize(html), "<br/>", "、", -1)
						} else {
							tvProgram.Themesong += "、" + strings.Replace(p.Sanitize(html), "<br/>", "、", -1)
						}
					}
				}
				// if t.Find("th").Text() == "脚本" {
				// 	tvProgram.Dramatist = t.Find("td").Text()
				// }

			}
		})
		// fmt.Println(tvProgram)
	})
	return tvProgram
}

func UpdateTvProgramsInformation() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 1000000
	var offset int64

	var newTvProgram models.TvProgram
	l, _ := models.GetAllTvProgram(query, fields, sortby, order, offset, limit)
	for _, tvProgram := range l {
		// newTvProgram = tvProgram
		tvInfo := GetTvProgramInformation(tvProgram.(models.TvProgram).Title)
		newTvProgram = models.TvProgram{
			Id:                 tvProgram.(models.TvProgram).Id,
			Title:              tvProgram.(models.TvProgram).Title,
			Content:            tvProgram.(models.TvProgram).Content,
			ImageUrl:           tvProgram.(models.TvProgram).ImageUrl,
			ImageUrlReference:  tvProgram.(models.TvProgram).ImageUrlReference,
			MovieUrl:           tvProgram.(models.TvProgram).MovieUrl,
			MovieUrlReference:  tvProgram.(models.TvProgram).MovieUrlReference,
			Cast:               tvInfo.Cast,
			Category:           tvInfo.Category,
			Dramatist:          tvInfo.Dramatist,
			Supervisor:         tvInfo.Supervisor,
			Director:           tvInfo.Director,
			Production:         tvInfo.Production,
			Year:               tvProgram.(models.TvProgram).Year,
			Season:             tvProgram.(models.TvProgram).Season,
			Week:               tvProgram.(models.TvProgram).Week,
			Hour:               tvProgram.(models.TvProgram).Hour,
			Themesong:          tvInfo.Themesong,
			CreateUserId:       tvProgram.(models.TvProgram).CreateUserId,
			Star:               tvProgram.(models.TvProgram).Star,
			CountStar:          tvProgram.(models.TvProgram).CountStar,
			CountWatched:       tvProgram.(models.TvProgram).CountWatched,
			CountWantToWatch:   tvProgram.(models.TvProgram).CountWantToWatch,
			CountClicked:       tvProgram.(models.TvProgram).CountClicked,
			CountAuthorization: tvProgram.(models.TvProgram).CountAuthorization,
		}
		if err := models.UpdateTvProgramById(&newTvProgram); err != nil {
			fmt.Println("Miss")
		} else {
			fmt.Println(newTvProgram)
		}
	}
}
