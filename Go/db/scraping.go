package db

import (
	"app/models"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
)

func CategoryReshape(category string) (newCategory string) {
	switch category {
	case "刑事", "検事", "スピンオフ刑事", "刑事ミステリー", "刑事コメディー", "刑事推理", "警察学園":
		newCategory = "刑事・検事"
	case "社会派", "ビジネス", "企業", "オフィス", "ロマンティック・コメディ経済", "空港", "会社", "業界":
		newCategory = "企業・オフィス"
	case "学園", "青春", "学園コメディ", "学園コメディー", "学園恋愛", "学園アクション", "青春ホラー":
		newCategory = "学園・青春"
	case "ホーム", "ヒューマン", "人間", "SF人間", "ホームヒューマン", "女性":
		newCategory = "ホーム・ヒューマン"
	case "ラブコメディ", "SF・恋愛", "ラブストーリー":
		newCategory = "恋愛"
	case "ミステリー", "推理サスペンス", "サスペンス", "SFサスペンス", "コメディーミステリー", "ロマンチックミステリー", "犯罪サスペンス", "ホームサスペンス", "逃亡サスペンス", "サスペンス推理", "サバイバルサスペンス", "クライム・サスペンス", "純愛ミステリー", "学園青春サスペンス":
		newCategory = "ミステリー・サスペンス"
	case "大河", "伝記", "SF歴史", "戦国":
		newCategory = "時代劇"
	case "法律", "法廷もの", "法廷ものコメディー", "裁判":
		newCategory = "弁護士"
	case "探偵", "推理", "サイコスリラー", "推理アクション":
		newCategory = "探偵・推理"
	case "シリアス・コメディ", "ファンタジー", "ファンタジーコメディ", "コメディ", "パロディ", "音楽コメディ", "冒険コメディ", "コメディヒューマン", "ケータイ発", "ロマンティック・コメディ", "冒険コメディー", "ヒューマンコメディー", "ホラーコメディ":
		newCategory = "コメディ・パロディ"
	case "経済", "金融", "政治コメディ":
		newCategory = "政治"
	case "料理・人間", "料理":
		newCategory = "グルメ"
	case "サスペンス犯罪", "犯罪", "復讐":
		newCategory = "犯罪・復讐"
	case "医療アクション", "医療恋愛":
		newCategory = "医療"
	case "スポーツコメディ":
		newCategory = "スポーツ"
	case "スパイコメディ", "アクションサスペンス":
		newCategory = "アクション"
	case "SF・ファンタジー", "特撮":
		newCategory = "SF"
	}
	return newCategory
}

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
				html, _ := u.Html()
				content := strings.Replace(p.Sanitize(html), "<br/>", " ", -1)
				content = strings.Replace(content, "\n", " ", -1)
				data = append(data, content)
			})
			wikiURL, _ := t.Find("a").Attr("href")
			if len(data) == 5 {
				title, _ := t.Find("a").Attr("title")
				tvProgram.Title = title
				tvProgram.Star = 5
				category := strings.Replace(data[1], "ドラマ", "", -1)
				tvProgram.Category = CategoryReshape(category)
				tvProgram.Production = data[2]
				tvProgram.Cast = data[4]
				tvProgram.ImageUrl = SetRandomImageURL()
				tvProgram.ImageUrlReference = ""
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
				}
				tvProgram.Week = &weekStruct

				hourBlock := strings.Split(strings.Split(data[3], "-")[0], "曜")
				var floatHour float32 = 100
				if len(hourBlock) == 2 {
					startTime := strings.TrimSpace(hourBlock[1])
					hourStart := strings.Split(startTime, ":")
					hour, _ := strconv.Atoi(hourStart[0])
					mins, _ := strconv.Atoi(hourStart[1])
					if 15 > mins && mins >= 0 {
						floatHour = float32(hour) + 0.0
					} else if 45 > mins && mins >= 15 {
						floatHour = float32(hour) + 0.5
					} else if 60 > mins && mins >= 45 {
						floatHour = float32(hour) + 1.0
					}
					// 無記入のとき
					if startTime == ":00" {
						floatHour = 100
					}
					tvProgram.Hour = floatHour
				}
				GetTvProgramInformation(tvProgram)
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
	var topCast string
	tableNum := 0
	newTvProgram := tvProgram
	dataAddFlag := false
	s.Each(func(_ int, u *goquery.Selection) {
		seasonNum := 0
		doramaFlag := false
		if tableNum == 0 {
			newTvProgram = *new(models.TvProgram)
		}
		newTvProgram.Title = doc.Find("h1").Text()
		newTvProgram.Star = 5
		newTvProgram.WikiReference = tvProgram.WikiReference
		newTvProgram.ImageUrl = SetRandomImageURL()
		u.Find("tbody > tr").Each(func(_ int, t *goquery.Selection) {
			c, _ := t.Find("td").Attr("class")
			if c == "category" {
				if strings.Contains(t.Find("td").Text(), "ドラマ") || strings.Contains(t.Find("td").Text(), "医療ミステリ") {
					doramaFlag = true
					tableNum += 1
				}
			}
			color, _ := t.Find("th").Attr("style")
			th := t.Find("th").Text()
			// 同一テーブルに複数のシーズンが表記されている場合
			if strings.Contains(color, "background-color: #FDEBD0") {
				if !strings.Contains(th, "話から") {
					seasonNum += 1
				}
				if seasonNum != 1 {
					// fmt.Println("there\n", newTvProgram)
					dataAddFlag = true
					if _, err := models.AddTvProgram(&newTvProgram); err != nil {
						fmt.Println(err)
					}
				}
				newTvProgram.Id = 0
				newTvProgram.Themesong = ""
				newTvProgram.Cast = topCast
				newTvProgram.Star = 5
				newTvProgram.Title = doc.Find("h1").Text()
				if strings.Contains(th, newTvProgram.Title) {
					newTvProgram.Title = th
				} else if strings.Contains(newTvProgram.Title, th) {
				} else {
					newTvProgram.Title += "（" + th + "）"
				}
			}
			if doramaFlag {
				html, _ := t.Find("td").Html()
				content := strings.Replace(p.Sanitize(html), "<br/>", " ", -1)
				content = strings.Replace(content, "\n", "", -1)
				switch th {
				case "ジャンル":
					if tvProgram.Category != "" {
						content = tvProgram.Category
					} else {
						content = strings.Replace(content, "ドラマ", "", -1)
					}
					newTvProgram.Category = content
				case "脚本":
					newTvProgram.Dramatist = content
				case "演出":
					newTvProgram.Director = content
				case "監督":
					newTvProgram.Supervisor = content
				case "出演者":
					if seasonNum == 0 {
						topCast = content
					}
					if seasonNum != 0 && newTvProgram.Cast != "" {
						newTvProgram.Cast += " " + content
					} else {
						newTvProgram.Cast = content
					}
				case "制作", "製作":
					if tvProgram.Production != "" {
						newTvProgram.Production = tvProgram.Production
					} else {
						newTvProgram.Production = content
					}
				case "オープニング":
					content = strings.Replace(content, "、「", "「", -1)
					content = strings.Replace(content, " 「", "「", -1)
					newTvProgram.Themesong = content
				case "エンディング":
					if strings.TrimSpace(t.Find("td").Text()) != "同上" {
						content = strings.Replace(content, "、「", "「", -1)
						content = strings.Replace(content, " 「", "「", -1)
						if newTvProgram.Themesong == "" {
							newTvProgram.Themesong = content
						} else {
							if !strings.Contains(newTvProgram.Themesong, content) {
								newTvProgram.Themesong += " " + content
							}
						}
					}
				case "放送国・地域":
					if strings.TrimSpace(content) != "日本" {
						doramaFlag = false
					}
				case "放送期間":
					re := regexp.MustCompile("(\\d{4})")
					contents := strings.Split(content, "年")
					year, _ := strconv.Atoi(re.FindStringSubmatch(contents[0])[0])
					newTvProgram.Year = year
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
					newTvProgram.Season = &seasonStruct
				case "放送時間":
					if content != "同上" {
						content = strings.Replace(content, "毎週", "", -1)
						content = strings.Replace(content, "曜日", "曜", -1)
						content = strings.Replace(content, "月 - 金", "平日", -1)
						contents := strings.Split(content, "曜")
						weekStruct := *new(models.Week)
						if len(contents) == 2 {
							weekStruct.Name = contents[0]
							newTvProgram.Week = &weekStruct
							contents = strings.Split(contents[1], "-")
						} else {
							weekStruct.Name = "?"
							newTvProgram.Week = &weekStruct
							contents = strings.Split(contents[0], "-")
						}
						content = strings.TrimSpace(contents[0])
						contents = strings.Split(content, ":")
						var floatHour float32 = 100
						if len(contents) == 2 {
							hour, _ := strconv.Atoi(contents[0])
							mins, _ := strconv.Atoi(contents[1])
							if 15 > mins && mins >= 0 {
								floatHour = float32(hour) + 0.0
							} else if 45 > mins && mins >= 15 {
								floatHour = float32(hour) + 0.5
							} else if 60 > mins && mins >= 45 {
								floatHour = float32(hour) + 1.0
							}
							newTvProgram.Hour = floatHour
						} else {
							contents = strings.Split(content, "時")
							if len(contents) >= 2 {
								hour, _ := strconv.Atoi(contents[0])
								newTvProgram.Hour = float32(hour)
							} else {
								newTvProgram.Hour = floatHour
							}
						}
					}
				}
			}
		})

		if doramaFlag {
			// fmt.Println(newTvProgram.Title, newTvProgram.Category)
			if _, err := models.AddTvProgram(&newTvProgram); err != nil {
				fmt.Println(err)
			}
			dataAddFlag = true
		}
	})
	if !dataAddFlag {
		// fmt.Println("---------------", tvProgram.Title)
		if _, err := models.AddTvProgram(&tvProgram); err != nil {
			fmt.Println(err)
		}
	}
}

