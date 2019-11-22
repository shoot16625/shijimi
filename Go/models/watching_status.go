package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type WatchingStatus struct {
	Id          int64 `orm:"auto"`
	UserId      int64
	TvProgramId int64
	Watched     bool      `orm:"default(false)"`
	WantToWatch bool      `orm:"default(false)"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(WatchingStatus))
}

// AddWatchingStatus insert a new WatchingStatus into database and returns
// last inserted Id on success.
func AddWatchingStatus(m *WatchingStatus) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetWatchingStatusById retrieves WatchingStatus by Id. Returns error if
// Id doesn't exist
func GetWatchingStatusById(id int64) (v *WatchingStatus, err error) {
	o := orm.NewOrm()
	v = &WatchingStatus{Id: id}
	if err = o.QueryTable(new(WatchingStatus)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllWatchingStatus retrieves all WatchingStatus matches certain condition. Returns empty list if
// no records exist
func GetAllWatchingStatus(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(WatchingStatus))
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

	var l []WatchingStatus
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

// UpdateWatchingStatus updates WatchingStatus by Id and returns error if
// the record to be updated doesn't exist
func UpdateWatchingStatusById(m *WatchingStatus) (err error) {
	o := orm.NewOrm()
	v := WatchingStatus{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteWatchingStatus deletes WatchingStatus by Id and returns error if
// the record to be deleted doesn't exist
func DeleteWatchingStatus(id int64) (err error) {
	o := orm.NewOrm()
	v := WatchingStatus{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&WatchingStatus{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// テレビ登録削除時の処理
func DeleteWatchingStatusByTvProgramId(id int64) {
	o := orm.NewOrm()
	num, _ := o.QueryTable(new(WatchingStatus)).Filter("TvProgramId", id).Delete()
	fmt.Println("delete WatchingStatus", num)
}

func GetWatchingStatusByUserAndTvProgram(userID int64, tvProgramID int64) (v *WatchingStatus, err error) {
	o := orm.NewOrm()
	v = &WatchingStatus{UserId: userID, TvProgramId: tvProgramID}
	if err = o.QueryTable(new(WatchingStatus)).Filter("UserId", userID).Filter("TvProgramId", tvProgramID).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}
