package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"strconv"

	"github.com/astaxie/beego/orm"
)


type TvProgram struct {
	Id       int64  `orm:"auto"`
	Title    string `orm:"size(128);unique"`
	Content  string `orm:"size(500);null"`
	ImageUrl string `orm:"size(500);null"`
	ImageUrlReference string `orm:"size(128);null"`
	MovieUrl string `orm:"size(500);null"`
	MovieUrlReference string `orm:"size(128);null"`
	Cast    string `orm:"size(128);null"`
	Category string `orm:"size(128);null"`
	Dramatist string `orm:"size(128);null"`
	Supervisor string `orm:"size(128);null"`
	Director string `orm:"size(128);null"`
	Production string `orm:"size(128);null"`
	Year int `orm:"null"`
	Season *Season `orm:"rel(fk);null"`
	Week *Week `orm:"rel(fk);null"`
	Hour float32 `orm:"null`
	Themesong string `orm:"size(128);null"`
	CreateUserId int64 `orm:"default(0)"`
	Star float32 `orm:"default(2.5)"`
	CountStar   int32 `orm:"default(0)"`
	CountWatched   int32 `orm:"default(0)"`
	CountWantToWatch   int32 `orm:"default(0)"`
	CountClicked   int32 `orm:"default(0)"`
	CountAuthorization int32 `orm:"default(0)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}


type Season struct	{
	Name string `orm:"pk"`
	Id	int
}

type Week struct	{
	Name string `orm:"pk"`
	Id	int
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
		for _, value := range strings.Split(v, " "){
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
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
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

func SearchTvProgramAll(str string) (v []TvProgram, err error) {
	o := orm.NewOrm()
	cond_all := orm.NewCondition()
	str = strings.Replace(str, "　", " ", -1)
  for _, v := range strings.Split(str, " ") {
  	cond := orm.NewCondition()
		v_float, _ := strconv.ParseFloat(v, 32)
		if v_float == 0 {
			// Hourに条件「文字式」を入れると自動的に0になっちゃうので，回避
			v_float = 100
		}
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
		cond = cond.Or("Hour", v_float)

		cond_all = cond_all.AndCond(cond)
  }

	if _, err = o.QueryTable(new(TvProgram)).SetCond(cond_all).OrderBy("-Year", "-Season__Id", "Week__Id", "Hour").All(&v); err == nil {
		return v, nil
	}
	return nil, err
}

func SearchTvProgram(query map[string][]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TvProgram))
	// fmt.Println(query)
	// query
	// cond_only := orm.NewCondition()
	cond_all := orm.NewCondition()
	for k, v := range query {
		cond_only := orm.NewCondition()
		for _, value := range v {
			if k == "Title" {
					cond_only = cond_only.And("Title__icontains", value)
			} else if k == "Staff" { 
					cond_only = cond_only.Or("Cast__icontains", value)
					cond_only = cond_only.Or("Dramatist__icontains", value)
					cond_only = cond_only.Or("Supervisor__icontains", value)
					cond_only = cond_only.Or("Director__icontains", value)
					cond_only = cond_only.Or("Production__icontains", value)
			} else if k == "Themesong" { 
					cond_only = cond_only.Or("Themesong__icontains", value)
			} else if k == "Year" { 
					cond_only = cond_only.Or("Year", value)
			} else if k == "Week" { 
					cond_only = cond_only.Or("Week__Name", value)
			} else if k == "Hour" {
					cond_only = cond_only.Or("Hour", value)
			} else if k == "Season" {
					cond_only = cond_only.Or("Season", value)
			} else if k == "Category" {
					cond_only = cond_only.Or("Category__icontains", value)
			}
		}
		// fmt.Println(k,v)
		cond_all = cond_all.AndCond(cond_only)
	}
	qs = qs.SetCond(cond_all)
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
	season_name := [4]string{"春","夏","秋","冬"}
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
				season = season_name[i]
			}
	}
	return season
}
