package controllers

import (
	"app/models"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

//  UserController operations for User
type UserController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Create", c.Create)
	c.Mapping("Index", c.Index)
	c.Mapping("Show", c.Show)
	c.Mapping("ShowReview", c.ShowReview)
	c.Mapping("ShowWatchedTv", c.ShowWatchedTv)
	c.Mapping("ShowWtwTv", c.ShowWtwTv)
	c.Mapping("Edit", c.Edit)
	c.Mapping("Login", c.Login)
	c.Mapping("Logout", c.Logout)
	c.Mapping("ForgetUsernamePage", c.ForgetUsernamePage)
	c.Mapping("ForgetPasswordPage", c.ForgetPasswordPage)
	c.Mapping("ForgetUsername", c.ForgetUsername)
	c.Mapping("ForgetPassword", c.ForgetPassword)
}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	// var v models.User
	// age, _ := c.GetInt("age")
	// fmt.Println(c.GetString("age"))
	hashPass, _ := models.PasswordHash(c.GetString("password"))
	hashSecondpass, _ := models.PasswordHash(c.GetString("SecondPassword"))
	IconURL := c.GetString("IconURL")
	if !strings.Contains(IconURL, "http") {
		rand.Seed(time.Now().UnixNano())
		r := strconv.Itoa(rand.Intn(13) + 1)
		if len(r) == 1 {
			r = "0" + r
		}
		IconURL = "/static/img/user_img/s256_f_" + r + ".png"
	}
	// json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	v := models.User{
		Username:       c.GetString("username"),
		Password:       hashPass,
		SecondPassword: hashSecondpass,
		Age:            c.GetString("age"),
		Address:        c.GetString("address"),
		Gender:         c.GetString("gender"),
		Job:            c.GetString("job"),
		IconUrl:        IconURL,
		Marital:        c.GetString("marital"),
		BloodType:      c.GetString("bloodType"),
	}
	if _, err := models.AddUser(&v); err == nil {
		session := c.StartSession()
		session.Set("username", c.GetString("username"))
		session.Set("UserId", v.Id)
		c.Redirect("/tv/user/show", 302)
	} else {
		v.Password = c.GetString("password")
		v.SecondPassword = c.GetString("SecondPassword")
		c.Data["User"] = v
		c.TplName = "user/create.tpl"
	}
}

