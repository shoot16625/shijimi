package routers

import (
	"app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.TvProgramController{})
	ns := beego.NewNamespace("/tv",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
			beego.NSRouter("/create", &controllers.UserController{}, "*:Create"),
			beego.NSRouter("/index", &controllers.UserController{}, "*:Index"),
			beego.NSRouter("/login", &controllers.UserController{}, "*:Login"),
			beego.NSRouter("/login_admin_page", &controllers.UserController{}, "*:LoginAdminPage"),
			beego.NSRouter("/login_admin", &controllers.UserController{}, "*:LoginAdmin"),
			beego.NSRouter("/logout", &controllers.UserController{}, "*:Logout"),
			beego.NSRouter("/forget_username_page", &controllers.UserController{}, "*:ForgetUsernamePage"),
			beego.NSRouter("/forget_username", &controllers.UserController{}, "*:ForgetUsername"),
			beego.NSRouter("/forget_password_page", &controllers.UserController{}, "*:ForgetPasswordPage"),
			beego.NSRouter("/forget_password", &controllers.UserController{}, "*:ForgetPassword"),
			beego.NSRouter("/show", &controllers.UserController{}, "*:Show"),
			beego.NSRouter("/show_review", &controllers.UserController{}, "*:ShowReview"),
			beego.NSRouter("/show_watched_tv", &controllers.UserController{}, "*:ShowWatchedTv"),
			beego.NSRouter("/show_wtw_tv", &controllers.UserController{}, "*:ShowWtwTv"),
			beego.NSRouter("/edit", &controllers.UserController{}, "*:Edit"),
			beego.NSRouter("/show/:id", &controllers.UserController{}, "*:Show"),
			beego.NSRouter("/show_review/:id", &controllers.UserController{}, "*:ShowReview"),
		),
		beego.NSNamespace("/comment",
			beego.NSInclude(
				&controllers.CommentController{},
			),
			beego.NSRouter("/update/:id/:top", &controllers.CommentController{}, "*:GetNewComments"),
		),
		beego.NSNamespace("/tv_program",
			beego.NSInclude(
				&controllers.TvProgramController{},
			),
			beego.NSRouter("/index", &controllers.TvProgramController{}, "*:Index"),
			beego.NSRouter("/edit/:id", &controllers.TvProgramController{}, "*:EditPage"),
			// beego.NSRouter("/create", &controllers.TvProgramController{},"*:Create"),
			beego.NSRouter("/create_page", &controllers.TvProgramController{}, "*:CreatePage"),
			beego.NSRouter("/get_tv_info", &controllers.TvProgramController{}, "*:GetTvProgramWikiInfo"),
			beego.NSRouter("/get_movie_info", &controllers.TvProgramController{}, "*:GetMovieWikiInfo"),
			beego.NSRouter("/comment/:id", &controllers.CommentController{}, "*:Show"),
			beego.NSRouter("/comment/search_comment/:id", &controllers.CommentController{}, "*:SearchComment"),
			beego.NSRouter("/review/:id", &controllers.ReviewCommentController{}, "*:Show"),
			beego.NSRouter("/review/search_comment/:id", &controllers.ReviewCommentController{}, "*:SearchComment"),
			beego.NSRouter("/search/", &controllers.TvProgramController{}, "*:Search"),
			beego.NSRouter("/search_tv_program/", &controllers.TvProgramController{}, "*:SearchTvProgram"),
		),
		beego.NSNamespace("/watching_status",
			beego.NSInclude(
				&controllers.WatchingStatusController{},
			),
		),
		beego.NSNamespace("/review_comment",
			beego.NSInclude(
				&controllers.ReviewCommentController{},
			),
		),
		beego.NSNamespace("/review_comment_like",
			beego.NSInclude(
				&controllers.ReviewCommentLikeController{},
			),
		),
		beego.NSNamespace("/comment_like",
			beego.NSInclude(
				&controllers.CommentLikeController{},
			),
		),
		beego.NSNamespace("/login_history",
			beego.NSInclude(
				&controllers.LoginHistoryController{},
			),
		),
		beego.NSNamespace("/browsing_history",
			beego.NSInclude(
				&controllers.BrowsingHistoryController{},
			),
		),
		beego.NSNamespace("/search_history",
			beego.NSInclude(
				&controllers.SearchHistoryController{},
			),
		),
		beego.NSNamespace("/tv_program_update_history",
			beego.NSInclude(
				&controllers.TvProgramUpdateHistoryController{},
			),
		),
		beego.NSNamespace("/foot_print_to_user",
			beego.NSInclude(
				&controllers.FootPrintToUserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
