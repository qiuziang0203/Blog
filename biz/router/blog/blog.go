// Code generated by hertz generator. DO NOT EDIT.

package blog

import (
	blog "Blog/biz/handler/blog"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_login := root.Group("/login", _loginMw()...)
		{
			_likecomment := _login.Group("/LikeComment", _likecommentMw()...)
			_likecomment.POST("/", append(_likecomment0Mw(), blog.LikeComment)...)
		}
		{
			_admin := _login.Group("/admin", _adminMw()...)
			{
				_addannouncement := _admin.Group("/addAnnouncement", _addannouncementMw()...)
				_addannouncement.OPTIONS("/", append(_addannouncement0Mw(), blog.AddAnnouncement)...)
				_addannouncement.POST("/", append(_addannouncement1Mw(), blog.AddAnnouncement)...)
			}
			{
				_approbalblog := _admin.Group("/approbalBlog", _approbalblogMw()...)
				_approbalblog.OPTIONS("/", append(_approvalblogMw(), blog.ApprovalBlog)...)
				_approbalblog.PUT("/", append(_approvalblog0Mw(), blog.ApprovalBlog)...)
			}
			{
				_delannouncement := _admin.Group("/delAnnouncement", _delannouncementMw()...)
				_delannouncement.OPTIONS("/", append(_delannouncement0Mw(), blog.DelAnnouncement)...)
				_delannouncement.PUT("/", append(_delannouncement1Mw(), blog.DelAnnouncement)...)
			}
			{
				_forbiddenuser := _admin.Group("/forbiddenUser", _forbiddenuserMw()...)
				_forbiddenuser.OPTIONS("/", append(_forbiddenuser0Mw(), blog.ForbiddenUser)...)
				_forbiddenuser.PUT("/", append(_forbiddenuser1Mw(), blog.ForbiddenUser)...)
			}
			{
				_getforbidden := _admin.Group("/getForbidden", _getforbiddenMw()...)
				_getforbidden.GET("/", append(_getforbidden0Mw(), blog.GetForbidden)...)
				_getforbidden.OPTIONS("/", append(_getforbidden1Mw(), blog.GetForbidden)...)
			}
			{
				_getpending := _admin.Group("/getpending", _getpendingMw()...)
				_getpending.GET("/", append(_getpending0Mw(), blog.GetPending)...)
				_getpending.OPTIONS("/", append(_getpending1Mw(), blog.GetPending)...)
			}
		}
		{
			_changedraft := _login.Group("/changeDraft", _changedraftMw()...)
			_changedraft.PUT("/", append(_chengedraftMw(), blog.ChengeDraft)...)
		}
		{
			_changeuserinfo := _login.Group("/changeUserInfo", _changeuserinfoMw()...)
			_changeuserinfo.PUT("/", append(_changeuserinfo0Mw(), blog.ChangeUserInfo)...)
		}
		{
			_commentblog := _login.Group("/commentBlog", _commentblogMw()...)
			_commentblog.OPTIONS("/", append(_commentblog0Mw(), blog.CommentBlog)...)
			_commentblog.POST("/", append(_commentblog1Mw(), blog.CommentBlog)...)
		}
		{
			_createblog := _login.Group("/createblog", _createblogMw()...)
			_createblog.OPTIONS("/", append(_createblog0Mw(), blog.CreateBlog)...)
			_createblog.POST("/", append(_createblog1Mw(), blog.CreateBlog)...)
		}
		{
			_delblog := _login.Group("/delBlog", _delblogMw()...)
			_delblog.PUT("/", append(_delblog0Mw(), blog.DelBlog)...)
		}
		{
			_favorblog := _login.Group("/favorBlog", _favorblogMw()...)
			_favorblog.OPTIONS("/", append(_favorblog0Mw(), blog.FavorBlog)...)
			_favorblog.PUT("/", append(_favorblog1Mw(), blog.FavorBlog)...)
		}
		{
			_followuser := _login.Group("/followUser", _followuserMw()...)
			_followuser.OPTIONS("/", append(_followuser0Mw(), blog.FollowUser)...)
			_followuser.POST("/", append(_followuser1Mw(), blog.FollowUser)...)
		}
		{
			_getdraft := _login.Group("/getDraft", _getdraftMw()...)
			_getdraft.GET("/", append(_getdraft0Mw(), blog.GetDraft)...)
		}
		{
			_getfavor := _login.Group("/getFavor", _getfavorMw()...)
			_getfavor.GET("/", append(_getfavor0Mw(), blog.GetFavor)...)
			_getfavor.OPTIONS("/", append(_getfavor1Mw(), blog.GetFavor)...)
		}
		{
			_getfollow := _login.Group("/getFollow", _getfollowMw()...)
			_getfollow.GET("/", append(_getfollow0Mw(), blog.GetFollow)...)
			_getfollow.OPTIONS("/", append(_getfollow1Mw(), blog.GetFollow)...)
		}
		{
			_getuser := _login.Group("/getUser", _getuserMw()...)
			_getuser.GET("/", append(_getuserinfoMw(), blog.GetUserInfo)...)
		}
		{
			_likeblog := _login.Group("/likeBlog", _likeblogMw()...)
			_likeblog.OPTIONS("/", append(_likeblog0Mw(), blog.LikeBlog)...)
			_likeblog.PUT("/", append(_likeblog1Mw(), blog.LikeBlog)...)
		}
		{
			_logout := _login.Group("/logout", _logoutMw()...)
			_logout.GET("/", append(_logout0Mw(), blog.Logout)...)
			_logout.OPTIONS("/", append(_logout1Mw(), blog.Logout)...)
		}
		{
			_recomment := _login.Group("/reComment", _recommentMw()...)
			_recomment.POST("/", append(_recomment0Mw(), blog.ReComment)...)
		}
		{
			_upload := _login.Group("/upload", _uploadMw()...)
			_upload.OPTIONS("/", append(_upload0Mw(), blog.Upload)...)
			_upload.POST("/", append(_upload1Mw(), blog.Upload)...)
		}
	}
	{
		_unlogin := root.Group("/unlogin", _unloginMw()...)
		{
			_getalltype := _unlogin.Group("/getAllType", _getalltypeMw()...)
			_getalltype.GET("/", append(_getalltype0Mw(), blog.GetAllType)...)
			_getalltype.OPTIONS("/", append(_getalltype1Mw(), blog.GetAllType)...)
		}
		{
			_getannouncement := _unlogin.Group("/getAnnouncement", _getannouncementMw()...)
			_getannouncement.GET("/", append(_getannouncement0Mw(), blog.GetAnnouncement)...)
			_getannouncement.OPTIONS("/", append(_getannouncement1Mw(), blog.GetAnnouncement)...)
		}
		{
			_getblogbyid := _unlogin.Group("/getBlogByID", _getblogbyidMw()...)
			_getblogbyid.GET("/", append(_getblogbyid0Mw(), blog.GetBlogByID)...)
			_getblogbyid.OPTIONS("/", append(_getblogbyid1Mw(), blog.GetBlogByID)...)
		}
		{
			_getblogbykeyword := _unlogin.Group("/getBlogByKeyword", _getblogbykeywordMw()...)
			_getblogbykeyword.GET("/", append(_getblogbykeyword0Mw(), blog.GetBlogByKeyword)...)
			_getblogbykeyword.OPTIONS("/", append(_getblogbykeyword1Mw(), blog.GetBlogByKeyword)...)
		}
		{
			_getblogbytype := _unlogin.Group("/getBlogByType", _getblogbytypeMw()...)
			_getblogbytype.GET("/", append(_getblogbytype0Mw(), blog.GetBlogByType)...)
			_getblogbytype.OPTIONS("/", append(_getblogbytype1Mw(), blog.GetBlogByType)...)
		}
		{
			_getblogcomment := _unlogin.Group("/getBlogComment", _getblogcommentMw()...)
			_getblogcomment.GET("/", append(_getblogcomment0Mw(), blog.GetBlogComment)...)
			_getblogcomment.OPTIONS("/", append(_getblogcomment1Mw(), blog.GetBlogComment)...)
		}
		{
			_gethotblog := _unlogin.Group("/getHotBlog", _gethotblogMw()...)
			_gethotblog.GET("/", append(_gethotblog0Mw(), blog.GetHotBlog)...)
			_gethotblog.OPTIONS("/", append(_gethotblog1Mw(), blog.GetHotBlog)...)
		}
		{
			_getrecomment := _unlogin.Group("/getReComment", _getrecommentMw()...)
			_getrecomment.GET("/", append(_getrecomment0Mw(), blog.GetReComment)...)
			_getrecomment.OPTIONS("/", append(_getrecomment1Mw(), blog.GetReComment)...)
		}
		{
			_getuserblog := _unlogin.Group("/getUserBlog", _getuserblogMw()...)
			_getuserblog.GET("/", append(_getuserblog0Mw(), blog.GetUserBlog)...)
			_getuserblog.OPTIONS("/", append(_getuserblog1Mw(), blog.GetUserBlog)...)
		}
		{
			_getuserbyid := _unlogin.Group("/getUserByID", _getuserbyidMw()...)
			_getuserbyid.GET("/", append(_getuserinfobyidMw(), blog.GetUserInfoByID)...)
			_getuserbyid.OPTIONS("/", append(_getuserinfobyid0Mw(), blog.GetUserInfoByID)...)
		}
		{
			_getuserbykeyword := _unlogin.Group("/getUserByKeyword", _getuserbykeywordMw()...)
			_getuserbykeyword.GET("/", append(_getuserbykeyword0Mw(), blog.GetUserByKeyword)...)
			_getuserbykeyword.OPTIONS("/", append(_getuserbykeyword1Mw(), blog.GetUserByKeyword)...)
		}
		{
			_passwordlogin := _unlogin.Group("/passwordLogin", _passwordloginMw()...)
			_passwordlogin.OPTIONS("/", append(_passwordlogin0Mw(), blog.PasswordLogin)...)
			_passwordlogin.POST("/", append(_passwordlogin1Mw(), blog.PasswordLogin)...)
		}
		{
			_register := _unlogin.Group("/register", _registerMw()...)
			_register.OPTIONS("/", append(_register0Mw(), blog.Register)...)
			_register.POST("/", append(_register1Mw(), blog.Register)...)
		}
	}
}