package db

import (
	"app/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

func ExecInitSQL() {
	o := orm.NewOrm()
	o.Using("default")
	for i, u := range [11]string{"月", "火", "水", "木", "金", "土", "日", "平日", "スペシャル", "映画", "?"} {
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
	// 管理者用 ID：1
	UserSQL("doramaba-admin", "doramaba-password", 199505, "男性", "愛知県", "学生", "doramaba-password", "/static/img/shijimi-transparence.png", "未婚", "A型")

	// uchida用 ID：2
	UserSQL("ちゃお倉木", "password", 199505, "男性", "愛知県", "学生", "乃木小学校", "https://img.cinematoday.jp/a/N0077397/_size_640x/_v_1445346612/1.jpg", "未婚", "A型")

	// お問い合わせ用 ID：1
	TvProgramSQL("お問い合わせ専用", "サービス改善のため、忌憚のないご意見・ご感想をお待ちしております。3回に1回くらい褒めていただけると幸いです。", "/static/img/shijimi-transparence.png", "", "", "", "大学院生", "コメディ・パロディ", "松江育ち", "単独開発Help!!", "日記感覚で", "shijimi", 2019, "秋", "milet", "日", 100, 2.5, 0, 0, 0)
}

func ExecTestSQL() {
	for i := 1; i < 3; i++ {
		UserSQL("test-user-"+strconv.Itoa(i), "password", 199001, "男性", "愛知県", "学生", "乃木", "http://blog-imgs-34.fc2.com/m/i/n/minamijima/Mx1mTAr153j4wz96j8npWxBF_500.jpg", "未婚", "A型")
	}

	for i := 1; i < 5; i++ {
		TvProgramSQL("TestTest:"+strconv.Itoa(i), "hogehogehogehogehogehogehogehogehogehogehogehogehogehogehogehogehogehogehogehogehogehogehogehoge", "https://1.bp.blogspot.com/-dkBk4bYQrTk/XVKfloSYxiI/AAAAAAABUC8/j6K3SGQG0WMxKFn71LzznPz0SPgI5ufGQCLcBGAs/s1600/bird_sekisei_inko_blue.png", "いらすとや", "https://www.youtube.com/embed/AIMjbleH394", "milet「us」MUSIC VIDEO（日本テレビ系水曜ドラマ『偽装不倫』主題歌）", "TestA、TestB、TestC、TestD、TestE、TestF", "恋愛、不倫、コメディ・パロディ", "TestG、TestH", "TestI、TestJ", "TestK、TestL", "日テレ", 2019, "秋", "milet 「us」", "月", float32(i+18), 2.5, 0, 0, 0)

		for j := 1; j < 10; j++ {
			CommentSQL("hogehoge\r\nfugafuga\r\n"+strconv.Itoa(i), 3, int64(i+1), 0)
			CommentSQL("hogehogefugafugahogehogefugafugahogehogefugafuga\r\n"+strconv.Itoa(i), 4, int64(i+1), 0)
		}
		ReviewCommentSQL("レビューネタバレありコメント\nレビューは一人一回\n", 3, int64(i+1), 0, true, "神曲", int32(6))
		ReviewCommentSQL("レビューネタバレなしコメント\n", 4, int64(i+1), 0, false, "泣きっぱなし、演技すごい", int32(4))
		fmt.Println("update:", i)
	}
}

func ExecDemoSQL() {
	UserSQL("ユーザA", "password", 199001, "男性", "愛知県", "学生", "乃木", "http://flat-icon-design.com/f/f_object_161/s512_f_object_161_0bg.png", "未婚", "A型")
	UserSQL("Bさん", "password", 199001, "男性", "愛知県", "学生", "乃木", "http://flat-icon-design.com/f/f_object_105/s512_f_object_105_0bg.png", "未婚", "A型")
	for j := 1; j < 20; j++ {
		CommentSQL("コメントを投稿（250字まで）\r\nコメントを投稿（250字まで）\r\n"+strconv.Itoa(j), 3, 19, int32(j))
		CommentSQL("桑野さん最高すぎる！\r\n"+strconv.Itoa(j), 4, 19, int32(j*3))
	}
	ReviewCommentSQL("レビューを投稿（450字まで）\nネタバレありです\nレビューは一人一回まで\n評価は10段階\nおすすめポイントタグ", 3, 19, 3, true, "神曲、ゆる～い", int32(6))
	ReviewCommentSQL("再放送4回みた。ELTは熱いよね！！\nネタバレはありません\n", 4, 19, 20, false, "泣きっぱなし、演技すごい", int32(8))
}

func TvProgramSQL(title string, content string, imageURL string, imageURLreference string, movieURL string, movieURLreference string, cast string, category string, dramatist string, supervisor string, director string, production string, year int, season string, themesong string, week string, hour float32, star float32, countstar int32, countWatched int32, countWantToWatch int32) {
	o := orm.NewOrm()
	o.Using("default")
	v := new(models.TvProgram)
	v.Title = title
	v.Content = content
	v.ImageURL = imageURL
	v.ImageURLReference = imageURLreference
	v.MovieURL = movieURL
	v.MovieURLReference = movieURLreference
	v.Cast = cast
	v.Category = category
	v.Dramatist = dramatist
	v.Supervisor = supervisor
	v.Director = director
	v.Production = production
	v.Year = year
	u := *new(models.Season)
	u.Name = season
	v.Season = &u
	w := *new(models.Week)
	w.Name = week
	v.Week = &w
	v.Themesong = themesong
	v.Hour = hour
	v.Star = star
	v.CountStar = countstar
	v.CountWatched = countWatched
	v.CountWantToWatch = countWantToWatch
	o.Insert(v)
}

func CommentSQL(content string, userID int64, tvProgramID int64, countlike int32) {
	o := orm.NewOrm()
	o.Using("default")
	v := new(models.Comment)
	v.Content = content
	v.UserId = userID
	v.TvProgramId = tvProgramID
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

func ReviewCommentSQL(content string, userID int64, tvProgramID int64, countlike int32, spoiler bool, FavoritePoint string, star int32) {
	o := orm.NewOrm()
	o.Using("default")
	v := new(models.ReviewComment)
	v.Content = content
	v.UserId = userID
	v.TvProgramId = tvProgramID
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

func UserSQL(username string, password string, age int, gender string, address string, job string, secondPassword string, IconURL string, marital string, bloodType string) {
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
	hashSecondpass, _ := models.PasswordHash(secondPassword)
	v.SecondPassword = hashSecondpass
	v.IconURL = IconURL
	v.Marital = marital
	v.BloodType = bloodType
	o.Insert(v)
}

func WatchingStatusSQL(userID int64, tvProgramID int64, watched bool, WantToWatch bool) {
	o := orm.NewOrm()
	o.Using("default")
	v := new(models.WatchingStatus)
	v.UserId = userID
	v.TvProgramId = tvProgramID
	v.Watched = watched
	v.WantToWatch = WantToWatch
	o.Insert(v)
}
