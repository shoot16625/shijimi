package controllers

import (
	"app/db"
	"app/models"
	"math/rand"

	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

//  TvProgramController operations for TvProgram
type TvProgramController struct {
	beego.Controller
}

// URLMapping ...
func (c *TvProgramController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Index", c.Index)
	c.Mapping("EditPage", c.EditPage)
	c.Mapping("Search", c.Search)
	c.Mapping("SearchTvProgram", c.SearchTvProgram)
	c.Mapping("CreatePage", c.CreatePage)
	c.Mapping("CreateTvProgramAuto", c.CreateTvProgramAuto)
	c.Mapping("GetTvProgramWikiInfo", c.GetTvProgramWikiInfo)
	c.Mapping("GetMovieWikiInfo", c.GetMovieWikiInfo)
}

// Post ...
// @Title Post
// @Description create TvProgram
// @Param	body		body 	models.TvProgram	true		"body for TvProgram content"
// @Success 201 {int} models.TvProgram
// @Failure 403 body is empty
// @router / [post]
func (c *TvProgramController) Post() {
	session := c.StartSession()
	if session.Get("UserId") == nil {
		c.Redirect("/", 302)
		return
	}
	// ログインユーザのみ作成（html側でも縛りあり）
	year, _ := c.GetInt("year")
	season := *new(models.Season)
	season.Name = models.RegexpWords(c.GetString("season"), `\(.+\)`, "")
	week := *new(models.Week)
	week.Name = c.GetString("week")
	var hour float64 = 100
	if c.GetString("hour") != "指定なし" {
		hourString := c.GetString("hour")
		hourString = strings.Replace(hourString, ":00", "", -1)
		hourString = strings.Replace(hourString, ":30", ".5", -1)
		hour, _ = strconv.ParseFloat(hourString, 32)
	}
	movieURL := models.ReshapeMovieURL(c.GetString("MovieURL"))
	if movieURL == "" {
		movieURL = db.GetYoutubeURL(c.GetString("title"))
	}
	imageURL := models.CheckImageURL(c.GetString("ImageURL"), c.GetString("title"))
	imageURLReference := models.ReshapeImageURLReference(imageURL)

	var v models.TvProgram
	v = models.TvProgram{
		Title:             c.GetString("title"),
		Content:           c.GetString("content"),
		ImageUrl:          imageURL,
		ImageUrlReference: imageURLReference,
		MovieUrl:          movieURL,
		WikiReference:     c.GetString("WikiReference"),
		Cast:              models.RegexpWords(c.GetString("cast"), "、|，|　", ","),
		Category:          strings.Join(c.GetStrings("category"), ","),
		Dramatist:         models.RegexpWords(c.GetString("dramatist"), "、|，|　", ","),
		Supervisor:        models.RegexpWords(c.GetString("supervisor"), "、|，|　", ","),
		Director:          models.RegexpWords(c.GetString("director"), "、|，|　", ","),
		Production:        models.RegexpWords(c.GetString("production"), "、|，|　| ", ","),
		Year:              year,
		Star:              5,
		Season:            &season,
		Week:              &week,
		Hour:              float32(hour),
		Themesong:         models.RegexpWords(c.GetString("themesong"), "、|，|　", ","),
		CreateUserId:      session.Get("UserId").(int64),
	}
	if _, err := models.AddTvProgram(&v); err == nil {
		w, _ := models.GetUserById(v.CreateUserId)
		w.CountEditTvProgram++
		_ = models.UpdateUserById(w)
		c.Redirect("/tv/tv_program/comment/"+strconv.FormatInt(v.Id, 10), 302)
	} else {
		c.Data["TvProgram"] = v
		c.Data["GetWikiInfo"] = false
		v, _ := models.GetUserById(session.Get("UserId").(int64))
		c.Data["User"] = v
		c.TplName = "tv_program/create.tpl"
	}
}

// GetOne ...
// @Title Get One
// @Description get TvProgram by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.TvProgram
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TvProgramController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetTvProgramById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get TvProgram
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.TvProgram
// @Failure 403
// @router / [get]
func (c *TvProgramController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}
	sortby = append(sortby, "Id")
	order = append(order, "desc")

	l, err := models.GetAllTvProgram(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the TvProgram
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TvProgram	true		"body for TvProgram content"
// @Success 200 {object} models.TvProgram
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TvProgramController) Put() {
	session := c.StartSession()
	if session.Get("UserId") == nil {
		c.Redirect("/", 302)
		return
	}
	// if session.Get("UserId") != nil {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	year, _ := c.GetInt("year")
	season := *new(models.Season)
	season.Name = models.RegexpWords(c.GetString("season"), `\(.+\)`, "")
	week := *new(models.Week)
	week.Name = c.GetString("week")
	var hour float64 = 100
	if c.GetString("hour") != "指定なし" {
		hourString := c.GetString("hour")
		hourString = strings.Replace(hourString, ":00", "", -1)
		hourString = strings.Replace(hourString, ":30", ".5", -1)
		hour, _ = strconv.ParseFloat(hourString, 32)
	}
	movieURL := models.ReshapeMovieURL(c.GetString("MovieURL"))
	imageURL := models.CheckImageURL(c.GetString("ImageURL"), c.GetString("title"))
	imageURLReference := models.ReshapeImageURLReference(imageURL)
	oldTvInfo, _ := models.GetTvProgramById(id)
	v := *oldTvInfo
	v.Title = c.GetString("title")
	v.Content = c.GetString("content")
	v.ImageUrl = imageURL
	v.ImageUrlReference = imageURLReference
	v.MovieUrl = movieURL
	v.WikiReference = c.GetString("WikiReference")
	v.Cast = c.GetString("cast")
	v.Category = strings.Join(c.GetStrings("category"), ",")
	v.Dramatist = c.GetString("dramatist")
	v.Supervisor = c.GetString("supervisor")
	v.Director = c.GetString("director")
	v.Production = c.GetString("production")
	v.Season = &season
	v.Week = &week
	v.Year = year
	v.Hour = float32(hour)
	v.Themesong = c.GetString("themesong")
	v.CountUpdated++

	if err := models.UpdateTvProgramById(&v); err == nil {
		userID := session.Get("UserId").(int64)
		w := models.TvProgramUpdateHistory{
			UserId:      userID,
			TvProgramId: id,
		}
		_, _ = models.AddTvProgramUpdateHistory(&w)
		z, _ := models.GetUserById(userID)
		z.CountEditTvProgram++
		_ = models.UpdateUserById(z)
		c.Redirect("/tv/tv_program/comment/"+idStr, 302)
	} else {
		c.Data["TvProgram"] = v
		// なぜか効かない
		c.Data["TitleFlag"] = 1
		c.Redirect("/tv/tv_program/edit/"+idStr, 302)
	}
}

// Delete ...
// @Title Delete
// @Description delete the TvProgram
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TvProgramController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	session := c.StartSession()
	// 管理者のみ削除できるようjs側で制御
	if session.Get("UserId") != nil {
		if err := models.DeleteTvProgram(id); err == nil {
			models.DeleteWatchingStatusByTvProgramId(id)
			models.DeleteBrowsingHistoryByTvProgramId(id)
			models.DeleteTvProgramUpdateHistoryByTvProgramId(id)
			c.Data["Status"] = "テレビ番組を削除しました"
		} else {
			c.Data["Status"] = "テレビ番組の削除に失敗しました"
		}
	} else {
		c.Data["Status"] = "テレビ番組の削除に失敗しました"
	}
	var Info struct {
		CntUsers      int64
		CntTvPrograms int64
	}
	Info.CntUsers = models.GetUserCount()
	Info.CntTvPrograms = models.GetTvProgramCount()
	c.Data["Info"] = Info
	c.TplName = "user/logout.tpl"
}

