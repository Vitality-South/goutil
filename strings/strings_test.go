package strings

import (
	"testing"
	"unicode/utf8"
)

// testMask tests one of the masking functions such as MaskLeft or MaskRight to ensure correctness
func testMask(t *testing.T, f func(rune, int, string) string, r rune, s, maskedAll, keep1, keep2, keep4, keep8, keep16, keep32, lenRuneLess1, lenRune, lenRunePlus1 string, checkRuneCount bool) {
	if q := f(r, -16, s); (q != maskedAll) || ((utf8.RuneCountInString(q) != utf8.RuneCountInString(s)) && checkRuneCount) {
		t.Errorf("expected [%s], got [%s]", maskedAll, q)
	}

	if q := f(r, -4, s); (q != maskedAll) || ((utf8.RuneCountInString(q) != utf8.RuneCountInString(s)) && checkRuneCount) {
		t.Errorf("expected [%s], got [%s]", maskedAll, q)
	}

	if q := f(r, -1, s); (q != maskedAll) || ((utf8.RuneCountInString(q) != utf8.RuneCountInString(s)) && checkRuneCount) {
		t.Errorf("expected [%s], got [%s]", maskedAll, q)
	}

	if q := f(r, 0, s); (q != maskedAll) || ((utf8.RuneCountInString(q) != utf8.RuneCountInString(s)) && checkRuneCount) {
		t.Errorf("expected [%s], got [%s]", maskedAll, q)
	}

	if q := f(r, 1, s); (q != keep1) || ((utf8.RuneCountInString(q) != utf8.RuneCountInString(s)) && checkRuneCount) {
		t.Errorf("expected [%s], got [%s]", keep1, q)
	}

	if q := f(r, 2, s); (q != keep2) || ((utf8.RuneCountInString(q) != utf8.RuneCountInString(s)) && checkRuneCount) {
		t.Errorf("expected [%s], got [%s]", keep2, q)
	}

	if q := f(r, 4, s); (q != keep4) || ((utf8.RuneCountInString(q) != utf8.RuneCountInString(s)) && checkRuneCount) {
		t.Errorf("expected [%s], got [%s]", keep4, q)
	}

	if q := f(r, 8, s); (q != keep8) || ((utf8.RuneCountInString(q) != utf8.RuneCountInString(s)) && checkRuneCount) {
		t.Errorf("expected [%s], got [%s]", keep8, q)
	}

	if q := f(r, 16, s); (q != keep16) || ((utf8.RuneCountInString(q) != utf8.RuneCountInString(s)) && checkRuneCount) {
		t.Errorf("expected [%s], got [%s]", keep16, q)
	}

	if q := f(r, 32, s); (q != keep32) || ((utf8.RuneCountInString(q) != utf8.RuneCountInString(s)) && checkRuneCount) {
		t.Errorf("expected [%s], got [%s]", keep32, q)
	}

	if q := f(r, utf8.RuneCountInString(s)-1, s); (q != lenRuneLess1) || ((utf8.RuneCountInString(q) != utf8.RuneCountInString(lenRuneLess1)) && checkRuneCount) {
		t.Errorf("expected [%s], got [%s]", lenRuneLess1, q)
	}

	if q := f(r, utf8.RuneCountInString(s), s); (q != lenRune) || ((utf8.RuneCountInString(q) != utf8.RuneCountInString(lenRune)) && checkRuneCount) {
		t.Errorf("expected [%s], got [%s]", lenRune, q)
	}

	if q := f(r, utf8.RuneCountInString(s)+1, s); (q != lenRunePlus1) || ((utf8.RuneCountInString(q) != utf8.RuneCountInString(lenRunePlus1)) && checkRuneCount) {
		t.Errorf("expected [%s], got [%s]", lenRunePlus1, q)
	}
}

