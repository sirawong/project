package entities

type PageOption struct {
	Page    int      `form:"page"`
	PerPage int      `form:"per_page"`
	Filters []string `form:"filters"`
	Search  []string `form:"search"`
	Sorts   []string `form:"sorts"`
}
