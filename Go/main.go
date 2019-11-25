package main

import (
	"app/db"
	_ "app/routers"

	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	_ "github.com/go-sql-driver/mysql"
)

const location = "Asia/Tokyo"

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	if beego.BConfig.RunMode == "prod" {
		port, _ := strconv.Atoi(os.Getenv("PORT"))
		fmt.Println(port)
		beego.BConfig.Listen.HTTPPort = port
		beego.BConfig.Listen.HTTPSPort = port
	}
	beego.BConfig.WebConfig.StaticDir["/manifest.json"] = "manifest.json"
	beego.BConfig.WebConfig.StaticDir["/serviceWorker.js"] = "serviceWorker.js"
	beego.Run()
}

func init() {
	// タイムゾーンを日本に設定
	loc, e := time.LoadLocation(location)
	if e != nil {
		fmt.Println(e)
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
	orm.RegisterDriver(beego.AppConfig.String("driver"), orm.DRMySQL)
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	sqlconn := user + ":" + pass + "@tcp(db:3306)/" + dbName
	orm.RegisterDataBase("default", beego.AppConfig.String("driver"), sqlconn+"?charset=utf8mb4&loc=Asia%2FTokyo")
	// orm.RegisterDataBase("default", beego.AppConfig.String("driver"), beego.AppConfig.String("sqlconn")+"?charset=utf8mb4&loc=Asia%2FTokyo")
	// データを初期化して起動
	// err := orm.RunSyncdb("default", true, false)
	// データの変更点を追加して起動
	err := orm.RunSyncdb("default", false, false)
	if err != nil {
		fmt.Println(err)
	}

	// tpl内で使える関数
	// 時刻表記
	beego.AddFuncMap("dateformatJst", func(in time.Time) string {
		return in.Format("2006-01-02 15:04:05")
	})
	// 年齢計算
	beego.AddFuncMap("birthday2Age", func(birthday string) (age string) {
		t := time.Now()
		birth, _ := time.Parse("2006-01-02", birthday)
		ageInt := t.Year() - birth.Year()
		t2 := time.Date(birth.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
		duration := t2.Sub(birth)
		if int(duration.Hours()) < 0 {
			ageInt--
		}
		age = strconv.Itoa(ageInt)
		return age
	})

	// formからDELETE・PUTをPOSTとできるようにする
	var FilterMethod = func(ctx *context.Context) {
		if ctx.Input.Query("_method") != "" && ctx.Input.IsPost() {
			ctx.Request.Method = strings.ToUpper(ctx.Input.Query("_method"))
		}
	}
	beego.InsertFilter("*", beego.BeforeRouter, FilterMethod)

	// クッキーを使えるようにする
	sessionconf := &session.ManagerConfig{
		CookieName:      "ShiJimi_Cookie",
		Gclifetime:      7200,
		Secure:          true,
		EnableSetCookie: true,
	}
	beego.GlobalSessions, _ = session.NewManager("memory", sessionconf)
	go beego.GlobalSessions.GC()

	// 初期データの投入
	db.EmptyInitSQL()

	// db.ExecInitSQL()
	// db.AddRecentTvInfo()
	// db.AddTvProgramsInformation()
	// db.GetMovieWalkers()
	// db.ExecDemoSQL()
	// db.ExecDemoHerokuSQL()
	// fmt.Println(db.GetTvProgramInformationByURL("https://ja.wikipedia.org/wiki/%E3%82%B3%E3%83%BC%E3%83%89%E3%83%BB%E3%83%96%E3%83%AB%E3%83%BC_-%E3%83%89%E3%82%AF%E3%82%BF%E3%83%BC%E3%83%98%E3%83%AA%E7%B7%8A%E6%80%A5%E6%95%91%E5%91%BD-"))
}
