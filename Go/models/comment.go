package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id          int64  `orm:"auto"`
	Content     string `orm:"type(longtext)"`
	TvProgramId int64
	UserId      int64
	CountLike   int       `orm:"default(0)"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
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
	var maxLimit int64 = 1000
	if maxLimit < limit {
		limit = maxLimit
	}
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
				fmt.Println(value)
				t, _ := time.Parse("2006-01-02 15:04", value)
				t = t.Local()
				t = t.Add(-9 * time.Hour)
				condOnly = condOnly.And("created__gte", t)
			} else if k == "AfterTime" {
				t, _ := time.Parse("2006-01-02 15:04", value)
				t = t.Local()
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
	var maxLimit int64 = 1000
	if maxLimit < limit {
		limit = maxLimit
	}
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

func GetCommentByTvprogramId(id int64) (v []Comment, err error) {
	o := orm.NewOrm()
	if _, err = o.QueryTable(new(Comment)).Filter("TvProgramId", id).OrderBy("-Created").All(&v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetCommentByUserId(id int64) (v []Comment, err error) {
	var limit int64 = 1000
	o := orm.NewOrm()
	if _, err = o.QueryTable(new(Comment)).Filter("UserId", id).Limit(limit).OrderBy("-Created").All(&v); err == nil {
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

func DeleteCommentsByUserId(id int64) {
	o := orm.NewOrm()
	num, _ := o.QueryTable(new(Comment)).Filter("UserId", id).Delete()
	fmt.Println("delete comment", num)
}
