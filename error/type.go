package error

type ErrorType string

var (
	ErrorTypeValidate       ErrorType = "ERROR1"
	ErrorTypeLogic          ErrorType = "ERROR2"
	ErrorTypeAuthentication ErrorType = "ERROR3"
)

type CodeValidate struct {
	Code        string
	Description string
}

var MapCodeValidate = map[string]CodeValidate{
	"required": {
		Code:        "01",
		Description: "required",
	},
	"required_if": {
		Code:        "01",
		Description: "required",
	},
	"max": {
		Code:        "21",
		Description: "maxLength",
	},
	"maxLengthWithoutRequire": {
		Code:        "21",
		Description: "maxLength",
	},
	"maxElementFileIds": {
		Code:        "95",
		Description: "maxElementFileIds",
	},
	"minLength": {
		Code:        "22",
		Description: "minLength",
	},
	"prohibitedCharacter": {
		Code:        "23",
		Description: "prohibitedCharacter",
	},
	"fixedStringLength": {
		Code:        "24",
		Description: "fixedStringLength",
	},
	"oneof": {
		Code:        "27",
		Description: "enum",
	},
	"masterComplementType": {
		Code:        "27",
		Description: "enum",
	},
	"inputRestriction": {
		Code:        "27",
		Description: "enum",
	},
	"numberFormatOnly": {
		Code:        "30",
		Description: "numberFormatOnly",
	},
	"gte": {
		Code:        "33",
		Description: "numberMin",
	},
	"lte": {
		Code:        "33",
		Description: "numberMax",
	},
	"minMaxMaintain": {
		Code:        "33",
		Description: "numberRange",
	},
	"fixedDecimalNumber": {
		Code:        "34",
		Description: "fixedDecimalNumber",
	},
	"dateFormatOnly": {
		Code:        "41",
		Description: "dateFormatOnly",
	},
	"dateInRangeOnly": {
		Code:        "42",
		Description: "dateInRangeOnly",
	},
	"pastDateNotAllowed": {
		Code:        "43",
		Description: "pastDateNotAllowed",
	},
	"ltfield": {
		Code:        "44",
		Description: "pastDateNotAllowed",
	},
	"urlFormatOnly": {
		Code:        "51",
		Description: "urlFormatOnly",
	},
	"creditCardNumberFormatOnly": {
		Code:        "52",
		Description: "creditCardNumberFormatOnly",
	},
	"customPassword": {
		Code:        "53",
		Description: "passwordFormatOnly",
	},
	"email": {
		Code:        "54",
		Description: "emailFormatOnly",
	},
	"customEmail": {
		Code:        "54",
		Description: "emailFormatOnly",
	},
	"postalCodeFormat": {
		Code:        "55",
		Description: "postalCodeFormat",
	},
	"telephoneNumberFormatOnly": {
		Code:        "56",
		Description: "telephoneNumberFormatOnly",
	},
	"minElement": {
		Code:        "60",
		Description: "minElement",
	},
	"fullWidthKatakanaOnly": {
		Code:        "71",
		Description: "fullWidthKatakanaOnly",
	},
	"fullWidthOnly": {
		Code:        "72",
		Description: "fullWidthOnly",
	},
	"halfWidthKatakanaOnly": {
		Code:        "73",
		Description: "halfWidthKatakanaOnly",
	},
	"halfWidthOnly": {
		Code:        "74",
		Description: "halfWidthOnly",
	},
	"halfWidthNumericOnly": {
		Code:        "75",
		Description: "halfWidthNumericOnly",
	},
	"halfWidthAlphaNumericStringOnly": {
		Code:        "76",
		Description: "halfWidthAlphaNumericStringOnly",
	},
	"uppercaseHalfWidthAlphaNumericStringOnly": {
		Code:        "77",
		Description: "uppercaseHalfWidthAlphaNumericStringOnly",
	},
	"byteInRangeOnly": {
		Code:        "91",
		Description: "byteInRangeOnly",
	},
	"fixedByteLength": {
		Code:        "92",
		Description: "fixedByteLength",
	},
	"maxFileSize": {
		Code:        "93",
		Description: "maxFileSize",
	},
	"prohibitedFileType": {
		Code:        "94",
		Description: "prohibitedFileType",
	},
	"maxFileNumber": {
		Code:        "95",
		Description: "maxFileNumber",
	},
	"attachFileRequired": {
		Code:        "96",
		Description: "attachFileRequired",
	},
	"imageAndPDFFileTypeOnly": {
		Code:        "97",
		Description: "imageAndPDFFileTypeOnly",
	},
	"imageFileTypeOnly": {
		Code:        "98",
		Description: "imageFileTypeOnly",
	},
	"fileType": {
		Code:        "27",
		Description: "enum",
	},
	"callbackContact": {
		Code:        "56",
		Description: "callBackContactFormat",
	},
}