// 番組一覧・おすすめページ
func (c *TvProgramController) Index() {
	var l []interface{}
	session := c.StartSession()
	// おすすめTV取得
	if session.Get("UserId") != nil {
		userID := session.Get("UserId").(int64)
		l = models.GetRecommendTvProgramsByUser(userID)
	}
	// fmt.Println(l)
	// おすすめTV取得できなかったら，放送中順
	if l == nil {

		var fields []string
		var sortby []string
		var order []string
		var query = make(map[string]string)
		var limit int64 = 100
		var offset int64

		sortby = append(sortby, "Year")
		sortby = append(sortby, "Season__Id")
		sortby = append(sortby, "Week__Id")
		sortby = append(sortby, "Hour")
		order = append(order, "desc")
		order = append(order, "desc")
		order = append(order, "asc")
		order = append(order, "asc")
		query["Year"] = strconv.Itoa(time.Now().Year())
		season := models.GetOnAirSeason()
		v, _ := models.GetSeasonByName(season)
		query["Season__Id__lte"] = strconv.Itoa(v.Id)
		l, _ = models.GetAllTvProgram(query, fields, sortby, order, offset, limit)
		if len(l) < int(limit) {
			query = make(map[string]string)
			query["Year__lt"] = strconv.Itoa(time.Now().Year())
			m, err := models.GetAllTvProgram(query, fields, sortby, order, offset, limit-int64(len(l)))
			if err == nil {
				l = append(l, m...)
			}
		}
	}
	c.Data["TvProgram"] = l

	if session.Get("UserId") != nil {
		userID := session.Get("UserId").(int64)
		var ratings []models.WatchingStatus
		for _, tvProgram := range l {
			r, err := models.GetWatchingStatusByUserAndTvProgram(userID, tvProgram.(models.TvProgram).Id)
			if err != nil {
				ratings = append(ratings, *new(models.WatchingStatus))
			} else {
				ratings = append(ratings, *r)
			}
		}
		c.Data["WatchStatus"] = ratings
		v, _ := models.GetUserById(userID)
		c.Data["User"] = v
	}
	// 分析部分
	t := time.Now()
	yesterday := t.Add(-24 * time.Hour)
	if browsingLog, err := models.GetTopBrowsingHistory(yesterday.Format("2006-01-02 15:04:05")); err == nil {
		c.Data["ViewTvProgramIn24"] = browsingLog
	}
	if goodStarTvProgram, err := models.GetTopStarPoint(); err == nil {
		c.Data["goodStarTvProgramOnAir"] = goodStarTvProgram
	}

	c.TplName = "tv_program/index.tpl"
}

