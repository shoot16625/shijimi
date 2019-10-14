package controllers

import (
	"app/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"fmt"

	"github.com/astaxie/beego"
)

//  ReviewCommentController operations for ReviewComment
type ReviewCommentController struct {
	beego.Controller
}

// URLMapping ...
func (c *ReviewCommentController) URLMapping() {
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
// @Description create ReviewComment
// @Param	body		body 	models.ReviewComment	true		"body for ReviewComment content"
// @Success 201 {int} models.ReviewComment
// @Failure 403 body is empty
// @router / [post]
func (c *ReviewCommentController) Post() {
	var v models.ReviewComment
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	_,err := models.GetReviewCommentByUserIdAndTvProgramId(v.UserId, v.TvProgramId)
	if err != nil{
		fmt.Println("first review!!")
	if _, err := models.AddReviewComment(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get ReviewComment by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ReviewComment
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ReviewCommentController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetReviewCommentById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get ReviewComment
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ReviewComment
// @Failure 403
// @router / [get]
func (c *ReviewCommentController) GetAll() {
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

	l, err := models.GetAllReviewComment(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the ReviewComment
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ReviewComment	true		"body for ReviewComment content"
// @Success 200 {object} models.ReviewComment
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ReviewCommentController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.ReviewComment{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateReviewCommentById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the ReviewComment
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ReviewCommentController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteReviewComment(id); err == nil {
		c.Data["json"] = "OK"
		u, err := models.GetReviewCommentLikeByComment(id)
		if err == nil{
			for _, value := range u{
					_ = models.DeleteReviewCommentLike(value.Id)
			}
	} else {
		c.Data["json"] = err.Error()
	}
	} else {
		c.Data["json"] = err.Error()
	}
	c.Redirect("/tv/user/show_review", 302)
	c.ServeJSON()
}

func (c *ReviewCommentController) Show() {

	idStr := c.Ctx.Input.Param(":id")
	tv_program_id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetTvProgramById(tv_program_id)
	if err != nil {
		c.Data["TvProgram"] = err.Error()
	} else {
		c.Data["TvProgram"] = v
	}

	l, err := models.GetReviewCommentByTvprogramId(tv_program_id)
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

	// var rating_tv_program []models.ReviewComment
	// for _, comment := range l {
	// 	u, err := models.GetReviewCommentByUserIdAndTvProgramId(comment.UserId, tv_program_id)
	// 	if err != nil{
	// 		u = new(models.ReviewComment)
	// 	}
	// 	fmt.Println(u)
	// // fmt.Println(u)
	// 	rating_tv_program = append(rating_tv_program, *u)
	// }
	// c.Data["RatingTvProgram"] = rating_tv_program


	session := c.StartSession()
	if session.Get("UserId")==nil{
		fmt.Println("you are not user, so your tv_Like break.")
	}	else {
		user_id := session.Get("UserId").(int64)
		// fmt.Println(user_id, "user id")
		w, err := models.GetWatchingStatusByUserAndTvProgram(user_id, tv_program_id)
		if err != nil {
		// fmt.Println(w)
			c.Data["WatchStatus"] = new(models.WatchingStatus)
		} else {
		// fmt.Println(w)
			c.Data["WatchStatus"] = w
		}

		x,_ := models.GetUserById(user_id)
		c.Data["User"] = x

		var review_comment_likes []models.ReviewCommentLike
		for _, comment := range l {
			u, err := models.GetReviewCommentLikeByCommentAndUser(comment.Id, user_id)
			if err != nil{
				u = new(models.ReviewCommentLike)
			}
		// fmt.Println(u)
			review_comment_likes = append(review_comment_likes, *u)
		}
		c.Data["CommentLike"] = review_comment_likes
	}
	c.TplName = "review_comment/show.tpl"
}


func (c *ReviewCommentController) SearchComment() {

	idStr := c.Ctx.Input.Param(":id")
	tv_program_id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetTvProgramById(tv_program_id)
	if err != nil {
		c.Data["TvProgram"] = err.Error()
	} else {
		c.Data["TvProgram"] = v
	}

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string][]string)
	var limit int64 = 100
	var offset int64
	var word []string
	var tv_id []string
	type SearchWords struct {
			Word string
			Category string
			Spoiler string
			Star string
			Limit int64
			Sortby string
	}
	if v := c.GetString("word"); v != "" {
		// word = strings.Replace(v, "　", " ", -1)
		word = strings.Split(strings.Replace(v, "　", " ", -1), " ")
		query["Word"] = word
	}
	if v := c.GetStrings("FavoritePoint"); v != nil {
		query["FavoritePoint"] = v
	}

	if v := c.GetStrings("star"); v != nil {
		query["Star"] = v
	}
	if v := c.GetStrings("spoiler"); v != nil {
		query["Spoiler"] = v
	}
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	tv_id = append(tv_id, strconv.FormatInt(tv_program_id, 10))
	query["TvProgramId"] = tv_id

	if v := c.GetString("sortby"); v != "" {
		sort_elem := v
		if sort_elem == "新しい順" {
			sortby = append(sortby, "Created")
			order = append(order, "desc")
		} else if sort_elem == "古い順" {
			sortby = append(sortby, "Created")
			order = append(order, "asc")
		} else if sort_elem == "いいねが多い順" {
			sortby = append(sortby, "CountLike")
			order = append(order, "desc")
		} else if sort_elem == "評価が高い順" {
			sortby = append(sortby, "Star")
			order = append(order, "desc")
		} else if sort_elem == "評価が引く順" {
			sortby = append(sortby, "Star")
			order = append(order, "asc")
		}
	}
	fmt.Println(query)

	var s SearchWords
	s = SearchWords {
		Word: c.GetString("word"),
		Category: strings.Join(c.GetStrings("FavoritePoint"), "、"),
		Spoiler: strings.Join(c.GetStrings("spoiler"), "、"),
		Star: strings.Join(c.GetStrings("star"), "、"),
		Limit: limit,
		Sortby: c.GetString("sortby"),
	}
	c.Data["SearchWords"] = s
	fmt.Println("SearchWords",s)

	session := c.StartSession()
	if session.Get("UserId")!=nil{
		var u models.SearchHistory
		u = models.SearchHistory {
			UserId: session.Get("UserId").(int64),
			Word: strings.Join(word, "、"),
			Category: s.Category,
			Spoiler: s.Spoiler,
			Star: s.Star,
			Limit: s.Limit,
			Sortby:s.Sortby,
		}
		_,_ = models.AddSearchHistory(&u)
	}

	l, err := models.SearchReviewComment(query, fields, sortby, order, offset, limit)
	c.Data["Comment"] = l
	// fmt.Println(l[0].(models.Comment).Id)
	var users []models.User
	for _, comment := range l {
		u, _ := models.GetUserById(comment.(models.ReviewComment).UserId)
		users = append(users, *u)
	}
	c.Data["Users"] = users
	// session := c.StartSession()
// 閲覧数カウント
	if session.Get(tv_program_id)==nil{
		fmt.Println("first tv click")
		if session.Get("UserId")!=nil{
			user_id := session.Get("UserId").(int64)
		var b models.BrowsingHistory
		b = models.BrowsingHistory{
			UserId: user_id,
			TvProgramId: tv_program_id,
		}
		_, err = models.AddBrowsingHistory(&b)
		if err==nil{
			fmt.Println("browsing_history", b)
		}
	}
		v.CountClicked++
		_ = models.UpdateTvProgramById(v)
		session.Set(tv_program_id, true)
	}

	if session.Get("UserId")==nil{
		fmt.Println("you are not user, so your tv_Like break.")
	}	else {
		user_id := session.Get("UserId").(int64)
		w, err := models.GetWatchingStatusByUserAndTvProgram(user_id, tv_program_id)
		if err != nil {
			c.Data["WatchStatus"] = new(models.WatchingStatus)
		} else {
			c.Data["WatchStatus"] = w
		}

		x,_ := models.GetUserById(user_id)
		c.Data["User"] = x

		var comment_likes []models.ReviewCommentLike
		for _, comment := range l {
			u, err := models.GetReviewCommentLikeByCommentAndUser(comment.(models.ReviewComment).Id, user_id)
			if err != nil{
				u = new(models.ReviewCommentLike)
			}
			comment_likes = append(comment_likes, *u)
		}
		c.Data["CommentLike"] = comment_likes

	}
	c.TplName = "review_comment/show.tpl"
}