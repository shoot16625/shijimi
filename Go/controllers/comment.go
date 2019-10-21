package controllers

import (
	"app/models"
	"encoding/json"
	"errors"
	"fmt"
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
	if session.Get("UserId") == nil {
		fmt.Println("you are not user, so permission denyed.")
		c.Redirect("/tv/tv_program/comment/"+c.GetString("TvProgramId"), 302)
	} else {
		fmt.Println("permission clear.")
		var v models.Comment
		// fmt.Println(c.Ctx.Input.RequestBody)
		// if true {
		json.Unmarshal(c.Ctx.Input.RequestBody, &v)

		// } else {
		// 		idStr := c.GetString("TvProgramId")
		// 		id, _ := strconv.ParseInt(idStr, 0, 64)
		// 		v = models.Comment{
		// 			// Content: models.ReplacNewLine(c.GetString("content"), "<br>"),
		// 			Content: c.GetString("content"),
		// 			TvProgramId: id,
		// 			UserId: session.Get("UserId").(int64),
		// 		}}
		// fmt.Println(v)
		if _, err := models.AddComment(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
		// fmt.Println(c.Data["json"])
		c.Redirect("/tv/tv_program/comment/"+c.GetString("TvProgramId"), 302)
		c.ServeJSON()
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

// Put ...
// @Title Put
// @Description update the Comment
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Comment	true		"body for Comment content"
// @Success 200 {object} models.Comment
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CommentController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Comment{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateCommentById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
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
	if err := models.DeleteComment(id); err == nil {
		c.Data["json"] = "OK"
		u, err := models.GetCommentLikeByComment(id)
		if err == nil {
			for _, value := range u {
				_ = models.DeleteCommentLike(value.Id)
			}
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.Redirect("/tv/user/show", 302)
	c.ServeJSON()
}

func (c *CommentController) Show() {

	idStr := c.Ctx.Input.Param(":id")
	tvProgramID, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetTvProgramById(tvProgramID)
	if err != nil {
		c.Data["TvProgram"] = err.Error()
	} else {
		c.Data["TvProgram"] = v
	}

	l, err := models.GetCommentByTvprogramId(tvProgramID)
	if err != nil {
		c.Data["Comment"] = nil
	} else {
		c.Data["Comment"] = l
	}

	var users []models.User
	for _, comment := range l {
		u, _ := models.GetUserById(comment.UserId)
		users = append(users, *u)
	}
	c.Data["Users"] = users
	session := c.StartSession()
	// 閲覧数カウント
	if session.Get(tvProgramID) == nil {
		fmt.Println("first tv click")
		if session.Get("UserId") != nil {
			userID := session.Get("UserId").(int64)
			var b models.BrowsingHistory
			b = models.BrowsingHistory{
				UserId:      userID,
				TvProgramId: tvProgramID,
			}
			_, err = models.AddBrowsingHistory(&b)
			if err == nil {
				fmt.Println("browsing_history", b)
			}
		}
		v.CountClicked++
		_ = models.UpdateTvProgramById(v)
		session.Set(tvProgramID, true)
	}

	if session.Get("UserId") == nil {
		fmt.Println("you are not user, so your tv_Like break.")
	} else {
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

	idStr := c.Ctx.Input.Param(":id")
	tvProgramID, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetTvProgramById(tvProgramID)
	if err != nil {
		c.Data["TvProgram"] = err.Error()
	} else {
		c.Data["TvProgram"] = v
	}

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 100
	var offset int64
	var word string
	type SearchWords struct {
		Word   string
		Limit  int64
		Sortby string
	}

	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}

	if v := c.GetString("word"); v != "" {
		query["Content__icontains"] = v
		word = strings.Replace(v, "　", " ", -1)
		word = strings.Replace(word, " ", "、", -1)
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

	var s SearchWords
	s = SearchWords{
		Word:   c.GetString("word"),
		Limit:  limit,
		Sortby: c.GetString("sortby"),
	}
	c.Data["SearchWords"] = s
	session := c.StartSession()
	if session.Get("UserId") != nil {
		var u models.SearchHistory
		u = models.SearchHistory{
			UserId: session.Get("UserId").(int64),
			Word:   word,
			Limit:  s.Limit,
			Sortby: s.Sortby,
			Item:   "comment",
		}
		_, _ = models.AddSearchHistory(&u)
	}

	l, err := models.GetAllComment(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["Comment"] = nil
	} else {
		c.Data["Comment"] = l
	}
	// fmt.Println(l[0].(models.Comment).Id)
	var users []models.User
	for _, comment := range l {
		u, _ := models.GetUserById(comment.(models.Comment).UserId)
		users = append(users, *u)
	}
	c.Data["Users"] = users
	// session := c.StartSession()
	// 閲覧数カウント
	if session.Get(tvProgramID) == nil {
		fmt.Println("first tv click")
		if session.Get("UserId") != nil {
			userID := session.Get("UserId").(int64)
			var b models.BrowsingHistory
			b = models.BrowsingHistory{
				UserId:      userID,
				TvProgramId: tvProgramID,
			}
			_, err = models.AddBrowsingHistory(&b)
			if err == nil {
				fmt.Println("browsing_history", b)
			}
		}
		v.CountClicked++
		_ = models.UpdateTvProgramById(v)
		session.Set(tvProgramID, true)
	}

	if session.Get("UserId") == nil {
		fmt.Println("you are not user, so your tv_Like break.")
	} else {
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
