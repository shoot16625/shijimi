package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type BrowsingHistory struct {
	Id          int64 `orm:"auto"`
	UserId      int64
	TvProgramId int64
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(BrowsingHistory))
}

// AddBrowsingHistory insert a new BrowsingHistory into database and returns
// last inserted Id on success.
func AddBrowsingHistory(m *BrowsingHistory) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBrowsingHistoryById retrieves BrowsingHistory by Id. Returns error if
// Id doesn't exist
func GetBrowsingHistoryById(id int64) (v *BrowsingHistory, err error) {
	o := orm.NewOrm()
	v = &BrowsingHistory{Id: id}
	if err = o.QueryTable(new(BrowsingHistory)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBrowsingHistory retrieves all BrowsingHistory matches certain condition. Returns empty list if
// no records exist
func GetAllBrowsingHistory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BrowsingHistory))
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

	var l []BrowsingHistory
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

// UpdateBrowsingHistory updates BrowsingHistory by Id and returns error if
// the record to be updated doesn't exist
func UpdateBrowsingHistoryById(m *BrowsingHistory) (err error) {
	o := orm.NewOrm()
	v := BrowsingHistory{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBrowsingHistory deletes BrowsingHistory by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBrowsingHistory(id int64) (err error) {
	o := orm.NewOrm()
	v := BrowsingHistory{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&BrowsingHistory{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// テレビ登録削除時の処理
func DeleteBrowsingHistoryByTvProgramId(id int64) {
	o := orm.NewOrm()
	num, _ := o.QueryTable(new(BrowsingHistory)).Filter("TvProgramId", id).Delete()
	fmt.Println("delete BrowsingHistory", num)
}

// n時間の間で閲覧数の多かった番組を取得
func GetTopBrowsingHistory(t string) (l []orm.Params, err error) {
	o := orm.NewOrm()
	sql := "SELECT tp.*, COUNT(bh.tv_program_id) AS Num FROM browsing_history AS bh JOIN tv_program AS tp ON tp.id = bh.tv_program_id WHERE bh.Created > '" + t + "' GROUP BY bh.tv_program_id ORDER BY Num DESC LIMIT 3"
	if _, err := o.Raw(sql).Values(&l); err == nil {
		return l, err
	} else {
		return nil, err
	}
}
