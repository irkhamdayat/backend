package constant

import "strings"

type CtxKey string

const (
	DB          CtxKey = "db"
	Lang               = "lang"
	LangID             = "ID"
	LangEN             = "EN"
	LangDefault        = LangEN
)

var (
	Langs = []string{LangID, LangEN}
)

const (
	AlphaNumeric        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	AlphaCapitalNumeric = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numeric             = "0123456789"
	FileExtensionJPG    = ".jpg"
	FileExtensionJPEG   = ".jpeg"
	FileExtensionPNG    = ".png"
	FileExtensionHEIC   = ".heic"
	FileExtensionPDF    = ".pdf"
	FileExtensionZIP    = ".zip"
	FileExtensionXLSX   = ".xlsx"
)

var FileTypeDocument = []string{
	FileExtensionXLSX,
}

var FileTypeImage = []string{
	FileExtensionHEIC,
	FileExtensionPNG,
	FileExtensionJPG,
	FileExtensionJPEG,
}

var FileTypeRichMedia = []string{
	FileExtensionPNG,
	FileExtensionJPG,
	FileExtensionJPEG,
}

var FileTypeImageAndPDF = append(FileTypeImage, FileExtensionPDF)

const (
	January   = "January"
	February  = "February"
	March     = "March"
	April     = "April"
	May       = "May"
	June      = "June"
	July      = "July"
	August    = "August"
	September = "September"
	October   = "October"
	November  = "November"
	December  = "December"
)

var (
	MonthList = []string{
		January, February, March, April, May, June, July, August, September, October, November, December,
	}
)

var (
	MapIntMonthToStrMonthLowerCase = map[int]string{
		1:  strings.ToLower(January),
		2:  strings.ToLower(February),
		3:  strings.ToLower(March),
		4:  strings.ToLower(April),
		5:  strings.ToLower(May),
		6:  strings.ToLower(June),
		7:  strings.ToLower(July),
		8:  strings.ToLower(August),
		9:  strings.ToLower(September),
		10: strings.ToLower(October),
		11: strings.ToLower(November),
		12: strings.ToLower(December),
	}
)
