package input

type QR struct {
	Content string `json:"content" binding:"required"`
}

type Addr struct {
	Line1      string `json:"line1" binding:"required"`
	Line2      string `json:"line2,omitempty"`
	Name       string `json:"name" binding:"required"`
	PostNumber string `json:"post_number,omitempty"`
	Vertical   bool   `json:"rotation,omitempty"`
}

type Ord struct {
	ID    int     `json:"id" binding:"required"`
	Items []*Item `json:"items"`
}

type Item struct {
	Name string `json:"name"`
	Cnt  int    `json:"cnt"`
}
