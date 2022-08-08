package brightdata

import (
	cache2 "github.com/patrickmn/go-cache"
	"net/url"
	"time"
)

type brightDataResponse struct {
	General struct {
		BasicView    bool      `json:"basic_view"`
		Country      string    `json:"country"`
		CountryCode  string    `json:"country_code"`
		Gl           string    `json:"gl"`
		Language     string    `json:"language"`
		Location     string    `json:"location"`
		Mobile       bool      `json:"mobile"`
		PageTitle    string    `json:"page_title"`
		Query        string    `json:"query"`
		ResultsCnt   int64     `json:"results_cnt"`
		SearchEngine string    `json:"search_engine"`
		SearchTime   float64   `json:"search_time"`
		SearchType   string    `json:"search_type"`
		Timestamp    time.Time `json:"timestamp"`
	} `json:"general"`

	Organic []Organic `json:"organic"`

	Videos []struct {
		Link string `json:"link"`
	} `json:"videos"`
	Images []struct {
		Link string `json:"link"`
	} `json:"images"`
	Twitter struct {
		Items []struct {
			Link string `json:"link"`
		} `json:"items"`
		Title string `json:"title"`
	} `json:"twitter"`
	FeaturedSnippets []struct {
		Link string `json:"link"`
	} `json:"featured_snippets"`
	PopularProducts struct {
		Items []struct {
			Link string `json:"link"`
		} `json:"items"`
		Title string `json:"title"`
	} `json:"popular_products"`
	Recipes struct {
		Items []struct {
			Link string `json:"link"`
		} `json:"items"`
		Title string `json:"title"`
	} `json:"recipes"`
	TopStories struct {
		Items []struct {
			Link string `json:"link"`
		} `json:"items"`
		Title string `json:"title"`
	} `json:"top_stories"`
	Jobs struct {
		Items []struct {
			Site string `json:"site"`
		} `json:"items"`
	} `json:"jobs"`
	Apps struct {
		Items []struct {
			Link string `json:"link"`
		} `json:"items"`
	} `json:"apps"`
	Shopping []struct {
		Link string `json:"link"`
	} `json:"shopping"`
	ProductOffers []struct {
		Link string `json:"link"`
	} `json:"product_offers"`
	TopKnowledgeCarousel struct {
		List []struct {
			Link string `json:"link"`
		} `json:"list"`
	} `json:"top_knowledge_carousel"`
	Knowledge struct {
		Description     string `json:"description"`
		DescriptionLink string `json:"description_link"`
	} `json:"knowledge"`
	HotelsSelection struct {
		Link string `json:"link"`
	} `json:"hotels_selection"`
	SnackPackMap struct {
		Link string `json:"link"`
	} `json:"snack_pack_map"`
	SnackPack []struct {
		Link string `json:"link"`
	} `json:"snack_pack"`
	TopAds []struct {
		Link string `json:"link"`
	} `json:"top_ads"`
	BottomAds []struct {
		Link string `json:"link"`
	} `json:"bottom_ads"`
	TopPla []struct {
		Link string `json:"link"`
	} `json:"top_pla"`
	BottomPla []struct {
		Link string `json:"link"`
	} `json:"bottom_pla"`
	JackpotPla []struct {
		Link string `json:"link"`
	} `json:"jackpot_pla"`
	PeopleAlsoAsk []struct {
		AnswerLink string `json:"answer_link"`
		Answers    []struct {
			Title string `json:"title,omitempty"`
			Link  string `json:"link,omitempty"`
		} `json:"answers"`
		Question string `json:"question"`
	} `json:"people_also_ask"`
	News []struct {
		Link string `json:"link"`
	} `json:"news"`
	LocalServicesAds struct {
		Items []struct {
			Link string `json:"link"`
		} `json:"items"`
	} `json:"local_services_ads"`
}

type BrightDataRedisOptions struct {
	Addr        string
	DB          int
	Password    string
	CacheExpiry time.Duration
}

type BrightDataClientOptions struct {
	LuminatiUrl   string
	UseRedisCache bool
	RedisOpts     BrightDataRedisOptions
}

type BrightDataClient struct {
	options    *BrightDataClientOptions
	proxy      *url.URL
	cache      *BDRedisClient
	localCache *cache2.Cache
}

type SearchOptions struct {
	Keyword     string
	Country     string
	Desktop     bool
	CheckDomain string
}

type (
	// Serps defines the collection to be returned from
	// the client.
	Serps struct {
		Organic    []Organic `json:"serps"`
		Features   []string  `json:"features"`
		CheckedUrl Domain    `json:"checked_url"`
	}
	// Domain are URL specific results returned by
	// Serps.CheckURL
	Domain struct {
		Query   Query   `json:"query"`
		Results []Query `json:"results"`
	}
	// Query defines the first top level
	Query struct {
		Rank        int      `json:"position"`
		Link        string   `json:"url"`
		Description string   `json:"text"`
		Features    []string `json:"features"`
	}

	Organic struct {
		Description string `json:"description"`
		DisplayLink string `json:"display_link"`
		Extensions  []struct {
			Type string `json:"type"`
			Link string `json:"link,omitempty"`
		} `json:"extensions"`
		Link       string `json:"link"`
		Rank       int    `json:"rank"`
		Subresults []struct {
			Link string `json:"link"`
		} `json:"subresults"`
		Title  string `json:"title"`
		Videos []struct {
			Link string `json:"link"`
		} `json:"videos"`
	}
)
