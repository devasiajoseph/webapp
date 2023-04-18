package news

type Object struct {
	NewsID     int    `json:"news_id" db:"news_id"`
	DomainID   int    `json:"domain_id" db:"domain_id"`
	Title      string `json:"title" db:"title"`
	MinContent string `json:"min_content" db:"min_content"`
	Content    string `json:"content" db:"content"`
	CoverPhoto string `json:"cover_photo" db:"cover_photo"`
	Thumbnail  string `json:"thumbnail" db:"thumbnail"`
}

var sqlInsert = ""
var sqlDelete = ""
var sqlUpdate = ""

func (obj *Object) Create() error {
	return nil
}

func (obj *Object) Update() error {
	return nil
}

func (obj *Object) Save() error {
	return nil
}

func (obj *Object) Delete() error {
	return nil
}
