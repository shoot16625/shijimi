package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type PointHistory struct {
	Id         int64 `orm:"auto"`
	UserId     int64
	MoneyPoint int
	Created    time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(PointHistory))
}

// AddPointHistory insert a new PointHistory into database and returns
// last inserted Id on success.
func AddPointHistory(m *PointHistory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPointHistoryById retrieves PointHistory by Id. Returns error if
// Id doesn't exist
func GetPointHistoryById(id int64) (v *PointHistory, err error) {
	o := orm.NewOrm()
	v = &PointHistory{Id: id}
	if err = o.QueryTable(new(PointHistory)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPointHistory retrieves all PointHistory matches certain condition. Returns empty list if
// no records exist
func GetAllPointHistory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PointHistory))
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

	var l []PointHistory
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

// UpdatePointHistory updates PointHistory by Id and returns error if
// the record to be updated doesn't exist
func UpdatePointHistoryById(m *PointHistory) (err error) {
	o := orm.NewOrm()
	v := PointHistory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePointHistory deletes PointHistory by Id and returns error if
// the record to be deleted doesn't exist
func DeletePointHistory(id int64) (err error) {
	o := orm.NewOrm()
	v := PointHistory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PointHistory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// 最新のログインポイント付与履歴
func GetLoginPointHistoryByUserId(id int64) (v []PointHistory, err error) {
	limit := 1
	o := orm.NewOrm()
	if _, err = o.QueryTable(new(PointHistory)).Filter("UserId", id).Filter("MoneyPoint", 1).OrderBy("-Id").Limit(limit).All(&v); err == nil {
		// fmt.Println(v)
		return v, nil
	}
	return nil, err
}

// 本日初めてのログインかどうかチェック
func TodayFirstLoginCheck(userID int64) bool {
	flag := false
	v, _ := GetLoginPointHistoryByUserId(userID)
	if len(v) == 0 {
		flag = true
	} else {
		lastLoginPointAddTime := v[0].Created
		t := time.Now()
		u := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
		durationA := u.Sub(lastLoginPointAddTime)
		durationB := t.Sub(u)
		// fmt.Println(durationA, durationB)
		if durationA > 0 && durationB > 0 {
			flag = true
		}
	}
	return flag
}
