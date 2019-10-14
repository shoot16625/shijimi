package main

import (
	"app/models"
	_ "app/routers"

	// "app/db"
	"fmt"
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
	orm.RegisterDataBase("default", beego.AppConfig.String("driver"), beego.AppConfig.String("sqlconn")+"?charset=utf8&loc=Asia%2FTokyo")
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
		CookieName: "ShiJimi",
		Gclifetime: 7200,
	}
	beego.GlobalSessions, _ = session.NewManager("memory", sessionconf)
	go beego.GlobalSessions.GC()

	// 初期データの投入
	execInitSQL()
	execSQL()
	// db.Scraping()

}

func execInitSQL() {
	o := orm.NewOrm()
	o.Using("default")
	for i, u := range [9]string{"月", "火", "水", "木", "金", "土", "日", "スペシャル", "映画"} {
		v := new(models.Week)
		v.Name = u
		v.Id = i + 1
		o.Insert(v)
	}
	for i, u := range [4]string{"冬", "春", "夏", "秋"} {
		v := new(models.Season)
		v.Name = u
		v.Id = i + 1
		o.Insert(v)
	}
}

func execSQL() {
	for i := 1; i < 5; i++ {
		TvProgramSQL("偽装不倫_"+strconv.Itoa(i), "濱鐘子は32歳で独身。周りからは「パラサイトシングル」や「婚活疲れした派遣社員」など不名誉なレッテルが貼られる始末。2年間の婚活も成就せず、自分はモテない女だと自覚するようになってしまう。「おひとり様」生活がしっくりくるようになり、1年間の派遣社員生活を終え、婚活に別れを告げるために1人旅をしようと計画する。 そんな中、1人旅をするために乗り込んだ飛行機の中で1人の若者に出会い、自分は「既婚者」だと言い様のない嘘を吐いてしまう。それでも若者は「この旅行の間だけでも不倫をしましょう」と迫ってきて、鐘子は恋の楽しさを覚える深みにはまっていく。", "https://www.ntv.co.jp/gisouhurin/images/ogp.jpg", "ntv", "https://www.youtube.com/embed/6Wxy-Gr3VgM", "7月期新水曜ドラマ『偽装不倫』7月10日（水）よる10時スタート／プロローグ［日本テレビ］", "杏、宮沢氷魚、谷原章介、仲間由紀恵", "恋愛、アラサー、OL", "衛藤凛", "", "鈴木勇馬、南雲聖一", "日テレ", 2019, "秋", "milet 「us」", "水", float32(i), 3, 0, 0, 0)

		for j := 1; j < 11; j++ {
			CommentSQL("面白いドラマ\r\n今後も楽しみ＿"+strconv.Itoa(i), 1, int64(j), 0)
			// CommentSQL("good tv　"+strconv.Itoa(i), 1, int64(j), 0)
			CommentSQL("Yes＿"+strconv.Itoa(i), 2, int64(j), 0)
		}
		UserSQL("shuto"+strconv.Itoa(i), "password", 22, "男性", "愛知県", "学生", "乃木"+strconv.Itoa(i), "http://blog-imgs-34.fc2.com/m/i/n/minamijima/Mx1mTAr153j4wz96j8npWxBF_500.jpg", "未婚")

		// WatchingStatusSQL(1,int64(i),true,false)
		// CommentLikeSQL(1, 0, true)
		// ReviewCommentLikeSQL(1, int64(i), true)
		ReviewCommentSQL("レビューネタバレあり\n"+strconv.Itoa(i), 1, int64(i), 0, true, "神曲", int32(1))
		ReviewCommentSQL("レビューネタバレなし\n"+strconv.Itoa(i), 2, int64(i), 0, false, "泣きっぱなし、演技すごい", int32(1))
	}
	TvProgramSQL("偽装不倫", "濱鐘子は3", "https://www.ntv.co.jp/gisouhurin/images/ogp.jpg", "ntv", "https://www.youtube.com/embed/u22OYxxAnhs", "7月期新水曜ドラマ『偽装不倫』7月10日（水）よる10時スタート／プロローグ［日本テレビ］", "杏、宮沢氷魚、谷原章介、仲間由紀恵", "恋愛、アラサー、OL", "衛藤凛", "", "鈴木勇馬、南雲聖一", "TBS", 2020, "秋", "milet 「us」", "火", 9, 3, 0, 0, 0)
	TvProgramSQL("偽装不倫a", "濱鐘子は3", "https://www.ntv.co.jp/gisouhurin/images/ogp.jpg", "ntv", "https://www.youtube.com/embed/u22OYxxAnhs", "7月期新水曜ドラマ『偽装不倫』7月10日（水）よる10時スタート／プロローグ［日本テレビ］", "杏、宮沢氷魚、谷原章介、仲間由紀恵", "恋愛、アラサー、OL", "衛藤凛", "", "鈴木勇馬、南雲聖一", "TBS", 2020, "夏", "milet 「us」", "火", 9, 3, 0, 0, 0)
	TvProgramSQL("偽装不倫b", "濱鐘子は3", "https://www.ntv.co.jp/gisouhurin/images/ogp.jpg", "ntv", "https://www.youtube.com/embed/u22OYxxAnhs", "7月期新水曜ドラマ『偽装不倫』7月10日（水）よる10時スタート／プロローグ［日本テレビ］", "杏、宮沢氷魚、谷原章介、仲間由紀恵", "恋愛、アラサー、OL", "衛藤凛", "", "鈴木勇馬、南雲聖一", "TBS", 2020, "春", "milet 「us」", "火", 9, 3, 0, 0, 0)
	TvProgramSQL("偽装不倫c", "濱鐘子は3", "https://www.ntv.co.jp/gisouhurin/images/ogp.jpg", "ntv", "https://www.youtube.com/embed/u22OYxxAnhs", "7月期新水曜ドラマ『偽装不倫』7月10日（水）よる10時スタート／プロローグ［日本テレビ］", "杏、宮沢氷魚、谷原章介、仲間由紀恵", "恋愛、アラサー、OL", "衛藤凛", "", "鈴木勇馬、南雲聖一", "TBS", 2020, "冬", "milet 「us」", "火", 9, 3, 0, 0, 0)
	TvProgramSQL("偽装不倫d", "濱鐘子は3", "https://www.ntv.co.jp/gisouhurin/images/ogp.jpg", "ntv", "https://www.youtube.com/embed/u22OYxxAnhs", "7月期新水曜ドラマ『偽装不倫』7月10日（水）よる10時スタート／プロローグ［日本テレビ］", "杏、宮沢氷魚、谷原章介、仲間由紀恵", "恋愛、アラサー、OL", "衛藤凛", "", "鈴木勇馬、南雲聖一", "TBS", 2020, "秋", "milet 「us」", "火", 10, 3, 0, 0, 0)
	TvProgramSQL("偽装不倫e", "濱鐘子は3", "https://www.ntv.co.jp/gisouhurin/images/ogp.jpg", "ntv", "https://www.youtube.com/embed/u22OYxxAnhs", "7月期新水曜ドラマ『偽装不倫』7月10日（水）よる10時スタート／プロローグ［日本テレビ］", "杏、宮沢氷魚、谷原章介、仲間由紀恵", "恋愛、アラサー、OL", "衛藤凛", "", "鈴木勇馬、南雲聖一", "TBS", 2020, "秋", "milet 「us」", "水", 10, 3, 0, 0, 0)
	TvProgramSQL("偽装不倫f", "濱鐘子は3", "https://www.ntv.co.jp/gisouhurin/images/ogp.jpg", "ntv", "https://www.youtube.com/embed/u22OYxxAnhs", "7月期新水曜ドラマ『偽装不倫』7月10日（水）よる10時スタート／プロローグ［日本テレビ］", "杏、宮沢氷魚、谷原章介、仲間由紀恵", "恋愛、アラサー、OL", "衛藤凛", "", "鈴木勇馬、南雲聖一", "TBS", 2020, "秋", "milet 「us」", "木", 10, 3, 0, 0, 0)
	TvProgramSQL("偽装不倫g", "濱鐘子は3", "https://www.ntv.co.jp/gisouhurin/images/ogp.jpg", "ntv", "https://www.youtube.com/embed/u22OYxxAnhs", "7月期新水曜ドラマ『偽装不倫』7月10日（水）よる10時スタート／プロローグ［日本テレビ］", "杏、宮沢氷魚、谷原章介、仲間由紀恵", "恋愛、アラサー、OL", "衛藤凛", "", "鈴木勇馬、南雲聖一", "TBS", 2020, "秋", "milet 「us」", "木", 10, 3, 0, 0, 0)
	TvProgramSQL("偽装不倫h", "濱鐘子は3", "https://www.ntv.co.jp/gisouhurin/images/ogp.jpg", "ntv", "https://www.youtube.com/embed/u22OYxxAnhs", "7月期新水曜ドラマ『偽装不倫』7月10日（水）よる10時スタート／プロローグ［日本テレビ］", "杏、宮沢氷魚、谷原章介、仲間由紀恵", "恋愛、アラサー、OL", "衛藤凛", "", "鈴木勇馬、南雲聖一", "TBS", 2020, "秋", "milet 「us」", "金", 10, 3, 0, 0, 0)
	TvProgramSQL("偽装不倫i", "濱鐘子は3", "https://www.ntv.co.jp/gisouhurin/images/ogp.jpg", "ntv", "https://www.youtube.com/embed/u22OYxxAnhs", "7月期新水曜ドラマ『偽装不倫』7月10日（水）よる10時スタート／プロローグ［日本テレビ］", "杏、宮沢氷魚、谷原章介、仲間由紀恵", "恋愛、アラサー、OL", "衛藤凛", "", "鈴木勇馬、南雲聖一", "TBS", 2020, "秋", "milet 「us」", "土", 10, 3, 0, 0, 0)
	// for i := 1; i < 3; i++ {
	// 	TvProgramSQL("Heaven? 〜ご苦楽レストラン〜_"+strconv.Itoa(i), "フレンチレストランで働いている伊賀観は、ある日ふとした事件がきっかけで黒須仮名子と出会う。黒須は新たにフレン。", "http://www.tbs.co.jp/Heaven_tbs/img/ogp02.jpg?58465464654", "tbs", "https://www.youtube.com/embed/6Wxy-Gr3VgM", "[新ドラマ]『Heaven?～ご苦楽レストラン～』7/9(火)スタート!! 石原さとみがフレンチレストランのオーナー役に!!【TBS】", "石原さとみ、福士蒼汰、志尊淳", "コメディ、フレンチ", "吉田恵里香", "", "木村ひさし、松木彩、村尾嘉昭", "TBS", int(2019), "夏", "あいみょん 「真夏の夜の匂いがする」", "火", float32(i), float32(3), int32(0), int32(0), int32(0))
	// }
	// WatchingStatusSQL(1,1,true,false)
}

