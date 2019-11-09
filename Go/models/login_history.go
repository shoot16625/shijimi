package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type LoginHistory struct {
	Id      int64 `orm:"auto"`
	UserId  int64
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(LoginHistory))
}

// AddLoginHistory insert a new LoginHistory into database and returns
// last inserted Id on success.
func AddLoginHistory(m *LoginHistory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLoginHistoryById retrieves LoginHistory by Id. Returns error if
// Id doesn't exist
func GetLoginHistoryById(id int64) (v *LoginHistory, err error) {
	o := orm.NewOrm()
	v = &LoginHistory{Id: id}
	if err = o.QueryTable(new(LoginHistory)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllLoginHistory retrieves all LoginHistory matches certain condition. Returns empty list if
// no records exist
func GetAllLoginHistory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LoginHistory))
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

	var l []LoginHistory
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

// UpdateLoginHistory updates LoginHistory by Id and returns error if
// the record to be updated doesn't exist
func UpdateLoginHistoryById(m *LoginHistory) (err error) {
	o := orm.NewOrm()
	v := LoginHistory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLoginHistory deletes LoginHistory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLoginHistory(id int64) (err error) {
	o := orm.NewOrm()
	v := LoginHistory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&LoginHistory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// 本日初めてのログインかどうかチェック
func GetLoginHistoryByUserId(userID int64) bool {
	t := time.Now()
	u := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	d := u.Format("2006-01-02 15:04:05")
	sql := "select * from login_history where user_id = " + strconv.FormatInt(userID, 10) + " AND Created > '" + d + "'LIMIT 1"
	o := orm.NewOrm()
	var l []orm.Params
	if _, err := o.Raw(sql).Values(&l); err == nil {
		if len(l) != 0 {
			fmt.Println("本日ログイン済み", l)
			return false
		}
	}
	return true
}
