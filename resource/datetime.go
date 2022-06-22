package resource

import (
	"regexp"
	"time"
)

var looksLikeFourDigitTZ = regexp.MustCompile(`\d\d\d\d$`)

func ParseDateTime(s string) (time.Time, error) {
	// https://datatracker.ietf.org/doc/html/rfc7643#section-2.3.5
	// xsd:dateTime (https://www.w3.org/TR/xmlschema11-2/#dateTime)

	// base pattern = YYYY-MM-DDThh:mm:ss(.nnn)
	// time zone can be 'Z' or +/-ttt(:tt)
	// Z07:00 handles Z or +/-07:00, but does not handle +/-0700,
	// so we're going to have to use hueristics here

	if looksLikeFourDigitTZ.MatchString(s) {
		const tzfmt = "2006-01-02T15:04:05.999999999Z0700"
		return time.Parse(tzfmt, s)
	}
	const tzfmt = "2006-01-02T15:04:05.999999999Z07:00"
	return time.Parse(tzfmt, s)
}