// func UpdateTvProgramsInformation() {
// 	var fields []string
// 	var sortby []string
// 	var order []string
// 	var query = make(map[string]string)
// 	var limit int64 = 1000000
// 	var offset int64

// 	l, _ := models.GetAllTvProgram(query, fields, sortby, order, offset, limit)
// 	for _, tvProgram := range l {
// 		query["TvProgramId"] = strconv.FormatInt(tvProgram.(models.TvProgram).Id, 10)
// 		// 誰も情報を更新していない番組のみ
// 		if h, err := models.GetAllTvProgramUpdateHistory(query, fields, sortby, order, offset, limit); err == nil && len(h) == 0 {
// 			GetTvProgramInformation(tvProgram.(models.TvProgram))
// 		} else {
// 			// updateしたuserIDがすべて0番＝admin
// 			var n int64 = 0
// 			for _, v := range h {
// 				n += v.(models.TvProgramUpdateHistory).UserId
// 			}
// 			if n == 0 {
// 				GetTvProgramInformation(tvProgram.(models.TvProgram))
// 			}
// 		}

// 	}
// }

// Add dorama information in wiki lists.
func AddTvProgramsInformation() {
	wikis := []string{"日本のテレビドラマ一覧_(2010年代)", "日本のテレビドラマ一覧_(2000年代)"}
	// wikis := []string{"日本のテレビドラマ一覧_(2010年代)"}
	for _, v := range wikis {
		GetWikiDoramas("https://ja.wikipedia.org/wiki/" + v)
	}

}

