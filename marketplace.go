package discogs

import "strconv"

const (
	priceSuggestionsURI = "/price_suggestions/"
	releaseStatsURI     = "/stats/"
)

type marketPlaceService struct {
	url      string
	currency string
}

type MarketPlaceService interface {
	// The best price suggestions according to grading
	PriceSuggestions(releaseID int) (*PriceListing, error)
	// Short summary of marketplace listings
	ReleaseStats(releaseID int) (*Stats, error)
}

// Listing is a marketplace listing with the user's currency and a price value
type Listing struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}

// PriceListings are Listings per grading quality
type PriceListing struct {
	VeryGood     Listing `json:"Very Good (VG),omitempty"`
	GoodPlus     Listing `json:"Good Plus (G+),omitempty"`
	NearMint     Listing `json:"Near Mint (NM or M-)"`
	Good         Listing `json:"Good (G),omitempty"`
	VeryGoodPlus Listing `json:"Very Good Plus (VG+),omitempty"`
	Mint         Listing `json:"Mint (M),omitempty"`
	Fair         Listing `json:"Fair (F),omitempty"`
	Poor         Listing `json:"Poor (P),omitempty"`
}

// Stats returns the marketplace stats summary for a release containing
type Stats struct {
	LowestPrice Listing `json:"lowest_price"`
	ForSale     int     `json:"num_for_sale"`
	Blocked     bool    `json:"blocked_from_sale"`
}

func (s *marketPlaceService) ReleaseStats(releaseID int) (*Stats, error) {
	var stats *Stats
	err := request(s.url+releaseStatsURI+strconv.Itoa(releaseID), nil, &stats)
	return stats, err
}

func (s *marketPlaceService) PriceSuggestions(releaseID int) (*PriceListing, error) {
	var listings *PriceListing
	err := request(s.url+priceSuggestionsURI+strconv.Itoa(releaseID), nil, &listings)
	return listings, err
}

func newMarketPlaceService(url string, currency string) MarketPlaceService {
	return &marketPlaceService{
		url:      url,
		currency: currency,
	}
}
