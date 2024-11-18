# 문자열 String

## 문자열 구조

### string 구조
```
type StringHeader struct {
    Data    uintptr     // 문자열의 데이터가 있는 메모리 주소를 나타내는 일종의 포인터
    Len     int         // 문자열의 길이
}
```
- string 타입 -> Go 언어에서 제공하는 내장 타입
- StringHeader 구조체로 강제 타입 변환을 해 내부 구현을 확인 가능
- 문자열을 복사할 때 실제 데이터를 복사하지 않고, 같은 데이터를 참조


string 합 연산이 빈번하면 strings.Builder 사용