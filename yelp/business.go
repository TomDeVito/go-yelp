package yelp

// A SearchResult is returned from the Search API. It includes
// the region, the total number of results, and a list of matching businesses.
// The business objects returned by this query are shallow - they will not include
// deep results such as reviews.
type SearchResult struct {
	Region     Region     // Suggested bounds in a map to display results in
	Total      int        // Total number of business results
	Businesses []Business // The list of business entries (see Business)
}

// Region provides the location of a business obtained from search.
type Region struct {
	Span   Span   // Span of suggested map bounds
	Center Center // Center position of map bounds
}

// Span provides the variance of the location from the region in the search result.
type Span struct {
	LatitudeDelta  float32 `json:"latitude_delta" gorethink:"latitude_delta"`   // Latitude width of map bounds
	LongitudeDelta float32 `json:"longitude_delta" gorethink:"longitude_delta"` // Longitude height of map bounds
}

// Center provides the coordinate where the business is most likely to be located.
type Center struct {
	Latitude  float32 // Latitude position of map bounds center
	Longitude float32 // Longitude position of map bounds center
}

// Deal defines a set of special offerings from the business.
type Deal struct {
	ID                     string       // Deal identifier
	Title                  string       // Deal title
	URL                    string       // Deal URL
	ImageURL               string       `json:"image_URL" gorethink:"image_URL"`                             // Deal image URL
	CurrencyCode           string       `json:"currency_code" gorethink:"currency_code"`                     // ISO_4217 Currency Code
	TimeStart              float32      `json:"time_start" gorethink:"time_start"`                           // Deal start time (Unix timestamp)
	TimeEnd                float32      `json:"time_end" gorethink:"time_end"`                               // Deal end time (optional: this field is present only if the Deal ends)
	IsPopular              bool         `json:"is_popular" gorethink:"is_popular"`                           // Whether the Deal is popular (optional: this field is present only if true)
	WhatYouGet             string       `json:"what_you_get" gorethink:"what_you_get"`                       // Additional details for the Deal, separated by newlines
	ImportantRestrictions  string       `json:"Important_restrictions" gorethink:"Important_restrictions"`   // Important restrictions for the Deal, separated by newlines
	AdditionalRestrictions string       `json:"Additional_restrictions" gorethink:"Additional_restrictions"` // Deal additional restrictions
	Options                []DealOption //Deal options

}

// DealOption provides options are optionally included on a deal.
type DealOption struct {
	Title                  string  // Deal option title
	PurchaseURL            string  `json:"Purchase_URL" gorethink:"Purchase_URL"` // Deal option URL for purchase
	Price                  float32 // Deal option price (in cents)
	FormattedPrice         string  `json:"Formatted_price" gorethink:"Formatted_price"`                   // Deal option price (formatted, e.g. "$6")
	OriginalPrice          float32 `json:"Original_price" gorethink:"Original_price"`                     // Deal option original price (in cents)
	FormattedOriginalPrice string  `json:"Formatted_original_price" gorethink:"Formatted_original_price"` // Deal option original price (formatted, e.g. "$12")
	IsQuantityLimited      bool    `json:"Is_quantity_limited" gorethink:"Is_quantity_limited"`           // Whether the deal option is limited or unlimited
	RemainingCount         float32 `json:"Remaining_count" gorethink:"Remaining_count"`                   // The remaining deal options available for purchase (optional: this field is only present if the deal is limited)
}

// GiftCertificate defines optional data available on Businesses.
type GiftCertificate struct {
	ID             string                   // Gift certificate identifier
	URL            string                   // Gift certificate landing page URL
	ImageURL       string                   `json:"Image_URL" gorethink:"Image_URL"`             //	Gift certificate image URL
	CurrencyCode   string                   `json:"Currency_code" gorethink:"Currency_code"`     // ISO_4217 Currency Code
	UnusedBalances string                   `json:"Unused_balances" gorethink:"Unused_balances"` // Whether unused balances are returned as cash or store credit
	Options        []GiftCertificateOptions //	Gift certificate options
}

// GiftCertificateOptions can define a set of pricing options for a gift certificate.
type GiftCertificateOptions struct {
	Price          float32 `json:"GiftCertificatePrice" gorethink:"GiftCertificatePrice"` //	Gift certificate option price (in cents)
	FormattedPrice string  `json:"Formatted_price" gorethink:"Formatted_price"`           //	Gift certificate option price (formatted, e.g. "$50")
}

// Review data contains a list of user reviews for a given Business (when queried using the Business API).
type Review struct {
	ID                  string  `json:"Rating_ID" gorethink:"Rating_ID"`                           // Review identifier
	Rating              float32 `json:"Rating_rating" gorethink:"Rating_rating"`                   // Rating from 1-5
	RatingImageURL      string  `json:"Rating_image_URL" gorethink:"Rating_image_URL"`             // URL to star rating image for this business (size = 84x17)
	RatingImageLargeURL string  `json:"Rating_image_large_URL" gorethink:"Rating_image_large_URL"` // URL to large version of rating image for this business (size = 166x30)
	Excerpt             string  `json:"Rating_excerpt" gorethink:"Rating_excerpt"`                 // Review excerpt
	TimeCreated         float32 `json:"Time_created" gorethink:""gorethink:"Time_created"`         // Time created (Unix timestamp)
	User                User    // User who wrote the review
}