func TestMaskLeft(t *testing.T) {
	// test single byte rune with single byte character string
	testMask(t, MaskLeft, '*', "1234567890", "**********", "*********0", "********90", "******7890", "**34567890", "1234567890", "1234567890", "*234567890", "1234567890", "1234567890", true)

	// test multi-byte rune with single byte character string
	testMask(t, MaskLeft, '⌘', "1234567890", "⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘", "⌘⌘⌘⌘⌘⌘⌘⌘⌘0", "⌘⌘⌘⌘⌘⌘⌘⌘90", "⌘⌘⌘⌘⌘⌘7890", "⌘⌘34567890", "1234567890", "1234567890", "⌘234567890", "1234567890", "1234567890", true)

	// test single byte rune and multi-byte character string
	testMask(t, MaskLeft, '*', "⌘日本語世界⌘日本語", "**********", "*********語", "********本語", "******⌘日本語", "**本語世界⌘日本語", "⌘日本語世界⌘日本語", "⌘日本語世界⌘日本語", "*日本語世界⌘日本語", "⌘日本語世界⌘日本語", "⌘日本語世界⌘日本語", true)

	// test multi-byte rune and multi-byte character string
	testMask(t, MaskLeft, '⌗', "⌘日本語世界⌘日本語", "⌗⌗⌗⌗⌗⌗⌗⌗⌗⌗", "⌗⌗⌗⌗⌗⌗⌗⌗⌗語", "⌗⌗⌗⌗⌗⌗⌗⌗本語", "⌗⌗⌗⌗⌗⌗⌘日本語", "⌗⌗本語世界⌘日本語", "⌘日本語世界⌘日本語", "⌘日本語世界⌘日本語", "⌗日本語世界⌘日本語", "⌘日本語世界⌘日本語", "⌘日本語世界⌘日本語", true)

	// test zero value rune and multi-byte character string
	var z rune
	testMask(t, MaskLeft, z, "⌘日本語世界⌘日本語", string([]rune{z, z, z, z, z, z, z, z, z, z}), string([]rune{z, z, z, z, z, z, z, z, z, '語'}), string([]rune{z, z, z, z, z, z, z, z, '本', '語'}), string([]rune{z, z, z, z, z, z, '⌘', '日', '本', '語'}), string([]rune{z, z, '本', '語', '世', '界', '⌘', '日', '本', '語'}), "⌘日本語世界⌘日本語", "⌘日本語世界⌘日本語", string([]rune{z, '日', '本', '語', '世', '界', '⌘', '日', '本', '語'}), "⌘日本語世界⌘日本語", "⌘日本語世界⌘日本語", true)

	// test zero value rune and zero value string
	var x string
	testMask(t, MaskLeft, z, x, x, x, x, x, x, x, x, x, x, x, true)

	// test non-zero value rune and zero value string
	testMask(t, MaskLeft, '⌘', x, x, x, x, x, x, x, x, x, x, x, true)

	// test invalid rune and multi-byte character string
	y := rune(0xfffffff)
	testMask(t, MaskLeft, y, "⌘日本語世界⌘日本語", string([]rune{y, y, y, y, y, y, y, y, y, y}), string([]rune{y, y, y, y, y, y, y, y, y, '語'}), string([]rune{y, y, y, y, y, y, y, y, '本', '語'}), string([]rune{y, y, y, y, y, y, '⌘', '日', '本', '語'}), string([]rune{y, y, '本', '語', '世', '界', '⌘', '日', '本', '語'}), "⌘日本語世界⌘日本語", "⌘日本語世界⌘日本語", string([]rune{y, '日', '本', '語', '世', '界', '⌘', '日', '本', '語'}), "⌘日本語世界⌘日本語", "⌘日本語世界⌘日本語", true)

	// test single byte rune with invalid UTF-8 string
	bad := []byte{32, 237, 159, 193, '\xF5', '\xF5', '\xFF', '\xD0', '\xDF', 32, 104, 101, 108, 108, 111, 32, 116, 104, 101, 114, 101, 32, 52, 50, 32, 32}
	testMask(t, MaskLeft, '*', string(bad), "**************************", "************************* ", "************************  ", "**********************42  ", "******************ere 42  ", "**********hello there 42  ", " �������� hello there 42  ", "*�������� hello there 42  ", " �������� hello there 42  ", " �������� hello there 42  ", true)

	// test multi-byte rune with invalid UTF-8 string
	testMask(t, MaskLeft, '⌘', string(bad), "⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘", "⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘ ", "⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘  ", "⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘42  ", "⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘ere 42  ", "⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘hello there 42  ", " �������� hello there 42  ", "⌘�������� hello there 42  ", " �������� hello there 42  ", " �������� hello there 42  ", true)

	// test zero value rune with invalid UTF-8 string
	testMask(t, MaskLeft, z, string(bad), string([]rune{z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z}), string([]rune{z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, ' '}), string([]rune{z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, ' ', ' '}), string([]rune{z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, '4', '2', ' ', ' '}), string([]rune{z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, 'e', 'r', 'e', ' ', '4', '2', ' ', ' '}), string([]rune{z, z, z, z, z, z, z, z, z, z, 'h', 'e', 'l', 'l', 'o', ' ', 't', 'h', 'e', 'r', 'e', ' ', '4', '2', ' ', ' '}), " �������� hello there 42  ", string([]rune{z, '�', '�', '�', '�', '�', '�', '�', '�', ' ', 'h', 'e', 'l', 'l', 'o', ' ', 't', 'h', 'e', 'r', 'e', ' ', '4', '2', ' ', ' '}), " �������� hello there 42  ", " �������� hello there 42  ", true)

	// test invalid rune with invalid UTF-8 string
	testMask(t, MaskLeft, y, string(bad), string([]rune{y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y}), string([]rune{y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, ' '}), string([]rune{y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, ' ', ' '}), string([]rune{y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, '4', '2', ' ', ' '}), string([]rune{y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, 'e', 'r', 'e', ' ', '4', '2', ' ', ' '}), string([]rune{y, y, y, y, y, y, y, y, y, y, 'h', 'e', 'l', 'l', 'o', ' ', 't', 'h', 'e', 'r', 'e', ' ', '4', '2', ' ', ' '}), " �������� hello there 42  ", string([]rune{y, '�', '�', '�', '�', '�', '�', '�', '�', ' ', 'h', 'e', 'l', 'l', 'o', ' ', 't', 'h', 'e', 'r', 'e', ' ', '4', '2', ' ', ' '}), " �������� hello there 42  ", " �������� hello there 42  ", true)
}

