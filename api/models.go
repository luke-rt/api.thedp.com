package api

type Article struct {
	Id             string     `bson:"_id"`
	UuId           string     `bson:"uuid"`
	Slug           string     `bson:"slug"`
	SeoTitle       string     `bson:"seotitle"`
	SeoDescription string     `bson:"seodescription"`
	SeoImage       string     `bson:"seoimage"`
	Headline       string     `bson:"headline"`
	Subhead        string     `bson:"subhead"`
	Abstract       string     `bson:"abstract"`
	Content        string     `bson:"content"`
	Infobox        string     `bson:"infobox"`
	Template       string     `bson:"template"`
	ShortToken     string     `bson:"shorttoken"`
	Status         string     `bson:"status"`
	Weight         string     `bson:"weight"`
	MediaId        string     `bson:"mediaid"`
	CreatedAt      string     `bson:"createdat"`
	ModifiedAt     string     `bson:"modifiedat"`
	PublishedAt    string     `bson:"publishedat"`
	Metadata       []Metadata `bson:"metadata"`
	Hits           string     `bson:"hits"`
	NormalizedTags string     `bson:"normalizedtags"`
	CeoId          string     `bson:"ceoid"`
	SstsId         string     `bson:"sstsid"`
	SstsPath       string     `bson:"sstsPath"`
	Tags           []Tag      `bson:"tags"`
	Authors        []Author   `bson:"authors"`
	DominantMedia  Media      `bson:"dominantmedia"`
	CTime          int        `bson:"ctime"`
	MTime          int        `bson:"mtime"`
	PTime          int        `bson:"ptime"`
}

type Tag struct {
	Id       string     `bson:"id"`
	UuId     string     `bson:"uuid"`
	Name     string     `bson:"name"`
	Slug     string     `bson:"slug"`
	Metadata []Metadata `bson:"metadata"`
	CeoId    string     `bson:"ceoid"`
}

type Author struct {
	Id       string     `bson:"id"`
	UuId     string     `bson:"uuid"`
	Name     string     `bson:"name"`
	Email    string     `bson:"email"`
	Slug     string     `bson:"slug"`
	Bio      string     `bson:"bio"`
	Tagline  string     `bson:"tagline"`
	Metadata []Metadata `bson:"metadata"`
	CeoId    string     `bson:"ceoid"`
	Status   string     `bson:"status"`
}

type Media struct {
	Id               string     `bson:"id"`
	Title            string     `bson:"title"`
	Content          string     `bson:"content"`
	Authors          []Author   `bson:"authors"`
	BaseName         string     `bson:"base_name"`
	Extension        string     `bson:"extension"`
	SeoTitle         string     `bson:"seo_title"`
	SeoDescription   string     `bson:"seo_description"`
	SeoImage         string     `bson:"seo_image"`
	Source           string     `bson:"source"`
	Metadata         []Metadata `bson:"metadata"`
	NormalizedTags   string     `bson:"normalized_tags"`
	UuId             string     `bson:"uuid"`
	ClickThrough     string     `bson:"click_through"`
	SVGPreview       string     `bson:"svg_preview"`
	SstsId           string     `bson:"ssts_id"`
	SstsPath         string     `bson:"ssts_path"`
	Type             string     `bson:"type"`
	CeoId            string     `bson:"ceo_id"`
	Weight           string     `bson:"weight"`
	Hits             string     `bson:"hits"`
	AttachmentUuId   string     `bson:"attachment_uuid"`
	Status           string     `bson:"status"`
	CreatedAt        string     `bson:"created_at"`
	ModifiedAt       string     `bson:"modified_at"`
	PublishedAt      string     `bson:"published_at"`
	PreviewExtension string     `bson:"preview_extension"`
	Height           string     `bson:"height"`
	Width            string     `bson:"width"`
	Transcoded       string     `bson:"transcoded"`
}

type Metadata struct {
	Label string `bson:"label"`
	// TODO: mongodb stores it as either a string or an empty array
	// make more type safe or make mongodb always store as a string
	Value interface{} `bson:"value"`
}
