package en_kr

const KR_START = '가'
const KR_END = '힣'

// Mapping tables for compatibility Jamo to standard Jamo conversion.
var compatToCho = map[rune]rune{
	'ㄱ': 0x1100, 'ㄲ': 0x1101, 'ㄴ': 0x1102, 'ㄷ': 0x1103, 'ㄸ': 0x1104,
	'ㄹ': 0x1105, 'ㅁ': 0x1106, 'ㅂ': 0x1107, 'ㅃ': 0x1108, 'ㅅ': 0x1109,
	'ㅆ': 0x110A, 'ㅇ': 0x110B, 'ㅈ': 0x110C, 'ㅉ': 0x110D, 'ㅊ': 0x110E,
	'ㅋ': 0x110F, 'ㅌ': 0x1110, 'ㅍ': 0x1111, 'ㅎ': 0x1112,
}

// Mapping tables for compatibility Jamo to standard Jamo conversion.
var compatToJung = map[rune]rune{
	'ㅏ': 0x1161, 'ㅐ': 0x1162, 'ㅑ': 0x1163, 'ㅒ': 0x1164, 'ㅓ': 0x1165,
	'ㅔ': 0x1166, 'ㅕ': 0x1167, 'ㅖ': 0x1168, 'ㅗ': 0x1169, 'ㅘ': 0x116A,
	'ㅙ': 0x116B, 'ㅚ': 0x116C, 'ㅛ': 0x116D, 'ㅜ': 0x116E, 'ㅝ': 0x116F,
	'ㅞ': 0x1170, 'ㅟ': 0x1171, 'ㅠ': 0x1172, 'ㅡ': 0x1173, 'ㅢ': 0x1174,
	'ㅣ': 0x1175,
}

// Mapping tables for compatibility Jamo to standard Jamo conversion.
var compatToJong = map[rune]rune{
	'ㄱ': 0x11A8, 'ㄲ': 0x11A9, 'ㄳ': 0x11AA, 'ㄴ': 0x11AB, 'ㄵ': 0x11AC,
	'ㄶ': 0x11AD, 'ㄷ': 0x11AE, 'ㄹ': 0x11AF, 'ㄺ': 0x11B0, 'ㄻ': 0x11B1,
	'ㄼ': 0x11B2, 'ㄽ': 0x11B3, 'ㄾ': 0x11B4, 'ㄿ': 0x11B5, 'ㅀ': 0x11B6,
	'ㅁ': 0x11B7, 'ㅂ': 0x11B8, 'ㅄ': 0x11B9, 'ㅅ': 0x11BA, 'ㅆ': 0x11BB,
	'ㅇ': 0x11BC, 'ㅈ': 0x11BD, 'ㅊ': 0x11BE, 'ㅋ': 0x11BF, 'ㅌ': 0x11C0,
	'ㅍ': 0x11C1, 'ㅎ': 0x11C2,
}

// Mapping tables for standard Jamo to compatibility Jamo conversion.
var choToCompat = map[rune]rune{
	0x1100: 'ㄱ', 0x1101: 'ㄲ', 0x1102: 'ㄴ', 0x1103: 'ㄷ', 0x1104: 'ㄸ',
	0x1105: 'ㄹ', 0x1106: 'ㅁ', 0x1107: 'ㅂ', 0x1108: 'ㅃ', 0x1109: 'ㅅ',
	0x110A: 'ㅆ', 0x110B: 'ㅇ', 0x110C: 'ㅈ', 0x110D: 'ㅉ', 0x110E: 'ㅊ',
	0x110F: 'ㅋ', 0x1110: 'ㅌ', 0x1111: 'ㅍ', 0x1112: 'ㅎ',
}

// Mapping tables for standard Jamo to compatibility Jamo conversion.
var jungToCompat = map[rune]rune{
	0x1161: 'ㅏ', 0x1162: 'ㅐ', 0x1163: 'ㅑ', 0x1164: 'ㅒ', 0x1165: 'ㅓ',
	0x1166: 'ㅔ', 0x1167: 'ㅕ', 0x1168: 'ㅖ', 0x1169: 'ㅗ', 0x116A: 'ㅘ',
	0x116B: 'ㅙ', 0x116C: 'ㅚ', 0x116D: 'ㅛ', 0x116E: 'ㅜ', 0x116F: 'ㅝ',
	0x1170: 'ㅞ', 0x1171: 'ㅟ', 0x1172: 'ㅠ', 0x1173: 'ㅡ', 0x1174: 'ㅢ',
	0x1175: 'ㅣ',
}