func TestMaskRight(t *testing.T) {
	// (t *testing.T, f func(rune, int, string) string, r rune, s, maskedAll, keep1, keep2, keep4, keep8, keep16, keep32, lenRuneLess1, lenRune, lenRunePlus1 string, checkRuneCount bool) {
	// test single byte rune with single byte character string
	testMask(t, MaskRight, '*', "1234567890", "**********", "1*********", "12********", "1234******", "12345678**", "1234567890", "1234567890", "123456789*", "1234567890", "1234567890", true)

	// test multi-byte rune with single byte character string
	testMask(t, MaskRight, '⌘', "1234567890", "⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘", "1⌘⌘⌘⌘⌘⌘⌘⌘⌘", "12⌘⌘⌘⌘⌘⌘⌘⌘", "1234⌘⌘⌘⌘⌘⌘", "12345678⌘⌘", "1234567890", "1234567890", "123456789⌘", "1234567890", "1234567890", true)

	// test single byte rune and multi-byte character string
	testMask(t, MaskRight, '*', "⌘日本語世界⌘日本語", "**********", "⌘*********", "⌘日********", "⌘日本語******", "⌘日本語世界⌘日**", "⌘日本語世界⌘日本語", "⌘日本語世界⌘日本語", "⌘日本語世界⌘日本*", "⌘日本語世界⌘日本語", "⌘日本語世界⌘日本語", true)

	// test multi-byte rune and multi-byte character string
	testMask(t, MaskRight, '⌗', "⌘日本語世界⌘日本語", "⌗⌗⌗⌗⌗⌗⌗⌗⌗⌗", "⌘⌗⌗⌗⌗⌗⌗⌗⌗⌗", "⌘日⌗⌗⌗⌗⌗⌗⌗⌗", "⌘日本語⌗⌗⌗⌗⌗⌗", "⌘日本語世界⌘日⌗⌗", "⌘日本語世界⌘日本語", "⌘日本語世界⌘日本語", "⌘日本語世界⌘日本⌗", "⌘日本語世界⌘日本語", "⌘日本語世界⌘日本語", true)

	// test zero value rune and multi-byte character string
	var z rune
	testMask(t, MaskRight, z, "⌘日本語世界⌘日本語", string([]rune{z, z, z, z, z, z, z, z, z, z}), string([]rune{'⌘', z, z, z, z, z, z, z, z, z}), string([]rune{'⌘', '日', z, z, z, z, z, z, z, z}), string([]rune{'⌘', '日', '本', '語', z, z, z, z, z, z}), string([]rune{'⌘', '日', '本', '語', '世', '界', '⌘', '日', z, z}), "⌘日本語世界⌘日本語", "⌘日本語世界⌘日本語", string([]rune{'⌘', '日', '本', '語', '世', '界', '⌘', '日', '本', z}), "⌘日本語世界⌘日本語", "⌘日本語世界⌘日本語", true)

	// test zero value rune and zero value string
	var x string
	testMask(t, MaskRight, z, x, x, x, x, x, x, x, x, x, x, x, true)

	// test non-zero value rune and zero value string
	testMask(t, MaskRight, '⌘', x, x, x, x, x, x, x, x, x, x, x, true)

	// test invalid UTF-8 rune and multi-byte character string
	y := rune(0xfffffff)
	testMask(t, MaskRight, y, "⌘日本語世界⌘日本語", string([]rune{y, y, y, y, y, y, y, y, y, y}), string([]rune{'⌘', y, y, y, y, y, y, y, y, y}), string([]rune{'⌘', '日', y, y, y, y, y, y, y, y}), string([]rune{'⌘', '日', '本', '語', y, y, y, y, y, y}), string([]rune{'⌘', '日', '本', '語', '世', '界', '⌘', '日', y, y}), "⌘日本語世界⌘日本語", "⌘日本語世界⌘日本語", string([]rune{'⌘', '日', '本', '語', '世', '界', '⌘', '日', '本', y}), "⌘日本語世界⌘日本語", "⌘日本語世界⌘日本語", true)

	// test single byte rune with invalid UTF-8 string
	bad := []byte{32, 237, 159, 193, '\xF5', '\xF5', '\xFF', '\xD0', '\xDF', 32, 104, 101, 108, 108, 111, 32, 116, 104, 101, 114, 101, 32, 52, 50, 32, 32}
	testMask(t, MaskRight, '*', string(bad), "**************************", " *************************", " �************************", " ���**********************", " �������******************", " �������� hello **********", " �������� hello there 42  ", " �������� hello there 42 *", " �������� hello there 42  ", " �������� hello there 42  ", true)

	// test multi-byte rune with invalid UTF-8 string
	testMask(t, MaskRight, '⌘', string(bad), "⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘", " ⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘", " �⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘", " ���⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘", " �������⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘", " �������� hello ⌘⌘⌘⌘⌘⌘⌘⌘⌘⌘", " �������� hello there 42  ", " �������� hello there 42 ⌘", " �������� hello there 42  ", " �������� hello there 42  ", true)

	// test zero value rune with invalid UTF-8 string
	testMask(t, MaskRight, z, string(bad), string([]rune{z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z}), string([]rune{' ', z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z}), string([]rune{' ', '�', z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z}), string([]rune{' ', '�', '�', '�', z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z}), string([]rune{' ', '�', '�', '�', '�', '�', '�', '�', z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z, z}), string([]rune{' ', '�', '�', '�', '�', '�', '�', '�', '�', ' ', 'h', 'e', 'l', 'l', 'o', ' ', z, z, z, z, z, z, z, z, z, z}), " �������� hello there 42  ", string([]rune{' ', '�', '�', '�', '�', '�', '�', '�', '�', ' ', 'h', 'e', 'l', 'l', 'o', ' ', 't', 'h', 'e', 'r', 'e', ' ', '4', '2', ' ', z}), " �������� hello there 42  ", " �������� hello there 42  ", true)

	// test invalid rune with invalid UTF-8 string
	testMask(t, MaskRight, y, string(bad), string([]rune{y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y}), string([]rune{' ', y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y}), string([]rune{' ', '�', y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y}), string([]rune{' ', '�', '�', '�', y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y}), string([]rune{' ', '�', '�', '�', '�', '�', '�', '�', y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y, y}), string([]rune{' ', '�', '�', '�', '�', '�', '�', '�', '�', ' ', 'h', 'e', 'l', 'l', 'o', ' ', y, y, y, y, y, y, y, y, y, y}), " �������� hello there 42  ", string([]rune{' ', '�', '�', '�', '�', '�', '�', '�', '�', ' ', 'h', 'e', 'l', 'l', 'o', ' ', 't', 'h', 'e', 'r', 'e', ' ', '4', '2', ' ', y}), " �������� hello there 42  ", " �������� hello there 42  ", true)
}