func TvProgramSQL(title string, content string, imageurl string, imageurlreference string, movieurl string, movieurlreference string, cast string, category string, dramatist string, supervisor string, director string, production string, year int, season string, themesong string, week string, hour float32, star float32, countstar int32, countwatched int32, countwanttowatch int32) {
	o := orm.NewOrm()
	o.Using("default")
	v := new(models.TvProgram)
	v.Title = title
	v.Content = content
	v.ImageUrl = imageurl
	v.ImageUrlReference = imageurlreference
	v.MovieUrl = movieurl
	v.MovieUrlReference = movieurlreference
	v.Cast = cast
	v.Category = category
	v.Dramatist = dramatist
	v.Supervisor = supervisor
	v.Director = director
	v.Production = production
	v.Year = year
	// v.Season = season
	u := *new(models.Season)
	u.Name = season
	v.Season = &u
	// v.Season.Name = season
	v.Themesong = themesong
	w := *new(models.Week)
	w.Name = week
	v.Week = &w
	v.Hour = hour
	// v.Timezone = timezone
	// fmt.Println(3)
	// u := *new(models.Timezone)
	// u.Week = "月"
	// v.Timezone = &u
	v.Star = star
	v.CountStar = countstar
	v.CountWatched = countwatched
	v.CountWantToWatch = countwanttowatch
	o.Insert(v)
}

