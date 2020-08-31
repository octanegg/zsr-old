package octane

// Event .
type Event struct {
	Name      string  `json:"name" bson:"name"`
	StartDate string  `json:"startDate" bson:"startDate"`
	EndDate   string  `json:"endDate" bson:"endDate"`
	Region    string  `json:"region" bson:"region"`
	Mode      int     `json:"mode" bson:"mode"`
	Prize     Prize   `json:"prize" bson:"prize"`
	Tier      string  `json:"tier" bson:"tier"`
	Stages    []Stage `json:"stages" bson:"stages"`
}

// Stage .
type Stage struct {
	Name       string     `json:"name" bson:"name"`
	Format     string     `json:"format" bson:"format"`
	Region     string     `json:"region" bson:"region"`
	StartDate  string     `json:"startDate" bson:"startDate"`
	EndDate    string     `json:"endDate" bson:"endDate"`
	Prize      Prize      `json:"prize" bson:"prize"`
	Liquipedia string     `json:"liquipedia" bson:"liquipedia"`
	Qualifier  bool       `json:"qualifier" bson:"qualifier"`
	Substages  []Substage `json:"substages" bson:"substages"`
}

// Substage .
type Substage struct {
	Name   string `json:"name" bson:"name"`
	Format string `json:"format" bson:"format"`
}

// Prize .
type Prize struct {
	Amount   float32 `json:"amount" bson:"amount"`
	Currency string  `json:"currency" bson:"currency"`
}
