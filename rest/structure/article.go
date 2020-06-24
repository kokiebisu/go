package structure

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
