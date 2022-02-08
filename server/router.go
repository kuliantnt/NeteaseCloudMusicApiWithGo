package server

import (
	"os"

	"github.com/lianlian/singo/api"
	"github.com/lianlian/singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
//  @return *gin.Engine
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/")
	{
		//ping
		v1.POST("ping", api.Ping)

		v1.GET("login/status", api.LoginStatus)
		v1.GET("login/cellphone", api.LoginCellphone)
		v1.GET("login", api.LoginEmail)
		v1.GET("login/refresh", api.LoginRefresh)
		v1.GET("login/qr/key", api.LoginQrkey)
		v1.GET("captcha/sent", api.CaptchaSent)
		v1.GET("captcha/verify", api.CaptchaVerify)
		v1.GET("register/cellphone", api.RegisterCellphone)
		v1.GET("cellphone/existence/check", api.CellphoneExistenceCheck)
		v1.GET("activate/init/profile", api.ActivateInitProfile)
		v1.GET("rebind", api.Rebind)
		v1.GET("logout", api.Logout)
		v1.GET("user/detail", api.UserDetail)
		v1.GET("user/subcount", api.UserSubcount)
		v1.GET("user/update", api.UserUpdate)
		v1.GET("countries/code/list", api.CountriesCodeList)
		v1.GET("user/playlist", api.UserPlaylist)
		v1.GET("user/playlist/update", api.PlaylistUpdate)
		v1.GET("playlist/desc/update", api.PlaylistDescUpdate)
		v1.GET("playlist/name/update", api.PlaylistNameUpdate)
		v1.GET("playlist/tags/update", api.PlaylistTagsUpdate)
		v1.GET("playlist/order/update", api.PlaylistOrderUpdate)
		v1.GET("song/order/update", api.SongOrderUpdate)
		v1.GET("user/dj", api.UserDj)
		v1.GET("user/follows", api.UserFollows)
		v1.GET("user/followeds", api.UserFolloweds)
		v1.GET("user/event", api.UserEvent)
		v1.GET("event/forward", api.EventForward)
		v1.GET("event/del", api.EventDel)
		v1.GET("share/resource", api.ShareResource)
		v1.GET("comment/event", api.CommentEvent)
		v1.GET("follow", api.Follow)
		v1.GET("user/record", api.UserRecord)
		v1.GET("hot/topic", api.HotTopic)
		v1.GET("comment/hotwall/list", api.CommentHotwallList)
		v1.GET("playmode/intelligence/list", api.PlaymodeIntelligenceList)
		v1.GET("event", api.Event)
		v1.GET("artist/list", api.ArtistList)
		v1.GET("artist/sub", api.ArtistSub)
		v1.GET("artist/top/song", api.ArtistTopSong)
		v1.GET("artist/sublist", api.ArtistSublist)
		v1.GET("video/sub", api.VideoSub)
		v1.GET("mv/sub", api.MvSub)
		v1.GET("mv/sublist", api.MvSublist)
		v1.GET("playlist/catlist", api.PlaylistCatlist)
		v1.GET("playlist/hot", api.PlaylistHot)
		v1.GET("top/playlist", api.TopPlaylist)
		v1.GET("top/playlist/highquality", api.TopPlaylistHighquality)
		v1.GET("related/playlist", api.RelatedPlaylist)
		v1.GET("playlist/detail", api.PlaylistDetail)
		v1.GET("song/url", api.SongUrl)
		v1.GET("check/music", api.CheckMusic)
		v1.GET("search", api.Search)
		v1.GET("search/default", api.SearchDefault)
		v1.GET("search/hot", api.SearchHot)
		v1.GET("search/hot/detail", api.SearchHotDetail)
		v1.GET("search/suggest", api.SearchSuggest)
		v1.GET("search/multimatch", api.SearchMultimatch)
		v1.GET("playlist/create", api.PlaylistCreate)
		v1.GET("playlist/delete", api.PlaylistDelete)
		v1.GET("playlist/subscribe", api.PlaylistSubscribe)
		v1.GET("playlist/subscribers", api.PlaylistSubscribers)
		v1.GET("playlist/tracks", api.PlaylistTracks)
		v1.GET("lyric", api.Lyric)
		v1.GET("top/song", api.TopSong)
		v1.GET("homepage/block/page", api.HomepageBlockPage)
		v1.GET("homepage/dragon/ball", api.HomepageDragonBall)
		v1.GET("comment/music", api.CommentMusic)
		v1.GET("comment/album", api.CommentAlbum)
		v1.GET("comment/playlist", api.CommentPlaylist)
		v1.GET("comment/mv", api.CommentMv)
		v1.GET("comment/dj", api.CommentDj)
		v1.GET("comment/video", api.CommentVideo)
		v1.GET("comment/hot", api.CommentHot)
		v1.GET("comment/like", api.CommentLike)
		v1.GET("comment", api.Comment)
		v1.GET("banner", api.Banner)
		v1.GET("resource/like", api.ResourceLike)
		v1.GET("song/detail", api.SongDetail)
		v1.GET("album", api.Album)
		v1.GET("album/detail/dynamic", api.AlbumDetailDynamic)
		v1.GET("album/sub", api.AlbumSub)
		v1.GET("album/sublist", api.AlbumSublist)
		v1.GET("artists", api.Artists)
		v1.GET("artist/mv", api.ArtistMv)
		v1.GET("artist/album", api.ArtistAlbum)
		v1.GET("artist/desc", api.ArtistDesc)
		v1.GET("simi/artist", api.SimiArtist)
		v1.GET("simi/playlist", api.SimiPlaylist)
		v1.GET("simi/mv", api.SimiMv)
		v1.GET("simi/song", api.SimiSong)
		v1.GET("simi/user", api.SimiUser)
		v1.GET("recommend/resource", api.RecommendResource)
		v1.GET("recommend/songs", api.RecommendSongs)
		v1.GET("history/recommend/songs", api.HistoryRecommendSongs)
		v1.GET("history/recommend/songs/detail", api.HistoryRecommendDongsDetail)
		v1.GET("personal_fm", api.PersonalFm)
		v1.GET("daily_signin", api.DailySignin)
		v1.GET("like", api.Like)
		v1.GET("likelist", api.LikeList)
		v1.GET("fm_trash", api.FmTrash)
		v1.GET("top/album", api.TopAlbum)
		v1.GET("album/new", api.AlbumNew)
		v1.GET("scrobble", api.Scrobble)
		v1.GET("top/artists", api.TopArtists)
		v1.GET("mv/all", api.MvAll)
		v1.GET("mv/first", api.MvFirst)
		v1.GET("mv/exclusive/rcmd", api.MvExclusiveRcmd)
		v1.GET("personalized/mv", api.PersonalizedMv)
		v1.GET("personalized", api.Personalized)
		v1.GET("personalized/newsong", api.PersonalizedNewsong)
		v1.GET("personalized/djprogram", api.PersonalizedDjprogram)
		v1.GET("program/recommend", api.ProgramRecommend)
		v1.GET("personalized/privatecontent", api.PersonalizedPrivatecontent)
		v1.GET("personalized/privatecontent/list", api.PersonalizedPrivatecontentList)
		v1.GET("top/mv", api.TopMv)
		v1.GET("mv/detail", api.MvDetail)
		v1.GET("mv/detail/info", api.MvDetailInfo)
		v1.GET("mv/url", api.MvUrl)
		v1.GET("video/group/list", api.VideoGroupList)
		v1.GET("video/category/list", api.VideoCategoryList)
		v1.GET("video/group", api.VideoGroup)
		v1.GET("video/timeline/all", api.VideoTimelineAll)
		v1.GET("video/timeline/recommend", api.VideoTimelineRecommend)
		v1.GET("related/allvideo", api.RelatedAllVideo)
		v1.GET("video/detail", api.VideoDetail)
		v1.GET("video/detail/info", api.VideoDetailInfo)
		v1.GET("video/url", api.VideoUrl)
		v1.GET("toplist", api.Toplist)
		v1.GET("toplist/detail", api.ToplistDetail)
		v1.GET("toplist/artist", api.ToplistArtist)
		v1.GET("user/cloud", api.UserCloud)
		v1.GET("user/cloud/detail", api.UserCloudDetail)
		v1.GET("user/cloud/del", api.UserCloudDel)
		v1.GET("dj/banner", api.DjBanner)
		v1.GET("user/audio", api.UserAudio)
		v1.GET("dj/hot", api.DjHot)
		v1.GET("dj/program/toplist", api.DjProgramToplist)
		v1.GET("dj/toplist/pay", api.DjToplistPay)
		v1.GET("dj/program/toplist/hours", api.DjProgramToplistHours)
		v1.GET("dj/toplist/hours", api.DjToplistHours)
		v1.GET("dj/toplist/newcomer", api.DjToplistNewcomer)
		v1.GET("dj/toplist/popular", api.DjToplistPopular)
		v1.GET("dj/toplist", api.DjToplist)
		v1.GET("dj/radio/hot", api.DjRadioHot)
		v1.GET("dj/recommend", api.DjRecommend)
		v1.GET("dj/catelist", api.DjCatelist)
		v1.GET("dj/recommend/type", api.DjRecommendType)
		v1.GET("dj/sub", api.DjSub)
		v1.GET("dj/sublist", api.DjSublist)
		v1.GET("dj/paygift", api.DjPaygift)
		v1.GET("dj/category/excludehot", api.DjCategoryExcludehot)
		v1.GET("dj/category/recommend", api.DjCategoryRecommend)
		v1.GET("dj/today/perfered", api.DjTodayPerfered)
		v1.GET("dj/detail", api.DjDetail)
		v1.GET("dj/program", api.DjProgram)
		v1.GET("dj/program/detail", api.DjProgramDetail)
		v1.GET("msg/private", api.MsgPrivate)
		v1.GET("send/text", api.SendText)
		v1.GET("msg/private/history", api.MsgPrivateHistory)
		v1.GET("send/playlist", api.SendPlaylist)
		v1.GET("msg/comments", api.MsgComments)
		v1.GET("msg/forwards", api.MsgForwards)
		v1.GET("msg/notices", api.MsgNotices)
		v1.GET("setting", api.Setting)
		v1.GET("album/list", api.AlbumList)
		v1.GET("album/songsaleboard", api.AlbumSongsaleboard)
		v1.GET("album/list/style", api.AlbumListStyle)
		v1.GET("album/detail", api.AlbumDetail)
		v1.GET("digitalAlbum/purchased", api.DigitalAlbumPurchased)
		v1.GET("digitalAlbum/ordering", api.DigitalAlbumOrdering)
		v1.GET("batch", api.Batch)
		v1.GET("comment/floor", api.CommentFloor)
		v1.GET("artist/songs", api.ArtistSongs)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			//auth.GET("user/me", api.UserMe)
			//auth.DELETE("user/logout", api.UserLogout)
		}
	}
	return r
}
