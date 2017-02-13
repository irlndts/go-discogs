package discogs

type Video struct {
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Embed       bool   `json:"embed"`
	Title       string `json:"title"`
	Uri         string `json:"uri"`
}

type ArtistSource struct {
	Anv          string `json:"anv"`
	Id           int    `json:"id"`
	Join         string `json:"join"`
	Name         string `json:"name:`
	Resource_url string `json:"resource_url"`
	Role         string `json:"role"`
	Tracks       string `json:"tracks"`
}

type Image struct {
	Height       int    `json:"height"`
	Width        int    `json:"width"`
	Resource_url string `json:"resource_url"`
	Type         string `json:"type"`
	Uri          string `json:"uri"`
	Uri150       string `json:"uri150"`
}

type Track struct {
	Duration     string         `json:"duration"`
	Position     string         `json:"position"`
	Title        string         `json:"title"`
	Type         string         `json:"type_"`
	Extraartists []ArtistSource `json:"extraartists"`
}

type LabelSource struct {
	Catno        string `json:"catno"`
	Entity_type  string `json:"entity_type"`
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Resource_url string `json:"resource_url"`
}

type Identifier struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Format struct {
	Descriptions []string `json:"descriptions"`
	Name         string   `json:"name"`
	Qty          string   `json:"qty"`
}

type Company struct {
	Catno            string `json:"catno"`
	Entity_type      string `json:"entity_type"`
	Entity_type_name string `json:"entity_type_name"`
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Resource_url     string `json:"resource_url"`
}

type Community struct {
	Contributors []Contributor `json:"contributors"`
	Data_quality string        `json:"string"`
	Have         int           `json:"have"`
	Rating       Rating        `json:"rating"`
	Status       string        `json:"status"`
	Submitter    Submitter     `json:"submitter"`
	Want         int           `json:"want"`
}

type Submitter struct {
	Resource_url string `json:"resource_url"`
	Username     string `json:"username"`
}

type Rating struct {
	Average float32 `json:"average"`
	Count   int     `json:"count"`
}

type Contributor struct {
	Resource_url string `json:"resource_url"`
	Username     string `json:"username"`
}

type Page struct {
	Per_page int  `json:"per_page"`
	Items    int  `json:"items"`
	Page     int  `json:"page"`
	Urls     URLS `json:"urls"`
	Pages    int  `json:"pages"`
}

type URLS struct {
	Last string `json:"last"`
	Next string `json:"next"`
}

type Version struct {
	Catno        string `json:"catno"`
	Country      string `json:"country"`
	Format       string `json:"format"`
	Id           int    `json:"id"`
	Label        string `json:"label"`
	Released     string `json:"released"`
	Resource_url string `json:"resource_url"`
	Status       string `json:"status"`
	Thumb        string `json:"thumb"`
	Title        string `json:"title"`
}

type Member struct {
	Active       bool   `json:"active"`
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Resource_url string `json:"resource_url"`
}

type Sublable struct {
	Resource_url string `json:"url"`
	Id           int    `json:"id"`
	Name         string `json:"name"`
}

type ReleaseSource struct {
	Artist       string `json:"artist"`
	Catno        string `json:"catno"`
	Format       string `json:"format"`
	Id           int    `json:"id"`
	Resource_url string `json:"resource_url"`
	Status       string `json:"status"`
	Thumb        string `json:"thumb"`
	Title        string `json:"title"`
	Year         int    `json:"year"`
	Main_release int    `json:"main_release"`
	Role         string `json:"role"`
	Type         string `json:"type"`
}
