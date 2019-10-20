package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CommentLike struct {
	Id        int64 `orm:"auto"`
	UserId    int64
	CommentId int64
	Like      bool
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(CommentLike))
}

// AddCommentLike insert a new CommentLike into database and returns
// last inserted Id on success.
func AddCommentLike(m *CommentLike) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCommentLikeById retrieves CommentLike by Id. Returns error if
// Id doesn't exist
func GetCommentLikeById(id int64) (v *CommentLike, err error) {
	o := orm.NewOrm()
	v = &CommentLike{Id: id}
	if err = o.QueryTable(new(CommentLike)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCommentLike retrieves all CommentLike matches certain condition. Returns empty list if
// no records exist
func GetAllCommentLike(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CommentLike))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
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

	var l []CommentLike
	qs = qs.OrderBy(sortFields...).RelatedSel()
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

// UpdateCommentLike updates CommentLike by Id and returns error if
// the record to be updated doesn't exist
func UpdateCommentLikeById(m *CommentLike) (err error) {
	o := orm.NewOrm()
	v := CommentLike{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCommentLike deletes CommentLike by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCommentLike(id int64) (err error) {
	o := orm.NewOrm()
	v := CommentLike{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CommentLike{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetCommentLikeByCommentAndUser(comment_id int64, user_id int64) (v *CommentLike, err error) {
	o := orm.NewOrm()
	v = &CommentLike{CommentId: comment_id, UserId: user_id}
	if err = o.QueryTable(new(CommentLike)).Filter("CommentId", comment_id).Filter("UserId", user_id).RelatedSel().One(v); err == nil {
			// fmt.Println(v)
		return v, nil
	}
			// fmt.Println(v)
	return nil, err
}

func GetCommentLikeByComment(comment_id int64) (v []CommentLike, err error) {
	o := orm.NewOrm()
	if _,err = o.QueryTable(new(CommentLike)).Filter("CommentId", comment_id).All(&v); err == nil {
		return v, nil
	}
	return nil, err
}