package controllers

import (
	"app/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	// "time"

	"github.com/astaxie/beego"
)

//  CommentController operations for Comment
type CommentController struct {
	beego.Controller
}

// URLMapping ...
func (c *CommentController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetNewComments", c.GetNewComments)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Show", c.Show)
	c.Mapping("SearchComment", c.SearchComment)
}

// Post ...
// @Title Post
// @Description create Comment
// @Param	body		body 	models.Comment	true		"body for Comment content"
// @Success 201 {int} models.Comment
// @Failure 403 body is empty
// @router / [post]
func (c *CommentController) Post() {
	session := c.StartSession()
	if session.Get("UserId") != nil {
		var v models.Comment
		json.Unmarshal(c.Ctx.Input.RequestBody, &v)
		if _, err := models.AddComment(&v); err == nil {
			c.Data["json"] = v
			w, _ := models.GetTvProgramById(v.TvProgramId)
			w.CountComment++
			_ = models.UpdateTvProgramById(w)
			z, _ := models.GetUserById(v.UserId)
			z.CountComment++
			_ = models.UpdateUserById(z)
		} else {
			c.Data["json"] = err.Error()
		}
		c.Redirect("/tv/tv_program/comment/"+strconv.FormatInt(v.TvProgramId, 10), 302)
	}
}

