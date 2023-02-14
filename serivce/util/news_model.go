package util

type News struct {
	NewsID          string        `json:"news_id"`
	SourceName      string        `json:"source_name"`
	SiteDomain      string        `json:"site_domain"`
	Author          []interface{} `json:"author"`
	URL             string        `json:"url"`
	Time            string        `json:"time"`
	SiteBoardName   string        `json:"site_board_name"`
	BoardTheme      string        `json:"board_theme"`
	IfFrontPosition int           `json:"if_front_position"`
	Type            string        `json:"type"`
	CrawlTime       string        `json:"crawl_time"`
	Lang            string        `json:"lang"`
	Direction       string        `json:"direction"`
	CommentCount    int           `json:"comment_count"`
	ForwardCount    int           `json:"forward_count"`
	LikeCount       int           `json:"like_count"`
	ReadCount       int           `json:"read_count"`
	OriginalTags    []interface{} `json:"original_tags"`
	IfRepost        int           `json:"if_repost"`
	RepostSource    string        `json:"repost_source"`
	InsertTime      string        `json:"insert_time"`
	Title           string        `json:"title"`
	Content         []Content     `json:"content"`
}

type NewNewsStu struct {
	NewsID          string        `json:"news_id"`
	SourceName      string        `json:"source_name"`
	SiteDomain      string        `json:"site_domain"`
	Author          []interface{} `json:"author"`
	URL             string        `json:"url"`
	Time            string        `json:"time"`
	SiteBoardName   string        `json:"site_board_name"`
	BoardTheme      string        `json:"board_theme"`
	IfFrontPosition int           `json:"if_front_position"`
	Type            string        `json:"type"`
	CrawlTime       string        `json:"crawl_time"`
	Lang            string        `json:"lang"`
	Direction       string        `json:"direction"`
	CommentCount    int           `json:"comment_count"`
	ForwardCount    int           `json:"forward_count"`
	LikeCount       int           `json:"like_count"`
	ReadCount       int           `json:"read_count"`
	OriginalTags    []interface{} `json:"original_tags"`
	IfRepost        int           `json:"if_repost"`
	RepostSource    string        `json:"repost_source"`
	InsertTime      string        `json:"insert_time"`
	Title           string        `json:"title"`
	Content         []Content     `json:"content"`
	OriNewsID       string        `json:"ori_news_id"`
	SpecialName     string        `json:"special_name"`
	SpecialKeyword  string        `json:"special_keyword"`
}
type Data struct {
	Name        interface{} `json:"name"`
	Md5Src      string      `json:"md5src"`
	Description string      `json:"description"`
	Src         string      `json:"src"`
}
type Content struct {
	Type string `json:"type"`
	Data Data   `json:"data"`
}
