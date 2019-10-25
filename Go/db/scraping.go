package db

import (
	"app/models"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
)

// Scraping TvPrograms by wiki list.
func GetWikiDoramas(referencePath string) {
	doc, err := goquery.NewDocument(referencePath)
	if err != nil {
		fmt.Print("url scarapping failed\n")
		return
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
				category := strings.Replace(data[1], "ドラマ", "", -1)
				tvProgram.Category = category
				switch category {
				case "刑事", "検事", "スピンオフ刑事", "刑事ミステリー", "刑事コメディー", "刑事推理", "警察学園":
					tvProgram.Category = "刑事・検事"
				case "社会派", "ビジネス", "企業", "オフィス", "ロマンティック・コメディ経済", "空港", "会社", "業界":
					tvProgram.Category = "企業・オフィス"
				case "学園", "青春", "学園コメディ", "学園コメディー", "学園恋愛", "学園アクション", "青春ホラー":
					tvProgram.Category = "学園・青春"
				case "ホーム", "ヒューマン", "人間", "SF人間", "ホームヒューマン", "女性":
					tvProgram.Category = "ホーム・ヒューマン"
				case "ラブコメディ", "SF・恋愛", "ラブストーリー":
					tvProgram.Category = "恋愛"
				case "ミステリー", "推理サスペンス", "サスペンス", "SFサスペンス", "コメディーミステリー", "ロマンチックミステリー", "犯罪サスペンス", "ホームサスペンス", "逃亡サスペンス", "サスペンス推理", "サバイバルサスペンス", "クライム・サスペンス", "純愛ミステリー", "学園青春サスペンス":
					tvProgram.Category = "ミステリー・サスペンス"
				case "大河", "伝記", "SF歴史", "戦国":
					tvProgram.Category = "時代劇"
				case "法律", "法廷もの", "法廷ものコメディー", "裁判":
					tvProgram.Category = "弁護士"
				case "探偵", "推理", "サイコスリラー", "推理アクション":
					tvProgram.Category = "探偵・推理"
				case "シリアス・コメディ", "ファンタジー", "ファンタジーコメディ", "コメディ", "パロディ", "音楽コメディ", "冒険コメディ", "コメディヒューマン", "ケータイ発", "ロマンティック・コメディ", "冒険コメディー", "ヒューマンコメディー", "ホラーコメディ":
					tvProgram.Category = "コメディ・パロディ"
				case "経済", "金融", "政治コメディ":
					tvProgram.Category = "政治"
				case "料理・人間", "料理":
					tvProgram.Category = "グルメ"
				case "サスペンス犯罪", "犯罪", "復讐":
					tvProgram.Category = "犯罪・復讐"
				case "医療アクション", "医療恋愛":
					tvProgram.Category = "医療"
				case "スポーツコメディ":
					tvProgram.Category = "スポーツ"
				case "スパイコメディ", "アクションサスペンス":
					tvProgram.Category = "アクション"
				case "SF・ファンタジー", "特撮":
					tvProgram.Category = "SF"
				}
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
				} else if data[3] == "参照" || weekNameLen == 0 {
					weekStruct.Name = "?"
				} else if weekNameLen > 3 {
					weekStruct.Name = "スペシャル"
				} else if weekNameLen == 3 && strings.Contains(weekName, " ") {
					weekStruct.Name = strings.Split(weekName, " ")[1]
				} else {
					weekStruct.Name = "?"
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
				}
			}
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
				case "ジャンル":
					content = strings.Replace(content, "ドラマ", "", -1)
					if content != "連続" && tvProgram.Category == "" {
						newTvProgram.Category = content
					}
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
				case "制作", "製作":
					if tvProgram.Production == "" {
						newTvProgram.Production = content
					}
				case "オープニング":
					content = strings.Replace(content, "、「", "「", -1)
					newTvProgram.Themesong = content
				case "エンディング":
					if strings.TrimSpace(t.Find("td").Text()) != "同上" {
						content = strings.Replace(content, "、「", "「", -1)
						if newTvProgram.Themesong == "" {
							newTvProgram.Themesong = content
						} else {
							newTvProgram.Themesong += "、" + content
						}
					}
				case "放送国・地域":
					if strings.TrimSpace(t.Find("td").Text()) != "日本" {
						doramaFlag = false
					}
				case "放送期間":
					contents := strings.Split(content, "年")
					if tvProgram.Year == 0 {
						year, _ := strconv.Atoi(contents[0])
						newTvProgram.Year = year
					}
					contents = strings.Split(contents[1], "月")
					if tvProgram.Season == nil {
						month, _ := strconv.Atoi(contents[0])
						seasonName := ""
						if month <= 3 {
							seasonName = "冬"
						} else if month <= 6 {
							seasonName = "春"
						} else if month <= 9 {
							seasonName = "夏"
						} else if month <= 12 {
							seasonName = "秋"
						}
						seasonStruct := *new(models.Season)
						seasonStruct.Name = seasonName
						newTvProgram.Season = &seasonStruct
					}
				case "放送時間":
					contents := strings.Split(content, "曜")
					if tvProgram.Week == nil {
						weekStruct := *new(models.Week)
						weekStruct.Name = contents[0]
						newTvProgram.Week = &weekStruct
					}
					if tvProgram.Hour == 100 {
						contents = strings.Split(contents[1], "-")
						content = strings.TrimSpace(contents[0])
						contents = strings.Split(content, ":")
						var floatHour float32 = 100
						if len(contents) == 2 {
							hours, _ := strconv.Atoi(contents[0])
							mins, _ := strconv.Atoi(contents[1])
							if 15 > mins && mins >= 0 {
								floatHour = float32(hours) + 0.0
							} else if 45 > mins && mins >= 15 {
								floatHour = float32(hours) + 0.5
							} else if 60 > mins && mins >= 45 {
								floatHour = float32(hours) + 1.0
							}
							newTvProgram.Hour = floatHour
						}
					}
				}
			}
		})
		// fmt.Println(newTvProgram.Season.Name, newTvProgram.Week.Name, newTvProgram.Year, newTvProgram.Hour, newTvProgram.Production, newTvProgram.Category)

		if doramaFlag {
			if err := models.UpdateTvProgramById(&newTvProgram); err != nil {
				fmt.Println(err)
			}
			w := models.TvProgramUpdateHistory{
				UserId:      0,
				TvProgramId: tvProgram.Id,
			}
			_, _ = models.AddTvProgramUpdateHistory(&w)
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
			// updateしたuserIDがすべて0番＝admin
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

func GetMovieWalker(year string, month string) {
	referencePath := "https://movie.walkerplus.com/list/" + year + "/" + month
	doc, err := goquery.NewDocument(referencePath)
	if err != nil {
		fmt.Print("url scarapping failed\n")
		return
	}
	yearInt, _ := strconv.Atoi(year)
	seasonName := ""
	monthInt, _ := strconv.Atoi(month)
	if monthInt <= 3 {
		seasonName = "冬"
	} else if monthInt <= 6 {
		seasonName = "春"
	} else if monthInt <= 9 {
		seasonName = "夏"
	} else if monthInt <= 12 {
		seasonName = "秋"
	}
	var floatHour float32 = 100
	doc.Find(".movie").Each(func(_ int, m *goquery.Selection) {
		var tvProgram models.TvProgram
		seasonStruct := *new(models.Season)
		seasonStruct.Name = seasonName
		tvProgram.Season = &seasonStruct
		weekStruct := *new(models.Week)
		weekStruct.Name = "映画"
		tvProgram.Week = &weekStruct
		tvProgram.Title = m.Find("h3").Text()
		tvProgram.Year = yearInt
		tvProgram.Hour = floatHour
		id, _ := m.Find("h3 > a").Attr("href")
		id = strings.Replace(id, "/", "", -1)
		id = strings.Replace(id, "mv", "", -1)
		tvProgram.ImageURL = "https://movie.walkerplus.com/api/resizeimage/content/" + id + "?w=300"
		tvProgram.Content = m.Find(".info > p").Text()
		director := strings.TrimSpace(m.Find(".info > .directorList > dd").Text())
		director = strings.Replace(director, " ", "", -1)
		director = strings.Replace(director, "\n\n\n\n", "、", -1)
		tvProgram.Director = director
		cast := strings.TrimSpace(m.Find(".info > .roleList > dd").Text())
		cast = strings.Replace(cast, " ", "", -1)
		cast = strings.Replace(cast, "\n\n\n\n", "、", -1)
		tvProgram.Cast = cast
		if _, err := models.AddTvProgram(&tvProgram); err != nil {
			fmt.Println(err)
		}
		// fmt.Println(tvProgram.Title)
	})
}

func GetMovieWalkers() {
	var start int = 2000
	var end int = time.Now().Year()
	y := 0
	for {
		year := strconv.Itoa(start + y)
		for m := 1; m <= 12; m++ {
			month := strconv.Itoa(m)
			if len(month) == 1 {
				month = "0" + month
			}
			fmt.Println(year, month)
			GetMovieWalker(year, month)
		}
		y++
		if (end - start + 1) == y {
			break
		}
	}
}

// Scraping TvProgram Information directly by wikiReferenceURL.
func GetTvProgramInformationByURL(wikiReferenceURL string) (tvProgram models.TvProgram) {
	doc, err := goquery.NewDocument(wikiReferenceURL)
	if err != nil {
		fmt.Print("URL scarapping failed\n")
		return tvProgram
	}
	p := bluemonday.NewPolicy()
	p.AllowElements("br").AllowElements("td")
	s := doc.Find("table.infobox")
	doramaFlag := false
	s.Each(func(_ int, u *goquery.Selection) {
		// doramaFlag = false
		// if doramaFlag {
		// 	continue
		// }
		tvProgram = *new(models.TvProgram)
		tvProgram.Title = doc.Find("h1").Text()
		tvProgram.WikiReference = wikiReferenceURL
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
				case "ジャンル":
					content = strings.Replace(content, "ドラマ", "", -1)
					tvProgram.Category = content
				case "脚本":
					tvProgram.Dramatist = content
				case "演出":
					tvProgram.Director = content
				case "監督":
					tvProgram.Supervisor = content
				case "出演者":
					tvProgram.Cast = content
				case "制作", "製作":
					tvProgram.Production = content
				case "オープニング":
					content = strings.Replace(content, "、「", "「", -1)
					tvProgram.Themesong = content
				case "エンディング":
					if strings.TrimSpace(t.Find("td").Text()) != "同上" {
						if tvProgram.Themesong == "" {
							tvProgram.Themesong = content
						} else {
							tvProgram.Themesong += "、" + content
						}
					}
				case "放送国・地域":
					if strings.TrimSpace(t.Find("td").Text()) != "日本" {
						doramaFlag = false
					}
				case "放送期間":
					contents := strings.Split(content, "年")
					year, _ := strconv.Atoi(contents[0])
					tvProgram.Year = year
					contents = strings.Split(contents[1], "月")
					month, _ := strconv.Atoi(contents[0])
					seasonName := ""
					if month <= 3 {
						seasonName = "冬"
					} else if month <= 6 {
						seasonName = "春"
					} else if month <= 9 {
						seasonName = "夏"
					} else if month <= 12 {
						seasonName = "秋"
					}
					seasonStruct := *new(models.Season)
					seasonStruct.Name = seasonName
					tvProgram.Season = &seasonStruct
				case "放送時間":
					contents := strings.Split(content, "曜")
					weekStruct := *new(models.Week)
					weekStruct.Name = contents[0]
					tvProgram.Week = &weekStruct
					contents = strings.Split(contents[1], "-")
					content = strings.TrimSpace(contents[0])
					contents = strings.Split(content, ":")
					var floatHour float32 = 100
					if len(contents) == 2 {
						hours, _ := strconv.Atoi(contents[0])
						mins, _ := strconv.Atoi(contents[1])
						if 15 > mins && mins >= 0 {
							floatHour = float32(hours) + 0.0
						} else if 45 > mins && mins >= 15 {
							floatHour = float32(hours) + 0.5
						} else if 60 > mins && mins >= 45 {
							floatHour = float32(hours) + 1.0
						}
						tvProgram.Hour = floatHour
					}
				}
			}
		})

		if doramaFlag {
			fmt.Println(tvProgram.Season.Name, tvProgram.Week.Name, tvProgram.Year, tvProgram.Hour, tvProgram.Production, tvProgram.Category)
			return
			// fmt.Println(tvProgram)
			// if _, err := models.AddTvProgram(&tvProgram); err != nil {
			// 	fmt.Println(err)
			// }
		}
	})
	if doramaFlag {
		return tvProgram
	} else {
		tvProgram = *new(models.TvProgram)
		return tvProgram
	}
}