// Get movie information.
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
		tvProgram.ImageUrl = "https://movie.walkerplus.com/api/resizeimage/content/" + id + "?w=300"
		tvProgram.Content = m.Find(".info > p").Text()
		director := strings.TrimSpace(m.Find(".info > .directorList > dd").Text())
		director = strings.Replace(director, " ", "", -1)
		director = strings.Replace(director, "\n\n\n\n", " ", -1)
		tvProgram.Director = director
		cast := strings.TrimSpace(m.Find(".info > .roleList > dd").Text())
		cast = strings.Replace(cast, " ", "", -1)
		cast = strings.Replace(cast, "\n\n\n\n", " ", -1)
		tvProgram.Cast = cast
		tvProgram.Star = 5
		if _, err := models.AddTvProgram(&tvProgram); err != nil {
			fmt.Println(err)
		}
		// fmt.Println(tvProgram.Title)
	})
}

// Get movies.
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
// Use tv create page in top-page.
func GetTvProgramInformationByURL(wikiReferenceURL string) (newTvProgram models.TvProgram) {
	doc, err := goquery.NewDocument(wikiReferenceURL)
	if err != nil {
		fmt.Print("URL scarapping failed\n")
		return newTvProgram
	}
	p := bluemonday.NewPolicy()
	p.AllowElements("br").AllowElements("td")
	s := doc.Find("table.infobox")
	doramaFlag := false

	s.Each(func(_ int, u *goquery.Selection) {
		if doramaFlag {
		} else {
			newTvProgram = *new(models.TvProgram)
			newTvProgram.Title = doc.Find("h1").Text()
			newTvProgram.WikiReference = wikiReferenceURL
			newTvProgram.ImageUrl = SetRandomImageURL()
			u.Find("tbody > tr").Each(func(_ int, t *goquery.Selection) {
				c, _ := t.Find("td").Attr("class")
				if c == "category" {
					if strings.Contains(t.Find("td").Text(), "ドラマ") || strings.Contains(t.Find("td").Text(), "医療ミステリ") {
						doramaFlag = true
					}
				}
				if doramaFlag {
					th := t.Find("th").Text()
					html, _ := t.Find("td").Html()
					content := strings.Replace(p.Sanitize(html), "<br/>", " ", -1)
					content = strings.Replace(content, "\n", "", -1)
					switch th {
					case "ジャンル":
						content = strings.Replace(content, "ドラマ", "", -1)
						newTvProgram.Category = CategoryReshape(content)
					case "脚本":
						newTvProgram.Dramatist = content
					case "演出":
						newTvProgram.Director = content
					case "監督":
						newTvProgram.Supervisor = content
					case "出演者":
						newTvProgram.Cast = content
					case "制作", "製作":
						newTvProgram.Production = content
					case "オープニング":
						content = strings.Replace(content, "、「", "「", -1)
						content = strings.Replace(content, " 「", "「", -1)
						newTvProgram.Themesong = content
					case "エンディング":
						if strings.TrimSpace(t.Find("td").Text()) != "同上" {
							content = strings.Replace(content, "、「", "「", -1)
							content = strings.Replace(content, " 「", "「", -1)
							if newTvProgram.Themesong == "" {
								newTvProgram.Themesong = content
							} else {
								if !strings.Contains(newTvProgram.Themesong, content) {
									newTvProgram.Themesong += " " + content
								}
							}
						}
					case "放送国・地域":
						if strings.TrimSpace(content) != "日本" {
							doramaFlag = false
						}
					case "放送期間":
						re := regexp.MustCompile("(\\d{4})")
						contents := strings.Split(content, "年")
						year, _ := strconv.Atoi(re.FindStringSubmatch(contents[0])[0])
						newTvProgram.Year = year
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
						newTvProgram.Season = &seasonStruct
					case "放送時間":
						if content != "同上" {
							content = strings.Replace(content, "毎週", "", -1)
							content = strings.Replace(content, "曜日", "曜", -1)
							content = strings.Replace(content, "月 - 金", "平日", -1)
							contents := strings.Split(content, "曜")
							weekStruct := *new(models.Week)
							if len(contents) == 2 {
								weekStruct.Name = contents[0]
								newTvProgram.Week = &weekStruct
								contents = strings.Split(contents[1], "-")
							} else {
								weekStruct.Name = "?"
								newTvProgram.Week = &weekStruct
								contents = strings.Split(contents[0], "-")
							}
							content = strings.TrimSpace(contents[0])
							contents = strings.Split(content, ":")
							var floatHour float32 = 100
							if len(contents) == 2 {
								hour, _ := strconv.Atoi(contents[0])
								mins, _ := strconv.Atoi(contents[1])
								if 15 > mins && mins >= 0 {
									floatHour = float32(hour) + 0.0
								} else if 45 > mins && mins >= 15 {
									floatHour = float32(hour) + 0.5
								} else if 60 > mins && mins >= 45 {
									floatHour = float32(hour) + 1.0
								}
								newTvProgram.Hour = floatHour
							} else {
								contents = strings.Split(content, "時")
								if len(contents) >= 2 {
									hour, _ := strconv.Atoi(contents[0])
									newTvProgram.Hour = float32(hour)
								} else {
									newTvProgram.Hour = floatHour
								}
							}
						}
					}
				}
			})
		}
	})
	if doramaFlag {
		return newTvProgram
	} else {
		newTvProgram = *new(models.TvProgram)
		return newTvProgram
	}
}

