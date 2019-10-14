package controllers

import (
"app/models"
	// "app/sessions"
	// "encoding/json"
"errors"
"strconv"
"strings"
"fmt"
	// "net/http"

"github.com/astaxie/beego"
	// "github.com/astaxie/beego/session"
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
	var v models.User
	age, _ := c.GetInt("age")
	hash_pass, _ := models.PasswordHash(c.GetString("password"))
	hash_secondpass, _ := models.PasswordHash(c.GetString("SecondPassword"))
	icon_url := c.GetString("IconUrl")
	if icon_url == "" {
		icon_url = "http://flat-icon-design.com/f/f_object_174/s512_f_object_174_0bg.png"
	}
	// json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	v = models.User{
		Username: c.GetString("username"),
		Password: hash_pass,
		SecondPassword: hash_secondpass,
		Age: age,
		Address: c.GetString("address"),
		Gender: c.GetString("gender"),
		Job: c.GetString("job"),
		IconUrl: icon_url,
		Marital: c.GetString("marital"),
	}
	
	// fmt.Println(v)
	if _, err := models.AddUser(&v); err == nil {
		fmt.Println("user create and login!!")
			session := c.StartSession()
			session.Set("username", c.GetString("username"))
			session.Set("UserId", v.Id)
			// c.Redirect("/tv/tv_program/index", 302)
			c.Redirect("/tv/user/show", 302)
	} else {
		// c.Data["json"] = err.Error()
		v.Password = c.GetString("password")
		v.SecondPassword = c.GetString("SecondPassword")
		c.Data["User"] = v
		c.TplName = "user/create.tpl"
		// c.Redirect("/tv/user/create", 302)
	}
	// c.ServeJSON()
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
	fmt.Println("update user profile")
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	age, _ := c.GetInt("age")
	hash_pass, _ := models.PasswordHash(c.GetString("password"))
	hash_secondpass, _ := models.PasswordHash(c.GetString("SecondPassword"))
	if c.GetString("password") == "" {
		u, _ := models.GetUserById(id)
		hash_pass = u.Password
		hash_secondpass = u.SecondPassword
	}
	var v models.User
	v = models.User{
		Id: id,
		Username: c.GetString("username"),
		Password: hash_pass,
		SecondPassword: hash_secondpass,
		Age: age,
		Address: c.GetString("address"),
		Gender: c.GetString("gender"),
		Job: c.GetString("job"),
		IconUrl: c.GetString("IconUrl"),
		Marital: c.GetString("marital"),
	}
	fmt.Println(v)
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
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
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
	// fmt.Println(id)
	if id != 0 {
		v,_ := models.GetUserById(id)
		c.Data["User"] = v
		w,_ := models.GetCommentByUserId(id)
		c.Data["Comment"] = w

		var comment_likes []models.CommentLike
		var MyUserId int64 = 0
		if session.Get("UserId") == nil {
			c.Data["MyUserId"] = nil
		} else {
		MyUserId = session.Get("UserId").(int64)
		c.Data["MyUserId"] = MyUserId
			for _, comment := range w {
			u, err := models.GetCommentLikeByCommentAndUser(comment.Id, MyUserId)
			if err != nil{
				u = new(models.CommentLike)
			}
			comment_likes = append(comment_likes, *u)
		}
		c.Data["CommentLike"] = comment_likes
	}
		// fmt.Println(comment_likes)

		c.TplName = "user/user_page.tpl"
	} else {
	if session.Get("UserId") == nil{
		c.Redirect("/", 302)
	} else {
		UserId := session.Get("UserId").(int64)
		v,_ := models.GetUserById(UserId)
		c.Data["User"] = v
		w,_ := models.GetCommentByUserId(UserId)
		c.Data["Comment"] = w

		var comment_likes []models.CommentLike
		for _, comment := range w {
			u, err := models.GetCommentLikeByCommentAndUser(comment.Id, UserId)
			if err != nil{
				u = new(models.CommentLike)
			}
			comment_likes = append(comment_likes, *u)
		}
		c.Data["CommentLike"] = comment_likes
	}
		c.TplName = "user/show.tpl"
	}
}