// 番組登録情報の編集ページへ移動
func (c *TvProgramController) EditPage() {
	session := c.StartSession()
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if session.Get("UserId") != nil {
		v, err := models.GetTvProgramById(id)
		if err != nil {
			c.Redirect("/", 302)
			return
		}
		c.Data["TvProgram"] = v
		userID := session.Get("UserId").(int64)
		w, _ := models.GetUserById(userID)
		c.Data["User"] = w
		c.TplName = "tv_program/edit.tpl"
	} else {
		c.Redirect("/tv/tv_program/comment/"+idStr, 302)
	}
}

// トップページの処理
func (c *TvProgramController) Get() {
	session := c.StartSession()
	// c.Data["UserId"] = session.Get("UserId")
	if session.Get("UserId") != nil {
		userID := session.Get("UserId").(int64)
		w, _ := models.GetUserById(userID)
		c.Data["User"] = w
	}
	var fields []string
	var sortby []string
	var order []string
	var limit int64 = 20
	var offset int64
	var query = make(map[string]string)
	sortby = append(sortby, "Hour")
	order = append(order, "asc")
	// 現在放送中のドラマ
	query["Year"] = strconv.Itoa(time.Now().Year())
	query["Season"] = models.GetOnAirSeason()
	week := [7]string{"月", "火", "水", "木", "金", "土", "日"}
	weekName := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	for i, v := range week {
		query["Week.Name"] = v
		w, err := models.GetAllTvProgram(query, fields, sortby, order, offset, limit)
		if err == nil {
			c.Data["TvProgram"+weekName[i]] = w
		}
	}
	// 現在放送中の映画
	limit = 200
	query["Week.Name"] = "映画"
	w, err := models.GetAllTvProgram(query, fields, sortby, order, offset, limit)
	if err == nil && len(w) >= 1 {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(w))
		ran := 15
		if len(w) <= n+ran {
			n = len(w) - ran
		}
		if n <= 0 {
			n = 0
		}
		if ran > len(w) {
			ran = len(w)
		}
		c.Data["TvProgramMovie"] = w[n : n+ran]
	}

	// 公開予定の映画
	limit = 400
	query = make(map[string]string)
	query["Week.Name"] = "映画"
	year := time.Now().Year()
	season := models.GetOnAirSeason()
	if season == "秋" {
		query["Year"] = strconv.Itoa(year + 1)
	} else {
		query["Year"] = strconv.Itoa(year)
		v, _ := models.GetSeasonByName(season)
		query["Season__Id__gt"] = strconv.Itoa(v.Id)
	}
	w, err = models.GetAllTvProgram(query, fields, sortby, order, offset, limit)
	if err == nil && len(w) >= 1 {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(w))
		ran := 15
		if len(w) <= n+ran {
			n = len(w) - ran
		}
		if n <= 0 {
			n = 0
		}
		if ran > len(w) {
			ran = len(w)
		}
		c.Data["TvProgramNewMovie"] = w[n : n+ran]
	}
	c.TplName = "tv_program/top_page.tpl"
}

