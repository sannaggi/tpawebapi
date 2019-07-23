package collections

type Place struct {
	Name          string   `json:"name"`
	Price         int      `json:"price"`
	Rating        []Review `json:"rating"`
	Category      string   `json:"category"`
	HostID        string   `json:"hostid"`
	AverageRating float64  `json:"averagerating"`
	RatingCount   int      `json:"ratingcount"`
}

type Review struct {
	UserID string  `json:"userid"`
	Review string  `json:"review"`
	Rating float64 `json:"rating"`
}