func TestToValidUTF8Trimmed(t *testing.T) {
	good := "hello there 42 "
	bad := []byte{32, 237, 159, 193, '\xF5', '\xF5', '\xFF', '\xD0', '\xDF', 32, 104, 101, 108, 108, 111, 32, 116, 104, 101, 114, 101, 32, 52, 50, 32, 32}
	great := "all things in life"
	awful := []byte{237, 159, 193, '\xF5', '\xF5', '\xFF', '\xD0', '\xDF', 32}
	nogood := []byte{237, 159, 193, '\xF5', '\xF5', '\xFF', '\xD0', '\xDF', 32, 32}
	terrible := []byte{237, 159, 193, '\xF5', '\xF5', '\xFF', '\xD0', '\xDF'}
	gem := []byte{32, 237, 159, 193, 52, '\xF5', '\xF5', '\xFF', 50, '\xD0', '\xDF', 32}
	secret := []byte{32, 237, 159, 193, 52, '\xF5', '\xF5', '\xFF', 32, 50, '\xD0', '\xDF', 32}

	if s := ToValidUTF8Trimmed(good); s != "hello there 42" {
		t.Errorf("expected hello there 42, got %s", s)
	}

	if s := ToValidUTF8Trimmed(string(bad)); s != "hello there 42" {
		t.Errorf("expected hello there 42, got %s", s)
	}

	if s := ToValidUTF8Trimmed(great); s != great {
		t.Errorf("expected all things in life, got %s", s)
	}

	if s := ToValidUTF8Trimmed(string(awful)); s != "" {
		t.Errorf("expected empty string, got %s", s)
	}

	if s := ToValidUTF8Trimmed(string(nogood)); s != "" {
		t.Errorf("expected empty string, got %s", s)
	}

	if s := ToValidUTF8Trimmed(string(terrible)); s != "" {
		t.Errorf("expected empty string, got %s", s)
	}

	if s := ToValidUTF8Trimmed(string(gem)); s != "42" {
		t.Errorf("expected 42, got %s", s)
	}

	if s := ToValidUTF8Trimmed(string(secret)); s != "4 2" {
		t.Errorf("expected 4 2, got %s", s)
	}
}

