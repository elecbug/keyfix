package en_kr_test

import (
	"testing"

	. "github.com/elecbug/keyfix/en_kr"
)

func TestEnToKr(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		capslock bool
	}{
		{"ghkdrma ekfkawnl dhs tptkddmf qlcsoek.", "황금 다람쥐 온 세상을 빛내다.", false},
		{"DKSSUDGKTPDY", "안녕하세요", true},
		{"zjavbxj vmfhrmfoald", "컴퓨터 프로그래밍", false},
		{"XPTMXM ZPDLTM", "테스트 케이스", true},
		{"gksrmfrhk duddj qusghks xptmxm", "한글과 영어 변환 테스트", false},
		{"TNTWKEH EHLFrK? 12345!", "숫자도 될까? 12345!", true},
		{"RnpfqEnfgthofgfpg", "꿻뚫쇓렣", false},

		{"Hello, World!", "ㅗ디ㅣㅐ, 째깅!", false},
		{"HELLO, WORLD!", "ㅗ디ㅣㅐ, 재깅!", true},
		{"Hello, my name is Test.", "ㅗ디ㅣㅐ, ㅡㅛ ㅜ믇 ㅑㄴ ㅆㄷㄴㅅ.", false},
	}

	for _, test := range tests {
		result, err := NewRawTypo(test.input, test.capslock).Convert()
		if err != nil {
			t.Errorf("EnToKr(%s) returned error: %v", test.input, err)
			continue
		}
		if result != test.expected {
			t.Errorf("EnToKr(%s) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

func TestKrToEn(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		capslock bool
	}{
		{"황금 다람쥐 온 세상을 빛내다.", "ghkdrma ekfkawnl dhs tptkddmf qlcsoek.", false},
		{"안녕하세요", "dkssudgktpdy", false},
		{"컴퓨터 프로그래밍", "zjavbxj vmfhrmfoald", false},
		{"테스트 케이스", "xptmxm zpdltm", false},
		{"한글과 영어 변환 테스트", "gksrmfrhk duddj qusghks xptmxm", false},
		{"숫자도 될까? 12345!", "tntwkeh ehlfRk? 12345!", false},
		{"꿻뚫쇓렣", "rNPFQeNFGTHOFGFPG", true},

		{"ㅗ디ㅣㅐ, 째깅!", "hello, World!", false},
		{"ㅗ디ㅣㅐ, 재깅!", "HELLO, WORLD!", true},
		{"ㅗ디ㅣㅐ, ㅡㅛ ㅜ믇 ㅑㄴ ㅆㄷㄴㅅ.", "hello, my name is Test.", false},
	}

	for _, test := range tests {
		result, err := NewRawTypo(test.input, test.capslock).Convert()
		if err != nil {
			t.Errorf("KrToEn(%s) returned error: %v", test.input, err)
			continue
		}
		if result != test.expected {
			t.Errorf("KrToEn(%s) = %s; expected %s", test.input, result, test.expected)
		}
	}
}
