package fetch

type BookTex struct {
	ISBN          string
	Authors       string
	Publisher     string
	PublishedDate string
	URL           string
}

// google books search object response
type GBSearch struct {
	Items []items
	Kind string
	TotalItems int
}

type items struct {
	Kind string
	Id string
	Etag string
	SelfLink string
	AccessInfo accessInfo
	VolumeInfo volumeInfo
	SalesInfo salesInfo
	SearchInfo searchInfo
}

type volumeInfo struct {
	Title string
	Authors []string
	Publisher string
	PublishDate string
	Description string
	IndustryIdentifiers []identifers
	ReadingModes readingModes
	PageCount int64
	PrintType string
	Categories []string
	AverageRating float64
	RatingsCount int64
	MaturityRating string
	AllowAnonLogging bool
	ContentVersion string
	PanelizationSummary panelizationSummary
	ImageLinks imageLinks
	Language language
}

type identifers struct {
	Type string
	Identifier string
}
type readingModes struct {
	text bool
	image bool
}

type panelizationSummary struct {
	containsEpubBubbles bool
	containsImageBubbles bool
}

type imageLinks struct {
	smallThumbnail string
	thumbnail string
}

type language struct {
	PreviewLink string
	InfoLink string
	CanonicalVolumeLink string
}

type salesInfo struct {
	Country string
	Saleability string
	IsEbook bool
}

type accessInfo struct {
	Country string
	Viewability string
	Embeddable bool
	PublicDomain bool
	TextToSpeechPermission string
	Epub document
	Pdf document
	WebReaderLink string
	AccessViewStatus string
	QuoteSharingAllowed string
}

type document struct {
	IsAvailable bool
}

type searchInfo struct {
	TextSnippet string
}