func CommentSQL(content string, userID int64, tvprogramID int64, countlike int32) {
	o := orm.NewOrm()
	o.Using("default")
	v := new(models.Comment)
	v.Content = content
	v.UserId = userID
	v.TvProgramId = tvprogramID
	v.CountLike = countlike
	o.Insert(v)
}

func CommentLikeSQL(userID int64, commentID int64, like bool) {
	o := orm.NewOrm()
	o.Using("default")
	v := new(models.CommentLike)
	v.UserId = userID
	v.CommentId = commentID
	v.Like = like
	o.Insert(v)
}

func ReviewCommentSQL(content string, userID int64, tvprogramID int64, countlike int32, spoiler bool, FavoritePoint string, star int32) {
	o := orm.NewOrm()
	o.Using("default")
	v := new(models.ReviewComment)
	v.Content = content
	v.UserId = userID
	v.TvProgramId = tvprogramID
	v.FavoritePoint = FavoritePoint
	v.CountLike = countlike
	v.Spoiler = spoiler
	v.Star = star
	o.Insert(v)
}

func ReviewCommentLikeSQL(userID int64, reviewcommentID int64, like bool) {
	o := orm.NewOrm()
	o.Using("default")
	v := new(models.ReviewCommentLike)
	v.UserId = userID
	v.ReviewCommentId = reviewcommentID
	v.Like = like
	o.Insert(v)
}

func UserSQL(username string, password string, age int, gender string, address string, job string, secondpassword string, iconurl string, marital string) {
	o := orm.NewOrm()
	o.Using("default")
	v := new(models.User)
	v.Username = username
	hashPass, _ := models.PasswordHash(password)
	v.Password = hashPass
	v.Age = age
	v.Gender = gender
	v.Address = address
	v.Job = job
	hashSecondpass, _ := models.PasswordHash(secondpassword)
	v.SecondPassword = hashSecondpass
	v.IconUrl = iconurl
	v.Marital = marital
	o.Insert(v)
}

func WatchingStatusSQL(userID int64, tvprogramID int64, watched bool, WantToWatch bool) {
	o := orm.NewOrm()
	o.Using("default")
	v := new(models.WatchingStatus)
	v.UserId = userID
	v.TvProgramId = tvprogramID
	v.Watched = watched
	v.WantToWatch = WantToWatch
	o.Insert(v)
}