func TestStripHTMLTags(t *testing.T) {
	if s := StripHTMLTags(""); s != "" {
		t.Errorf("expected empty string, got %s", s)
	}

	if s := StripHTMLTags("abc"); s != "abc" {
		t.Errorf("expected abc, got %s", s)
	}

	if s := StripHTMLTags("4 > 3"); s != "4 &gt; 3" {
		t.Errorf("expected 4 &gt; 3, got %s", s)
	}

	if s := StripHTMLTags("bob<style>sam</style><script>alert('hello');</script>"); s != "bob" {
		t.Errorf("expected bob, got %s", s)
	}

	if s := StripHTMLTags("<p>BOB</p>"); s != "BOB" {
		t.Errorf("expected BOB, got %s", s)
	}
}

// testTruncate tests one of the truncation functions such as TruncateLeft or TruncateRight to ensure correctness
func testTruncate(t *testing.T, f func(string, int, string) string, length int, s, ellipsis, expected string) {
	// bound expected length to minimum zero
	l := func() int {
		if length < 0 {
			return 0
		}

		return length
	}()

	if q := f(s, length, ellipsis); (utf8.RuneCountInString(q) > l) || q != expected {
		t.Errorf("expected [%s], got [%s] with rune count %d", expected, q, utf8.RuneCountInString(q))
	}
}

