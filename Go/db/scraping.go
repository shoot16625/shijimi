package db

import (
	"app/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
)

// Scraping TvPrograms by wiki list.
func GetWikiDoramas(referencePath string, fromYear int) {
	doc, err := goquery.NewDocument(referencePath)
	if err != nil {
		fmt.Print("url scarapping failed\n")
		return
	}
	var year []string
	doc.Find("span.mw-headline").Each(func(_ int, s *goquery.Selection) {
		text, _ := s.Attr("id")
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
		y, _ := strconv.Atoi(year[n])
		if fromYear <= y {
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
					content := strings.Replace(p.Sanitize(html), "<br/>", ",", -1)
					content = strings.Replace(content, "\n", ",", -1)
					content = strings.Replace(content, ",（", "（", -1)
					content = strings.Replace(content, "、", ",", -1)
					content = models.RegexpWords(content, ", | ,", ",")
					content = models.RegexpWords(content, `[\(|（](P*S.[0-9|\-| |、]+)+[\)|）]`, "")
					content = models.RegexpWords(content, `下記詳細|参照|スタッフ参照|ほか|（.*特別出演.*）|（第[1-9]部）|（主演として.+）|\[注 *[1-9]\]|以下五十音順`, "")
					content = strings.TrimSpace(content)
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
					tvProgram.ImageUrl = models.SetRandomImageURL()
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
		}
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
		dramaFlag := false
		if tableNum == 0 {
			newTvProgram = *new(models.TvProgram)
		}
		newTvProgram.Title = ReshapeTitle(doc.Find("h1").Text())
		newTvProgram.Star = 5
		newTvProgram.WikiReference = tvProgram.WikiReference
		u.Find("tbody > tr").Each(func(_ int, t *goquery.Selection) {
			c, _ := t.Find("td").Attr("class")
			if c == "category" {
				td := t.Find("td").Text()
				if strings.Contains(td, "ドラマ") || strings.Contains(td, "医療ミステリ") || strings.Contains(td, "コメディ") || strings.Contains(td, "時代劇") {
					dramaFlag = true
					tableNum += 1
					if strings.Contains(td, "ケータイドラマ") {
						newTvProgram.Year = 2000
						weekStruct := *new(models.Week)
						weekStruct.Name = "?"
						newTvProgram.Week = &weekStruct
						seasonStruct := *new(models.Season)
						seasonStruct.Name = "秋"
						newTvProgram.Season = &seasonStruct
						newTvProgram.Hour = 100
					}
				}
			}
			color, _ := t.Find("th").Attr("style")
			th := t.Find("th").Text()
			// 同一テーブルに複数のシーズンが表記されている場合
			if strings.Contains(color, "background-color: #FDEBD0") && dramaFlag {
				if !strings.Contains(th, "話から") {
					seasonNum += 1
				}
				if seasonNum != 1 {
					dataAddFlag = true
					newTvProgram.ImageUrl = models.GetImageURL(newTvProgram.Title)
					newTvProgram.ImageUrlReference = models.ReshapeImageURLReference(newTvProgram.ImageUrl)
					newTvProgram.MovieUrl = GetYoutubeURL(newTvProgram.Title)
					fmt.Println(newTvProgram.Title)
					if _, err := models.AddTvProgram(&newTvProgram); err != nil {
						fmt.Println(err)
					}
				}
				newTvProgram.Id = 0
				newTvProgram.Themesong = ""
				newTvProgram.Cast = topCast
				newTvProgram.Star = 5
				newTvProgram.Title = ReshapeTitle(doc.Find("h1").Text())
				if strings.Contains(th, newTvProgram.Title) {
					newTvProgram.Title = th
				} else if strings.Contains(newTvProgram.Title, th) {
				} else {
					newTvProgram.Title += "（" + th + "）"
				}
				if strings.Contains(newTvProgram.Title, "（再放送）") {
					dramaFlag = false
				}
				newTvProgram.Title = ReshapeTitle(newTvProgram.Title)
			}
			if dramaFlag {
				html, _ := t.Find("td").Html()
				content := strings.Replace(p.Sanitize(html), "<br/>", ",", -1)
				content = ReshapeText(content)
				switch th {
				case "ジャンル":
					if tvProgram.Category != "" {
						content = tvProgram.Category
					} else {
						content = strings.Replace(content, "ドラマ", "", -1)
					}
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
						newTvProgram.Cast += "," + content
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
					content = ReshapeThemesong(content)
					newTvProgram.Themesong = content
				case "エンディング":
					if strings.TrimSpace(t.Find("td").Text()) != "同上" {
						content = ReshapeThemesong(content)

						if newTvProgram.Themesong == "" {
							newTvProgram.Themesong = content
						} else {
							if !strings.Contains(newTvProgram.Themesong, content) {
								newTvProgram.Themesong += "," + content
							}
						}
					}
				case "放送国・地域":
					if strings.TrimSpace(content) != "日本" {
						dramaFlag = false
					}
				case "放送期間":
					if content != "未定" {
						re := regexp.MustCompile("(\\d{4})")
						contents := strings.Split(content, "年")
						year, _ := strconv.Atoi(re.FindStringSubmatch(contents[0])[0])
						newTvProgram.Year = year

						contents = strings.Split(contents[1], "月")
						month, _ := strconv.Atoi(contents[0])
						seasonName := ReshapeHour(month)
						seasonStruct := *new(models.Season)
						seasonStruct.Name = seasonName
						newTvProgram.Season = &seasonStruct
					}
				case "放送時間":
					if content != "同上" {
						contents := ReshapeWeek(content)
						weekStruct := *new(models.Week)
						if len(contents) == 2 {
							if len(contents[0]) > 6 {
								contents[0] = "?"
							}
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

		if dramaFlag {
			newTvProgram.ImageUrl = models.GetImageURL(newTvProgram.Title)
			newTvProgram.ImageUrlReference = models.ReshapeImageURLReference(newTvProgram.ImageUrl)
			newTvProgram.MovieUrl = GetYoutubeURL(newTvProgram.Title)
			fmt.Println(newTvProgram.Title)
			if _, err := models.AddTvProgram(&newTvProgram); err != nil {
				fmt.Println(err)
			}
			dataAddFlag = true
		}
	})
	if !dataAddFlag {
		// fmt.Println("---------------", tvProgram.Title)
		tvProgram.ImageUrl = models.GetImageURL(tvProgram.Title)
		tvProgram.ImageUrlReference = models.ReshapeImageURLReference(tvProgram.ImageUrl)
		tvProgram.MovieUrl = GetYoutubeURL(tvProgram.Title)
		fmt.Println(tvProgram.Title)
		if _, err := models.AddTvProgram(&tvProgram); err != nil {
			fmt.Println(err)
		}
	}
}

// Add drama information in wiki lists.
// change here
func AddTvProgramsInformation(wikis []string, fromYear int) {
	for _, v := range wikis {
		GetWikiDoramas("https://ja.wikipedia.org/wiki/"+v, fromYear)
	}

}

// 映画情報の取得onWikibyURL
func GetMovieInformationByURL(wikiReferenceURL string) (newTvProgram models.TvProgram) {
	doc, err := goquery.NewDocument(wikiReferenceURL)
	if err != nil {
		fmt.Print("URL scarapping failed\n")
		return newTvProgram
	}
	p := bluemonday.NewPolicy()
	p.AllowElements("br").AllowElements("td")
	p.AllowElements("br").AllowElements("th")
	s := doc.Find("table.infobox")
	dramaFlag := false
	contentVolume := 0

	s.Each(func(_ int, u *goquery.Selection) {
		c, _ := u.Attr("style")
		if c == "width:22em; width:20em" {
			tmp := len(u.Find("tbody > tr > td").Text())
			if contentVolume < tmp {
				contentVolume = tmp
				dramaFlag = true
			} else {
				dramaFlag = false
			}
		}
		if dramaFlag {
			newTvProgram = *new(models.TvProgram)
			html, _ := u.Find("tbody > tr > th").First().Html()
			newTvProgram.Title = strings.Replace(p.Sanitize(html), "<br/>", " ", -1)
			newTvProgram.WikiReference = wikiReferenceURL
			newTvProgram.ImageUrl = models.GetImageURL(newTvProgram.Title)
			u.Find("tbody > tr").Each(func(_ int, t *goquery.Selection) {
				th := t.Find("th").Text()
				html, _ := t.Find("td").Html()
				content := strings.Replace(p.Sanitize(html), "<br/>", ",", -1)
				content = ReshapeText(content)
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
				case "制作会社", "製作会社":
					newTvProgram.Production = content
				case "主題歌":
					content = ReshapeThemesong(content)
					newTvProgram.Themesong = content
				case "公開":
					if content != "未定" {
						contents := strings.Split(content, ",")
						content = contents[0]
						re := regexp.MustCompile("(\\d{4})")
						contents = strings.Split(content, "年")
						year, _ := strconv.Atoi(re.FindStringSubmatch(contents[0])[0])
						newTvProgram.Year = year
						contents = strings.Split(contents[1], "月")
						month, _ := strconv.Atoi(contents[0])
						seasonName := ReshapeHour(month)
						seasonStruct := *new(models.Season)
						seasonStruct.Name = seasonName
						newTvProgram.Season = &seasonStruct
					}
				}
				weekStruct := *new(models.Week)
				weekStruct.Name = "映画"
				newTvProgram.Week = &weekStruct
				var floatHour float32 = 100
				newTvProgram.Hour = floatHour
			})
		}
	})
	if contentVolume > 0 {
		newTvProgram.MovieUrl = GetYoutubeURL(newTvProgram.Title)
		return newTvProgram
	} else {
		newTvProgram = *new(models.TvProgram)
		return newTvProgram
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
	monthInt, _ := strconv.Atoi(month)
	seasonName := ReshapeHour(monthInt)
	var floatHour float32 = 100
	selection := doc.Find("div.un_screeningListTab")
	selection.Find("li").Each(func(_ int, m *goquery.Selection) {
		var tvProgram models.TvProgram
		seasonStruct := *new(models.Season)
		seasonStruct.Name = seasonName
		tvProgram.Season = &seasonStruct
		weekStruct := *new(models.Week)
		weekStruct.Name = "映画"
		tvProgram.Week = &weekStruct
		tvProgram.Title = m.Find(".bl_movieListItem_ttl").Text()
		tvProgram.Year = yearInt
		tvProgram.Hour = floatHour
		id, _ := m.Find("div > ul > li > a").Attr("href")
		id = strings.Replace(id, "/", "", -1)
		id = strings.Replace(id, "mv", "", -1)
		imageURL := "https://movie.walkerplus.com/api/resizeimage/content/" + id + "?w=300"
		imageURL = models.CheckImageURL(imageURL, tvProgram.Title)
		tvProgram.ImageUrl = imageURL
		tvProgram.ImageUrlReference = models.ReshapeImageURLReference(imageURL)
		tvProgram.MovieUrl = GetYoutubeURL(tvProgram.Title)
		tvProgram.Content = m.Find(".bl_movieList_desc > p").Text()
		// director := strings.TrimSpace(m.Find(".info > .directorList > dd").Text())
		// director = strings.Replace(director, " ", "", -1)
		// director = strings.Replace(director, "\n\n\n\n", ",", -1)
		// tvProgram.Director = director
		// cast := strings.TrimSpace(m.Find(".info > .roleList > dd").Text())
		// cast = strings.Replace(cast, " ", "", -1)
		// cast = strings.Replace(cast, "\n\n\n\n", ",", -1)
		// tvProgram.Cast = cast
		tvProgram.Star = 5
		if _, err := models.AddTvProgram(&tvProgram); err != nil {
			fmt.Println(err)
		}
	})
}

// Get movies.
// ex. 2020, 2020, 3, 6 -> 2020/03-2020/06
func GetMovieWalkers(start_year, end_year, start_month, end_month int) {
	y := 0
	for {
		year := strconv.Itoa(start_year + y)
		for m := start_month; m <= end_month; m++ {
			month := strconv.Itoa(m)
			if len(month) == 1 {
				month = "0" + month
			}
			fmt.Println(year, month)
			GetMovieWalker(year, month)
		}
		if (end_year - start_year) == y {
			break
		}
		y++
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
	dramaFlag := false

	s.Each(func(_ int, u *goquery.Selection) {
		if dramaFlag {
		} else {
			newTvProgram = *new(models.TvProgram)
			newTvProgram.Title = ReshapeTitle(doc.Find("h1").Text())
			newTvProgram.WikiReference = wikiReferenceURL
			// newTvProgram.ImageUrl = SetRandomImageURL()
			newTvProgram.ImageUrl = models.GetImageURL(newTvProgram.Title)
			u.Find("tbody > tr").Each(func(_ int, t *goquery.Selection) {
				c, _ := t.Find("td").Attr("class")
				if c == "category" {
					td := t.Find("td").Text()
					if strings.Contains(td, "ドラマ") || strings.Contains(td, "医療ミステリ") || strings.Contains(td, "コメディ") || strings.Contains(td, "時代劇") {
						dramaFlag = true
					}
				}
				if dramaFlag {
					th := t.Find("th").Text()
					html, _ := t.Find("td").Html()
					content := strings.Replace(p.Sanitize(html), "<br/>", ",", -1)
					content = ReshapeText(content)
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
						content = ReshapeThemesong(content)

						newTvProgram.Themesong = content
					case "エンディング":
						if strings.TrimSpace(t.Find("td").Text()) != "同上" {
							content = ReshapeThemesong(content)

							if newTvProgram.Themesong == "" {
								newTvProgram.Themesong = content
							} else {
								if !strings.Contains(newTvProgram.Themesong, content) {
									newTvProgram.Themesong += "," + content
								}
							}
						}
					case "放送国・地域":
						if strings.TrimSpace(content) != "日本" {
							dramaFlag = false
						}
					case "放送期間":
						if content != "未定" {
							re := regexp.MustCompile("(\\d{4})")
							contents := strings.Split(content, "年")
							year, _ := strconv.Atoi(re.FindStringSubmatch(contents[0])[0])
							newTvProgram.Year = year

							contents = strings.Split(contents[1], "月")
							month, _ := strconv.Atoi(contents[0])
							seasonName := ReshapeHour(month)
							seasonStruct := *new(models.Season)
							seasonStruct.Name = seasonName
							newTvProgram.Season = &seasonStruct
						}
					case "放送時間":
						if content != "同上" {
							contents := ReshapeWeek(content)
							weekStruct := *new(models.Week)
							if len(contents) == 2 {
								if len(contents[0]) > 6 {
									contents[0] = "?"
								}
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
	if dramaFlag {
		newTvProgram.MovieUrl = GetYoutubeURL(newTvProgram.Title)
		return newTvProgram
	} else {
		newTvProgram = *new(models.TvProgram)
		return newTvProgram
	}
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
		dramaFlag := false
		if tableNum == 0 {
			newTvProgram = *new(models.TvProgram)
		}
		newTvProgram.Title = ReshapeTitle(doc.Find("h1").Text())
		newTvProgram.Star = 5
		newTvProgram.WikiReference = wikiReferenceURL
		u.Find("tbody > tr").Each(func(_ int, t *goquery.Selection) {
			c, _ := t.Find("td").Attr("class")
			if c == "category" {
				td := t.Find("td").Text()
				if strings.Contains(td, "ドラマ") || strings.Contains(td, "医療ミステリ") || strings.Contains(td, "コメディ") || strings.Contains(td, "時代劇") {
					dramaFlag = true
					tableNum += 1
				}
			}
			color, _ := t.Find("th").Attr("style")
			th := t.Find("th").Text()
			// 同一テーブルに複数のシーズンが表記されている場合
			if strings.Contains(color, "background-color: #FDEBD0") && dramaFlag {
				if !strings.Contains(th, "話から") {
					seasonNum += 1
				}
				if seasonNum != 1 {
					newTvProgram.ImageUrl = models.GetImageURL(newTvProgram.Title)
					newTvProgram.ImageUrlReference = models.ReshapeImageURLReference(newTvProgram.ImageUrl)
					newTvProgram.MovieUrl = GetYoutubeURL(newTvProgram.Title)
					fmt.Println(newTvProgram.Title)
					if _, err := models.AddTvProgram(&newTvProgram); err != nil {
						fmt.Println(err)
					}
				}
				newTvProgram.Id = 0
				newTvProgram.Themesong = ""
				newTvProgram.Cast = topCast
				newTvProgram.Star = 5
				newTvProgram.Title = ReshapeTitle(doc.Find("h1").Text())
				if strings.Contains(th, newTvProgram.Title) {
					newTvProgram.Title = th
				} else if strings.Contains(newTvProgram.Title, th) {
				} else {
					newTvProgram.Title += "（" + th + "）"
				}
				if strings.Contains(newTvProgram.Title, "（再放送）") {
					dramaFlag = false
				}
				newTvProgram.Title = ReshapeTitle(newTvProgram.Title)
			}
			if dramaFlag {
				html, _ := t.Find("td").Html()
				content := strings.Replace(p.Sanitize(html), "<br/>", ",", -1)
				content = ReshapeText(content)
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
						newTvProgram.Cast += "," + content
					} else {
						newTvProgram.Cast = content
					}
				case "制作", "製作":
					newTvProgram.Production = content
				case "オープニング":
					content = ReshapeThemesong(content)

					newTvProgram.Themesong = content
				case "エンディング":
					if strings.TrimSpace(t.Find("td").Text()) != "同上" {
						content = ReshapeThemesong(content)
						if newTvProgram.Themesong == "" {
							newTvProgram.Themesong = content
						} else {
							if !strings.Contains(newTvProgram.Themesong, content) {
								newTvProgram.Themesong += "," + content
							}
						}
					}
				case "放送国・地域":
					if strings.TrimSpace(content) != "日本" {
						dramaFlag = false
					}
				case "放送期間":
					if content != "未定" {
						re := regexp.MustCompile("(\\d{4})")
						contents := strings.Split(content, "年")
						year, _ := strconv.Atoi(re.FindStringSubmatch(contents[0])[0])
						newTvProgram.Year = year

						contents = strings.Split(contents[1], "月")
						month, _ := strconv.Atoi(contents[0])
						seasonName := ReshapeHour(month)
						seasonStruct := *new(models.Season)
						seasonStruct.Name = seasonName
						newTvProgram.Season = &seasonStruct
					}
				case "放送時間":
					if content != "同上" {
						contents := ReshapeWeek(content)
						weekStruct := *new(models.Week)
						if len(contents) == 2 {
							if len(contents[0]) > 6 {
								contents[0] = "?"
							}
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
		if dramaFlag {
			newTvProgram.ImageUrl = models.GetImageURL(newTvProgram.Title)
			newTvProgram.ImageUrlReference = models.ReshapeImageURLReference(newTvProgram.ImageUrl)
			newTvProgram.MovieUrl = GetYoutubeURL(newTvProgram.Title)
			fmt.Println(newTvProgram.Title)
			if _, err := models.AddTvProgram(&newTvProgram); err != nil {
				fmt.Println(err)
			}
		}
	})
}

// 映画情報の更新 in main.go
func GetMovieInformationByURLOnGo(wikiReferenceURL string, newTvProgram models.TvProgram) {
	if newTvProgram.MovieUrl == "" {
		newTvProgram.MovieUrl = GetYoutubeURL(newTvProgram.Title)
	}
	if newTvProgram.ImageUrl == "" || newTvProgram.ImageUrl == "https://d2ueuvlup6lbue.cloudfront.net/assets/common/img_cover-placehold-633a19fbbf6951dbb8debea06f29fefcc0666e8138e79c5ffd8a486da95432ae.svg" {
		newTvProgram.ImageUrl = models.GetImageURL(newTvProgram.Title)
		newTvProgram.ImageUrlReference = models.ReshapeImageURLReference(newTvProgram.ImageUrl)
	}
	doc, err := goquery.NewDocument(wikiReferenceURL)
	if err != nil {
		fmt.Print("URL scarapping failed\n")
	} else {
		p := bluemonday.NewPolicy()
		p.AllowElements("br").AllowElements("td")
		p.AllowElements("br").AllowElements("th")
		s := doc.Find("table.infobox")
		dramaFlag := false
		contentVolume := 0

		s.Each(func(_ int, u *goquery.Selection) {
			c, _ := u.Attr("style")
			if c == "width:22em; width:20em" {
				tmp := len(u.Find("tbody > tr > td").Text())
				if contentVolume < tmp {
					contentVolume = tmp
					dramaFlag = true
				} else {
					dramaFlag = false
				}
			}
			if dramaFlag {
				newTvProgram.WikiReference = wikiReferenceURL
				newTvProgram.ImageUrl = models.GetImageURL(newTvProgram.Title)
				u.Find("tbody > tr").Each(func(_ int, t *goquery.Selection) {
					th := t.Find("th").Text()
					html, _ := t.Find("td").Html()
					content := strings.Replace(p.Sanitize(html), "<br/>", ",", -1)
					content = ReshapeText(content)
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
					case "制作会社", "製作会社":
						if len(content) < 32 {
							newTvProgram.Production = content
						}
					case "主題歌":
						content = ReshapeThemesong(content)
						newTvProgram.Themesong = content
					}
				})
			}
		})
	}
	// fmt.Println(newTvProgram)
	if err := models.UpdateTvProgramById(&newTvProgram); err != nil {
		fmt.Println(err)
	}
}

// https://developers.google.com/youtube/v3/determine_quota_cost
func GetYoutubeURL(str string) (URL string) {
	title := strings.Replace(str, " ", "", -1)
	title = url.QueryEscape(title)
	keyNum := 5
	apikey := ""
	rand.Seed(time.Now().UnixNano())
	seed := rand.Intn(keyNum)
	if seed == 0 {
		// shijimi00
		apikey = "hogehoge"
	} else if seed == 1 {
		// shijimi01
		apikey = "hogehoge"
	} else if seed == 2 {
		// shijimi
		apikey = "hogehoge"
	} else if seed == 3 {
		// playTag55
		apikey = "hogehoge"
	} else if seed >= 4 {
		// playTag55
		apikey = "hogehoge"
	}

	query := "https://www.googleapis.com/youtube/v3/search?type=video&part=snippet&maxResults=1&order=viewCount&videoDuration=short&q=" + title + "&key=" + apikey
	resp, err := http.Get(query)
	if err != nil {
		fmt.Println("youtube api error")
		return URL
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("youtube api error")
		return URL
	}
	data := make(map[string]interface{})
	if err := json.Unmarshal(body, &data); err == nil {
		// quotaが制限を超えていなかったら結果を抽出
		if _, ok := data["items"]; ok {
			for _, item := range data["items"].([]interface{}) {
				youtubeID := item.(map[string]interface{})["id"].(map[string]interface{})["videoId"].(string)
				URL = "https://www.youtube.com/embed/" + youtubeID
			}
		}
	}
	return URL
}

// タイトルの整形
func ReshapeTitle(str string) string {
	content := models.RegexpWords(str, ` *[\(|（].*[テレビドラマ|連続ドラマ|時代劇|漫画|小説][\)|）]`, "")
	content = strings.Replace(content, "　", " ", -1)
	return content
}

// 曜日
func ReshapeWeek(str string) []string {
	content := strings.Replace(str, "毎週", "", -1)
	content = strings.Replace(content, "曜日", "曜", -1)
	content = strings.Replace(content, "月 - 金", "平日", -1)
	contents := strings.Split(content, "曜")
	return contents
}

// 放送時間
func ReshapeHour(month int) string {
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
	return seasonName
}

// 主題歌
func ReshapeThemesong(str string) string {
	content := strings.Replace(str, "『", "「", -1)
	content = strings.Replace(content, "』", "」", -1)
	content = models.RegexpWords(content, ",「| 「", "「")
	return content
}

// 内容
func ReshapeText(str string) string {
	content := strings.Replace(str, "\n", "", -1)
	content = strings.Replace(content, ",（", "（", -1)
	content = models.RegexpWords(content, ", | ,", ",")
	content = strings.Replace(content, "　", " ", -1)
	var contents []string
	for _, v := range strings.Split(content, ",") {
		v = models.RegexpWords(v, `[\(|（](P*S.[0-9|\-| |、]+)+[\)|）]|[\(|（].*[出演|シーズン|1st|2nd|3rd|原案]+.*[\)|）]|[\(|（].+[のみ|シリーズ]+[\)|）]|[\(|（][主演として|特別|脚本|SP\.|以上|当時]+.+[\)|）]|）]|[\(|（][音楽|MMJ|テレビ朝日|日本テレビ|関西テレビ|TBS|共同テレビ|CP|FCC|連続ドラマ]+[\)|）]|[\(|（].*第.+[部|作|話|期]+.*[\)|）]|[\(|（][1-9]* - [1-9]*[\)|）]|[\(|（][1-9]+[\)|）]`, "")
		// カッコでないもの
		v = models.RegexpWords(v, `\[注.* *[1-9]\]|\[[1-9]+\]|下記詳細|参照|スタッフ参照|ほか|以下五十音順|[0-9]+年版|第[1-9]+シリーズ|1st|2nd|3rd|【連続ドラマ】|【特別編】`, "")
		v = strings.TrimSpace(v)
		if v != "" {
			contents = append(contents, v)
		}
	}
	content = strings.Join(contents, ",")
	content = strings.TrimRight(content, "他")
	content = strings.TrimSpace(content)
	return content
}

// サーバ側でデータ投入
func AddRecentTvInfo(wikiTitles []string) {
	// wikiTitles := []string{"4分間のマリーゴールド", "モトカレマニア", "G線上のあなたと私", "同期のサクラ", "時効警察はじめました", "俺の話は長い", "グランメゾン東京", "ニッポンノワール-刑事Yの反乱-", "チート〜詐欺師の皆さん、ご注意ください〜", "リカ (小説)", "スカーレット (テレビドラマ)", "ブラック校則 (2019年の映画)", "左ききのエレン", "結婚できない男", "シャーロック_(テレビドラマ)", "絶対零度_(テレビドラマ)", "病室で念仏を唱えないでください", "やめるときも、すこやかなるときも", "10の秘密", "恋はつづくよどこまでも", "知らなくていいコト", "僕はどこから", "来世ではちゃんとします", "ケイジとケンジ〜所轄と地検の24時〜", "アライブ_がん専門医のカルテ", "ゆるキャン△", "駐在刑事", "女子高生の無駄づかい", "絶メシロード", "トップナイフ_(小説)", "アリバイ崩し承ります", "麒麟がくる", "テセウスの船", "シロでもクロでもない世界で、パンダは笑う。", "心の傷を癒すということ_(テレビドラマ)", "パパがも一度恋をした"}
	// wikiTitles := []string{"左ききのエレン", "10の秘密", "恋はつづくよどこまでも"}
	for _, v := range wikiTitles {
		url := "https://ja.wikipedia.org/wiki/" + v
		GetTvProgramInformationByURLOnGo(url)
	}
}

// サーバ側でデータ投入
func UpdateMovieInfo(year int) {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64
	query["Week.Name"] = "映画"
	if year != 0 {
		query["Year"] = strconv.Itoa(year)
	}
	tvPrograms, _ := models.GetAllTvProgram(query, fields, sortby, order, offset, limit)
	for _, tvProgram := range tvPrograms {
		url := "https://ja.wikipedia.org/wiki/" + tvProgram.(models.TvProgram).Title
		fmt.Println(tvProgram.(models.TvProgram).Title)
		GetMovieInformationByURLOnGo(url, tvProgram.(models.TvProgram))
	}
}

// カテゴリー
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
	case "探偵", "推理", "サイコスリラー", "推理アクション", "クイズ番組":
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
	case "テレビ", "連続":
		newCategory = "ホーム・ヒューマン"
	}
	return newCategory
}