// ツールバーの検索機能
func (c *TvProgramController) Search() {
	str := c.GetString("search-word")
	l, _ := models.SearchTvProgramAll(str)
	c.Data["TvProgram"] = l
	session := c.StartSession()
	if session.Get("UserId") != nil {
		userID := session.Get("UserId").(int64)
		var u models.SearchHistory
		str = strings.Replace(str, "　", ",", -1)
		str = strings.Replace(str, " ", ",", -1)
		u = models.SearchHistory{
			UserId: userID,
			Word:   str,
			Limit:  100,
			Item:   "tv",
		}
		_, _ = models.AddSearchHistory(&u)

		var ratings []models.WatchingStatus
		for _, tvProgram := range l {
			r, err := models.GetWatchingStatusByUserAndTvProgram(userID, tvProgram.Id)
			if err != nil {
				ratings = append(ratings, *new(models.WatchingStatus))
			} else {
				ratings = append(ratings, *r)
			}
			c.Data["WatchStatus"] = ratings
			v, _ := models.GetUserById(userID)
			c.Data["User"] = v
		}
	}
	// 分析部分
	t := time.Now()
	yesterday := t.Add(-24 * time.Hour)
	if browsingLog, err := models.GetTopBrowsingHistory(yesterday.Format("2006-01-02 15:04:05")); err == nil {
		c.Data["ViewTvProgramIn24"] = browsingLog
	}
	if goodStarTvProgram, err := models.GetTopStarPoint(); err == nil {
		c.Data["goodStarTvProgramOnAir"] = goodStarTvProgram
	}

	c.TplName = "tv_program/index.tpl"
}

