

struct RegisterRequest {
    1: required string username, // 注册用户名，最长32个字符
    2: required string password, // 密码，最长32个字符
    3: required string nickname, // 昵称
    4: required string email, //邮箱
    5: required string sex,
    6: required string sign,
    7: required string birth
}
struct RegisterResponse {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required map<string,string> data       //响应数据
}
struct PasswordLoginRequest {
    1: required string username, // 注册用户名，最长32个字符
    2: required string password, // 密码，最长32个字符
}
struct PasswordLoginResponse {
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required map<string,string> data       //响应数据
}
struct CreateBlogRequest{
    1: required string title // 题目
    2: required string text // 正文
    3: required i64 status //状态：0-保存为草稿 1-提交审核
    4: required set<string> blog_type // 博客分类
}
struct  CreateBlogResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required map<string,string> data       //响应数据
}

//获取专栏博客
struct GetBlogByTypeRequest{
    1: required string type_name,
    2: required i64 status, // 排序规则 0-按时间 1-按热度
    3: required i64 page_num // 查询页数
}
struct GetBlogByTypeResponse{
     1: required i32 status_code,   // 状态码，0-成功，其他值-失败
     2: required string status_msg, // 返回状态描述
     3: required set<map<string,string>> data       //响应数据
}
//关键词查询博客
struct GetBlogByKeywordRequest{
    1: required string keyword,
    2: required i64 status, // 排序规则 0-按时间 1-按热度
    3: required i64 page_num // 查询页数
}
struct GetBlogByKeywordResponse{
     1: required i32 status_code,   // 状态码，0-成功，其他值-失败
     2: required string status_msg, // 返回状态描述
     3: required set<map<string,string>> data       //响应数据
}

//获取待审批列表
struct GetPendingRequest{
    1: required i64 page_num // 查询页数
}
struct GetPendingResponse{
     1: required i32 status_code,   // 状态码，0-成功，其他值-失败
     2: required string status_msg, // 返回状态描述
     3: required set<map<string,string>> data       //响应数据
}

//审批博客
struct ApprovalBlogRequest{
    1: required i64 blog_id
}
struct ApprovalBlogResponse{
     1: required i32 status_code,   // 状态码，0-成功，其他值-失败
     2: required string status_msg, // 返回状态描述
     3: required map<string,string> data       //响应数据
}
// 退出登录
struct LogoutRequest{
}
struct LogoutResponse{
     1: required i32 status_code,   // 状态码，0-成功，其他值-失败
     2: required string status_msg, // 返回状态描述
     3: required map<string,string> data       //响应数据
}
//查看封禁列表
struct GetForbiddenRequest{
    1: required i64 page_num // 查询页数
}
struct  GetForbiddenResponse{
     1: required i32 status_code,   // 状态码，0-成功，其他值-失败
     2: required string status_msg, // 返回状态描述
     3: required set<map<string,string>> data       //响应数据
}
// 封禁与解封用户
struct ForbiddenUserRequest{
    1: required i64 id //被封禁用户id
    2: required i64 status// 0-封禁 1-解封
}
struct ForbiddenUserResponse{
     1: required i32 status_code,   // 状态码，0-成功，其他值-失败
     2: required string status_msg, // 返回状态描述
     3: required map<string,string> data       //响应数据
}
// 点赞与取消赞
struct LikeBlogRequest{
    1:required i64 blog_id
    2:required i64 status // 0-点赞 1-取消赞
}
struct LikeBlogResponse{
     1: required i32 status_code,   // 状态码，0-成功，其他值-失败
     2: required string status_msg, // 返回状态描述
     3: required map<string,string> data       //响应数据
}
// 文件上传
struct UploadRequest{
}

struct UploadResponse{
     1: required i32 status_code,   // 状态码，0-成功，其他值-失败
     2: required string status_msg, // 返回状态描述
     3: required map<string,string> data       //响应数据
}
// 收藏与取消收藏
struct FavorBlogRequest{
    1:required i64 blog_id
    2:required i64 status // 0-点赞 1-取消赞
}
struct FavorBlogResponse{
     1: required i32 status_code,   // 状态码，0-成功，其他值-失败
     2: required string status_msg, // 返回状态描述
     3: required map<string,string> data       //响应数据
}
//获取我的收藏列表
struct GetFavorRequest{
    1:required i64 page_num
}
struct GetFavorResponse{
     1: required i32 status_code,   // 状态码，0-成功，其他值-失败
     2: required string status_msg, // 返回状态描述
     3: required set<map<string,string>> data       //响应数据
}

struct GetUserByKeywordRequest{
    1: required string keyword
    2: required i64 page_num
}
struct GetUserByKeywordResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required set<map<string,string>> data       //响应数据
}