// GetOne ...
// @Title Get One
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetUserById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get User
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.User
// @Failure 403
// @router / [get]
func (c *UserController) GetAll() {
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

	l, err := models.GetAllUser(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for User content"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	// age, _ := c.GetInt("age")

	IconURL := c.GetString("IconURL")
	if !strings.Contains(IconURL, "http") && !strings.Contains(IconURL, "/static/img") {
		rand.Seed(time.Now().UnixNano())
		r := strconv.Itoa(rand.Intn(13) + 1)
		if len(r) == 1 {
			r = "0" + r
		}
		IconURL = "/static/img/user_img/s256_f_" + r + ".png"
	}
	oldUserInfo, _ := models.GetUserById(id)
	v := *oldUserInfo
	// hashPass := v.Password
	if c.GetString("password") != "" {
		hashPass, _ := models.PasswordHash(c.GetString("password"))
		v.Password = hashPass
	} else {
		v.Username = c.GetString("username")
		v.Age = c.GetString("age")
		v.Address = c.GetString("address")
		v.Gender = c.GetString("gender")
		v.Job = c.GetString("job")
		v.IconUrl = IconURL
		v.Marital = c.GetString("marital")
		v.BloodType = c.GetString("bloodType")
	}
	if err := models.UpdateUserById(&v); err == nil {
		c.Data["json"] = "OK"
		c.Redirect("show", 302)
	} else {
		c.Data["User"] = v
		c.Data["NameFlag"] = true
		c.TplName = "user/edit.tpl"
	}
}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserController) Delete() {
	session := c.StartSession()
	id := session.Get("UserId").(int64)
	if err := models.DeleteUser(id); err == nil {
		// 過去の投稿データを削除(いいねも削除)
		models.DeleteCommentsByUserId(id)
		models.DeleteReviewCommentsByUserId(id)
		session.Delete("UserId")
		session.Delete("Username")
		c.Data["Status"] = "ユーザを削除しました"
	} else {
		c.Data["Status"] = "退会に失敗しました"
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

func (c *UserController) Create() {
	c.TplName = "user/create.tpl"
}

func (c *UserController) Index() {
	// c.TplName = "user/index.tpl"
}

func (c *UserController) Show() {
	session := c.StartSession()
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	// コメント欄からプロフィール遷移してきた場合
	if id != 0 {
		v, _ := models.GetUserById(id)
		c.Data["User"] = v
		w, _ := models.GetCommentByUserId(id, 1000)
		c.Data["Comment"] = w

		var commentLikes []models.CommentLike
		var tvPrograms []models.TvProgram
		var myUserID int64 = 0
		if session.Get("UserId") == nil {
			c.Data["MyUserId"] = nil
			for _, comment := range w {
				v, err := models.GetTvProgramById(comment.TvProgramId)
				if err != nil {
					v = new(models.TvProgram)
				}
				tvPrograms = append(tvPrograms, *v)
			}
			c.Data["TvProgram"] = tvPrograms
		} else {
			myUserID = session.Get("UserId").(int64)
			c.Data["MyUserId"] = myUserID
			for _, comment := range w {
				u, err := models.GetCommentLikeByCommentAndUser(comment.Id, myUserID)
				if err != nil {
					u = new(models.CommentLike)
				}
				commentLikes = append(commentLikes, *u)
				v, err := models.GetTvProgramById(comment.TvProgramId)
				if err != nil {
					v = new(models.TvProgram)
				}
				tvPrograms = append(tvPrograms, *v)
			}
			c.Data["CommentLike"] = commentLikes
			c.Data["TvProgram"] = tvPrograms
			// foot_print_log：停止中
			// z := models.FootPrintToUser {
			// 	UserId:   myUserID,
			// 	ToUserId: id,
			// }
			// _, _ = models.AddFootPrintToUser(&z)
		}
		c.TplName = "user/user_comment.tpl"
	} else {
		if session.Get("UserId") == nil {
			c.Redirect("/", 302)
		} else {
			UserID := session.Get("UserId").(int64)
			v, _ := models.GetUserById(UserID)
			c.Data["User"] = v
			w, _ := models.GetCommentByUserId(UserID, 1000)
			c.Data["Comment"] = w

			var commentLikes []models.CommentLike
			var tvPrograms []models.TvProgram
			for _, comment := range w {
				u, err := models.GetCommentLikeByCommentAndUser(comment.Id, UserID)
				if err != nil {
					u = new(models.CommentLike)
				}
				commentLikes = append(commentLikes, *u)
				v, err := models.GetTvProgramById(comment.TvProgramId)
				if err != nil {
					v = new(models.TvProgram)
				}
				tvPrograms = append(tvPrograms, *v)
			}
			c.Data["CommentLike"] = commentLikes
			c.Data["TvProgram"] = tvPrograms
		}
		c.TplName = "user/show_comment.tpl"
	}
}

func (c *UserController) ShowReview() {
	session := c.StartSession()
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if id != 0 {
		v, _ := models.GetUserById(id)
		c.Data["User"] = v
		w, _ := models.GetReviewCommentByUserId(id, 100)
		c.Data["Comment"] = w
		if session.Get("UserId") == nil {
			c.Data["MyUserId"] = nil
			var tvPrograms []models.TvProgram
			for _, comment := range w {
				v, err := models.GetTvProgramById(comment.TvProgramId)
				if err != nil {
					v = new(models.TvProgram)
				}
				tvPrograms = append(tvPrograms, *v)
			}
			c.Data["TvProgram"] = tvPrograms
		} else {
			c.Data["MyUserId"] = session.Get("UserId").(int64)

			var commentLikes []models.ReviewCommentLike
			var tvPrograms []models.TvProgram
			for _, comment := range w {
				u, err := models.GetReviewCommentLikeByCommentAndUser(comment.Id, id)
				if err != nil {
					u = new(models.ReviewCommentLike)
				}
				commentLikes = append(commentLikes, *u)
				v, err := models.GetTvProgramById(comment.TvProgramId)
				if err != nil {
					v = new(models.TvProgram)
				}
				tvPrograms = append(tvPrograms, *v)
			}
			c.Data["CommentLike"] = commentLikes
			c.Data["TvProgram"] = tvPrograms
		}
		c.TplName = "user/user_review.tpl"
	} else {
		if session.Get("UserId") == nil {
			c.Redirect("/", 302)
		} else {
			userID := session.Get("UserId").(int64)
			v, _ := models.GetUserById(userID)
			c.Data["User"] = v
			w, _ := models.GetReviewCommentByUserId(userID, 100)
			c.Data["Comment"] = w

			var commentLikes []models.ReviewCommentLike
			var tvPrograms []models.TvProgram
			for _, comment := range w {
				u, err := models.GetReviewCommentLikeByCommentAndUser(comment.Id, userID)
				if err != nil {
					u = new(models.ReviewCommentLike)
				}
				commentLikes = append(commentLikes, *u)
				v, err := models.GetTvProgramById(comment.TvProgramId)
				if err != nil {
					v = new(models.TvProgram)
				}
				tvPrograms = append(tvPrograms, *v)
			}
			c.Data["CommentLike"] = commentLikes
			c.Data["TvProgram"] = tvPrograms
		}
		c.TplName = "user/show_review.tpl"
	}
}

func (c *UserController) ShowWatchedTv() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 1000
	var offset int64

	session := c.StartSession()
	userID := session.Get("UserId").(int64)

	sortby = append(sortby, "Updated")
	order = append(order, "desc")
	query["Watched"] = "1"
	query["UserId"] = strconv.FormatInt(userID, 10)
	v, _ := models.GetAllWatchingStatus(query, fields, sortby, order, offset, limit)
	c.Data["WatchStatus"] = v

	var tvPrograms []models.TvProgram
	for _, watched := range v {
		r, err := models.GetTvProgramById(watched.(models.WatchingStatus).TvProgramId)
		if err != nil {
			tvPrograms = append(tvPrograms, *new(models.TvProgram))
		} else {
			tvPrograms = append(tvPrograms, *r)
		}
	}
	c.Data["TvProgram"] = tvPrograms
	u, _ := models.GetUserById(userID)
	c.Data["User"] = u
	c.TplName = "user/show_watched.tpl"
}

func (c *UserController) ShowWtwTv() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 1000
	var offset int64

	session := c.StartSession()
	userID := session.Get("UserId").(int64)

	sortby = append(sortby, "Updated")
	order = append(order, "desc")
	query["WantToWatch"] = "1"
	query["UserId"] = strconv.FormatInt(userID, 10)
	v, _ := models.GetAllWatchingStatus(query, fields, sortby, order, offset, limit)
	c.Data["WatchStatus"] = v

	var tvPrograms []models.TvProgram
	for _, watched := range v {
		r, err := models.GetTvProgramById(watched.(models.WatchingStatus).TvProgramId)
		if err != nil {
			tvPrograms = append(tvPrograms, *new(models.TvProgram))
		} else {
			tvPrograms = append(tvPrograms, *r)
		}
	}
	c.Data["TvProgram"] = tvPrograms
	u, _ := models.GetUserById(userID)
	c.Data["User"] = u
	c.TplName = "user/show_wtw_tv.tpl"
}

