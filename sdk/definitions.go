package sdk

type ListParams struct {
	Limit     int
	Page      int
	DescOrder bool
}

type ListResponse []DogImage

type DogImage struct {
	ID         string     `json:"id"`
	URL        string     `json:"url"`
	Categories []Category `json:"categories"`
	Breeds     []Breed    `json:"breeds"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Breed struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Temperament      string `json:"temperament"`
	LifeSpan         string `json:"life_span"`
	AlternativeNames string `json:"alt_names"`
	WikipediaURL     string `json:"wikipedia_url"`
	Origin           string `json:"origin"`
	Weight           *HW    `json:"weight"`
	Height           *HW    `json:"height"`
	CountryCode      string `json:"country_code"`
}

type HW struct {
	Imperial string `json:"imperial"`
	Metric   string `json:"metric"`
}

type VoteBody struct {
	ImageId string `json:"image_id"`
	Value   int    `json:"value"`
}