struct FollowUserRequest{
    1: required i64 userid
    2: required i64 status
}
struct FollowUserResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required map<string,string> data       //响应数据
}
//获取关注列表
struct GetFollowRequest{
    1: required i64 page_num
}
struct GetFollowResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required set<map<string,string>> data       //响应数据
}
//评论博客
struct CommentBlogRequest{
    1:required i64 blog_id,
    2:required string text,
}
struct CommentBlogResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required map<string,string> data       //响应数据
}
//回复评论
struct ReCommentRequest{
     1:required i64 comment_id,
     2:required string text,
}
struct  ReCommentResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required map<string,string> data       //响应数据
}
//获取用户信息
struct GetUserInfoRequest{
}
struct GetUserInfoResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required map<string,string> data       //响应数据
}
//获取博客评论
struct GetBlogCommentRequest{
    1: required i64 blog_id
}
struct GetBlogCommentResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required set<map<string,string>> data       //响应数据
}
//获取评论回复
struct GetReCommentRequest{
    1: required i64 comment_id
}
struct GetReCommentResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required set<map<string,string>> data       //响应数据
}
// 获取草稿
struct GetDraftRequest{
    1: required i64 page_num
}
struct GetDraftResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required set<map<string,string>> data       //响应数据
}
//查看用户博客
struct GetUserBlogRequest{
    1: required i64 page_num
    2: required i64 user_id
}
struct GetUserBlogResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required set<map<string,string>> data       //响应数据
}
//修改用户信息
struct ChangeUserInfoRequest{
    1:required string avatar,
    2:required string nick_name,
    3:optional string password
    4:required string sex
    5:required string birth
    6:required string sign
}
struct ChangeUserInfoResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required map<string,string> data       //响应数据
}
//修改草稿
struct ChangeDraftRequest{
    1: required i64 blog_id,
    2: required string title,
    3: required string text,
    4: required i64 status
    5: required set<string> blog_type // 博客分类
}
struct ChangeDraftResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required map<string,string> data       //响应数据
}
struct LikeCommentRequest{
    1: required i64 comment_id
    2: required i64 status
}
struct LikeCommentResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required map<string,string> data       //响应数据
}
struct GetHotBlogRequest{
}
struct GetHotBlogResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required set<map<string,string>> data       //响应数据
}
struct GetAllTypeRequest{
}
struct GetAllTypeResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required set<map<string,string>> data       //响应数据
}
//添加公告
struct AddAnnouncementRequest{
    1: required string title
    2: required string text
}
struct AddAnnouncementResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required map<string,string> data       //响应数据
}
//查看公告
struct GetAnnouncementRequest{
}
struct GetAnnouncementResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required set<map<string,string>> data       //响应数据
}
//删除公告
struct DelAnnouncementRequest{
    1: required i64 id
}
struct DelAnnouncementResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required map<string,string> data       //响应数据
}
//根据id查询博客
struct GetBlogByIDRequest{
    1:required i64 blog_id
}
struct GetBlogByIDResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required map<string,string> data       //响应数据
}
//删除博客
struct DelBlogRequest{
    1:required i64 blog_id
}
struct DelBlogResponse{
    1: required i32 status_code,   // 状态码，0-成功，其他值-失败
    2: required string status_msg, // 返回状态描述
    3: required map<string,string> data       //响应数据
}
struct GetUserInfoByIDRequest{
    1:required i64 user_id
}
struct GetUserInfoByIDResponse{
        1: required i32 status_code,   // 状态码，0-成功，其他值-失败
        2: required string status_msg, // 返回状态描述
        3: required map<string,string> data       //响应数据
}

