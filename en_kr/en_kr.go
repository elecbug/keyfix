package en_kr

import (
	"fmt"
	"strings"
	"unicode"
)

// RawTypo represents a raw string for typo conversion between English and Korean.
type RawTypo struct {
	raw         string
	onCapslocks bool
}

// NewRawTypo creates a new RawTypo instance with the given raw string and Caps Lock handling option.
func NewRawTypo(raw string, onCapslocks bool) *RawTypo {
	return &RawTypo{
		raw:         raw,
		onCapslocks: onCapslocks,
	}
}

// Convert converts the raw string between English and Korean based on its content.
func (r *RawTypo) Convert() (string, error) {
	raw := r.raw

	if isEnglish(raw) {
		if r.onCapslocks {
			raw = convertCapslocks(raw)
		}
		raw = removeUnaffectedShift(raw)

		var b strings.Builder
		b.Grow(len(raw) * 2)

		for _, ch := range raw {
			if val, exists := enMap[string(ch)]; exists {
				b.WriteString(val)
			} else {
				// Preserve digits/punct/spaces/etc. that are not in the mapping
				b.WriteRune(ch)
			}
		}

		result := mergeRunes(b.String())
		return result, nil

	} else if isKorean(raw) {
		result := ""

		for _, ch := range raw {
			if KR_START <= ch && ch <= KR_END {
				start, mid, end, err := getKrComponents(ch)

				if err != nil {
					return r.raw, fmt.Errorf("failed to decompose Hangul syllable %c: %v", ch, err)
				}

				// fmt.Printf("Decomposed %c to %c %c %c", ch, start, mid, end)

				result += string(start) + string(mid)
				if end != 0 {
					result += string(end)
				}
			} else {
				// Preserve digits/punct/spaces/etc. that are not in the mapping
				result += string(ch)
			}
		}

		result = splitMergedRunes(result)

		for old, new := range krMap {
			result = strings.ReplaceAll(result, old, new)
		}

		if r.onCapslocks {
			result = convertCapslocks(result)
		}

		return result, nil
	}

	return r.raw, fmt.Errorf("input contains non-convertible characters")
}

// convertCapslocks converts a string typed with Caps Lock on to its correct case.
func convertCapslocks(s string) string {
	result := ""

	for _, ch := range s {
		if unicode.IsUpper(ch) {
			lowerCh := unicode.ToLower(ch)
			result += string(lowerCh)
		} else if unicode.IsLower(ch) {
			upperCh := unicode.ToUpper(ch)
			result += string(upperCh)
		} else {
			result += string(ch)
		}
	}

	return result
}

// removeUnaffectedShift removes the effect of Shift key on characters that are not affected by it.
func removeUnaffectedShift(s string) string {
	result := ""

	for _, ch := range s {
		if unicode.IsUpper(ch) {
			if _, exists := enMap[string(ch)]; !exists {
				result += string(unicode.ToLower(ch))
				continue
			}
		}

		result += string(ch)
	}

	return result
}

// isEnglish checks if the given string consists solely of English alphabetic characters.
func isEnglish(s string) bool {
	for _, ch := range s {
		if !unicode.IsLower(ch) && !unicode.IsUpper(ch) && !unicode.IsDigit(ch) && !unicode.IsSpace(ch) && !unicode.IsPunct(ch) && !unicode.IsSymbol(ch) {
			return false
		}
	}

	return true
}

// isKorean checks if the given string consists solely of Korean Hangul characters.
func isKorean(s string) bool {
	for _, ch := range s {
		if !unicode.Is(unicode.Hangul, ch) && !unicode.IsDigit(ch) && !unicode.IsSpace(ch) && !unicode.IsPunct(ch) && !unicode.IsSymbol(ch) {
			return false
		}
	}

	return true
}

// getKrUnicode constructs a Korean Hangul syllable from its components.
func getKrUnicode(start, mid, end rune) rune {
	return 0xAC00 + (start * 21 * 28) + (mid * 28) + end
}

// getKrComponents decomposes a Korean Hangul syllable into its components.
func getKrComponents(r rune) (rune, rune, rune, error) {
	syllableIndex := r - 0xAC00
	start := syllableIndex/(21*28) + choStart
	mid := (syllableIndex%(21*28))/28 + jungStart
	end := syllableIndex%28 + jongStart - 1

	if end < jongStart {
		end = 0 // No final consonant
	}

	if start < choStart || start > choEnd || mid < jungStart || mid > jungEnd || (end != 0 && (end < jongStart || end > jongEnd)) {
		return 0, 0, 0, fmt.Errorf("invalid Hangul syllable: %c", r)
	}

	return choToCompat[start], jungToCompat[mid], jongToCompat[end], nil
}