// 詳細検索機能
func (c *TvProgramController) SearchTvProgram() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string][]string)
	var limit int64 = 100
	var offset int64
	var title []string
	var staff []string
	var themesong []string

	type SearchWords struct {
		Title     string
		Staff     string
		Themesong string
		Year      string
		Week      string
		Hour      string
		Season    string
		Category  string
		Limit     int64
		Sortby    string
	}

	if v := c.GetString("title"); v != "" {
		v = models.ReshapeWordsA(v)
		title = strings.Split(v, ",")
		query["Title"] = title
	}
	if v := c.GetString("staff"); v != "" {
		v = models.ReshapeWordsA(v)
		staff = strings.Split(v, ",")
		query["Staff"] = staff
	}
	if v := c.GetString("themesong"); v != "" {
		v = models.ReshapeWordsA(v)
		themesong = strings.Split(v, ",")
		query["Themesong"] = themesong
	}
	if v := c.GetStrings("year"); v != nil {
		query["Year"] = v
	}
	if v := c.GetStrings("week"); v != nil {
		for _, value := range v {
			if strings.Contains(value, "映画以外") {
				dramas := []string{"月", "火", "水", "木", "金", "土", "日"}
				for _, w := range dramas {
					v = append(v, w)
				}
				break
			}
		}
		query["Week"] = v
	}
	if v := c.GetStrings("hour"); v != nil {
		for _, value := range v {
			value = strings.Replace(value, ":00", "", -1)
			value = strings.Replace(value, ":30", ".5", -1)
			query["Hour"] = append(query["Hour"], value)
		}
	}
	if v := c.GetStrings("season"); v != nil {
		for _, value := range v {
			value = models.RegexpWords(value, `\(.+\)`, "")
			query["Season"] = append(query["Season"], value)
		}
	}
	if v := c.GetStrings("category"); v != nil {
		query["Category"] = v
	}
	if v := c.GetString("sortby"); v != "" {
		sortElem := v
		if sortElem == "新しい順" {
			sortby = append(sortby, "Year")
			sortby = append(sortby, "Season__Id")
			sortby = append(sortby, "Week__Id")
			sortby = append(sortby, "Hour")
			order = append(order, "desc")
			order = append(order, "desc")
			order = append(order, "asc")
			order = append(order, "asc")
		} else if sortElem == "古い順" {
			sortby = append(sortby, "Year")
			sortby = append(sortby, "Season__Id")
			sortby = append(sortby, "Week__Id")
			sortby = append(sortby, "Hour")
			order = append(order, "asc")
			order = append(order, "asc")
			order = append(order, "asc")
			order = append(order, "asc")
		} else if sortElem == "タイトル順" {
			sortby = append(sortby, "Title")
			sortby = append(sortby, "Year")
			sortby = append(sortby, "Season__Id")
			sortby = append(sortby, "Week__Id")
			sortby = append(sortby, "Hour")
			order = append(order, "asc")
			order = append(order, "desc")
			order = append(order, "asc")
			order = append(order, "asc")
			order = append(order, "asc")
		} else if sortElem == "アーティスト順" {
			sortby = append(sortby, "Themesong")
			sortby = append(sortby, "Year")
			sortby = append(sortby, "Season__Id")
			sortby = append(sortby, "Week__Id")
			sortby = append(sortby, "Hour")
			order = append(order, "asc")
			order = append(order, "desc")
			order = append(order, "asc")
			order = append(order, "asc")
			order = append(order, "asc")
		} else if sortElem == "閲覧数が多い順" {
			sortby = append(sortby, "CountClicked")
			order = append(order, "desc")
		} else if sortElem == "見た人が多い順" {
			sortby = append(sortby, "CountWatched")
			order = append(order, "desc")
		} else if sortElem == "見たい人が多い順" {
			sortby = append(sortby, "CountWantToWatch")
			order = append(order, "desc")
		} else if sortElem == "評価が高い順" {
			sortby = append(sortby, "Star")
			order = append(order, "desc")
		} else if sortElem == "ツイートが多い順" {
			sortby = append(sortby, "CountComment")
			order = append(order, "desc")
		} else if sortElem == "レビューが多い順" {
			sortby = append(sortby, "CountReviewComment")
			order = append(order, "desc")
		}
	}

	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}

	var s SearchWords

	s = SearchWords{
		Title:     strings.Join(title, ","),
		Staff:     strings.Join(staff, ","),
		Themesong: strings.Join(themesong, ","),
		Year:      strings.Join(c.GetStrings("year"), ","),
		Week:      strings.Join(c.GetStrings("week"), ","),
		Hour:      strings.Join(c.GetStrings("hour"), ","),
		Season:    strings.Join(c.GetStrings("season"), ","),
		Category:  strings.Join(c.GetStrings("category"), ","),
		Limit:     limit,
		Sortby:    c.GetString("sortby"),
	}
	c.Data["SearchWords"] = s

	l, _ := models.SearchTvProgram(query, fields, sortby, order, offset, limit)
	c.Data["TvProgram"] = l
	session := c.StartSession()
	if session.Get("UserId") != nil {
		userID := session.Get("UserId").(int64)
		var u models.SearchHistory
		searchWords := []string{s.Title, s.Staff, s.Themesong}
		searchWord := strings.Join(searchWords, ",")
		searchWord = strings.Trim(searchWord, ",")
		searchWord = strings.Replace(searchWord, ",,", ",", 1)
		u = models.SearchHistory{
			UserId:   userID,
			Word:     searchWord,
			Year:     s.Year,
			Season:   s.Season,
			Week:     s.Week,
			Hour:     s.Hour,
			Category: s.Category,
			Limit:    s.Limit,
			Sortby:   s.Sortby,
			Item:     "tv",
		}
		_, _ = models.AddSearchHistory(&u)

		var ratings []models.WatchingStatus
		for _, tvProgram := range l {
			r, err := models.GetWatchingStatusByUserAndTvProgram(userID, tvProgram.(models.TvProgram).Id)
			if err != nil {
				ratings = append(ratings, *new(models.WatchingStatus))
			} else {
				ratings = append(ratings, *r)
			}
			c.Data["WatchStatus"] = ratings
			v, _ := models.GetUserById(userID)
			c.Data["User"] = v
		}
	}
	// 分析部分
	t := time.Now()
	yesterday := t.Add(-24 * time.Hour)
	if browsingLog, err := models.GetTopBrowsingHistory(yesterday.Format("2006-01-02 15:04:05")); err == nil {
		c.Data["ViewTvProgramIn24"] = browsingLog
	}
	if goodStarTvProgram, err := models.GetTopStarPoint(); err == nil {
		c.Data["goodStarTvProgramOnAir"] = goodStarTvProgram
	}

	c.TplName = "tv_program/index.tpl"
}