// Mapping tables for standard Jamo to compatibility Jamo conversion.
var jongToCompat = map[rune]rune{
	0x11A8: 'ㄱ', 0x11A9: 'ㄲ', 0x11AA: 'ㄳ', 0x11AB: 'ㄴ', 0x11AC: 'ㄵ',
	0x11AD: 'ㄶ', 0x11AE: 'ㄷ', 0x11AF: 'ㄹ', 0x11B0: 'ㄺ', 0x11B1: 'ㄻ',
	0x11B2: 'ㄼ', 0x11B3: 'ㄽ', 0x11B4: 'ㄾ', 0x11B5: 'ㄿ', 0x11B6: 'ㅀ',
	0x11B7: 'ㅁ', 0x11B8: 'ㅂ', 0x11B9: 'ㅄ', 0x11BA: 'ㅅ', 0x11BB: 'ㅆ',
	0x11BC: 'ㅇ', 0x11BD: 'ㅈ', 0x11BE: 'ㅊ', 0x11BF: 'ㅋ', 0x11C0: 'ㅌ',
	0x11C1: 'ㅍ', 0x11C2: 'ㅎ',
}

// Jamo Unicode ranges
const (
	// 초성 (initial consonants): ㄱ-ㅎ
	choStart = 0x1100
	choEnd   = 0x1112
	// 중성 (vowels): ㅏ-ㅣ
	jungStart = 0x1161
	jungEnd   = 0x1175
	// 종성 (final consonants): ㄱ-ㅎ
	jongStart = 0x11A8
	jongEnd   = 0x11C2
)

// Mapping tables for English to Korean and Korean to English character conversion.
var enMap = map[string]string{
	"a": "ㅁ",
	"b": "ㅠ",
	"c": "ㅊ",
	"d": "ㅇ",
	"e": "ㄷ",
	"f": "ㄹ",
	"g": "ㅎ",
	"h": "ㅗ",
	"i": "ㅑ",
	"j": "ㅓ",
	"k": "ㅏ",
	"l": "ㅣ",
	"m": "ㅡ",
	"n": "ㅜ",
	"o": "ㅐ",
	"p": "ㅔ",
	"q": "ㅂ",
	"r": "ㄱ",
	"s": "ㄴ",
	"t": "ㅅ",
	"u": "ㅕ",
	"v": "ㅍ",
	"w": "ㅈ",
	"x": "ㅌ",
	"y": "ㅛ",
	"z": "ㅋ",
	"R": "ㄲ",
	"E": "ㄸ",
	"Q": "ㅃ",
	"T": "ㅆ",
	"W": "ㅉ",
	"O": "ㅒ",
	"P": "ㅖ",
}

// Mapping table for Korean to English character conversion.
var krMap = map[string]string{
	"ㅁ": "a",
	"ㅠ": "b",
	"ㅊ": "c",
	"ㅇ": "d",
	"ㄷ": "e",
	"ㄹ": "f",
	"ㅎ": "g",
	"ㅗ": "h",
	"ㅑ": "i",
	"ㅓ": "j",
	"ㅏ": "k",
	"ㅣ": "l",
	"ㅡ": "m",
	"ㅜ": "n",
	"ㅐ": "o",
	"ㅔ": "p",
	"ㅂ": "q",
	"ㄱ": "r",
	"ㄴ": "s",
	"ㅅ": "t",
	"ㅕ": "u",
	"ㅍ": "v",
	"ㅈ": "w",
	"ㅌ": "x",
	"ㅛ": "y",
	"ㅋ": "z",
	"ㄲ": "R",
	"ㄸ": "E",
	"ㅃ": "Q",
	"ㅆ": "T",
	"ㅉ": "W",
	"ㅒ": "O",
	"ㅖ": "P",
}