// mergeRunes merges individual Korean Jamo runes into complete Hangul syllables.
func mergeRunes(s string) string {
	toChoIndex := func(r rune) rune {
		if r >= choStart && r <= choEnd {
			return r - choStart
		}
		if mapped, exists := compatToCho[r]; exists {
			return mapped - choStart
		}
		return -1
	}

	toJungIndex := func(r rune) rune {
		if r >= jungStart && r <= jungEnd {
			return r - jungStart
		}
		if mapped, exists := compatToJung[r]; exists {
			return mapped - jungStart
		}
		return -1
	}

	// Returns jong index in [1..27], 0 means "no final", -1 means "not a jong candidate"
	toJongIndex := func(r rune) rune {
		if r >= jongStart && r <= jongEnd {
			return r - jongStart + 1
		}
		if mapped, exists := compatToJong[r]; exists {
			return mapped - jongStart + 1
		}
		return -1
	}

	// Combine two jung indices into a single compound jung index.
	// Indices follow the standard Jungseong order:
	// ㅏ0 ㅐ1 ㅑ2 ㅒ3 ㅓ4 ㅔ5 ㅕ6 ㅖ7 ㅗ8 ㅘ9 ㅙ10 ㅚ11 ㅛ12 ㅜ13 ㅝ14 ㅞ15 ㅟ16 ㅠ17 ㅡ18 ㅢ19 ㅣ20
	combineJung := func(a, b rune) (rune, bool) {
		switch a {
		case 8: // ㅗ
			switch b {
			case 0: // ㅏ
				return 9, true // ㅘ
			case 1: // ㅐ
				return 10, true // ㅙ
			case 20: // ㅣ
				return 11, true // ㅚ
			}
		case 13: // ㅜ
			switch b {
			case 4: // ㅓ
				return 14, true // ㅝ
			case 5: // ㅔ
				return 15, true // ㅞ
			case 20: // ㅣ
				return 16, true // ㅟ
			}
		case 18: // ㅡ
			if b == 20 { // ㅣ
				return 19, true // ㅢ
			}
		}
		return -1, false
	}

	// Combine two jong indices into a single compound jong index.
	// jong indices are [1..27] with:
	// ㄱ1 ㄲ2 ㄳ3 ㄴ4 ㄵ5 ㄶ6 ㄷ7 ㄹ8 ㄺ9 ㄻ10 ㄼ11 ㄽ12 ㄾ13 ㄿ14 ㅀ15 ㅁ16 ㅂ17 ㅄ18 ㅅ19 ㅆ20 ㅇ21 ㅈ22 ㅊ23 ㅋ24 ㅌ25 ㅍ26 ㅎ27
	combineJong := func(a, b rune) (rune, bool) {
		switch a {
		case 1: // ㄱ
			if b == 19 { // ㅅ
				return 3, true // ㄳ
			}
		case 4: // ㄴ
			if b == 22 { // ㅈ
				return 5, true // ㄵ
			}
			if b == 27 { // ㅎ
				return 6, true // ㄶ
			}
		case 8: // ㄹ
			switch b {
			case 1: // ㄱ
				return 9, true // ㄺ
			case 16: // ㅁ
				return 10, true // ㄻ
			case 17: // ㅂ
				return 11, true // ㄼ
			case 19: // ㅅ
				return 12, true // ㄽ
			case 25: // ㅌ
				return 13, true // ㄾ
			case 26: // ㅍ
				return 14, true // ㄿ
			case 27: // ㅎ
				return 15, true // ㅀ
			}
		case 17: // ㅂ
			if b == 19 { // ㅅ
				return 18, true // ㅄ
			}
		}
		return -1, false
	}

	// Read a jung (possibly compound) starting at i. Returns (jungIdx, consumedCount, ok).
	readJung := func(runes []rune, i int) (rune, int, bool) {
		if i >= len(runes) {
			return -1, 0, false
		}
		j1 := toJungIndex(runes[i])
		if j1 < 0 {
			return -1, 0, false
		}
		// Try to combine with next jung
		if i+1 < len(runes) {
			j2 := toJungIndex(runes[i+1])
			if j2 >= 0 {
				if jc, ok := combineJung(j1, j2); ok {
					return jc, 2, true
				}
			}
		}
		return j1, 1, true
	}

	// Read a jong (possibly compound) starting at i, but do not steal the next syllable's 초성.
	// If a jong candidate is followed by a jung, we treat it as next syllable's 초성.
	readJong := func(runes []rune, i int) (rune, int, bool) {
		if i >= len(runes) {
			return 0, 0, false
		}
		j1 := toJongIndex(runes[i])
		if j1 < 0 {
			return 0, 0, false
		}
		// If next is a jung, this consonant should start the next syllable.
		if i+1 < len(runes) && toJungIndex(runes[i+1]) >= 0 {
			return 0, 0, false
		}

		// Try compound jong with the next consonant, again ensuring we don't steal next jung.
		if i+1 < len(runes) {
			j2 := toJongIndex(runes[i+1])
			if j2 >= 0 {
				// If consonant2 is followed by jung, consonant2 should be next syllable's 초성.
				if i+2 < len(runes) && toJungIndex(runes[i+2]) >= 0 {
					return j1, 1, true
				}
				if jc, ok := combineJong(j1, j2); ok {
					return jc, 2, true
				}
			}
		}
		return j1, 1, true
	}

	runes := []rune(s)

	var b strings.Builder
	b.Grow(len(runes) * 2)

	i := 0
	for i < len(runes) {
		r := runes[i]

		// Hard boundary: any non-Hangul rune is preserved as-is.
		// Digits/punct/spaces will always break composition.
		if !unicode.Is(unicode.Hangul, r) {
			b.WriteRune(r)
			i++
			continue
		}

		// If it is not a chosung candidate, just emit as-is.
		choIdx := toChoIndex(r)
		if choIdx < 0 {
			b.WriteRune(r)
			i++
			continue
		}

		// Need at least a jung to form a syllable.
		jungIdx, jungConsumed, ok := readJung(runes, i+1)
		if !ok {
			b.WriteRune(r)
			i++
			continue
		}

		// Optional jong
		jongIdx, jongConsumed, _ := readJong(runes, i+1+jungConsumed)

		syllable := getKrUnicode(choIdx, jungIdx, jongIdx)
		b.WriteRune(syllable)

		i += 1 + jungConsumed + jongConsumed
	}

	return b.String()
}