func (c *UserController) Edit() {
	session := c.StartSession()
	userID := session.Get("UserId").(int64)
	v, _ := models.GetUserById(userID)
	c.Data["User"] = v
	c.TplName = "user/edit.tpl"
}

func (c *UserController) Login() {
	session := c.StartSession()
	v, err := models.GetUserByUsername(c.GetString("username"))
	if err == nil && models.UserPassMach(v.Password, c.GetString("password")) {
		session.Set("username", c.GetString("username"))
		session.Set("UserId", v.Id)
		firstLoginToday := models.AddLoginPoint(v.Id)
		if firstLoginToday {
			c.Data["Status"] = "1ポイント獲得しました!!"
		} else {
			c.Data["Status"] = "ログインしました!!"
		}
		z := models.LoginHistory{
			UserId: v.Id,
		}
		_, _ = models.AddLoginHistory(&z)

		UserID := v.Id
		v, _ := models.GetUserById(UserID)
		c.Data["User"] = v
		w, _ := models.GetCommentByUserId(UserID, 1000)
		c.Data["Comment"] = w

		var commentLikes []models.CommentLike
		var tvPrograms []models.TvProgram
		for _, comment := range w {
			u, err := models.GetCommentLikeByCommentAndUser(comment.Id, UserID)
			if err != nil {
				u = new(models.CommentLike)
			}
			commentLikes = append(commentLikes, *u)
			v, err := models.GetTvProgramById(comment.TvProgramId)
			if err != nil {
				v = new(models.TvProgram)
			}
			tvPrograms = append(tvPrograms, *v)
		}
		c.Data["CommentLike"] = commentLikes
		c.Data["TvProgram"] = tvPrograms
		c.TplName = "user/show_comment.tpl"
	} else {
		c.Data["LoginError"] = true
		c.Data["UserId"] = session.Get("UserId")
		var fields []string
		var sortby []string
		var order []string
		var limit int64 = 100
		var offset int64
		var query = make(map[string]string)
		sortby = append(sortby, "Hour")
		order = append(order, "asc")
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
		query["Week.Name"] = "映画"
		w, err := models.GetAllTvProgram(query, fields, sortby, order, offset, limit)
		if err == nil {
			c.Data["TvProgramMovie"] = w
		}
		c.TplName = "tv_program/top_page.tpl"
	}
}