// GetOne ...
// @Title Get One
// @Description get Comment by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Comment
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CommentController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetCommentById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Comment
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Comment
// @Failure 403
// @router / [get]
func (c *CommentController) GetAll() {
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

	l, err := models.GetAllComment(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

func (c *CommentController) GetNewComments() {
	type CommentAndUser struct {
		Comments []interface{}
		Users    []models.User
	}
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 100
	var offset int64
	tvProgramID := c.Ctx.Input.Param(":id")
	topCommentId := c.Ctx.Input.Param(":top")
	sortby = append(sortby, "Created")
	order = append(order, "desc")
	query["TvProgramId"] = tvProgramID
	query["Id__gt"] = topCommentId
	l, err := models.GetAllComment(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = nil
	} else {
		c.Ctx.Output.SetStatus(200)
		var users []models.User
		for _, comment := range l {
			u, _ := models.GetUserById(comment.(models.Comment).UserId)
			users = append(users, *u)
		}
		commentAndUser := *new(CommentAndUser)
		commentAndUser.Comments = l
		commentAndUser.Users = users
		c.Data["json"] = commentAndUser
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Comment
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Comment	true		"body for Comment content"
// @Success 200 {object} models.Comment
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CommentController) Put() {
	// idStr := c.Ctx.Input.Param(":id")
	// id, _ := strconv.ParseInt(idStr, 0, 64)
	// v := models.Comment{Id: id}
	// json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	// if err := models.UpdateCommentById(&v); err == nil {
	// 	c.Data["json"] = "OK"
	// } else {
	// 	c.Data["json"] = err.Error()
	// }
	// c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Comment
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CommentController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	session := c.StartSession()
	if session.Get("UserId") != nil {
		if err := models.DeleteComment(id); err == nil {
			u, err := models.GetCommentLikeByComment(id)
			if err == nil {
				for _, value := range u {
					_ = models.DeleteCommentLike(value.Id)
				}
			}
		}
	}
	c.Redirect("/tv/user/show", 302)
}

func (c *CommentController) Show() {
	idStr := c.Ctx.Input.Param(":id")
	tvProgramID, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetTvProgramById(tvProgramID)
	if err != nil {
		c.Data["TvProgram"] = nil
	} else {
		c.Data["TvProgram"] = v
	}

	l, err := models.GetCommentByTvProgramId(tvProgramID, 200)
	if err != nil {
		c.Data["Comment"] = nil
	} else {
		c.Data["Comment"] = l
	}

	cnt := models.CountAllCommentNumByTvProgramId(tvProgramID)
	c.Data["CommentNum"] = cnt

	var users []models.User
	for _, comment := range l {
		u, err := models.GetUserById(comment.UserId)
		if err != nil {
			u = new(models.User)
		}
		users = append(users, *u)
	}
	c.Data["Users"] = users
	session := c.StartSession()
	// 閲覧数カウント
	if session.Get(tvProgramID) == nil {
		var userID int64 = 0
		if session.Get("UserId") != nil {
			userID = session.Get("UserId").(int64)
		}
		var b models.BrowsingHistory
		b = models.BrowsingHistory{
			UserId:      userID,
			TvProgramId: tvProgramID,
		}
		_, _ = models.AddBrowsingHistory(&b)

		v.CountClicked++
		_ = models.UpdateTvProgramById(v)
		session.Set(tvProgramID, true)
	}

	if session.Get("UserId") != nil {
		userID := session.Get("UserId").(int64)
		w, err := models.GetWatchingStatusByUserAndTvProgram(userID, tvProgramID)
		if err != nil {
			c.Data["WatchStatus"] = new(models.WatchingStatus)
		} else {
			c.Data["WatchStatus"] = w
		}

		x, _ := models.GetUserById(userID)
		c.Data["User"] = x

		var commentLikes []models.CommentLike
		for _, comment := range l {
			u, err := models.GetCommentLikeByCommentAndUser(comment.Id, userID)
			if err != nil {
				u = new(models.CommentLike)
			}
			commentLikes = append(commentLikes, *u)
		}
		c.Data["CommentLike"] = commentLikes

	}
	c.TplName = "comment/show.tpl"
}

func (c *CommentController) SearchComment() {

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 200
	var offset int64
	var word string
	var userName string
	type SearchWords struct {
		Word       string
		Username   string
		BeforeDate string
		BeforeTime string
		AfterDate  string
		AfterTime  string
		Limit      int64
		Sortby     string
	}
	idStr := c.Ctx.Input.Param(":id")
	tvProgramID, _ := strconv.ParseInt(idStr, 0, 64)
	s := SearchWords{
		Word:       c.GetString("word"),
		Username:   c.GetString("username"),
		BeforeDate: c.GetString("before-date"),
		BeforeTime: c.GetString("before-time"),
		AfterDate:  c.GetString("after-date"),
		AfterTime:  c.GetString("after-time"),
		Limit:      limit,
		Sortby:     c.GetString("sortby"),
	}
	// ゴミ箱ボタンを押下した場合はリダイレクト
	if s.BeforeDate == "" && s.BeforeTime == "" && s.AfterDate == "" && s.AfterTime == "" {
		c.Redirect("/tv/tv_program/comment/"+idStr, 302)
	}

	c.Data["SearchWords"] = s

	v, err := models.GetTvProgramById(tvProgramID)
	if err != nil {
		c.Data["TvProgram"] = err.Error()
	} else {
		c.Data["TvProgram"] = v
	}
	cnt := models.CountAllCommentNumByTvProgramId(tvProgramID)
	c.Data["CommentNum"] = cnt

	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	if v := c.GetString("word"); v != "" {
		v = models.ReshapeWordsA(v)
		query["Content"] = v
		word = v
	}
	if v := c.GetString("username"); v != "" {
		v = models.ReshapeWordsA(v)
		query["Username"] = v
		userName = v
	}

	if v := c.GetString("before-date"); v != "" {
		if v := c.GetString("before-time"); v != "" {
			query["BeforeTime"] = c.GetString("before-date") + " " + c.GetString("before-time")
		}
	}
	if v := c.GetString("after-date"); v != "" {
		if v := c.GetString("after-time"); v != "" {
			query["AfterTime"] = c.GetString("after-date") + " " + c.GetString("after-time")
		}
	}

	query["TvProgramId"] = strconv.FormatInt(tvProgramID, 10)

	if v := c.GetString("sortby"); v != "" {
		sortElem := v
		if sortElem == "新しい順" {
			sortby = append(sortby, "Created")
			order = append(order, "desc")
		} else if sortElem == "古い順" {
			sortby = append(sortby, "Created")
			order = append(order, "asc")
		} else if sortElem == "いいねが多い順" {
			sortby = append(sortby, "CountLike")
			order = append(order, "desc")
		}
	}

	session := c.StartSession()
	if session.Get("UserId") != nil {
		var u models.SearchHistory
		searchWords := []string{word, userName}
		searchWord := strings.Join(searchWords, ",")
		searchWord = strings.Trim(searchWord, ",")
		searchWord = strings.Replace(searchWord, ",,", ",", 1)
		u = models.SearchHistory{
			UserId: session.Get("UserId").(int64),
			Word:   searchWord,
			Limit:  s.Limit,
			Hour:   s.BeforeDate + " " + s.BeforeTime + "," + s.AfterDate + " " + s.AfterTime,
			Sortby: s.Sortby,
			Item:   "comment",
		}
		_, _ = models.AddSearchHistory(&u)
	}

	l, err := models.SearchComment(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["Comment"] = nil
	} else {
		c.Data["Comment"] = l
	}
	var users []models.User
	for _, comment := range l {
		u, _ := models.GetUserById(comment.(models.Comment).UserId)
		users = append(users, *u)
	}
	c.Data["Users"] = users

	if session.Get("UserId") != nil {
		userID := session.Get("UserId").(int64)
		w, err := models.GetWatchingStatusByUserAndTvProgram(userID, tvProgramID)
		if err != nil {
			c.Data["WatchStatus"] = new(models.WatchingStatus)
		} else {
			c.Data["WatchStatus"] = w
		}

		x, _ := models.GetUserById(userID)
		c.Data["User"] = x

		var commentLikes []models.CommentLike
		for _, comment := range l {
			u, err := models.GetCommentLikeByCommentAndUser(comment.(models.Comment).Id, userID)
			if err != nil {
				u = new(models.CommentLike)
			}
			commentLikes = append(commentLikes, *u)
		}
		c.Data["CommentLike"] = commentLikes
	}
	c.TplName = "comment/show.tpl"
}
