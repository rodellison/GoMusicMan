package models

//From: https://gist.github.com/tmaiaroto/4ec7668ae986335b0a6d
var (
	USC_ABBREV map[string]string
	USC        map[string]string
)

func init() {

	USC = map[string]string{
		"AL": "Alabama",
		"AK": "Alaska",
		"AZ": "Arizona",
		"AR": "Arkansas",
		"CA": "California",
		"CO": "Colorado",
		"CT": "Connecticut",
		"DE": "Delaware",
		"FL": "Florida",
		"GA": "Georgia",
		"HI": "Hawaii",
		"ID": "Idaho",
		"IL": "Illinois",
		"IN": "Indiana",
		"IA": "Iowa",
		"KS": "Kansas",
		"KY": "Kentucky",
		"LA": "Louisiana",
		"ME": "Maine",
		"MD": "Maryland",
		"MA": "Massachusetts",
		"MI": "Michigan",
		"MN": "Minnesota",
		"MS": "Mississippi",
		"MO": "Missouri",
		"MT": "Montana",
		"NE": "Nebraska",
		"NV": "Nevada",
		"NH": "New Hampshire",
		"NJ": "New Jersey",
		"NM": "New Mexico",
		"NY": "New York",
		"NC": "North Carolina",
		"ND": "North Dakota",
		"OH": "Ohio",
		"OK": "Oklahoma",
		"OR": "Oregon",
		"PA": "Pennsylvania",
		"RI": "Rhode Island",
		"SC": "South Carolina",
		"SD": "South Dakota",
		"TN": "Tennessee",
		"TX": "Texas",
		"UT": "Utah",
		"VT": "Vermont",
		"VA": "Virginia",
		"WA": "Washington",
		"WV": "West Virginia",
		"WI": "Wisconsin",
		"WY": "Wyoming",
		// Territories
		"AS": "American Samoa",
		"DC": "District of Columbia",
		"FM": "Federated States of Micronesia",
		"GU": "Guam",
		"MH": "Marshall Islands",
		"MP": "Northern Mariana Islands",
		"PW": "Palau",
		"PR": "Puerto Rico",
		"VI": "Virgin Islands",
		// Armed Forces (AE includes Europe, Africa, Canada, and the Middle East)
		"AA": "Armed Forces Americas",
		"AE": "Armed Forces Europe",
		"AP": "Armed Forces Pacific",
	}

	USC_ABBREV = ReverseUSCMap(USC)

}

func ReverseUSCMap(m map[string]string) map[string]string {
	n := make(map[string]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}
