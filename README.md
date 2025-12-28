# Keyfix
> Keyboard layout–based typo recovery between Korean and English input

한·영 키보드 레이아웃 오입력으로 발생한 텍스트를 복구하는 Go 라이브러리입니다.

[`keyfix`](https://github.com/elecbug/keyfix)는 사용자가 의도한 언어가 아닌,
잘못된 키보드 레이아웃(QWERTY ↔ 두벌식) 상태에서 입력된 문자열을 분석하여
올바른 텍스트로 변환합니다.

## 주요 기능

- 영문 키보드로 잘못 입력된 한글 문장 복구
- 한글 입력 상태에서 잘못 입력된 영문 문자열 복구 (역전환 지원)
- Caps Lock / Shift 오입력 보정
- 완성형 한글 ↔ 자모 기반 변환 처리
- 숫자, 공백, 특수문자 보존
- 규칙 기반 처리 (AI / 사전 의존 없음)

## 사용 예시
### 영문 → 한글 (키보드 레이아웃 오입력)

```go
typo := kr.NewRawTypo("dkssudgktpdy", false)
result, _ := typo.Convert()

fmt.Println(result)
// 출력: 안녕하세요
```

### 한글 → 영문 (역 전환)

```go
typo := kr.NewRawTypo("안녕하세요", false)
result, _ := typo.Convert()

fmt.Println(result)
// 출력: dkssudgktpdy
```

### 숫자 및 특수문자 보존 

```go
typo := kr.NewRawTypo("dkssudgktpdy123!!", false)
result, _ := typo.Convert()

fmt.Println(result)
// 출력: 안녕하세요123!!
```

## 기타

- 이 라이브러리는 번역기나 철자 교정기가 아닙니다. 사용자가 입력한 텍스트의 의미가 아니라 사용자의 입력 행위(키 입력 순서)를 기준으로 처리합니다.
- 따라서 다음과 같은 특징을 가집니다.
  - 사전, 언어 모델, 통계 기반 처리 없음
  - 항상 동일한 입력에 대해 동일한 출력
  - 오프라인 환경에서도 동작
  - IME / 입력기 레벨 유틸리티로 사용 가능

### 처리 범위

지원되는 입력 유형:
- 두벌식 한글 키보드
- QWERTY 영문 키보드
- 완성형 한글 (가–힣)
- 한글 자모 및 호환 자모
- 숫자, 공백, 기호 문자

처리하지 않는 영역:
- 철자 오류 교정
- 문법/의미 기반 변환
- 자동 언어 추론