// splitMergedRunes splits compound Jamo runes into their individual components.
func splitMergedRunes(s string) string {
	for _, ch := range s {
		switch ch {
		case 'ㅘ':
			s = strings.ReplaceAll(s, "ㅘ", "ㅗㅏ")
		case 'ㅙ':
			s = strings.ReplaceAll(s, "ㅙ", "ㅗㅐ")
		case 'ㅚ':
			s = strings.ReplaceAll(s, "ㅚ", "ㅗㅣ")
		case 'ㅝ':
			s = strings.ReplaceAll(s, "ㅝ", "ㅜㅓ")
		case 'ㅞ':
			s = strings.ReplaceAll(s, "ㅞ", "ㅜㅔ")
		case 'ㅟ':
			s = strings.ReplaceAll(s, "ㅟ", "ㅜㅣ")
		case 'ㅢ':
			s = strings.ReplaceAll(s, "ㅢ", "ㅡㅣ")
		case 'ㄳ':
			s = strings.ReplaceAll(s, "ㄳ", "ㄱㅅ")
		case 'ㄵ':
			s = strings.ReplaceAll(s, "ㄵ", "ㄴㅈ")
		case 'ㄶ':
			s = strings.ReplaceAll(s, "ㄶ", "ㄴㅎ")
		case 'ㄺ':
			s = strings.ReplaceAll(s, "ㄺ", "ㄹㄱ")
		case 'ㄻ':
			s = strings.ReplaceAll(s, "ㄻ", "ㄹㅁ")
		case 'ㄼ':
			s = strings.ReplaceAll(s, "ㄼ", "ㄹㅂ")
		case 'ㄽ':
			s = strings.ReplaceAll(s, "ㄽ", "ㄹㅅ")
		case 'ㄾ':
			s = strings.ReplaceAll(s, "ㄾ", "ㄹㅌ")
		case 'ㄿ':
			s = strings.ReplaceAll(s, "ㄿ", "ㄹㅍ")
		case 'ㅀ':
			s = strings.ReplaceAll(s, "ㅀ", "ㄹㅎ")
		case 'ㅄ':
			s = strings.ReplaceAll(s, "ㅄ", "ㅂㅅ")
		default:
			// Do nothing
		}
	}

	return s
}
