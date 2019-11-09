package models

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type TvProgram struct {
	Id                int64  `orm:"auto"`
	Title             string `orm:"size(128);unique"`
	Content           string `orm:"size(500);null"`
	ImageUrl          string `orm:"size(500);null"`
	ImageUrlReference string `orm:"size(200);null"`
	MovieUrl          string `orm:"size(500);null"`
	// MovieUrlReference  string    `orm:"size(200);null"`
	WikiReference      string  `orm:"size(500);null"`
	Cast               string  `orm:"size(256);null"`
	Category           string  `orm:"size(32);null"`
	Dramatist          string  `orm:"size(128);null"`
	Supervisor         string  `orm:"size(128);null"`
	Director           string  `orm:"size(128);null"`
	Production         string  `orm:"size(32);null"`
	Year               int     `orm:"default(2000)"`
	Season             *Season `orm:"rel(fk);null"`
	Week               *Week   `orm:"rel(fk);null"`
	Hour               float32 `orm:"default(100)`
	Themesong          string  `orm:"size(256);null"`
	CreateUserId       int64   `orm:"default(0)"`
	Star               float32 `orm:"default(5)"`
	CountStar          int
	CountWatched       int
	CountWantToWatch   int
	CountClicked       int
	CountUpdated       int
	CountComment       int
	CountReviewComment int
	// CountAuthorization int
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
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
func SearchTvProgramAll(str string) (l []TvProgram, err error) {
	var limit int64 = 100
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

	if _, err = o.QueryTable(new(TvProgram)).SetCond(condAll).Limit(limit).OrderBy("-Year", "-Season__Id", "Week__Id", "Hour").All(&l); err == nil {
		return l, nil
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
	var maxLimit int64 = 500
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

func GetOnAirSeason() (season string) {
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

// The number of TvPrograms.
func GetTvProgramCount() (cnt int64) {
	o := orm.NewOrm()
	cnt, _ = o.QueryTable(new(TvProgram)).Count()
	return cnt
}

// 現在放送中の番組で評価の高い番組を取得
func GetTopStarPoint() (l []TvProgram, err error) {
	t := time.Now()
	season := GetOnAirSeason()
	o := orm.NewOrm()
	if _, err = o.QueryTable(new(TvProgram)).Filter("Year", t.Year()).Filter("Season__Name", season).OrderBy("-Star").Limit(3).All(&l); err == nil {
		return l, err
	}
	return nil, err
}

type RecommendPoint struct {
	Index int64
	Point int
}

// get recommending TvPrograms.
func GetRecommendTvProgramsByUser(userID int64) (ml []interface{}) {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64
	w, _ := GetAllTvProgram(query, fields, sortby, order, offset, limit)
	limit = 10
	sortby = append(sortby, "Updated")
	order = append(order, "desc")
	query["Watched"] = "1"
	query["UserId"] = strconv.FormatInt(userID, 10)
	v, _ := GetAllWatchingStatus(query, fields, sortby, order, offset, limit)
	if len(v) == 0 {
		return nil
	}
	var Points []RecommendPoint
	for _, tvProgram := range w {
		x := RecommendPoint{
			Index: tvProgram.(TvProgram).Id,
			Point: 0,
		}
		Points = append(Points, x)
	}

	for _, watched := range v {
		if r, err := GetTvProgramById(watched.(WatchingStatus).TvProgramId); err == nil {
			casts := strings.Split(r.Cast, ",")
			// 「見た」番組に出演しているキャストの他の作品
			for _, cast := range casts {
				for index, tvProgram := range w {
					if strings.Contains(tvProgram.(TvProgram).Cast, cast) {
						if r.Id != tvProgram.(TvProgram).Id {
							Points[index].Point++
						}
					}
				}
			}
		}
	}

	sort.Slice(Points, func(i, j int) bool {
		return Points[i].Point > Points[j].Point
	})
	var displayNum int = 100
	if len(Points) < displayNum {
		displayNum = len(Points)
	}
	for _, recommendPoint := range Points[:displayNum] {
		r, _ := GetTvProgramById(recommendPoint.Index)
		ml = append(ml, *r)
	}
	return ml
}

// 正規表現で置換
func RegexpWords(str string, word string, repWord string) (res string) {
	rep := regexp.MustCompile(word)
	res = rep.ReplaceAllString(str, repWord)
	return res
}
func ReshapeWordsA(str string) (res string) {
	str = strings.Replace(str, "　", ",", -1)
	res = strings.Replace(str, " ", ",", -1)
	return res
}

// 入力されたMovieURLのチェック
func ReshapeMovieURL(str string) (res string) {
	if !strings.Contains(str, "https") {
		res = ""
	} else if strings.Contains(str, "https://www.youtube.com/watch?v=") {
		res = strings.Replace(str, "watch?v=", "embed/", -1)
	} else if strings.Contains(str, "https://youtu.be/") {
		res = strings.Replace(str, "youtu.be/", "www.youtube.com/embed/", -1)
	} else {
		res = ""
	}
	return res
}
func ReshapeImageURL(str string) (res string) {
	if !strings.Contains(str, "http") {
		rand.Seed(time.Now().UnixNano())
		r := strconv.Itoa(rand.Intn(10) + 1)
		if len(r) == 1 {
			r = "0" + r
		}
		res = "/static/img/tv_img/hanko_" + r + ".png"
	} else {
		res = str
	}
	return res
}
func ReshapeImageURLReference(str string) (res string) {
	res = ""
	if str != "" {
		if strings.Contains(str, "walkerplus") {
			res = "MovieWalker"
		} else if strings.Contains(str, "1.bp.blogspot.com") {
			res = "いらすとや"
		} else if strings.Contains(str, "/static/img") {
			res = ""
		} else {
			imageURLs := strings.Split(str, "/")
			res = imageURLs[2]
			res = strings.Replace(res, "www.", "", 1)
		}
	}
	return res
}
