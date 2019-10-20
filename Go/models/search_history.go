package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SearchHistory struct {
	Id     int64 `orm:"auto"`
	UserId int64
	Word string `orm:"size(60);null"`
	Year string `orm:"size(200);null"`
	Season string `orm:"size(60);null"`
	Week string `orm:"size(60);null"`
	Hour string `orm:"size(60);null"`
	Category string `orm:"size(60);null"`
	Spoiler string
	Star string`orm:"size(60);null"`
	Limit int64
	Sortby string `orm:"size(60);null"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(SearchHistory))
}

// AddSearchHistory insert a new SearchHistory into database and returns
// last inserted Id on success.
func AddSearchHistory(m *SearchHistory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSearchHistoryById retrieves SearchHistory by Id. Returns error if
// Id doesn't exist
func GetSearchHistoryById(id int64) (v *SearchHistory, err error) {
	o := orm.NewOrm()
	v = &SearchHistory{Id: id}
	if err = o.QueryTable(new(SearchHistory)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSearchHistory retrieves all SearchHistory matches certain condition. Returns empty list if
// no records exist
func GetAllSearchHistory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SearchHistory))
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

	var l []SearchHistory
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

// UpdateSearchHistory updates SearchHistory by Id and returns error if
// the record to be updated doesn't exist
func UpdateSearchHistoryById(m *SearchHistory) (err error) {
	o := orm.NewOrm()
	v := SearchHistory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSearchHistory deletes SearchHistory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSearchHistory(id int64) (err error) {
	o := orm.NewOrm()
	v := SearchHistory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SearchHistory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
