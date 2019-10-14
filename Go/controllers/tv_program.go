package controllers

import (
"app/models"
// "encoding/json"
"errors"
"strconv"
"strings"
"regexp"
// "reflect"
"time"
"fmt"

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
	// c.Mapping("Create", c.Create)
	c.Mapping("CreatePage", c.CreatePage)
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
	if session.Get("UserId")==nil{
		fmt.Println("you are not user, so permission denyed.")
		// c.Redirect("/tv/tv_program/index", 302)
	} else {
		year,_ := c.GetInt("year")
		rep := regexp.MustCompile(`\(.+\)`)
		season := *new(models.Season)
		season.Name = rep.ReplaceAllString(c.GetString("season"), "")
		week := *new(models.Week)
		week.Name = c.GetString("week")
		var hour float64 = 100
		if c.GetString("hour") != "指定なし" {
			hour_string := c.GetString("hour")
			hour_string = strings.Replace(hour_string, ":00", "", -1)
			hour_string = strings.Replace(hour_string, ":30", ".5", -1)
			hour,_ = strconv.ParseFloat(hour_string, 32)
		}
		movie_url := c.GetString("MovieUrl")
		if !strings.Contains(movie_url, "embed"){
			movie_url = strings.Replace(movie_url, "watch?v=", "embed/", -1)
		}
		image_url := c.GetString("IconUrl")
		if image_url == "" {
			image_url = "http://hankodeasobu.com/wp-content/uploads/animals_02.png"
		}
		var v models.TvProgram
		v = models.TvProgram{
			Title: c.GetString("title"),
			Content: c.GetString("content"),
			ImageUrl: image_url,
			ImageUrlReference: c.GetString("ImageUrlReference"),
			MovieUrl: movie_url,
			MovieUrlReference: c.GetString("MovieUrlReference"),
			Cast: strings.Replace(c.GetString("cast"), "　", "", -1),
			Category: strings.Join(c.GetStrings("category"), "、"),
			Dramatist: strings.Replace(c.GetString("dramatist"), "　", "", -1),
			Supervisor: strings.Replace(c.GetString("supervisor"), "　", "", -1),
			Director: strings.Replace(c.GetString("director"), "　", "", -1),
			Production: c.GetString("production"),
			Year: year,
			Season: &season,
			Week: &week,
			Hour: float32(hour),
			Themesong: c.GetString("themesong"),
			CreateUserId: session.Get("UserId").(int64),
		}

		if _, err := models.AddTvProgram(&v); err == nil {
			c.Redirect("/tv/tv_program/comment/"+ strconv.FormatInt(v.Id, 10), 302)
		} else {
			c.Data["TvProgram"] = v
			c.TplName = "tv_program/create.tpl"
		}
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
	var limit int64
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
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	fmt.Println(id)
	// v := models.TvProgram{Id: id}
	year,_ := c.GetInt("year")
	rep := regexp.MustCompile(`\(.+\)`)
	season := *new(models.Season)
	season.Name = rep.ReplaceAllString(c.GetString("season"), "")
	week := *new(models.Week)
	week.Name = c.GetString("week")
	var hour float64 = 100
	if c.GetString("hour") != "指定なし" {
		hour_string := c.GetString("hour")
		hour_string = strings.Replace(hour_string, ":00", "", -1)
		hour_string = strings.Replace(hour_string, ":30", ".5", -1)
		hour,_ = strconv.ParseFloat(hour_string, 32)
	}
	movie_url := c.GetString("MovieUrl")
	if !strings.Contains(movie_url, "embed"){
		movie_url = strings.Replace(movie_url, "watch?v=", "embed/", -1)
	}
	image_url := c.GetString("IconUrl")
	if image_url == "" {
		image_url = "http://hankodeasobu.com/wp-content/uploads/animals_02.png"
	}
	var v models.TvProgram
	v = models.TvProgram{
		Id: id,
		Title: c.GetString("title"),
		Content: c.GetString("content"),
		ImageUrl: image_url,
		ImageUrlReference: c.GetString("ImageUrlReference"),
		MovieUrl: movie_url,
		MovieUrlReference: c.GetString("MovieUrlReference"),
		Cast: strings.Replace(c.GetString("cast"), "　", "", -1),
		Category: strings.Join(c.GetStrings("category"), "、"),
		Dramatist: strings.Replace(c.GetString("dramatist"), "　", "", -1),
		Supervisor: strings.Replace(c.GetString("supervisor"), "　", "", -1),
		Director: strings.Replace(c.GetString("director"), "　", "", -1),
		Production: c.GetString("production"),
		Year: year,
		Season: &season,
		Week: &week,
		Hour: float32(hour),
		Themesong: c.GetString("themesong"),
		CreateUserId: session.Get("UserId").(int64),
	}
	// json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateTvProgramById(&v); err == nil {
		c.Data["json"] = "OK"
		c.Redirect("/tv/tv_program/comment/"+idStr, 302)
	} else {
		c.Data["json"] = err.Error()
		c.Redirect("/tv/tv_program/edit/"+idStr, 302)
	}
	c.ServeJSON()
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
	if err := models.DeleteTvProgram(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.Redirect("/tv/tv_program/index", 302)
	c.ServeJSON()
}

func (c *TvProgramController) Index() {
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
	l, _ := models.GetAllTvProgram(query, fields, sortby, order, offset, limit)
	c.Data["TvProgram"] = l

	session := c.StartSession()
	if session.Get("UserId")==nil{
		fmt.Println("you are not user, so your tv_Like break.")
	}	else {
		user_id := session.Get("UserId").(int64)
		var ratings []models.WatchingStatus
		for _, tv_program := range l {
			r, err := models.GetWatchingStatusByUserAndTvProgram(user_id, tv_program.(models.TvProgram).Id)
			if err != nil {
				ratings = append(ratings, *new(models.WatchingStatus))
			} else {
				ratings = append(ratings, *r)
			}
			c.Data["WatchStatus"] = ratings
			v,_ := models.GetUserById(user_id)
			c.Data["User"] = v
		}
	}
	c.TplName = "tv_program/index.tpl"
}

func (c *TvProgramController) EditPage() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, _ := models.GetTvProgramById(id)
	c.Data["TvProgram"] = v
	c.TplName = "tv_program/edit.tpl"
}