// イメージ画像をランダムに選ぶ
func SetRandomImageURL() (url string) {
	rand.Seed(time.Now().UnixNano())
	r := strconv.Itoa(rand.Intn(10) + 1)
	if len(r) == 1 {
		r = "0" + r
	}
	url = "/static/img/tv_img/hanko_" + r + ".png"
	return url
}

// Scraping TvProgram Information in main.go.
func GetTvProgramInformationByURLOnGo(wikiReferenceURL string) {
	newTvProgram := *new(models.TvProgram)
	doc, err := goquery.NewDocument(wikiReferenceURL)
	if err != nil {
		fmt.Print("URL scarapping failed\n")
		return
	}
	p := bluemonday.NewPolicy()
	p.AllowElements("br").AllowElements("td")
	s := doc.Find("table.infobox")
	var topCast string
	tableNum := 0
	s.Each(func(_ int, u *goquery.Selection) {
		seasonNum := 0
		doramaFlag := false
		if tableNum == 0 {
			newTvProgram = *new(models.TvProgram)
		}
		newTvProgram.Title = doc.Find("h1").Text()
		newTvProgram.Star = 5
		newTvProgram.WikiReference = wikiReferenceURL
		newTvProgram.ImageUrl = SetRandomImageURL()
		newTvProgram.ImageUrlReference = ""
		u.Find("tbody > tr").Each(func(_ int, t *goquery.Selection) {
			c, _ := t.Find("td").Attr("class")
			if c == "category" {
				if strings.Contains(t.Find("td").Text(), "ドラマ") || strings.Contains(t.Find("td").Text(), "医療ミステリ") {
					doramaFlag = true
					tableNum += 1
				}
			}
			color, _ := t.Find("th").Attr("style")
			th := t.Find("th").Text()
			// 同一テーブルに複数のシーズンが表記されている場合
			if strings.Contains(color, "background-color: #FDEBD0") {
				seasonNum += 1
				if seasonNum != 1 {
					if _, err := models.AddTvProgram(&newTvProgram); err != nil {
						fmt.Println(err)
					}
				}
				newTvProgram.Id = 0
				newTvProgram.Themesong = ""
				newTvProgram.Cast = topCast
				newTvProgram.Star = 5
				newTvProgram.Title = doc.Find("h1").Text()
				if strings.Contains(th, newTvProgram.Title) {
					newTvProgram.Title = th
				} else if strings.Contains(newTvProgram.Title, th) {
				} else {
					newTvProgram.Title += "（" + th + "）"
				}
			}
			if doramaFlag {
				html, _ := t.Find("td").Html()
				content := strings.Replace(p.Sanitize(html), "<br/>", " ", -1)
				content = strings.Replace(content, "\n", "", -1)
				switch th {
				case "ジャンル":
					content = strings.Replace(content, "ドラマ", "", -1)
					newTvProgram.Category = CategoryReshape(content)
				case "脚本":
					newTvProgram.Dramatist = content
				case "演出":
					newTvProgram.Director = content
				case "監督":
					newTvProgram.Supervisor = content
				case "出演者":
					if seasonNum == 0 {
						topCast = content
					}
					if seasonNum != 0 && newTvProgram.Cast != "" {
						newTvProgram.Cast += " " + content
					} else {
						newTvProgram.Cast = content
					}
				case "制作", "製作":
					newTvProgram.Production = content
				case "オープニング":
					content = strings.Replace(content, "、「", "「", -1)
					content = strings.Replace(content, " 「", "「", -1)
					newTvProgram.Themesong = content
				case "エンディング":
					if strings.TrimSpace(t.Find("td").Text()) != "同上" {
						content = strings.Replace(content, "、「", "「", -1)
						content = strings.Replace(content, " 「", "「", -1)
						if newTvProgram.Themesong == "" {
							newTvProgram.Themesong = content
						} else {
							if !strings.Contains(newTvProgram.Themesong, content) {
								newTvProgram.Themesong += " " + content
							}
						}
					}
				case "放送国・地域":
					if strings.TrimSpace(content) != "日本" {
						doramaFlag = false
					}
				case "放送期間":
					re := regexp.MustCompile("(\\d{4})")
					contents := strings.Split(content, "年")
					year, _ := strconv.Atoi(re.FindStringSubmatch(contents[0])[0])
					newTvProgram.Year = year
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
					newTvProgram.Season = &seasonStruct
				case "放送時間":
					if content != "同上" {
						content = strings.Replace(content, "毎週", "", -1)
						content = strings.Replace(content, "曜日", "曜", -1)
						content = strings.Replace(content, "月 - 金", "平日", -1)
						contents := strings.Split(content, "曜")
						weekStruct := *new(models.Week)
						if len(contents) == 2 {
							weekStruct.Name = contents[0]
							newTvProgram.Week = &weekStruct
							contents = strings.Split(contents[1], "-")
						} else {
							weekStruct.Name = "?"
							newTvProgram.Week = &weekStruct
							contents = strings.Split(contents[0], "-")
						}
						content = strings.TrimSpace(contents[0])
						contents = strings.Split(content, ":")
						var floatHour float32 = 100
						if len(contents) == 2 {
							hour, _ := strconv.Atoi(contents[0])
							mins, _ := strconv.Atoi(contents[1])
							if 15 > mins && mins >= 0 {
								floatHour = float32(hour) + 0.0
							} else if 45 > mins && mins >= 15 {
								floatHour = float32(hour) + 0.5
							} else if 60 > mins && mins >= 45 {
								floatHour = float32(hour) + 1.0
							}
							newTvProgram.Hour = floatHour
						} else {
							contents = strings.Split(content, "時")
							if len(contents) >= 2 {
								hour, _ := strconv.Atoi(contents[0])
								newTvProgram.Hour = float32(hour)
							} else {
								newTvProgram.Hour = floatHour
							}
						}
					}

				}
			}
		})
		if doramaFlag {
			if _, err := models.AddTvProgram(&newTvProgram); err != nil {
				fmt.Println(err)
			}
		}
	})
}

// サーバ側でデータ投入
func AddRecentTvInfo() {
	wikiTitles := []string{"4分間のマリーゴールド", "モトカレマニア", "G線上のあなたと私", "同期のサクラ", "時効警察はじめました", "俺の話は長い", "グランメゾン東京", "ニッポンノワール-刑事Yの反乱-", "チート〜詐欺師の皆さん、ご注意ください〜", "リカ (小説)", "スカーレット (テレビドラマ)", "ブラック校則 (2019年の映画)", "左ききのエレン", "結婚できない男", "シャーロック_(テレビドラマ)"}
	for _, v := range wikiTitles {
		url := "https://ja.wikipedia.org/wiki/" + v
		GetTvProgramInformationByURLOnGo(url)
	}
}
