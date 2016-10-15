package pl

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type pl struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyPositiveSuffix string
	currencyNegativePrefix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'pl' locale
func New() locales.Translator {
	return &pl{
		locale:                 "pl",
		pluralsCardinal:        []locales.PluralRule{2, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6, 2, 4, 5},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP ", "AED", "AFA ", "AFN", "ALK ", "ALL", "AMD", "ANG", "AOA", "AOK ", "AON ", "AOR ", "ARA ", "ARL ", "ARM ", "ARP ", "ARS", "ATS ", "AUD", "AWG", "AZM ", "AZN", "BAD ", "BAM", "BAN ", "BBD", "BDT", "BEC ", "BEF ", "BEL ", "BGL", "BGM", "BGN", "BGO ", "BHD", "BIF", "BMD", "BND", "BOB", "BOL ", "BOP ", "BOV ", "BRB ", "BRC ", "BRE ", "R$", "BRN ", "BRR ", "BRZ ", "BSD", "BTN", "BUK ", "BWP", "BYB ", "BYR", "BZD", "CAD", "CDF", "CHE ", "CHF", "CHW ", "CLE ", "CLF ", "CLP", "CNX ", "CNY", "COP", "COU ", "CRC", "CSD ", "CSK ", "CUC", "CUP", "CVE", "CYP ", "CZK", "DDM ", "DEM ", "DJF", "DKK", "DOP", "DZD", "ECS ", "ECV ", "EEK ", "EGP", "ERN", "ESA ", "ESB ", "ESP ", "ETB", "€", "FIM ", "FJD", "FKP", "FRF ", "GBP", "GEK ", "GEL", "GHC ", "GHS", "GIP", "GMD", "GNF", "GNS ", "GQE ", "GRD ", "GTQ", "GWE ", "GWP ", "GYD", "HKD", "HNL", "HRD ", "HRK", "HTG", "HUF", "IDR", "IEP ", "ILP ", "ILR ", "ILS", "INR", "IQD", "IRR", "ISJ ", "ISK", "ITL ", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH ", "KRO ", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL ", "LTL", "LTT ", "LUC ", "LUF ", "LUL ", "LVL", "LVR ", "LYD", "MAD", "MAF ", "MCF ", "MDC ", "MDL", "MGA", "MGF ", "MKD", "MKN ", "MLF ", "MMK", "MNT", "MOP", "MRO", "MTL ", "MTP ", "MUR", "MVP ", "MVR", "MWK", "MXN", "MXP ", "MXV ", "MYR", "MZE ", "MZM ", "MZN", "NAD", "NGN", "NIC ", "NIO", "NLG ", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI ", "PEN", "PES ", "PGK", "PHP", "PKR", "zł", "PLZ ", "PTE ", "PYG", "QAR", "RHD ", "ROL ", "RON", "RSD", "RUB", "RUR ", "RWF", "SAR", "SBD", "SCR", "SDD ", "SDG", "SDP ", "SEK", "SGD", "SHP", "SIT ", "SKK ", "SLL", "SOS", "SRD", "SRG ", "SSP", "STD", "SUR ", "SVC ", "SYP", "SZL", "THB", "TJR ", "TJS", "TMM ", "TMT", "TND", "TOP", "TPE ", "TRL ", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK ", "UGS ", "UGX", "USD", "USN ", "USS ", "UYI ", "UYP ", "UYU", "UZS", "VEB ", "VEF", "VND", "VNN ", "VUV", "WST", "FCFA", "XAG ", "XAU ", "XBA ", "XBB ", "XBC ", "XBD ", "EC$", "XDR ", "XEU ", "XFO ", "XFU ", "CFA", "XPD ", "CFPF", "XPT ", "XRE ", "XSU ", "XTS ", "XUA ", "XXX ", "YDD ", "YER", "YUD ", "YUM ", "YUN ", "YUR ", "ZAL ", "ZAR", "ZMK ", "ZMW", "ZRN ", "ZRZ ", "ZWD ", "ZWL ", "ZWR "},
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "sty", "lut", "mar", "kwi", "maj", "cze", "lip", "sie", "wrz", "paź", "lis", "gru"},
		monthsNarrow:           []string{"", "s", "l", "m", "k", "m", "c", "l", "s", "w", "p", "l", "g"},
		monthsWide:             []string{"", "stycznia", "lutego", "marca", "kwietnia", "maja", "czerwca", "lipca", "sierpnia", "września", "października", "listopada", "grudnia"},
		daysAbbreviated:        []string{"niedz.", "pon.", "wt.", "śr.", "czw.", "pt.", "sob."},
		daysNarrow:             []string{"N", "P", "W", "Ś", "C", "P", "S"},
		daysShort:              []string{"nie", "pon", "wto", "śro", "czw", "pią", "sob"},
		daysWide:               []string{"niedziela", "poniedziałek", "wtorek", "środa", "czwartek", "piątek", "sobota"},
		periodsAbbreviated:     []string{"AM", "PM"},
		periodsNarrow:          []string{"a", "p"},
		periodsWide:            []string{"AM", "PM"},
		erasAbbreviated:        []string{"p.n.e.", "n.e."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"p.n.e.", "n.e."},
		timezones:              map[string]string{"ACDT": "Czas środkowoaustralijski letni", "JST": "Japonia (czas standardowy)", "WAST": "Czas zachodnioafrykański letni", "∅∅∅": "Brasília (czas letni)", "CDT": "Czas środkowoamerykański letni", "PDT": "Czas pacyficzny letni", "TMT": "Turkmenistan (czas standardowy)", "AEDT": "Czas wschodnioaustraliski letni", "CLT": "Chile (czas standardowy)", "WITA": "Indonezja Środkowa", "AKST": "Alaska (czas standardowy)", "EST": "Czas wschodnioamerykański standardowy", "COST": "Kolumbia (czas letni)", "WESZ": "Czas zachodnioeuropejski letni", "HAT": "Nowa Fundlandia (czas letni)", "ChST": "Czamorro", "WIT": "Indonezja Wschodnia", "AKDT": "Alaska (czas letni)", "NZDT": "Nowa Zelandia (czas letni)", "COT": "Kolumbia (czas standardowy)", "MEZ": "Czas środkowoeuropejski standardowy", "BOT": "Boliwia", "SGT": "Singapur", "WIB": "Indonezja Zachodnia", "WARST": "Argentyna Zachodnia (czas letni)", "SRT": "Surinam", "HKT": "Hongkong (czas standardowy)", "ART": "Argentyna (czas standardowy)", "HNT": "Nowa Fundlandia (czas standardowy)", "HADT": "Hawaje-Aleuty (czas letni)", "EDT": "Czas wschodnioamerykański letni", "LHDT": "Lord Howe (czas letni)", "ACST": "Czas środkowoaustralijski standardowy", "CHADT": "Chatham (czas letni)", "HAST": "Hawaje-Aleuty (czas standardowy)", "HKST": "Hongkong (czas letni)", "AWST": "Czas zachodnioaustralijski standardowy", "WAT": "Czas zachodnioafrykański standardowy", "MYT": "Malezja", "JDT": "Japonia (czas letni)", "WEZ": "Czas zachodnioeuropejski standardowy", "GYT": "Gujana", "OEZ": "Czas wschodnioeuropejski standardowy", "OESZ": "Czas wschodnioeuropejski letni", "UYST": "Urugwaj (czas letni)", "LHST": "Lord Howe (czas standardowy)", "CST": "Czas środkowoamerykański standardowy", "AWDT": "Czas zachodnioaustralijski letni", "GMT": "Czas uniwersalny", "CHAST": "Chatham (czas standardowy)", "ACWDT": "Czas środkowo-zachodnioaustralijski letni", "SAST": "Czas południowoafrykański", "IST": "Czas indyjski", "TMST": "Turkmenistan (czas letni)", "MST": "Czas górski standardowy", "WART": "Argentyna Zachodnia (czas standardowy)", "AST": "Czas standardowy atlantycki", "ADT": "Czas atlantycki letni", "AEST": "Czas standardowy wschodnioaustraliski", "PST": "Czas pacyficzny standardowy", "EAT": "Czas wschodnioafrykański", "BT": "Bhutan", "MDT": "Czas górski letni", "NZST": "Nowa Zelandia (czas standardowy)", "GFT": "Gujana Francuska", "ACWST": "Czas środkowo-zachodnioaustralijski standardowy", "MESZ": "Czas środkowoeuropejski letni", "CAT": "Czas środkowoafrykański", "ARST": "Argentyna (czas letni)", "CLST": "Chile (czas letni)", "UYT": "Urugwaj (czas standardowy)", "VET": "Wenezuela", "ECT": "Ekwador"},
	}
}