func (c *TvProgramController) Get() {
	session := c.StartSession()
	c.Data["UserId"] = session.Get("UserId")
	var fields []string
	var sortby []string
	var order []string
	var limit int64 = 30
	var offset int64
	var query = make(map[string]string)
	sortby = append(sortby, "Hour")
	order = append(order, "asc")
	query["Year"] = strconv.Itoa(time.Now().Year())
	query["Season"] = models.GetOnairSeason()
	week := [7]string{"月","火","水","木","金","土","日"}
	week_name := [7]string{"mon","tue","wed","thu","fri","sat","san"}
	for i, v := range week {
		query["Week.Name"] = v
		w,err := models.GetAllTvProgram(query, fields, sortby, order, offset, limit)
		if err != nil {
			c.Data["TvProgram_"+week_name[i]] = nil
		} else {
			c.Data["TvProgram_"+week_name[i]] = w
		}
	}
	c.TplName = "tv_program/top-page.tpl"
}

func (c *TvProgramController) Search() {
	str := c.GetString("search_word")
	v, _ := models.SearchTvProgramAll(str)
	c.Data["TvProgram"] = v
	session := c.StartSession()
	if session.Get("UserId")!=nil{
		var u models.SearchHistory
		str = strings.Replace(str, "　", " ", -1)
		u = models.SearchHistory {
			UserId: session.Get("UserId").(int64),
			Word: strings.Replace(str, " ", "、", -1),
		}
		_,_ = models.AddSearchHistory(&u)
	}
	c.TplName = "tv_program/index.tpl"
}

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
		Title string
		Staff string
		Themesong string
		Year string
		Week string
		Hour string
		Season string
		Category string
		Limit int64
		Sortby string
	}

	if v := c.GetString("title"); v != "" {
		v = strings.Replace(v, "　", " ", -1)
		title = strings.Split(v, " ")
		query["Title"] = title
	}
	if v := c.GetString("staff"); v != "" {
		v = strings.Replace(v, "　", " ", -1)
		staff = strings.Split(v, " ")
		query["Staff"] = staff
	}
	if v := c.GetString("themesong"); v != "" {
		v = strings.Replace(v, "　", "", -1)
		themesong = strings.Split(v, " ")
		query["Themesong"] = themesong
	}
	if v := c.GetStrings("year"); v != nil {
		query["Year"] = v
	}
	if v := c.GetStrings("week"); v != nil {
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
		rep := regexp.MustCompile(`\(.+\)`)
		for _, value := range v {
			value = rep.ReplaceAllString(value, "")
			query["Season"] = append(query["Season"], value)
		}
	}
	if v := c.GetStrings("category"); v != nil {
		query["Category"] = v
	}
	if v := c.GetString("sortby"); v != "" {
		sort_elem := v
		if sort_elem == "新しい順" {
			sortby = append(sortby, "Year")
			sortby = append(sortby, "Season__Id")
			sortby = append(sortby, "Week__Id")
			sortby = append(sortby, "Hour")
			order = append(order, "desc")
			order = append(order, "desc")
			order = append(order, "asc")
			order = append(order, "asc")
		} else if sort_elem == "古い順" {
			sortby = append(sortby, "Year")
			sortby = append(sortby, "Season__Id")
			sortby = append(sortby, "Week__Id")
			sortby = append(sortby, "Hour")
			order = append(order, "asc")
			order = append(order, "asc")
			order = append(order, "asc")
			order = append(order, "asc")
		} else if sort_elem == "タイトル順" {
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
		} else if sort_elem == "アーティスト順" {
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
		} else if sort_elem == "閲覧数が多い順" {
			sortby = append(sortby, "CountClicked")
			order = append(order, "desc")
		} else if sort_elem == "見た人が多い順" {
			sortby = append(sortby, "CountWatched")
			order = append(order, "desc")
		} else if sort_elem == "見たい人が多い順" {
			sortby = append(sortby, "CountWantToWatch")
			order = append(order, "desc")
		}
	}

	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}

	var s SearchWords

	s = SearchWords {
		Title: strings.Join(title, "、"),
		Staff: strings.Join(staff, "、"),
		Themesong: strings.Join(themesong, "、"),
		Year: strings.Join(c.GetStrings("year"), "、"),
		Week: strings.Join(c.GetStrings("week"), "、"),
		Hour: strings.Join(c.GetStrings("hour"), "、"),
		Season: strings.Join(c.GetStrings("season"), "、"),
		Category: strings.Join(c.GetStrings("category"), "、"),
		Limit: limit,
		Sortby: c.GetString("sortby"),
	}
	c.Data["SearchWords"] = s
	session := c.StartSession()
	if session.Get("UserId")!=nil{
		var u models.SearchHistory
		u = models.SearchHistory {
			UserId: session.Get("UserId").(int64),
			Word: s.Title + " " + s.Staff + " " + s.Themesong,
			Year: s.Year,
			Season: s.Season,
			Week: s.Week,
			Hour: s.Hour,
			Category: s.Category,
			Limit: s.Limit,
			Sortby:s.Sortby,
		}
		_,_ = models.AddSearchHistory(&u)
	}

	// fmt.Println(fields, limit, offset, sortby, order, query)
	l, _ := models.SearchTvProgram(query, fields, sortby, order, offset, limit)
	c.Data["TvProgram"] = l
	// session := c.StartSession()
	if session.Get("UserId")==nil{
		fmt.Println("you are not user, so your tv_Like break.")
	}	else {
		user_id := session.Get("UserId").(int64)
		var ratings []models.WatchingStatus
		for _, tv_program := range l {
			r, err := models.GetWatchingStatusByUserAndTvProgram(user_id, tv_program.(models.TvProgram).Id)
			if err != nil {
				ratings = append(ratings, *new(models.WatchingStatus))
			} else {
				ratings = append(ratings, *r)
			}
			c.Data["WatchStatus"] = ratings
			v,_ := models.GetUserById(user_id)
			c.Data["User"] = v
		}
	}
	c.TplName = "tv_program/index.tpl"
}


