package animeflv

import "github.com/kuronosu/schema_scraper/pkg/config"

func S(_s string) *string { return &_s }
func I(_i int) *int       { return &_i }

var AnimeFLVBrowseSchema = &config.PageSchema{
	ID:         "AnimeFlvPage",
	Version:    "1.0.0",
	Cloudflare: true,
	List: config.ListSchema{
		ContainerSelector: "ul.pagination",
		ItemSelector:      "li:nth-last-child(2)>a",
	},
}

var AnimeFLVSchema = &config.PageSchema{
	ID:         "AnimeFlv",
	Version:    "1.0.0",
	Cloudflare: true,
	Detail: config.DetailSchema{
		Fields: []config.Field{
			{
				Name:     "title",
				Selector: "h1.Title",
			},
			{
				Name:     "synopsis",
				Selector: "div.Description",
			},
			{
				Name:     "status",
				Selector: "span.fa-tv",
			},
			{
				Name:     "type",
				Selector: "span.Type",
			},
			{
				Name:     "score",
				Selector: "span#votes_prmd",
				Type:     S("float"),
			},
			{
				Name:     "votes",
				Selector: "span#votes_nmbr",
				Type:     S("int"),
			},
			{
				Name:     "followers",
				Selector: "section.WdgtCn.Sm>div>div>span",
				Type:     S("int"),
			},
			{
				Name:     "cover",
				Selector: "div.Image>figure>img",
				Attr:     S("src"),
			},
			{
				Name:     "banner",
				Selector: "div.Bg",
				Attr:     S("style"),
				Regex: []config.Regex{
					{
						Pattern: "\\(([^)]+)\\)",
						Group:   1,
					},
				},
			},
			{
				Name:     "genres",
				Selector: "nav.Nvgnrs",
				Type:     S("array"),
				Item: &config.ListItem{
					Plain:    false,
					Selector: S("a"),
					Fields: []config.Field{
						{
							Name: "name",
						},
						{
							Name: "url",
							Attr: S("href"),
						},
					},
				},
			},
			{
				Name:     "related",
				Selector: "ul.ListAnmRel",
				Type:     S("array"),
				Item: &config.ListItem{
					Plain:    false,
					Selector: S("li"),
					Fields: []config.Field{
						{
							Name:     "name",
							Selector: "a",
						},
						{
							Name:     "url",
							Selector: "a",
							Attr:     S("href"),
						},
						{
							Name: "relation",
							Regex: []config.Regex{
								{
									Pattern: "\\(([^)]+)\\)",
									Group:   1,
								},
							},
						},
					},
				},
			},
			{
				Name:     "raw_episodes",
				Selector: "script",
				Contains: &config.FieldContains{
					Raw:    true,
					String: "var episodes",
				},
				Regex: []config.Regex{
					{
						Pattern: "var episodes = \\[(.*)\\];",
						Group:   1,
					},
				},
			},
			{
				Name:     "raw_data",
				Selector: "script",
				Contains: &config.FieldContains{
					Raw:    true,
					String: "var anime_info",
				},
				Regex: []config.Regex{
					{
						Pattern: "var anime_info = \\[(.*)\\];",
						Group:   1,
					},
				},
			},
		},
	},
	List: config.ListSchema{
		ContainerSelector: "ul.ListAnimes",
		ItemSelector:      "li>article>a",
		Prefix:            "https://www3.animeflv.net",
		IncludePrefix:     true,
		// Pagination: config.Pagination{
		// 	Next: config.PaginationLink{
		// 		Limit:    1,
		// 		Selector: "ul.pagination>li:last-child>a",
		// 		Prefix:   "https://www3.animeflv.net",
		// 	},
		// },
	},
}