func TestTruncateRight(t *testing.T) {
	testTruncate(t, TruncateRight, -16, "Hi there!", "…", "")
	testTruncate(t, TruncateRight, -1, "Hi there!", "…", "")
	testTruncate(t, TruncateRight, 0, "Hi there!", "…", "")
	testTruncate(t, TruncateRight, 1, "Hi there!", "…", "…")
	testTruncate(t, TruncateRight, 2, "Hi there!", "…", "H…")
	testTruncate(t, TruncateRight, 3, "Hi there!", "…", "Hi…")
	testTruncate(t, TruncateRight, 4, "Hi there!", "…", "Hi …")
	testTruncate(t, TruncateRight, 5, "Hi there!", "…", "Hi t…")
	testTruncate(t, TruncateRight, 6, "Hi there!", "…", "Hi th…")
	testTruncate(t, TruncateRight, 7, "Hi there!", "…", "Hi the…")
	testTruncate(t, TruncateRight, 8, "Hi there!", "…", "Hi ther…")
	testTruncate(t, TruncateRight, 9, "Hi there!", "…", "Hi there!")
	testTruncate(t, TruncateRight, 10, "Hi there!", "…", "Hi there!")
	testTruncate(t, TruncateRight, 64, "Hi there!", "…", "Hi there!")

	testTruncate(t, TruncateRight, -16, "Hi there!", "...", "")
	testTruncate(t, TruncateRight, -1, "Hi there!", "...", "")
	testTruncate(t, TruncateRight, 0, "Hi there!", "...", "")
	testTruncate(t, TruncateRight, 1, "Hi there!", "...", ".")
	testTruncate(t, TruncateRight, 2, "Hi there!", "...", "..")
	testTruncate(t, TruncateRight, 3, "Hi there!", "...", "...")
	testTruncate(t, TruncateRight, 4, "Hi there!", "...", "H...")
	testTruncate(t, TruncateRight, 5, "Hi there!", "...", "Hi...")
	testTruncate(t, TruncateRight, 6, "Hi there!", "...", "Hi ...")
	testTruncate(t, TruncateRight, 7, "Hi there!", "...", "Hi t...")
	testTruncate(t, TruncateRight, 8, "Hi there!", "...", "Hi th...")
	testTruncate(t, TruncateRight, 9, "Hi there!", "...", "Hi there!")
	testTruncate(t, TruncateRight, 10, "Hi there!", "...", "Hi there!")
	testTruncate(t, TruncateRight, 64, "Hi there!", "...", "Hi there!")

	testTruncate(t, TruncateRight, 7, "⌘日本語世界⌘日本語", "…", "⌘日本語世界…")
	testTruncate(t, TruncateRight, 7, "⌘日本語世界⌘日本語", "...", "⌘日本語...")
	testTruncate(t, TruncateRight, 7, "⌘日本語世界⌘日本語", "⌗⌗", "⌘日本語世⌗⌗")
}

func TestTruncateLeft(t *testing.T) {
	testTruncate(t, TruncateLeft, -16, "Hi there!", "…", "")
	testTruncate(t, TruncateLeft, -1, "Hi there!", "…", "")
	testTruncate(t, TruncateLeft, 0, "Hi there!", "…", "")
	testTruncate(t, TruncateLeft, 1, "Hi there!", "…", "…")
	testTruncate(t, TruncateLeft, 2, "Hi there!", "…", "…!")
	testTruncate(t, TruncateLeft, 3, "Hi there!", "…", "…e!")
	testTruncate(t, TruncateLeft, 4, "Hi there!", "…", "…re!")
	testTruncate(t, TruncateLeft, 5, "Hi there!", "…", "…ere!")
	testTruncate(t, TruncateLeft, 6, "Hi there!", "…", "…here!")
	testTruncate(t, TruncateLeft, 7, "Hi there!", "…", "…there!")
	testTruncate(t, TruncateLeft, 8, "Hi there!", "…", "… there!")
	testTruncate(t, TruncateLeft, 9, "Hi there!", "…", "Hi there!")
	testTruncate(t, TruncateLeft, 10, "Hi there!", "…", "Hi there!")
	testTruncate(t, TruncateLeft, 64, "Hi there!", "…", "Hi there!")

	testTruncate(t, TruncateLeft, -16, "Hi there!", "...", "")
	testTruncate(t, TruncateLeft, -1, "Hi there!", "...", "")
	testTruncate(t, TruncateLeft, 0, "Hi there!", "...", "")
	testTruncate(t, TruncateLeft, 1, "Hi there!", "...", ".")
	testTruncate(t, TruncateLeft, 2, "Hi there!", "...", "..")
	testTruncate(t, TruncateLeft, 3, "Hi there!", "...", "...")
	testTruncate(t, TruncateLeft, 4, "Hi there!", "...", "...!")
	testTruncate(t, TruncateLeft, 5, "Hi there!", "...", "...e!")
	testTruncate(t, TruncateLeft, 6, "Hi there!", "...", "...re!")
	testTruncate(t, TruncateLeft, 7, "Hi there!", "...", "...ere!")
	testTruncate(t, TruncateLeft, 8, "Hi there!", "...", "...here!")
	testTruncate(t, TruncateLeft, 9, "Hi there!", "...", "Hi there!")
	testTruncate(t, TruncateLeft, 10, "Hi there!", "...", "Hi there!")
	testTruncate(t, TruncateLeft, 64, "Hi there!", "...", "Hi there!")

	testTruncate(t, TruncateLeft, 7, "⌘日本語世界⌘日本語", "…", "…世界⌘日本語")
	testTruncate(t, TruncateLeft, 7, "⌘日本語世界⌘日本語", "...", "...⌘日本語")
	testTruncate(t, TruncateLeft, 7, "⌘日本語世界⌘日本語", "⌗⌗", "⌗⌗界⌘日本語")
}

