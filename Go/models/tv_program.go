package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type TvProgram struct {
	Id                 int64     `orm:"auto"`
	Title              string    `orm:"size(128);unique"`
	Content            string    `orm:"size(500);null"`
	ImageURL           string    `orm:"size(500);null"`
	ImageURLReference  string    `orm:"size(200);null"`
	MovieURL           string    `orm:"size(500);null"`
	MovieURLReference  string    `orm:"size(200);null"`
	WikiReference      string    `orm:"size(500);null"`
	Cast               string    `orm:"size(256);null"`
	Category           string    `orm:"size(32);null"`
	Dramatist          string    `orm:"size(128);null"`
	Supervisor         string    `orm:"size(128);null"`
	Director           string    `orm:"size(128);null"`
	Production         string    `orm:"size(32);null"`
	Year               int       `orm:"null"`
	Season             *Season   `orm:"rel(fk);null"`
	Week               *Week     `orm:"rel(fk);null"`
	Hour               float32   `orm:"default(100)`
	Themesong          string    `orm:"size(256);null"`
	CreateUserId       int64     `orm:"default(0)"`
	Star               float32   `orm:"default(2.5)"`
	CountStar          int32     `orm:"default(0)"`
	CountWatched       int32     `orm:"default(0)"`
	CountWantToWatch   int32     `orm:"default(0)"`
	CountClicked       int32     `orm:"default(0)"`
	CountAuthorization int32     `orm:"default(0)"`
	Created            time.Time `orm:"auto_now_add;type(datetime)"`
	Updated            time.Time `orm:"auto_now;type(datetime)"`
}

type Season struct {
	Name string `orm:"pk"`
	Id   int
}

type Week struct {
	Name string `orm:"pk"`
	Id   int
}

func init() {
	orm.RegisterModel(new(TvProgram))
	orm.RegisterModel(new(Season))
	orm.RegisterModel(new(Week))
}

// AddTvProgram insert a new TvProgram into database and returns
// last inserted Id on success.
func AddTvProgram(m *TvProgram) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func AddSeason(m *Season) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func AddWeek(m *Week) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// func AddTimezone(m *Timezone) (id int64, err error) {
// 	o := orm.NewOrm()
// 	id, err = o.Insert(m)
// 	return
// }

// GetTvProgramById retrieves TvProgram by Id. Returns error if
// Id doesn't exist
func GetTvProgramById(id int64) (v *TvProgram, err error) {
	o := orm.NewOrm()
	v = &TvProgram{Id: id}
	if err = o.QueryTable(new(TvProgram)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTvProgram retrieves all TvProgram matches certain condition. Returns empty list if
// no records exist
func GetAllTvProgram(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TvProgram))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		// k = strings.Replace(k, ".", "__", -1)
		// fmt.Println(k,v)
		// qs = qs.Filter(k, v)
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

	var l []TvProgram
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

// UpdateTvProgram updates TvProgram by Id and returns error if
// the record to be updated doesn't exist
func UpdateTvProgramById(m *TvProgram) (err error) {
	o := orm.NewOrm()
	v := TvProgram{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		_, _ = o.Update(m)
		// var num int64
		// if num, err = o.Update(m); err == nil {
		// 	fmt.Println("Number of records updated in database:", num)
		// }
	}
	return
}

// DeleteTvProgram deletes TvProgram by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTvProgram(id int64) (err error) {
	o := orm.NewOrm()
	v := TvProgram{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TvProgram{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// ツールバーの検索機能
func SearchTvProgramAll(str string) (v []TvProgram, err error) {
	o := orm.NewOrm()
	condAll := orm.NewCondition()
	str = strings.Replace(str, "　", " ", -1)
	for _, v := range strings.Split(str, " ") {
		cond := orm.NewCondition()
		cond = cond.Or("Title__icontains", v)
		cond = cond.Or("Cast__icontains", v)
		cond = cond.Or("Category__icontains", v)
		cond = cond.Or("Dramatist__icontains", v)
		cond = cond.Or("Supervisor__icontains", v)
		cond = cond.Or("Director__icontains", v)
		cond = cond.Or("Season__Name", v)
		cond = cond.Or("Themesong__icontains", v)
		cond = cond.Or("Week__Name", v)
		cond = cond.Or("Production__icontains", v)
		cond = cond.Or("Year", v)

		condAll = condAll.AndCond(cond)
	}

	if _, err = o.QueryTable(new(TvProgram)).SetCond(condAll).OrderBy("-Year", "-Season__Id", "Week__Id", "Hour").All(&v); err == nil {
		return v, nil
	}
	return nil, err
}

// 詳細検索機能
func SearchTvProgram(query map[string][]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TvProgram))
	condAll := orm.NewCondition()
	for k, v := range query {
		condOnly := orm.NewCondition()
		for _, value := range v {
			if k == "Title" {
				condOnly = condOnly.And("Title__icontains", value)
			} else if k == "Staff" {
				condOnly = condOnly.Or("Cast__icontains", value)
				condOnly = condOnly.Or("Dramatist__icontains", value)
				condOnly = condOnly.Or("Supervisor__icontains", value)
				condOnly = condOnly.Or("Director__icontains", value)
				condOnly = condOnly.Or("Production__icontains", value)
			} else if k == "Themesong" {
				condOnly = condOnly.Or("Themesong__icontains", value)
			} else if k == "Year" {
				condOnly = condOnly.Or("Year", value)
			} else if k == "Week" {
				condOnly = condOnly.Or("Week__Name", value)
			} else if k == "Hour" {
				condOnly = condOnly.Or("Hour", value)
			} else if k == "Season" {
				condOnly = condOnly.Or("Season__Name", value)
			} else if k == "Category" {
				condOnly = condOnly.Or("Category__icontains", value)
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

	var l []TvProgram
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

func GetOnairSeason() (season string) {
	seasonName := [4]string{"春", "夏", "秋", "冬"}
	var tmp int = 365
	t := time.Now()
	var seasons []time.Time
	seasons = append(seasons, time.Date(t.Year(), 7, 1, 0, 0, 0, 0, time.Local))
	seasons = append(seasons, time.Date(t.Year(), 10, 1, 0, 0, 0, 0, time.Local))
	seasons = append(seasons, time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, time.Local))
	seasons = append(seasons, time.Date(t.Year(), 4, 1, 0, 0, 0, 0, time.Local))
	for i := range seasons {
		duration := seasons[i].Sub(t)
		days := int(duration.Hours()) / 24
		if tmp > days && days > 2 {
			tmp = days
			season = seasonName[i]
		}
	}
	return season
}
