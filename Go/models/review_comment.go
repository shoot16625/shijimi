package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ReviewComment struct {
	Id            int64 `orm:"auto"`
	UserId        int64
	TvProgramId   int64
	Content       string `orm:"size(1000)"`
	CountLike     int    `orm:"default(0)"`
	Spoiler       bool
	Star          int       `orm:"default(5)"`
	FavoritePoint string    `orm:"size(100)";null"`
	Created       time.Time `orm:"auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(ReviewComment))
}

// AddReviewComment insert a new ReviewComment into database and returns
// last inserted Id on success.
func AddReviewComment(m *ReviewComment) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetReviewCommentById retrieves ReviewComment by Id. Returns error if
// Id doesn't exist
func GetReviewCommentById(id int64) (v *ReviewComment, err error) {
	o := orm.NewOrm()
	v = &ReviewComment{Id: id}
	if err = o.QueryTable(new(ReviewComment)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllReviewComment retrieves all ReviewComment matches certain condition. Returns empty list if
// no records exist
func GetAllReviewComment(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ReviewComment))
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

	var l []ReviewComment
	// var maxLimit int64 = 50
	// if maxLimit < limit {
	// 	limit = maxLimit
	// }
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

// UpdateReviewComment updates ReviewComment by Id and returns error if
// the record to be updated doesn't exist
func UpdateReviewCommentById(m *ReviewComment) (err error) {
	o := orm.NewOrm()
	v := ReviewComment{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteReviewComment deletes ReviewComment by Id and returns error if
// the record to be deleted doesn't exist
func DeleteReviewComment(id int64) (err error) {
	o := orm.NewOrm()
	v := ReviewComment{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ReviewComment{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetReviewCommentByTvProgramId(id int64, limit int64) (v []ReviewComment, err error) {
	o := orm.NewOrm()
	if _, err = o.QueryTable(new(ReviewComment)).Filter("TvProgramId", id).OrderBy("-Created").Limit(limit).All(&v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetReviewCommentByUserId(id int64, limit int64) (v []ReviewComment, err error) {
	o := orm.NewOrm()
	if _, err = o.QueryTable(new(ReviewComment)).Filter("UserId", id).Limit(limit).OrderBy("-Created").All(&v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetReviewCommentByUserIdAndTvProgramId(userID int64, tvProgramID int64) (v *ReviewComment, err error) {
	o := orm.NewOrm()
	v = &ReviewComment{UserId: userID, TvProgramId: tvProgramID}
	if err = o.QueryTable(new(ReviewComment)).Filter("UserId", userID).Filter("TvProgramId", tvProgramID).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// func GetRatingTvProgramByUserIdAndTvProgramId(userID int64, tvProgramID int64)(v *ReviewComment, err error){
// 	o := orm.NewOrm()
// 	v = &ReviewComment{UserId: userID, TvProgramId: tvProgramID}
// 	if err = o.QueryTable(new(ReviewComment)).Filter("UserId", userID).Filter("TvProgramId", tvProgramID).RelatedSel().One(v); err == nil {
// 		return v, nil
// 	}
// 	return nil, err
// }

func SearchReviewComment(query map[string][]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ReviewComment))
	// fmt.Println(query)
	// query
	// condOnly := orm.NewCondition()
	condAll := orm.NewCondition()
	for k, v := range query {
		condOnly := orm.NewCondition()
		for _, value := range v {
			fmt.Println(k, value)
			if k == "Word" {
				condOnly = condOnly.And("Content__icontains", value)
			} else if k == "Star" {
				condOnly = condOnly.Or("Star", value)
			} else if k == "Spoiler" {
				fmt.Println("1", v)
				if value == "ネタバレなし" {
					fmt.Println("2", v)
					condOnly = condOnly.Or("Spoiler", false)
				} else {
					fmt.Println("3", v)
					condOnly = condOnly.Or("Spoiler", true)
				}
			} else if k == "FavoritePoint" {
				condOnly = condOnly.Or("FavoritePoint__icontains", value)
			} else if k == "TvProgramId" {
				condOnly = condOnly.Or("TvProgramId", value)
			}
		}
		// fmt.Println(k,v)
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

	var l []ReviewComment
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

func GetAllReviewCommentByUserId(id int64) (v []ReviewComment, err error) {
	o := orm.NewOrm()
	if _, err = o.QueryTable(new(ReviewComment)).Filter("UserId", id).All(&v); err == nil {
		return v, nil
	}
	return nil, err
}

func DeleteReviewCommentsByUserId(id int64) {
	o := orm.NewOrm()
	num, _ := o.QueryTable(new(ReviewComment)).Filter("UserId", id).Delete()
	fmt.Println("delete review comment", num)
}

func CountAllReviewCommentNumByTvProgramId(id int64) (cnt int64) {
	o := orm.NewOrm()
	cnt, _ = o.QueryTable(new(ReviewComment)).Filter("TvProgramId", id).Count()
	return cnt
}
