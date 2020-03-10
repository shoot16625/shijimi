package models

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id          int64  `orm:"auto"`
	Content     string `orm:"size(500)""`
	TvProgramId int64
	UserId      int64
	CountLike   int       `orm:"default(0)"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	// Updated     time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Comment))
}

// AddComment insert a new Comment into database and returns
// last inserted Id on success.
func AddComment(m *Comment) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCommentById retrieves Comment by Id. Returns error if
// Id doesn't exist
func GetCommentById(id int64) (v *Comment, err error) {
	o := orm.NewOrm()
	v = &Comment{Id: id}
	if err = o.QueryTable(new(Comment)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllComment retrieves all Comment matches certain condition. Returns empty list if
// no records exist
func GetAllComment(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Comment))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		v = strings.Replace(v, "　", " ", -1)
		for _, value := range strings.Split(v, " ") {
			qs = qs.Filter(k, value)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Comment
	qs = qs.OrderBy(sortFields...).RelatedSel()
	// var maxLimit int64 = 200
	// if maxLimit < limit {
	// 	limit = maxLimit
	// }
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

func SearchComment(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Comment))
	condAll := orm.NewCondition()
	for k, v := range query {
		condOnly := orm.NewCondition()
		for _, value := range strings.Split(v, ",") {
			if k == "Content" {
				condOnly = condOnly.And("Content__icontains", value)
			} else if k == "Username" {
				if t, err := GetUserByUsername(value); err == nil {
					condOnly = condOnly.Or("UserId", t)
				}
			} else if k == "TvProgramId" {
				condOnly = condOnly.And("TvProgramId", value)
			} else if k == "BeforeTime" {
				// fmt.Println(value)
				t, _ := time.Parse("2006-01-02 15:04", value)
				t = t.Local()
				// herokuならコメントアウト
				t = t.Add(-9 * time.Hour)
				condOnly = condOnly.And("created__gte", t)
			} else if k == "AfterTime" {
				t, _ := time.Parse("2006-01-02 15:04", value)
				t = t.Local()
				// herokuならコメントアウト
				t = t.Add(-9 * time.Hour)
				condOnly = condOnly.And("created__lte", t)
			}
		}
		condAll = condAll.AndCond(condOnly)
	}
	qs = qs.SetCond(condAll)
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Comment
	qs = qs.OrderBy(sortFields...).RelatedSel()
	// var maxLimit int64 = 200
	// if maxLimit < limit {
	// 	limit = maxLimit
	// }
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateComment updates Comment by Id and returns error if
// the record to be updated doesn't exist
func UpdateCommentById(m *Comment) (err error) {
	o := orm.NewOrm()
	v := Comment{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteComment deletes Comment by Id and returns error if
// the record to be deleted doesn't exist
func DeleteComment(id int64) (err error) {
	o := orm.NewOrm()
	v := Comment{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Comment{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetCommentByTvProgramId(id int64, limit int64) (v []Comment, err error) {
	o := orm.NewOrm()
	if _, err = o.QueryTable(new(Comment)).Filter("TvProgramId", id).OrderBy("-Id").Limit(limit).All(&v); err == nil {
		return v, nil
	}
	return nil, err
}

func CountAllCommentNumByTvProgramId(id int64) (cnt int64) {
	o := orm.NewOrm()
	cnt, _ = o.QueryTable(new(Comment)).Filter("TvProgramId", id).Count()
	return cnt
}

func GetCommentByUserId(id int64, limit int64) (v []Comment, err error) {
	o := orm.NewOrm()
	if _, err = o.QueryTable(new(Comment)).Filter("UserId", id).OrderBy("-Id").Limit(limit).All(&v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllCommentByUserId(id int64) (v []Comment, err error) {
	o := orm.NewOrm()
	if _, err = o.QueryTable(new(Comment)).Filter("UserId", id).All(&v); err == nil {
		return v, nil
	}
	return nil, err
}

// コメント削除時の処理
func DeleteCommentsByUserId(id int64) {
	o := orm.NewOrm()
	// num, _ := o.QueryTable(new(Comment)).Filter("UserId", id).Delete()
	if v, err := GetAllCommentByUserId(id); err == nil {
		for _, w := range v {
			o.QueryTable(new(CommentLike)).Filter("CommentId", w.Id).Delete()
		}
	}
	num, _ := o.QueryTable(new(Comment)).Filter("UserId", id).Delete()
	// いいね情報も同時に削除する
	fmt.Println("delete comment", num)
}

func GetTwitter(keyword string) anaconda.SearchResponse {
	api := GetTwitterApi()
	v := url.Values{}
	v.Set("count", "1000")
	searchResult, _ := api.GetSearch(keyword, v)
	return searchResult
}

func GetTwitterApi() *anaconda.TwitterApi {
	api := anaconda.NewTwitterApiWithCredentials(beego.AppConfig.String("twitter-your-access-token"), beego.AppConfig.String("twitter-your-access-token-secret"), beego.AppConfig.String("twitter-your-consumer-key"), beego.AppConfig.String("twitter-your-consumer-secret"))
	return api
}

func NormalizeTwitter(searchResult anaconda.SearchResponse) (res []anaconda.Tweet) {
	for _, tweet := range searchResult.Statuses {
		// リツイートは除外
		// 宛先が指定されているtweetは除外
		if tweet.RetweetedStatus == nil && !(strings.HasPrefix(tweet.FullText, "@")) {
			res = append(res, tweet)
		}
	}
	return res
}

func ReshapeTweetJson(searchResult []anaconda.Tweet, tvProgramId int64) (res []Comment) {
	for index, tweet := range searchResult {
		var t Comment
		t.Content = tweet.FullText
		t.TvProgramId = tvProgramId
		t.UserId = 1
		t.Id = -int64(index)
		c, _ := time.Parse("Mon Jan 2 15:04:05 -0700 2006", tweet.CreatedAt)
		c = c.In(time.Local)
		t.Created = c
		res = append(res, t)
	}
	return res
}
