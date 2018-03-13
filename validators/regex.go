package validators

import "regexp"

// Code in this file is taken from asaskevich/govalidator

// Basic regular expressions for validating strings
const (
	// Email RegExp
	Email = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	// CreditCard RegExp
	CreditCard = "^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$"
	// ISBN10 RegExp
	ISBN10 = "^(?:[0-9]{9}X|[0-9]{10})$"
	// ISBN13 RegExp
	ISBN13 = "^(?:[0-9]{13})$"
	// UUID3 RegExp
	UUID3 = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
	// UUID4 RegExp
	UUID4 = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	// UUID5 RegExp
	UUID5 = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	// UUID RegExp
	UUID = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	// Alpha RegExp
	Alpha = "^[a-zA-Z]+$"
	// Alphanumeric RegExp
	Alphanumeric = "^[a-zA-Z0-9]+$"
	// Numeric RegExp
	Numeric = "^[0-9]+$"
	// Int RegExp
	Int = "^(?:[-+]?(?:0|[1-9][0-9]*))$"
	// Float RegExp
	Float = "^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$"
	// Hexadecimal RegExp
	Hexadecimal = "^[0-9a-fA-F]+$"
	// Hexcolor RegExp
	Hexcolor = "^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
	// RGBcolor RegExp
	RGBcolor = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"
	// ASCII RegExp
	ASCII = "^[\x00-\x7F]+$"
	// Multibyte RegExp
	Multibyte = "[^\x00-\x7F]"
	// FullWidth RegExp
	FullWidth = "[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	// HalfWidth RegExp
	HalfWidth = "[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	// Base64 RegExp
	Base64 = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	// PrintableASCII RegExp
	PrintableASCII = "^[\x20-\x7E]+$"
	// DataURI RegExp
	DataURI = "^data:.+\\/(.+);base64$"
	// Latitude RegExp
	Latitude = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	// Longitude RegExp
	Longitude = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	// DNSName RegExp
	DNSName = `^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`
	// IP RegExp
	IP = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	// URLSchema RegExp
	URLSchema = `((ftp|tcp|udp|wss?|https?):\/\/)`
	// URLUsername RegExp
	URLUsername = `(\S+(:\S*)?@)`
	// URLPath RegExp
	URLPath = `((\/|\?|#)[^\s]*)`
	// URLPort RegExp
	URLPort = `(:(\d{1,5}))`
	// URLIP RegExp
	URLIP = `([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))`
	// URLSubdomain RegExp
	URLSubdomain = `((www\.)|([a-zA-Z0-9]([-\.][-\._a-zA-Z0-9]+)*))`
	// URL RegExp
	URL = `^` + URLSchema + `?` + URLUsername + `?` + `((` + URLIP + `|(\[` + IP + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + URLSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + URLPort + `?` + URLPath + `?$`
	// SSN RegExp
	SSN = `^\d{3}[- ]?\d{2}[- ]?\d{4}$`
	// WinPath RegExp
	WinPath = `^[a-zA-Z]:\\(?:[^\\/:*?"<>|\r\n]+\\)*[^\\/:*?"<>|\r\n]*$`
	// UnixPath RegExp
	UnixPath = `^(/[^/\x00]*)+/?$`
	// Semver RegExp
	Semver = "^v?(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$"
	// tagName RegExp
	tagName = "valid"
	// hasLowerCase RegExp
	hasLowerCase = ".*[[:lower:]]"
	// hasUpperCase RegExp
	hasUpperCase = ".*[[:upper:]]"
)

// Used by IsFilePath func
const (
	// Unknown is unresolved OS type
	Unknown = iota
	// Win is Windows type
	Win
	// Unix is *nix OS types
	Unix
)

