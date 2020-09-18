package objects

// Form ...
type Form struct {
	Cik  string `json:"cik"`
	Name string `json:"name"`
}

// Cusip ...
type Cusip struct {
	Ticker  string `json:"ticker"`
	Cusip   string `json:"cusip"`
	Company string `json:"company"`
}

// Thirteen ...
type Thirteen struct {
	Date         string `json:"date"`
	FillingDate  string `json:"fillingDate"`
	AcceptedDate string `json:"acceptedDate"`
	Cik          string `json:"cik"`
	Cusip        string `json:"cusip"`
	Tickercusip  string `json:"tickercusip"`
	NameOfIssuer string `json:"nameOfIssuer"`
	Shares       int    `json:"shares"`
	TitleOfClass string `json:"titleOfClass"`
	Value        int    `json:"value"`
	Link         string `json:"link"`
	FinalLink    string `json:"finalLink"`
}
