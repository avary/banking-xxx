package lib

type regex struct {
	Pattern string
	Error   string
}

// NameRegex for Name
var NameRegex = regex{
	Pattern: `^[a-zA-Z]{2,20}$`,
	Error:   "Name must be between 1 and 20 characters and can only contain letters",
}

// CityRegex for city
var CityRegex = regex{
	Pattern: `^[a-zA-Z]{2,20}$`,
	Error:   "City must be between 2 and 20 characters long and can only contain letters",
}

// ZipRegex for zipcode
var ZipRegex = regex{
	Pattern: `^[0-9]{4}$`,
	Error:   "Zipcode must be exactly 4 digits long",
}

// DOBRegex for date of birth
var DOBRegex = regex{
	Pattern: `^[0-9]{4}-[0-9]{2}-[0-9]{2}$`,
	Error:   "Date of birth must be in the format YYYY-MM-DD",
}