// 新規テレビ登録ページへ移動
func (c *TvProgramController) CreatePage() {
	session := c.StartSession()
	if session.Get("UserId") != nil {
		userID := session.Get("UserId").(int64)
		w, _ := models.GetUserById(userID)
		c.Data["User"] = w
		c.TplName = "tv_program/create.tpl"
	} else {
		c.Redirect("/", 302)
	}
}

// 番組一括登録
// dramaはその年以降のデータを探索する
// 映画はその年（３ヶ月後の年）の１年分を探索する
func (c *TvProgramController) CreateTvProgramAuto() {
	session := c.StartSession()
	if session.Get("UserId") != nil {
		userID := session.Get("UserId").(int64)
		if userID == 1 {
			w, _ := models.GetUserById(userID)
			c.Data["User"] = w
			year := time.Now().Year()
			wikis := []string{"日本のテレビドラマ一覧_(2020年代)"}
			db.AddTvProgramsInformation(wikis, year)
			if time.Now().Month() > 9 {
				year++
			}
			db.GetMovieWalkers(year, year, 1, 12)
			c.TplName = "tv_program/create_tvprogram_auto_result.tpl"
		} else {
			c.Redirect("/", 302)
		}
	} else {
		c.Redirect("/", 302)
	}
}

// ドラマ情報をwikiから取得
func (c *TvProgramController) GetTvProgramWikiInfo() {
	session := c.StartSession()
	if session.Get("UserId") == nil {
		c.Redirect("/", 302)
		return
	}
	wikiReference := c.GetString("wikiReference")
	if !strings.Contains(wikiReference, "wikipedia") {
		wikiReference = "https://ja.wikipedia.org/wiki/" + wikiReference
	}
	tvProgram := db.GetTvProgramInformationByURL(wikiReference)
	c.Data["TvProgram"] = tvProgram
	c.Data["GetWikiInfo"] = true
	userID := session.Get("UserId").(int64)
	w, _ := models.GetUserById(userID)
	c.Data["User"] = w
	c.TplName = "tv_program/create.tpl"
}

// 映画情報をwikiから取得
func (c *TvProgramController) GetMovieWikiInfo() {
	session := c.StartSession()
	if session.Get("UserId") == nil {
		c.Redirect("/", 302)
		return
	}
	wikiReference := c.GetString("wikiReference")
	if !strings.Contains(wikiReference, "wikipedia") {
		wikiReference = "https://ja.wikipedia.org/wiki/" + wikiReference
	}
	tvProgram := db.GetMovieInformationByURL(wikiReference)
	c.Data["TvProgram"] = tvProgram
	c.Data["GetWikiInfo"] = true
	userID := session.Get("UserId").(int64)
	w, _ := models.GetUserById(userID)
	c.Data["User"] = w
	c.TplName = "tv_program/create.tpl"
}