func (c *UserController) Logout() {
	session := c.StartSession()
	userID := session.Get("UserId")
	if userID != nil {
		session.Delete("UserId")
		session.Delete("Username")
	}
	c.Data["Status"] = "ログアウトしました"
	var Info struct {
		CntUsers      int64
		CntTvPrograms int64
	}
	Info.CntUsers = models.GetUserCount()
	Info.CntTvPrograms = models.GetTvProgramCount()
	c.Data["Info"] = Info
	c.TplName = "user/logout.tpl"
}

func (c *UserController) ForgetUsernamePage() {
	c.TplName = "user/forget_username.tpl"
}

func (c *UserController) ForgetPasswordPage() {
	c.TplName = "user/forget_password.tpl"
}

func (c *UserController) ForgetUsername() {
	v, _ := models.GetUserByPasswords(c.GetString("password"), c.GetString("age"), c.GetString("SecondPassword"))
	if v == nil {
		c.Data["User"] = new(models.User)
	} else {
		c.Data["User"] = v
	}
	c.TplName = "user/forget_username.tpl"
}

func (c *UserController) ForgetPassword() {
	v, _ := models.GetUserByUsernameAndPassword(c.GetString("username"), c.GetString("age"), c.GetString("SecondPassword"))
	if v == nil {
		c.Data["User"] = new(models.User)
		c.TplName = "user/forget_password.tpl"
	} else {
		v.SecondPassword = c.GetString("SecondPassword")
		c.Data["User"] = v
		c.TplName = "user/reset_password.tpl"
	}
}
