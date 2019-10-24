package db

import (
	"app/models"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
)

// Scraping TvPrograms by wiki list.
func Scraping(referencePath string) {
	doc, err := goquery.NewDocument(referencePath)
	if err != nil {
		fmt.Print("url scarapping failed\n")
	}
	var year []string
	doc.Find("span.mw-headline").Each(func(_ int, s *goquery.Selection) {
		text, _ := s.Attr("id")
		text = text
		if len(year) < 10 {
			year = append(year, text[:4])
		}
	})
	p := bluemonday.NewPolicy()
	p.AllowElements("br").AllowElements("td")
	n := 0
	doc.Find("table.wikitable").Each(func(_ int, s *goquery.Selection) {
		season := string(s.Find("caption").Text())
		seasonName := ""
		if season[:1] == "4" {
			seasonName = "春"
		} else if season[:1] == "7" {
			seasonName = "夏"
		} else if season[:2] == "10" {
			seasonName = "秋"
		} else if season[:1] == "1" {
			seasonName = "冬"
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
				// data = append(data, u.Text())
				html, _ := u.Html()
				content := strings.Replace(p.Sanitize(html), "<br/>", "、", -1)
				content = strings.Replace(content, "\n", "、", -1)
				data = append(data, content)
			})
			wikiURL, _ := t.Find("a").Attr("href")
			if len(data) == 5 {
				tvProgram.Title = strings.Replace(data[0], "、", " ", -1)
				tvProgram.Category = strings.Replace(data[1], "ドラマ", "", -1)
				tvProgram.Production = data[2]
				tvProgram.Cast = data[4]
				tvProgram.ImageURL = "http://hankodeasobu.com/wp-content/uploads/animals_02.png"
				tvProgram.WikiReference = "https://ja.wikipedia.org" + wikiURL
				weekStruct := *new(models.Week)
				data[3] = strings.Replace(data[3], "平日", "平曜", -1)
				weekName := strings.Split(data[3], "曜")[0]
				weekNameLen := utf8.RuneCountInString(weekName)
				if weekNameLen == 1 {
					weekStruct.Name = weekName
					if weekName == "平" {
						weekStruct.Name = "平日"
					}
				} else if data[3] == "参照" || weekNameLen == 0 || weekNameLen > 3 {
					weekStruct.Name = "?"
				} else if weekNameLen == 3 && strings.Contains(weekName, " ") {
					weekStruct.Name = strings.Split(weekName, " ")[1]
				} else {
					weekStruct.Name = weekName
					// fmt.Println(weekName, weekNameLen)
				}
				tvProgram.Week = &weekStruct

				hourBlock := strings.Split(strings.Split(data[3], "-")[0], "曜")
				var floatHour float32 = 100
				if len(hourBlock) == 2 {
					startTime := strings.TrimSpace(hourBlock[1])
					hourStart := strings.Split(startTime, ":")
					hours, _ := strconv.Atoi(hourStart[0])
					mins, _ := strconv.Atoi(hourStart[1])
					if 15 > mins && mins >= 0 {
						floatHour = float32(hours) + 0.0
					} else if 45 > mins && mins >= 15 {
						floatHour = float32(hours) + 0.5
					} else if 60 > mins && mins >= 45 {
						floatHour = float32(hours) + 1.0
					}
					// 無記入のとき
					if startTime == ":00" {
						floatHour = 100
					}
					tvProgram.Hour = floatHour
				}
				if _, err := models.AddTvProgram(&tvProgram); err != nil {
					fmt.Println(err)
					// fmt.Println(weekName, data[0])
				}
			}
			// fmt.Println(tvProgram)
		})
		if season[:2] == "10" {
			n++
		}
	})
}

// Scraping TvProgram Information by wikiReferenceURL.
func GetTvProgramInformation(tvProgram models.TvProgram) {
	doc, err := goquery.NewDocument(tvProgram.WikiReference)
	if err != nil {
		fmt.Print("URL scarapping failed\n")
		return
	}
	// fmt.Println(tvProgram)
	p := bluemonday.NewPolicy()
	p.AllowElements("br").AllowElements("td")
	s := doc.Find("table.infobox")
	s.Each(func(_ int, u *goquery.Selection) {
		doramaFlag := false
		newTvProgram := tvProgram
		u.Find("tbody > tr").Each(func(_ int, t *goquery.Selection) {
			c, _ := t.Find("td").Attr("class")
			if c == "category" {
				if strings.Contains(t.Find("td").Text(), "ドラマ") || t.Find("td").Text() == "医療ミステリー" {
					doramaFlag = true
				}
			}
			if doramaFlag {
				th := t.Find("th").Text()
				html, _ := t.Find("td").Html()
				content := strings.Replace(p.Sanitize(html), "<br/>", "、", -1)
				content = strings.Replace(content, "\n", "", -1)
				switch th {
				case "脚本":
					newTvProgram.Dramatist = content
				case "演出":
					newTvProgram.Director = content
				case "監督":
					newTvProgram.Supervisor = content
				case "出演者":
					if len(tvProgram.Cast) < len(content) {
						newTvProgram.Cast = content
					}
				case "制作":
					if tvProgram.Production == "" {
						newTvProgram.Production = content
					}
				case "オープニング":
					newTvProgram.Themesong = content
				case "エンディング":
					if strings.TrimSpace(t.Find("td").Text()) != "同上" {
						if tvProgram.Themesong == "" {
							newTvProgram.Themesong = content
						} else {
							newTvProgram.Themesong += "、" + content
						}
					}
				case "放送国・地域":
					if strings.TrimSpace(t.Find("td").Text()) != "日本" {
						doramaFlag = false
					}
				}
			}
		})
		if doramaFlag {
			if err := models.UpdateTvProgramById(&newTvProgram); err != nil {
				fmt.Println(err)
			}
			// var w models.TvProgramUpdateHistory
			w := models.TvProgramUpdateHistory{
				UserId:      0,
				TvProgramId: tvProgram.Id,
			}
			_, _ = models.AddTvProgramUpdateHistory(&w)
			// fmt.Println(tvProgram.Title)
		}
	})
}

func UpdateTvProgramsInformation() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 1000000
	var offset int64

	l, _ := models.GetAllTvProgram(query, fields, sortby, order, offset, limit)
	for _, tvProgram := range l {
		query["TvProgramId"] = strconv.FormatInt(tvProgram.(models.TvProgram).Id, 10)
		if h, err := models.GetAllTvProgramUpdateHistory(query, fields, sortby, order, offset, limit); err == nil && len(h) == 0 {
			GetTvProgramInformation(tvProgram.(models.TvProgram))
		} else {
			var n int64 = 0
			for _, v := range h {
				n += v.(models.TvProgramUpdateHistory).UserId
			}
			if n == 0 {
				GetTvProgramInformation(tvProgram.(models.TvProgram))
			}
		}

	}
}
