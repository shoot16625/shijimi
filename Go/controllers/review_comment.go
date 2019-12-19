package controllers

import (
	"app/models"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

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
	session := c.StartSession()
	if session.Get("UserId") != nil {
		// 初投稿の場合のみ許可
		_, err := models.GetReviewCommentByUserIdAndTvProgramId(v.UserId, v.TvProgramId)
		if err != nil {
			if _, err := models.AddReviewComment(&v); err == nil {
				// c.Data["json"] = v
				if w, err := models.GetTvProgramById(v.TvProgramId); err == nil {
					if y, err := models.GetReviewCommentByTvProgramId(v.TvProgramId, 100000); err == nil {
						w.CountStar = len(y)
						var n int = 0
						for _, z := range y {
							n += z.Star
						}
						w.Star = float32(n / w.CountStar)
						w.CountReviewComment++
						if err := models.UpdateTvProgramById(w); err != nil {
							fmt.Println(err)
						}
						if w, err := models.GetUserById(v.UserId); err == nil {
							w.CountReviewComment++
							_ = models.UpdateUserById(w)
						}
					}
				} else {
					c.Redirect("/", 302)
					return
				}
			}
		}
	}
	c.Redirect("/tv/tv_program/review/"+strconv.FormatInt(v.TvProgramId, 10), 302)
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
	sortby = append(sortby, "Id")
	order = append(order, "desc")
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
	// idStr := c.Ctx.Input.Param(":id")
	// id, _ := strconv.ParseInt(idStr, 0, 64)
	// v := models.ReviewComment{Id: id}
	// json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	// if err := models.UpdateReviewCommentById(&v); err == nil {
	// 	c.Data["json"] = "OK"
	// } else {
	// 	c.Data["json"] = err.Error()
	// }
	// c.ServeJSON()
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
	session := c.StartSession()
	if session.Get("UserId") != nil {
		if err := models.DeleteReviewComment(id); err == nil {
			u, err := models.GetReviewCommentLikeByComment(id)
			if err == nil {
				for _, value := range u {
					_ = models.DeleteReviewCommentLike(value.Id)
				}
			}
		}
	}
	c.Redirect("/tv/user/show_review", 302)
}

func (c *ReviewCommentController) Show() {

	idStr := c.Ctx.Input.Param(":id")
	tvProgramID, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetTvProgramById(tvProgramID)
	if err != nil {
		c.Redirect("/", 302)
		return
	} else {
		c.Data["TvProgram"] = v
	}

	l, err := models.GetReviewCommentByTvProgramId(tvProgramID, 100)
	if err != nil {
		c.Data["Comment"] = nil
	} else {
		c.Data["Comment"] = l
	}
	cnt := models.CountAllReviewCommentNumByTvProgramId(tvProgramID)
	c.Data["CommentNum"] = cnt

	fpRanking := models.FavoritePointRankingByTvProgramId(tvProgramID)
	if len(fpRanking) > 3 {
		// お気に入りポイントトップ3を抽出
		fpRanking = fpRanking[:3]
	} else if fpRanking[0].Value == 0 {
		fpRanking = nil
	}
	c.Data["FavoritePointRanking"] = fpRanking

	var users []models.User
	for _, comment := range l {
		u, _ := models.GetUserById(comment.UserId)
		if err != nil {
			u = new(models.User)
		}
		users = append(users, *u)
	}
	c.Data["Users"] = users

	session := c.StartSession()
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

		var reviewCommentLikes []models.ReviewCommentLike
		for _, comment := range l {
			u, err := models.GetReviewCommentLikeByCommentAndUser(comment.Id, userID)
			if err != nil {
				u = new(models.ReviewCommentLike)
			}
			reviewCommentLikes = append(reviewCommentLikes, *u)
		}
		c.Data["CommentLike"] = reviewCommentLikes
	}
	c.TplName = "review_comment/show.tpl"
}

func (c *ReviewCommentController) SearchComment() {

	idStr := c.Ctx.Input.Param(":id")
	tvProgramID, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetTvProgramById(tvProgramID)
	if err != nil {
		c.Redirect("/", 302)
		return
	} else {
		c.Data["TvProgram"] = v
	}
	cnt := models.CountAllReviewCommentNumByTvProgramId(tvProgramID)
	c.Data["CommentNum"] = cnt

	fpRanking := models.FavoritePointRankingByTvProgramId(tvProgramID)
	if len(fpRanking) > 3 {
		// お気に入りポイントトップ3を抽出
		fpRanking = fpRanking[:3]
	} else if fpRanking[0].Value == 0 {
		fpRanking = nil
	}
	c.Data["FavoritePointRanking"] = fpRanking

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string][]string)
	var limit int64 = 100
	var offset int64
	var word []string
	var tvID []string
	type SearchWords struct {
		Word     string
		Category string
		Spoiler  string
		Star     string
		Limit    int64
		Sortby   string
	}
	if v := c.GetString("word"); v != "" {
		word = strings.Split(strings.Replace(v, "　", ",", -1), ",")
		query["Word"] = word
	}
	if v := c.GetStrings("favorite-point"); v != nil {
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
	tvID = append(tvID, strconv.FormatInt(tvProgramID, 10))
	query["TvProgramId"] = tvID

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
		} else if sortElem == "評価が高い順" {
			sortby = append(sortby, "Star")
			order = append(order, "desc")
		} else if sortElem == "評価が引く順" {
			sortby = append(sortby, "Star")
			order = append(order, "asc")
		}
	}
	// var s SearchWords
	s := SearchWords{
		Word:     c.GetString("word"),
		Category: strings.Join(c.GetStrings("favorite-point"), ","),
		Spoiler:  strings.Join(c.GetStrings("spoiler"), ","),
		Star:     strings.Join(c.GetStrings("star"), ","),
		Limit:    limit,
		Sortby:   c.GetString("sortby"),
	}
	c.Data["SearchWords"] = s

	session := c.StartSession()
	if session.Get("UserId") != nil {
		var u models.SearchHistory
		u = models.SearchHistory{
			UserId:   session.Get("UserId").(int64),
			Word:     strings.Join(word, ","),
			Category: s.Category,
			Spoiler:  s.Spoiler,
			Star:     s.Star,
			Limit:    s.Limit,
			Sortby:   s.Sortby,
			Item:     "review",
		}
		_, _ = models.AddSearchHistory(&u)
	}

	l, err := models.SearchReviewComment(query, fields, sortby, order, offset, limit)
	c.Data["Comment"] = l
	var users []models.User
	for _, comment := range l {
		u, _ := models.GetUserById(comment.(models.ReviewComment).UserId)
		users = append(users, *u)
	}
	c.Data["Users"] = users
	// 閲覧数カウント
	if session.Get(tvProgramID) == nil {
		if session.Get("UserId") != nil {
			userID := session.Get("UserId").(int64)
			b := models.BrowsingHistory{
				UserId:      userID,
				TvProgramId: tvProgramID,
			}
			_, err = models.AddBrowsingHistory(&b)
		}
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

		var commentLikes []models.ReviewCommentLike
		for _, comment := range l {
			u, err := models.GetReviewCommentLikeByCommentAndUser(comment.(models.ReviewComment).Id, userID)
			if err != nil {
				u = new(models.ReviewCommentLike)
			}
			commentLikes = append(commentLikes, *u)
		}
		c.Data["CommentLike"] = commentLikes

	}
	c.TplName = "review_comment/show.tpl"
}
