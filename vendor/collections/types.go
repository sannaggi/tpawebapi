package collections

type ReviewDetail struct {
	UserID string  `json:"userid"`
	Review string  `json:"review"`
	Rating float64 `json:"rating"`
}

type BasicReview struct {
	UserID string `json:"userid"`
	Review string `json:"review"`
}

type Amenity struct {
	Type   string `json:"type"`
	Detail string `json:"detail"`
}

type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