//未登录状态服务
service Unlogin{
    RegisterResponse Register(1:RegisterRequest request)(api.post="/unlogin/register/",api.options = "/unlogin/register/")
    PasswordLoginResponse PasswordLogin(PasswordLoginRequest requset)(api.post = "/unlogin/passwordLogin/",api.options = "/unlogin/passwordLogin/")
    GetBlogByTypeResponse GetBlogByType(GetBlogByTypeRequest requset)(api.get = "/unlogin/getBlogByType/",api.options = "/unlogin/getBlogByType/")
    GetBlogByKeywordResponse GetBlogByKeyword(GetBlogByKeywordRequest requset)(api.get = "/unlogin/getBlogByKeyword/",api.options = "/unlogin/getBlogByKeyword/")
    GetUserByKeywordResponse  GetUserByKeyword(GetUserByKeywordRequest reqest)(api.get = "/unlogin/getUserByKeyword/",api.options = "/unlogin/getUserByKeyword/")
    GetBlogCommentResponse GetBlogComment(GetBlogCommentRequest request)(api.get = "/unlogin/getBlogComment/",api.options = "/unlogin/getBlogComment/")
    GetUserBlogResponse GetUserBlog(GetUserBlogRequest request)(api.get = "/unlogin/getUserBlog/",api.options = "/unlogin/getUserBlog/")
    GetReCommentResponse GetReComment(GetReCommentRequest request)(api.get = "/unlogin/getReComment/",api.options = "/unlogin/getReComment/")
    GetHotBlogResponse GetHotBlog(GetHotBlogRequest request)(api.get = "/unlogin/getHotBlog/",api.options = "/unlogin/getHotBlog/")
    GetAllTypeResponse GetAllType(GetAllTypeRequest request)(api.get="/unlogin/getAllType/",api.options = "/unlogin/getAllType/")
    GetAnnouncementResponse GetAnnouncement(GetAnnouncementRequest request)(api.get="/unlogin/getAnnouncement/",api.options = "/unlogin/getAnnouncement/")
    GetBlogByIDResponse GetBlogByID(GetBlogByIDRequest request)(api.get="/unlogin/getBlogByID/",api.options = "/unlogin/getBlogByID/")
    GetUserInfoByIDResponse GetUserInfoByID(GetUserInfoByIDRequest request)(api.get="/unlogin/getUserByID/",api.options = "/unlogin/getUserByID/")
}
// 登录状态服务
service Login{
    CreateBlogResponse CreateBlog(1:CreateBlogRequest request)(api.post="/login/createblog/",api.options="/login/createblog/")
    LogoutResponse Logout(1:LogoutRequest request)(api.get="/login/logout/",api.options="/login/logout/")
    UploadResponse Upload(1:UploadRequest request)(api.post="/login/upload/",api.options="/login/upload/")
    LikeBlogResponse LikeBlog(1:LikeBlogRequest request)(api.put="/login/likeBlog/",api.options="/login/likeBlog/")
    FavorBlogResponse FavorBlog(1:LikeBlogRequest request)(api.put="/login/favorBlog/",api.options="/login/favorBlog/")
    GetFavorResponse GetFavor(1:GetFavorRequest requset)(api.get="/login/getFavor/",api.options="/login/getFavor/")
    FollowUserResponse FollowUser(1:FollowUserRequest request)(api.post="/login/followUser/",api.options="/login/followUser/")
    GetFollowResponse GetFollow(1:GetFollowRequest request)(api.get="/login/getFollow/",api.options="/login/getFollow/")
    CommentBlogResponse CommentBlog(1:CommentBlogRequest request)(api.post="/login/commentBlog/",api.options="/login/commentBlog/")
    GetUserInfoResponse GetUserInfo(1:GetUserInfoRequest request)(api.get="/login/getUser/",api.potions="/login/getUser/")
    GetDraftResponse GetDraft(1:GetDraftRequest request)(api.get="/login/getDraft/",api.potions="/login/getDraft/")
    ChangeUserInfoResponse ChangeUserInfo(1:ChangeUserInfoRequest request)(api.put="/login/changeUserInfo/",api.potions="/login/changeUserInfo/")
    ChangeDraftResponse ChengeDraft(1:ChangeDraftRequest request)(api.put="/login/changeDraft/",api.potions="/login/changeDraft/")
    ReCommentResponse ReComment(1:ReCommentRequest request)(api.post="/login/reComment/",api.potions="/login/reComment/")
    LikeCommentResponse LikeComment(1:LikeCommentRequest request)(api.post="/login/LikeComment/",api.potions="/login/reComment/")
    DelBlogResponse DelBlog(1:DelBlogRequest request)(api.put="/login/delBlog/",api.potions="/login/delBlog/")
}
//管理员服务
service Admin{
    GetPendingResponse GetPending(1:GetPendingRequest request)(api.get="/login/admin/getpending/",api.options="/login/admin/getpending/")
    ApprovalBlogResponse ApprovalBlog(1:ApprovalBlogRequest request)(api.put="/login/admin/approbalBlog/",api.options="/login/admin/approbalBlog/")
    ForbiddenUserResponse ForbiddenUser(1:ForbiddenUserRequest request)(api.put="/login/admin/forbiddenUser/",api.options="/login/admin/forbiddenUser/")
    GetForbiddenResponse  GetForbidden(1:GetForbiddenRequest request)(api.get="/login/admin/getForbidden/",api.options="/login/admin/getForbidden/")
    AddAnnouncementResponse AddAnnouncement(1:AddAnnouncementRequest request)(api.post="/login/admin/addAnnouncement/",api.options="/login/admin/addAnnouncement/")
    DelAnnouncementResponse DelAnnouncement(1:DelAnnouncementRequest request)(api.put="/login/admin/delAnnouncement/",api.options="/login/admin/delAnnouncement/")
}