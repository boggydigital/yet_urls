package rutube_urls

type PlayOptions struct {
	Id           int           `json:"id"`
	ThumbnailUrl string        `json:"thumbnail_url"`
	Captions     []interface{} `json:"captions"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Viewer       interface{}   `json:"viewer"`
	BlackRabbit  bool          `json:"black_rabbit"`
	Skin         interface{}   `json:"skin"`
	Appearance   struct {
		Color                   string      `json:"color"`
		AutoStart               bool        `json:"auto_start"`
		ShowSubscription        interface{} `json:"show_subscription"`
		ShowAuthor              bool        `json:"show_author"`
		ShowTitle               bool        `json:"show_title"`
		ShowHd                  bool        `json:"show_hd"`
		ShowLogotype            bool        `json:"show_logotype"`
		ShowFullTab             bool        `json:"show_full_tab"`
		ShowAvatar              bool        `json:"show_avatar"`
		ShowTopPanel            bool        `json:"show_top_panel"`
		ShowShareButton         bool        `json:"show_share_button"`
		ShowSocialButtons       bool        `json:"show_social_buttons"`
		ShowEndscreen           bool        `json:"show_endscreen"`
		ShowRelated             bool        `json:"show_related"`
		ShowCategoryRecommended bool        `json:"show_category_recommended"`
		ShowUserRecommended     bool        `json:"show_user_recommended"`
		ShowPgRating            bool        `json:"show_pg_rating"`
		AvailablePauseClick     bool        `json:"available_pause_click"`
		AvailableTitleClick     bool        `json:"available_title_click"`
		AvailableLogoClick      bool        `json:"available_logo_click"`
		AvailableAuthorClick    bool        `json:"available_author_click"`
		AvailableRelatedClick   bool        `json:"available_related_click"`
		PluginModifying         bool        `json:"plugin_modifying"`
		ForbidHtmlControls      bool        `json:"forbid_html_controls"`
		ShowAdultStub           bool        `json:"show_adult_stub"`
		AvailableLinkStub       bool        `json:"available_link_stub"`
		AvailableSpecialSource  bool        `json:"available_special_source"`
		AdvertSkinType          interface{} `json:"advert_skin_type"`
		ShowPreloader           bool        `json:"show_preloader"`
		EnableEmbedLogo         bool        `json:"enable_embed_logo"`
		CookiesMail             bool        `json:"cookies_mail"`
		ForbidHtmlFullscreen    bool        `json:"forbid_html_fullscreen"`
		AllowAutoSeek           bool        `json:"allow_auto_seek"`
		ForbidTimelinePreview   bool        `json:"forbid_timeline_preview"`
		AllowPushSubscribe      bool        `json:"allow_push_subscribe"`
		EnablePlaylist          bool        `json:"enable_playlist"`
		ShowLogoWithControls    bool        `json:"show_logo_with_controls"`
		MiniPlayer              bool        `json:"mini_player"`
		CinemaMode              bool        `json:"cinema_mode"`
	} `json:"appearance"`
	VideoBalancer struct {
		Default string `json:"default"`
		M3U8    string `json:"m3u8"`
	} `json:"video_balancer"`
	VideoBalancerVertical struct {
	} `json:"video_balancer_vertical"`
	LiveStreams struct {
	} `json:"live_streams"`
	//Advert []struct {
	//	Name                   string      `json:"name"`
	//	UrlTemplate            string      `json:"url_template"`
	//	Start                  float64     `json:"start"`
	//	Count                  int         `json:"count"`
	//	Delay                  float64     `json:"delay"`
	//	OnlyFire               bool        `json:"only_fire"`
	//	Xmltimeout             float64     `json:"xmltimeout"`
	//	Adtimeout              float64     `json:"adtimeout"`
	//	IsServerRequest        bool        `json:"is_server_request"`
	//	Vast                   interface{} `json:"vast"`
	//	UrlTemplateAntiadblock string      `json:"url_template_antiadblock"`
	//} `json:"advert"`
	//Stat []struct {
	//	Name        string  `json:"name"`
	//	UrlTemplate string  `json:"url_template"`
	//	Start       float64 `json:"start"`
	//	Count       int     `json:"count"`
	//	Delay       float64 `json:"delay"`
	//	Method      string  `json:"method"`
	//	Body        string  `json:"body"`
	//	EvType      string  `json:"ev_type,omitempty"`
	//} `json:"stat"`
	//StatApi []struct {
	//	Name      string  `json:"name"`
	//	MethodStr string  `json:"method_str"`
	//	Start     float64 `json:"start"`
	//	Count     int     `json:"count"`
	//	Delay     float64 `json:"delay"`
	//	EvType    string  `json:"ev_type,omitempty"`
	//} `json:"stat_api"`
	//LogLevel  int `json:"log_level"`
	AclAccess struct {
		Allowed bool        `json:"allowed"`
		ErrCode interface{} `json:"err_code"`
		ErrText string      `json:"err_text"`
	} `json:"acl_access"`
	PurchaseTtl interface{} `json:"purchase_ttl"`
	//Yast           interface{} `json:"yast"`
	//YastLiveOnline interface{} `json:"yast_live_online"`
	ViewsHistory interface{} `json:"views_history"`
	//Meta         struct {
	//	AdvertRuleId       int           `json:"advert_rule_id"`
	//	AdvertLimitRuleId  int           `json:"advert_limit_rule_id"`
	//	QualityRuleId      int           `json:"quality_rule_id"`
	//	ThemeId            int           `json:"theme_id"`
	//	StatRuleIds        []int         `json:"stat_rule_ids"`
	//	AppearanceId       int           `json:"appearance_id"`
	//	AlternativeVideoId interface{}   `json:"alternative_video_id"`
	//	SpecialAccess      []interface{} `json:"special_access"`
	//} `json:"meta"`
	Referer string `json:"referer"`
	Fts     int    `json:"fts"`
	Limits  []struct {
		Name             string      `json:"name"`
		Limit            int         `json:"limit"`
		FirstPeriodFreq  interface{} `json:"first_period_freq"`
		EndFirstPeriod   interface{} `json:"end_first_period"`
		SecondPeriodFreq interface{} `json:"second_period_freq"`
	} `json:"limits"`
	IsSpatial              bool          `json:"is_spatial"`
	Mobile                 bool          `json:"mobile"`
	AllowP2P               bool          `json:"allow_p2p"`
	RemoveUnseekableBlocks bool          `json:"remove_unseekable_blocks"`
	DrmToken               interface{}   `json:"drm_token"`
	EffectiveVideo         string        `json:"effective_video"`
	Cuepoints              []interface{} `json:"cuepoints"`
	GpmdId                 int           `json:"gpmd_id"`
	GpmdPuid1              int           `json:"gpmd_puid1"`
	Author                 struct {
		Id        int         `json:"id"`
		Url       string      `json:"url"`
		AvatarUrl string      `json:"avatar_url"`
		Logo      interface{} `json:"logo"`
		Name      string      `json:"name"`
	} `json:"author"`
	HasVideo   bool   `json:"has_video"`
	VideoId    string `json:"video_id"`
	IsAdult    bool   `json:"is_adult"`
	IsHidden   bool   `json:"is_hidden"`
	IsLicensed bool   `json:"is_licensed"`
	VideoUrl   string `json:"video_url"`
	PgRating   struct {
		Age  int    `json:"age"`
		Logo string `json:"logo"`
	} `json:"pg_rating"`
	Duration int           `json:"duration"`
	Tags     []interface{} `json:"tags"`
	Category struct {
		Id          int    `json:"id"`
		CategoryUrl string `json:"category_url"`
		Name        string `json:"name"`
		ShortName   string `json:"short_name"`
		ForKids     bool   `json:"for_kids"`
	} `json:"category"`
	Html     string      `json:"html"`
	TvShowId interface{} `json:"tv_show_id"`
	Tv       struct {
	} `json:"tv"`
	Season        interface{} `json:"season"`
	Episode       interface{} `json:"episode"`
	Player        string      `json:"player"`
	IframeUrl     interface{} `json:"iframe_url"`
	CountLoadTime bool        `json:"count_load_time"`
	ViewId        string      `json:"view_id"`
	IsHarmHealth  bool        `json:"is_harm_health"`
	StreamType    interface{} `json:"stream_type"`
	OriginType    string      `json:"origin_type"`
}