func (c *UserController) ShowReview() {
	session := c.StartSession()
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	// fmt.Println(id)
	if id != 0 {
		v,_ := models.GetUserById(id)
		c.Data["User"] = v
		w,_ := models.GetReviewCommentByUserId(id)
		c.Data["Comment"] = w
		if session.Get("UserId") == nil {
			c.Data["MyUserId"] = nil
		} else {
		c.Data["MyUserId"] = session.Get("UserId").(int64)

		var comment_likes []models.ReviewCommentLike
		for _, comment := range w {
			u, err := models.GetReviewCommentLikeByCommentAndUser(comment.Id, id)
			if err != nil{
				u = new(models.ReviewCommentLike)
			}
			comment_likes = append(comment_likes, *u)
		}
		c.Data["CommentLike"] = comment_likes
	}
		c.TplName = "user/user_review.tpl"
	} else {
	if session.Get("UserId") == nil{
		c.Redirect("/", 302)
	} else {
		UserId := session.Get("UserId").(int64)
		v,_ := models.GetUserById(UserId)
		c.Data["User"] = v
		w,_ := models.GetReviewCommentByUserId(UserId)
		c.Data["Comment"] = w

		var comment_likes []models.ReviewCommentLike
		for _, comment := range w {
			u, err := models.GetReviewCommentLikeByCommentAndUser(comment.Id, UserId)
			if err != nil{
				u = new(models.ReviewCommentLike)
			}
			comment_likes = append(comment_likes, *u)
		}
		c.Data["CommentLike"] = comment_likes
	}
		c.TplName = "user/show_review.tpl"
	}
}

func (c *UserController) Edit() {
	session := c.StartSession()
	UserId := session.Get("UserId").(int64)
	v,_ := models.GetUserById(UserId)
	c.Data["User"] = v
	c.TplName = "user/edit.tpl"
}

func (c *UserController) Login() {
	session := c.StartSession()
	v,_ := models.GetUserByUsername(c.GetString("username"))
	if  v == nil {
		fmt.Println("not user")
		c.TplName = "user/login_error.tpl"
	}	else {
		if models.UserPassMach(v.Password, c.GetString("password")) {
			fmt.Println("good password")
			session.Set("username", c.GetString("username"))
			session.Set("UserId", v.Id)
			v := models.LoginHistory {
				UserId: v.Id,
			}
			if _, err := models.AddLoginHistory(&v); err == nil {
				fmt.Println(v)
			}
			c.Redirect("/tv/user/show", 302)
		}		else {
			fmt.Println("bad password")
			c.TplName = "user/login_error.tpl"
		}
	}
}

func (c *UserController) Logout() {
	session := c.StartSession()
	user_id := session.Get("UserId")
	if user_id != nil {
		session.Delete("UserId")
		session.Delete("Username")
		fmt.Println(user_id,":logout")
	}
	c.TplName = "user/logout.tpl"
}

func (c *UserController) ForgetUsernamePage() {
	c.TplName = "user/forget_username.tpl"
}

func (c *UserController) ForgetPasswordPage() {
	c.TplName = "user/forget_password.tpl"
}


func (c *UserController) ForgetUsername() {
	v,_ := models.GetUserByPasswords(c.GetString("password"),c.GetString("SecondPassword"))
	if  v == nil {
		fmt.Println("can`t find user")
		c.Data["User"] = new(models.User)
	} else {
		c.Data["User"] = v
	}
	c.TplName = "user/forget_username.tpl"
}

func (c *UserController) ForgetPassword() {
	fmt.Println(c.GetString("username"),c.GetString("SecondPassword"))
	v,_ := models.GetUserByUsernameAndPassword(c.GetString("username"),c.GetString("SecondPassword"))
	if  v == nil {
		fmt.Println("can`t find password")
		c.Data["User"] = new(models.User)
		c.TplName = "user/forget_password.tpl"
	} else {
		v.SecondPassword = c.GetString("SecondPassword")
		c.Data["User"] = v
		c.TplName = "user/reset_password.tpl"
	}
}