// User data is linked off of reviews.
type User struct {
	ID       string `json:"User_id" gorethink:"User_id"`               // User identifier
	ImageURL string `json:"User_image_URL" gorethink:"User_image_URL"` // User profile image URL
	Name     string `json:"User_name" gorethink:"User_name"`           // User name
}

// Coordinate data is used with location information.
type Coordinate struct {
	Latitude  float32 `json:"Coordinate_latitude" gorethink:"Coordinate_latitude"`   // Latitude of current location
	Longitude float32 `json:"Coordinate_longitude" gorethink:"Coordinate_longitude"` // Longitude of current location
}

// Location information defines the location of a given business.
type Location struct {
	Coordinate     Coordinate // Address for this business formatted for display. Includes all address fields, cross streets and city, state_code, etc.
	Address        []string   `json:"address" gorethink:"address"`                 // Address for this business. Only includes address fields.
	DisplayAddress []string   `json:"Display_address" gorethink:"Display_address"` // Display address for the business.
	City           string     `json:"State_city" gorethink:"State_city"`           // City for this business
	StateCode      string     `json:"State_code" gorethink:"State_code"`           // ISO 3166-2 state code for this business
	PostalCode     string     `json:"Postal_code" gorethink:"Postal_code"`         // Postal code for this business
	CountryCode    string     `json:"Country_code" gorethink:"Country_code"`       // ISO 3166-1 country code for this business
	CrossStreets   string     `json:"Cross_streets" gorethink:"Cross_streets"`     // Cross streets for this business
	Neighborhoods  []string   // List that provides neighborhood(s) information for business
	GeoAccuracy    float32    `json:"Geo_accuracy" gorethink:"Geo_accuracy"` // Geo accuracy for the location.
}

// Business information is returned in full from the business API, or shallow from the search API.
type Business struct {
	ID                string            `json:"id" gorethink:"id" `                      // Yelp ID for this business
	Name              string            `json:"name" gorethink:"name"`                   // Name of this business
	ImageURL          string            `json:"Image_URL" gorethink:"Image_URL"`         // URL of photo for this business
	URL               string            `json:"business_url" gorethink:"business_url"`   // URL for business page on Yelp
	MobileURL         string            `json:"Mobile_URL" gorethink:"Mobile_URL"`       // URL for mobile business page on Yelp
	Phone             string            `json:"phone_number" gorethink:"phone_number"`   // Phone number for this business with international dialing code (e.g. +442079460000)
	DisplayPhone      string            `json:"Display_phone" gorethink:"Display_phone"` // Phone number for this business formatted for display
	ReviewCount       int               `json:"Review_count" gorethink:"Review_count"`   // Number of reviews for this business
	Categories        [][]string        // Provides a list of category name, alias pairs that this business is associated with. The alias is provided so you can search with the category_filter.
	Distance          float32           `json:"distance" gorethink:"distance"`                          // Distance that business is from search location in meters, if a latitude/longitude is specified.
	Rating            float32           `json:"rating"  gorethink:"rating"`                             // Rating for this business (value ranges from 1, 1.5, ... 4.5, 5)
	RatingImgURL      string            `json:"Rating_img_URL" gorethink:"Rating_img_URL"`              // URL to star rating image for this business (size = 84x17)
	RatingImgURLSmall string            `json:"Rating_img_URL_small" gorethink:"Rating_img_URL_small" ` // URL to small version of rating image for this business (size = 50x10)
	RatingImgURLLarge string            `json:"Rating_img_URL_large" gorethink:"Rating_img_URL_large"`  // URL to large version of rating image for this business (size = 166x30)
	SnippetText       string            `json:"Snippet_text"  gorethink:"Snippet_text"`                 // Snippet text associated with this business
	SnippetImageURL   string            `json:"Snippet_image_URL" gorethink:"Snippet_image_URL"`        // URL of snippet image associated with this business
	Location          Location          // Location data for this business
	IsClaimed         bool              `json:"Is_claimed" gorethink:"Is_claimed"`               // Whether business has been claimed by a business owner
	IsClosed          bool              `json:"Is_closed" gorethink:"Is_closed"`                 // Whether business has been (permanently) closed
	MenuProvider      string            `json:"Menu_provider" gorethink:"Menu_provider"`         // Provider of the menu for this business
	MenuDateUpdated   float32           `json:"Menu_date_updated" gorethink:"Menu_date_updated"` // Last time this menu was updated on Yelp (Unix timestamp)
	Deals             []Deal            // Deal info for this business (optional: this field is present only if there’s a Deal)
	GiftCertificates  []GiftCertificate `json:"Gift_certificates" gorethink:"Gift_certificates"` // Gift certificate info for this business (optional: this field is present only if there are gift certificates available)
	Reviews           []Review          // Contains one review associated with business
}