// Locale returns the current translators string locale
func (pl *pl) Locale() string {
	return pl.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'pl'
func (pl *pl) PluralsCardinal() []locales.PluralRule {
	return pl.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'pl'
func (pl *pl) PluralsOrdinal() []locales.PluralRule {
	return pl.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'pl'
func (pl *pl) PluralsRange() []locales.PluralRule {
	return pl.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'pl'
func (pl *pl) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	iMod10 := i % 10
	iMod100 := i % 100

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	} else if v == 0 && iMod10 >= 2 && iMod10 <= 4 && iMod100 < 12 && iMod100 > 14 {
		return locales.PluralRuleFew
	} else if (v == 0 && i != 1 && iMod10 >= 0 && iMod10 <= 1) || (v == 0 && iMod10 >= 5 && iMod10 <= 9) || (v == 0 && iMod100 >= 12 && iMod100 <= 14) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'pl'
func (pl *pl) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'pl'
func (pl *pl) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := pl.CardinalPluralRule(num1, v1)
	end := pl.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (pl *pl) MonthAbbreviated(month time.Month) string {
	return pl.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (pl *pl) MonthsAbbreviated() []string {
	return pl.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (pl *pl) MonthNarrow(month time.Month) string {
	return pl.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (pl *pl) MonthsNarrow() []string {
	return pl.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (pl *pl) MonthWide(month time.Month) string {
	return pl.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (pl *pl) MonthsWide() []string {
	return pl.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (pl *pl) WeekdayAbbreviated(weekday time.Weekday) string {
	return pl.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (pl *pl) WeekdaysAbbreviated() []string {
	return pl.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (pl *pl) WeekdayNarrow(weekday time.Weekday) string {
	return pl.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (pl *pl) WeekdaysNarrow() []string {
	return pl.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (pl *pl) WeekdayShort(weekday time.Weekday) string {
	return pl.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (pl *pl) WeekdaysShort() []string {
	return pl.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (pl *pl) WeekdayWide(weekday time.Weekday) string {
	return pl.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (pl *pl) WeekdaysWide() []string {
	return pl.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'pl' and handles both Whole and Real numbers based on 'v'
func (pl *pl) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(pl.group) - 1; j >= 0; j-- {
					b = append(b, pl.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, pl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'pl' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (pl *pl) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pl.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, pl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, pl.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'pl'
func (pl *pl) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pl.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(pl.group) - 1; j >= 0; j-- {
					b = append(b, pl.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, pl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, pl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, pl.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'pl'
// in accounting notation.
func (pl *pl) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := pl.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, pl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(pl.group) - 1; j >= 0; j-- {
					b = append(b, pl.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, pl.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, pl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, pl.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, pl.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'pl'
func (pl *pl) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'pl'
func (pl *pl) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'pl'
func (pl *pl) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'pl'
func (pl *pl) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, pl.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, pl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'pl'
func (pl *pl) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'pl'
func (pl *pl) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'pl'
func (pl *pl) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'pl'
func (pl *pl) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, pl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, pl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := pl.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