func TestIsAffirmative(t *testing.T) {
	t.Parallel()

	testStructs := []struct {
		Input    string
		Expected bool
	}{
		{
			Input:    "",
			Expected: false,
		},
		{
			Input:    "Y",
			Expected: true,
		},
		{
			Input:    "yEaH",
			Expected: true,
		},
		{
			Input:    " y ",
			Expected: false,
		},
		{
			Input:    "no",
			Expected: false,
		},
		{Input: "absolutely", Expected: true},
		{Input: "affirmative", Expected: true},
		{Input: "all right", Expected: true},
		{Input: "amen", Expected: true},
		{Input: "aye", Expected: true},
		{Input: "beyond a doubt", Expected: true},
		{Input: "by all means", Expected: true},
		{Input: "certainly", Expected: true},
		{Input: "definitely", Expected: true},
		{Input: "even so", Expected: true},
		{Input: "exactly", Expected: true},
		{Input: "fine", Expected: true},
		{Input: "gladly", Expected: true},
		{Input: "good enough", Expected: true},
		{Input: "good", Expected: true},
		{Input: "granted", Expected: true},
		{Input: "i accept", Expected: true},
		{Input: "i concur", Expected: true},
		{Input: "i guess", Expected: true},
		{Input: "if you must", Expected: true},
		{Input: "indubitably", Expected: true},
		{Input: "just so", Expected: true},
		{Input: "most assuredly", Expected: true},
		{Input: "naturally", Expected: true},
		{Input: "of course", Expected: true},
		{Input: "ok", Expected: true},
		{Input: "okay", Expected: true},
		{Input: "positively", Expected: true},
		{Input: "precisely", Expected: true},
		{Input: "right on", Expected: true},
		{Input: "righto", Expected: true},
		{Input: "sure thing", Expected: true},
		{Input: "sure", Expected: true},
		{Input: "surely", Expected: true},
		{Input: "true", Expected: true},
		{Input: "undoubtedly", Expected: true},
		{Input: "unquestionably", Expected: true},
		{Input: "very well", Expected: true},
		{Input: "whatever", Expected: true},
		{Input: "willingly", Expected: true},
		{Input: "without fail", Expected: true},
		{Input: "y", Expected: true},
		{Input: "ya", Expected: true},
		{Input: "yea", Expected: true},
		{Input: "yeah", Expected: true},
		{Input: "yep", Expected: true},
		{Input: "yes", Expected: true},
		{Input: "yessir", Expected: true},
		{Input: "yup", Expected: true},
	}

	for i, testStruct := range testStructs {
		got := IsAffirmative(testStruct.Input)
		if got != testStruct.Expected {
			t.Errorf("Expected %t, got %t for %v on interation %d", testStruct.Expected, got, testStruct.Input, i)
		}
	}
}