var (
	// EmailRegExp defines a RegExp that handles checking for Email values
	EmailRegExp = regexp.MustCompile(Email)
	// CreditCardRegExp defines a RegExp that handles checking for CreditCard values
	CreditCardRegExp = regexp.MustCompile(CreditCard)
	// ISBN10RegExp defines a RegExp that handles checking for ISBN10 values
	ISBN10RegExp = regexp.MustCompile(ISBN10)
	// ISBN13RegExp defines a RegExp that handles checking for ISBN13 values
	ISBN13RegExp = regexp.MustCompile(ISBN13)
	// UUID3RegExp defines a RegExp that handles checking for UUID3 values
	UUID3RegExp = regexp.MustCompile(UUID3)
	// UUID4RegExp defines a RegExp that handles checking for UUID4 values
	UUID4RegExp = regexp.MustCompile(UUID4)
	// UUID5RegExp defines a RegExp that handles checking for UUID5 values
	UUID5RegExp = regexp.MustCompile(UUID5)
	// UUIDRegExp defines a RegExp that handles checking for UUID values
	UUIDRegExp = regexp.MustCompile(UUID)
	// AlphaRegExp defines a RegExp that handles checking for Alpha values
	AlphaRegExp = regexp.MustCompile(Alpha)
	// AlphanumericRegExp defines a RegExp that handles checking for Alphanumeric values
	AlphanumericRegExp = regexp.MustCompile(Alphanumeric)
	// NumericRegExp defines a RegExp that handles checking for Numeric values
	NumericRegExp = regexp.MustCompile(Numeric)
	// IntRegExp defines a RegExp that handles checking for Int values
	IntRegExp = regexp.MustCompile(Int)
	// FloatRegExp defines a RegExp that handles checking for Float values
	FloatRegExp = regexp.MustCompile(Float)
	// HexadecimalRegExp defines a RegExp that handles checking for Hexadecimal values
	HexadecimalRegExp = regexp.MustCompile(Hexadecimal)
	// HexcolorRegExp defines a RegExp that handles checking for Hexcolor values
	HexcolorRegExp = regexp.MustCompile(Hexcolor)
	// RGBcolorRegExp defines a RegExp that handles checking for RGBcolor values
	RGBcolorRegExp = regexp.MustCompile(RGBcolor)
	// ASCIIRegExp defines a RegExp that handles checking for ASCII values
	ASCIIRegExp = regexp.MustCompile(ASCII)
	// PrintableASCIIRegExp defines a RegExp that handles checking for PrintableASCII values
	PrintableASCIIRegExp = regexp.MustCompile(PrintableASCII)
	// MultibyteRegExp defines a RegExp that handles checking for Multibyte values
	MultibyteRegExp = regexp.MustCompile(Multibyte)
	// FullWidthRegExp defines a RegExp that handles checking for FullWidth values
	FullWidthRegExp = regexp.MustCompile(FullWidth)
	// HalfWidthRegExp defines a RegExp that handles checking for HalfWidth values
	HalfWidthRegExp = regexp.MustCompile(HalfWidth)
	// Base64RegExp defines a RegExp that handles checking for Base64 values
	Base64RegExp = regexp.MustCompile(Base64)
	// DataURIRegExp defines a RegExp that handles checking for DataURI values
	DataURIRegExp = regexp.MustCompile(DataURI)
	// LatitudeRegExp defines a RegExp that handles checking for Latitude values
	LatitudeRegExp = regexp.MustCompile(Latitude)
	// LongitudeRegExp defines a RegExp that handles checking for Longitude values
	LongitudeRegExp = regexp.MustCompile(Longitude)
	// DNSNameRegExp defines a RegExp that handles checking for DNSName values
	DNSNameRegExp = regexp.MustCompile(DNSName)
	// URLRegExp defines a RegExp that handles checking for URL values
	URLRegExp = regexp.MustCompile(URL)
	// SSNRegExp defines a RegExp that handles checking for SSN values
	SSNRegExp = regexp.MustCompile(SSN)
	// WinPathRegExp defines a RegExp that handles checking for WinPath values
	WinPathRegExp = regexp.MustCompile(WinPath)
	// UnixPathRegExp defines a RegExp that handles checking for UnixPath values
	UnixPathRegExp = regexp.MustCompile(UnixPath)
	// SemverRegExp defines a RegExp that handles checking for Semver values
	SemverRegExp = regexp.MustCompile(Semver)
	// HasLowerCaseRegExp defines a RegExp that handles checking for hasLowerCase values
	HasLowerCaseRegExp = regexp.MustCompile(hasLowerCase)
	// HasUpperCaseRegExp defines a RegExp that handles checking for hasUpperCase values
	HasUpperCaseRegExp = regexp.MustCompile(hasUpperCase)
)
