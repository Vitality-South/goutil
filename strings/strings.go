// Package strings implements helper functions for working with strings
package strings

import (
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

// MaskLeft masks a sensitive string, such as an API Key, so that it can be safely logged
// all of the data will be masked with the provided rune (mask) except the last (keep) runes
//
// Keep is the number of runes to keep. If keep is zero or less than, all of the text is masked
//
// # This function operates on runes to remain Unicode safe
//
// NOTE: len([]rune{}) is safe, otherwise use utf8.RuneCountInString(string) as equivalent for strings
func MaskLeft(mask rune, keep int, text string) string {
	rs := []rune(text)

	// bound keep to minimum zero and maximum len(rs)
	k := func() int {
		if keep < 0 {
			return 0
		}

		if keep > len(rs) {
			return len(rs)
		}

		return keep
	}()

	// if we are keeping the whole text, go ahead and return it
	if len(rs) <= k {
		return string(rs)
	}

	// mask out what we're not keeping
	b := strings.Repeat(string(mask), len(rs)-k)
	last := rs[len(rs)-k:]
	bs := b + string(last)

	return string(bs)
}

// MaskRight masks a sensitive string, such as an API Key, so that it can be safely logged
// all of the data will be masked with the provided rune (mask) except the first (keep) runes
//
// Keep is the number of runes to keep. If keep is zero or less than, all of the text is masked
//
// # This function operates on runes to remain Unicode safe
//
// NOTE: len([]rune{}) is safe, otherwise use utf8.RuneCountInString(string) as equivalent for strings
func MaskRight(mask rune, keep int, text string) string {
	rs := []rune(text)

	// bound keep to minimum zero and maximum len(rs)
	k := func() int {
		if keep < 0 {
			return 0
		}

		if keep > len(rs) {
			return len(rs)
		}

		return keep
	}()

	// if we are keeping the whole text, go ahead and return it
	if len(rs) <= k {
		return string(rs)
	}

	// mask out what we're not keeping
	b := strings.Repeat(string(mask), len(rs)-k)
	first := rs[:k]
	bs := string(first) + b

	return string(bs)
}

// ToValidUTF8Trimmed removes each run of invalid UTF-8 byte sequences from the string s,
// with all leading and trailing white space removed, as defined by Unicode.
func ToValidUTF8Trimmed(s string) string {
	v := strings.ToValidUTF8(s, "")
	return strings.TrimSpace(v)
}

// StripHTMLTags will strip HTML tags from a string input and return the stripped output
// Content between certain tags, like script and style tags, are also remove
// Certain characters are converted to HTML entities, such as > becoming &gt;
//
// Recommendation: use "github.com/microcosm-cc/bluemonday" if you want greater control
//
// This is pure
func StripHTMLTags(s string) string {
	p := bluemonday.StrictPolicy()
	return p.Sanitize(s)
}

// TruncateRight will truncate a string, keeping the string starting from the
// beginning, and return a new string that is exactly the "length" of runes
// provided, with the final part of the string that was replaced now being
// the provided "ellipsis". If "length" is greater than or equal to the
// length of "text" (in runes), then "text" is returned as is. If "length"
// is less than or equal to the length of "ellipsis", then only "length"
// runes of "ellipsis" is returned.
//
// This function operates on runes to remain Unicode safe.
func TruncateRight(text string, length int, ellipsis string) string {
	rs := []rune(text)

	// bound length to minimum zero and maximum len(rs)
	l := func() int {
		if length < 0 {
			return 0
		}

		if length > len(rs) {
			return len(rs)
		}

		return length
	}()

	// if we are keeping the whole text, go ahead and return it
	if len(rs) <= l {
		return string(rs)
	}

	// truncate the end of the string
	r := rs[:l]

	// copy ellipsis into the new string, but only enough of it that will fit
	// by overwriting the final runes of r with runes from ellipsis
	es := []rune(ellipsis)
	min := func() int {
		if len(r) < len(es) {
			return len(r)
		}

		return len(es)
	}()
	copy(r[len(r)-min:], es[:min])

	return string(r)
}

// TruncateLeft will truncate a string, keeping the string starting from the
// end, and return a new string that is exactly the "length" provided,
// with the beginning part of the string that was replaced now being
// the provided "ellipsis". If "length" is greater than or equal to the
// length of "text" (in runes), then "text" is returned as is.
//
// This function operates on runes to remain Unicode safe.
func TruncateLeft(text string, length int, ellipsis string) string {
	rs := []rune(text)

	// bound length to minimum zero and maximum len(rs)
	l := func() int {
		if length < 0 {
			return 0
		}

		if length > len(rs) {
			return len(rs)
		}

		return length
	}()

	// if we are keeping the whole text, go ahead and return it
	if len(rs) <= l {
		return string(rs)
	}

	// truncate the beginning of the string
	r := rs[len(rs)-l:]

	// copy ellipsis into the new string, but only enough of it that will fit
	// by overwriting the beginning runes of r with runes from ellipsis
	es := []rune(ellipsis)
	min := func() int {
		if len(r) < len(es) {
			return len(r)
		}

		return len(es)
	}()
	copy(r, es[:min])

	return string(r)
}

// IsAffirmative returns true if the string represents a true-ish value.
// This function is case  insensitive.
//
// Inspired by https://github.com/jonschlinkert/is-affirmative and
// https://github.com/jonschlinkert/affirmative
//
// These are the current affirmative words:
//
//	'absolutely',
//	'affirmative',
//	'all right',
//	'amen',
//	'aye',
//	'beyond a doubt',
//	'by all means',
//	'certainly',
//	'definitely',
//	'even so',
//	'exactly',
//	'fine',
//	'gladly',
//	'good enough',
//	'good',
//	'granted',
//	'i accept',
//	'i concur',
//	'i guess',
//	'if you must',
//	'indubitably',
//	'just so',
//	'most assuredly',
//	'naturally',
//	'of course',
//	'ok',
//	'okay',
//	'positively',
//	'precisely',
//	'right on',
//	'righto',
//	'sure thing',
//	'sure',
//	'surely',
//	'true',
//	'undoubtedly',
//	'unquestionably',
//	'very well',
//	'whatever',
//	'willingly',
//	'without fail',
//	'y',
//	'ya',
//	'yea',
//	'yeah',
//	'yep',
//	'yes',
//	'yessir',
//	'yup'
func IsAffirmative(s string) bool {
	switch strings.ToLower(s) {
	case "absolutely":
		return true
	case "affirmative":
		return true
	case "all right":
		return true
	case "amen":
		return true
	case "aye":
		return true
	case "beyond a doubt":
		return true
	case "by all means":
		return true
	case "certainly":
		return true
	case "definitely":
		return true
	case "even so":
		return true
	case "exactly":
		return true
	case "fine":
		return true
	case "gladly":
		return true
	case "good enough":
		return true
	case "good":
		return true
	case "granted":
		return true
	case "i accept":
		return true
	case "i concur":
		return true
	case "i guess":
		return true
	case "if you must":
		return true
	case "indubitably":
		return true
	case "just so":
		return true
	case "most assuredly":
		return true
	case "naturally":
		return true
	case "of course":
		return true
	case "ok":
		return true
	case "okay":
		return true
	case "positively":
		return true
	case "precisely":
		return true
	case "right on":
		return true
	case "righto":
		return true
	case "sure thing":
		return true
	case "sure":
		return true
	case "surely":
		return true
	case "true":
		return true
	case "undoubtedly":
		return true
	case "unquestionably":
		return true
	case "very well":
		return true
	case "whatever":
		return true
	case "willingly":
		return true
	case "without fail":
		return true
	case "y":
		return true
	case "ya":
		return true
	case "yea":
		return true
	case "yeah":
		return true
	case "yep":
		return true
	case "yes":
		return true
	case "yessir":
		return true
	case "yup":
		return true
	}

	return false
}
