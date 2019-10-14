package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ReviewCommentLike struct {
	Id              int64 `orm:"auto"`
	UserId          int64
	ReviewCommentId int64
	Like            bool
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(ReviewCommentLike))
}

// AddReviewCommentLike insert a new ReviewCommentLike into database and returns
// last inserted Id on success.
func AddReviewCommentLike(m *ReviewCommentLike) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetReviewCommentLikeById retrieves ReviewCommentLike by Id. Returns error if
// Id doesn't exist
func GetReviewCommentLikeById(id int64) (v *ReviewCommentLike, err error) {
	o := orm.NewOrm()
	v = &ReviewCommentLike{Id: id}
	if err = o.QueryTable(new(ReviewCommentLike)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllReviewCommentLike retrieves all ReviewCommentLike matches certain condition. Returns empty list if
// no records exist
func GetAllReviewCommentLike(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ReviewCommentLike))
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

	var l []ReviewCommentLike
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

// UpdateReviewCommentLike updates ReviewCommentLike by Id and returns error if
// the record to be updated doesn't exist
func UpdateReviewCommentLikeById(m *ReviewCommentLike) (err error) {
	o := orm.NewOrm()
	v := ReviewCommentLike{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteReviewCommentLike deletes ReviewCommentLike by Id and returns error if
// the record to be deleted doesn't exist
func DeleteReviewCommentLike(id int64) (err error) {
	o := orm.NewOrm()
	v := ReviewCommentLike{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ReviewCommentLike{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetReviewCommentLikeByCommentAndUser(comment_id int64, user_id int64) (v *ReviewCommentLike, err error) {
	o := orm.NewOrm()
	v = &ReviewCommentLike{ReviewCommentId: comment_id, UserId: user_id}
	if err = o.QueryTable(new(ReviewCommentLike)).Filter("ReviewCommentId", comment_id).Filter("UserId", user_id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetReviewCommentLikeByComment(comment_id int64) (v []ReviewCommentLike, err error) {
	o := orm.NewOrm()
	if _,err = o.QueryTable(new(ReviewCommentLike)).Filter("ReviewCommentId", comment_id).All(&v); err == nil {
		return v, nil
	}
	return nil, err
}