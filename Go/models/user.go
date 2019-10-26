package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id             int64  `orm:"auto"`
	Username       string `orm:"size(30);unique"`
	Password       string `orm:"size(500)" json:"-"`
	Age            int
	Gender         string    `orm:"size(20)"`
	Address        string    `orm:"size(20)"`
	Job            string    `orm:"size(20)"`
	SecondPassword string    `orm:"size(500)" json:"-"`
	IconURL        string    `orm:"size(255);null"`
	Marital        string    `orm:"size(20);null"`
	BloodType      string    `orm:"size(20);null"`
	Created        time.Time `orm:"auto_now_add;type(datetime)"`
	Updated        time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int64) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.QueryTable(new(User)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
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

	var l []User
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

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int64) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetUserByUsername(username string) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Username: username}
	if err = o.QueryTable(new(User)).Filter("Username", username).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// ユーザ名の検索
func GetUserByPasswords(password string, SecondPassword string) (v *User, err error) {
	o := orm.NewOrm()
	var u []User
	num := 0
	var d User
	if _, err = o.QueryTable(new(User)).All(&u); err == nil {
		for _, value := range u {
			if UserPassMach(value.Password, password) {
				if UserPassMach(value.SecondPassword, SecondPassword) {
					num++
					d = value
				}
			}
		}
		if num == 1 {
			return &d, nil
		}
	}
	return nil, err
}

// ユーザー名と第2パスワードを使ってパスワードの再設定
func GetUserByUsernameAndPassword(username string, SecondPassword string) (v *User, err error) {
	v, err = GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if UserPassMach(v.SecondPassword, SecondPassword) {
		return v, nil
	} else {
		return nil, err
	}
}

// パスワードのハッシュ化
func PasswordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// 入力パスワードが合っているか判定
func UserPassMach(hash, pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw)) == nil
}

// The number of users.
func GetUserCount() (cnt int64) {
	o := orm.NewOrm()
	cnt, _ = o.QueryTable(new(User)).Count()
	return cnt
}
