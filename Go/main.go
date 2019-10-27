package main

import (
	"app/db"
	_ "app/routers"

	"fmt"
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
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func init() {
	// タイムゾーンを日本に設定
	loc, e := time.LoadLocation(location)
	if e != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc

	orm.RegisterDriver(beego.AppConfig.String("driver"), orm.DRMySQL)
	orm.RegisterDataBase("default", beego.AppConfig.String("driver"), beego.AppConfig.String("sqlconn")+"?charset=utf8mb4&loc=Asia%2FTokyo")
	// データを初期化して起動
	err := orm.RunSyncdb("default", true, false)
	// データの変更点を追加して起動
	// err := orm.RunSyncdb("default", false, false)
	if err != nil {
		fmt.Println(err)
	}

	// tpl内で使える関数
	beego.AddFuncMap("dateformatJst", func(in time.Time) string {
		return in.Format("2006-01-02 15:04:05")
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
	db.ExecInitSQL()
	db.ExecTestSQL()
	// db.GetTvProgramInformationByURL("https://ja.wikipedia.org/wiki/G線上のあなたと私")
	db.AddRecentTvInfo()
	// db.GetWikiDoramas("https://ja.wikipedia.org/wiki/日本のテレビドラマ一覧_(2010年代)")
	// db.GetWikiDoramas("https://ja.wikipedia.org/wiki/日本のテレビドラマ一覧_(2000年代)")
	db.UpdateTvProgramsInformation()
	// db.GetMovieWalkers()
}