func (c *TvProgramController) CreatePage() {
	session := c.StartSession()
	if session.Get("UserId")==nil{
		fmt.Println("you are not user, so your permission denyed.")
		c.Redirect("/", 302)
	}	else {
		c.TplName = "tv_program/create.tpl"
	}
}

func (c *TvProgramController) Create() {

// 		year, _ := c.GetInt("year")
// 		season := new(models.Season)
// 		season.Name = c.GetString("season")
// 		fmt.Println(sseason)
// 		var v models.TvProgram
// 		v = models.TvProgram{
// 			Title: c.GetString("title"),
// 			Content: c.GetString("content"),
// 			ImageUrl: c.GetString("ImageUrl"),
// 			ImageUrlReference: c.GetString("ImageUrlReference"),
// 			MovieUrl: c.GetString("MovieUrl"),
// 			MovieUrlReference: c.GetString("MovieUrlReference"),
// 			Cast: c.GetString("cast"),
// 			Category: c.GetString("category"),
// 			Dramatist: c.GetString("dramatist"),
// 			Supervisor: c.GetString("supervisor"),
// 			Director: c.GetString("director"),
// 			Year: year,
// 			Season: season,
// 			Themesong: c.GetString("themesong"),
// 			Timezone: c.GetString("timezone"),
// 		}
// 		fmt.Println(v)
// if _, err := models.AddTvProgram(&v); err == nil {
// 			c.Ctx.Output.SetStatus(201)
// 			c.Data["json"] = v
// 		c.Redirect("/tv/tv_program/comment/"+strconv.FormatInt(v.Id, 10), 302)
// 		} else {
// 			c.Data["json"] = err.Error()
// 	c.TplName = "tv_program/create.tpl"
// 		}
// 		c.ServeJSON()
}