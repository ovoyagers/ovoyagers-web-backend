package usermodel

// Custom error messages for validation errors
var customErrorMessages = map[string]string{
	"CountryCode.iso3166_1_alpha2": "countryCode must be a valid ISO 3166-1 alpha-2 country code",
	"CountryCode.required":         "countryCode is required",
	"Phone.required":               "phone is required",
	"Phone.e164":                   "phone must be in E.164 format",
	"Name.required":                "name is required",
	"Name.min":                     "name must be at least 2 characters",
	"Name.max":                     "name must be at most 100 characters",
	"Email.required":               "email is required",
	"Email.email":                  "email must be a valid email address",
	"DOB.required":                 "dob is required",
	"Age.gte":                      "age must be at least 13",
	"Age.lte":                      "age must be at most 130",
	"Gender.required":              "gender is required",
	"Gender.oneof":                 "gender must be one of 'male', 'female', or 'other'",
	"TextBio.max":                  "Text bio must be at most 500 characters",
	"Fullname.required":            "fullname is required",
	"Fullname.min":                 "fullname must be at least 2 characters",
	"Fullname.max":                 "fullname must be at most 100 characters",
	"Fullname.alpha":               "fullname must only contain letters",
	"PreferredLanguages.required":  "preferredLanguages is required",
	"NativeLanguages.required":     "nativeLanguages is required",
	"PreferredLanguages.anyof":     "preferredLanguages must be any of 'English', 'Spanish', 'French', 'German', 'Telugu', 'Tamil', 'Hindi', 'Gujarati', or 'Marathi'",
	"NativeLanguages.anyof":        "nativeLanguages must be any of 'English', 'Spanish', 'French', 'German', 'Telugu', 'Tamil', 'Hindi', 'Gujarati', or 'Marathi'",
}
