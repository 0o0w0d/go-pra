## Structure

### 선언

```
type 타입명 struct {
    필드명 타입
    ...
    필드명 타입
}
```

1. type 키워드를 사용해 정의
2. 타입명의 첫 번째 글자가 대문자 -> 패키지 외부로 공개
3. 중괄호 {} 안에 구조체에 속한 필드 작성 (필드 = 필드명 + 타입)

---
 
### 초기화

#### 초깃값 생략
```
var person Person
```
- 초깃값 생략 시 모든 필드가 기본값으로 초기화

#### 모든 필드 초기화
```
var person Person = Person{"zion", 27, "경기도", 160.123}
```

```
var person Person = Person{
    "zion", 
    27, 
    "경기도", 
    160.123,            // 여러 줄로 초기화 시 마지막 값 뒤 쉼표
}
```

#### 일부 필드 초기화
```
var person Person = Person{ Age: 27, Address: "경기도" }
```
- 일부 필드 값만 초기화 시 `필드명: 필드값` 형식으로 초기화
- 초기화되지 않은 나머지 변수에는 타입의 기본값이 할당

---

### Embedded Field(Embedded Struct)

#### 내장 타입처럼 포함
```
type User struct {
    Name    string
    ID      string
    Age     int
}

type VIPUser struct {
    UserInfo    User
    VIPLevel    int
    Price       int
}
```

#### 포함된 필드 방식
```
type User struct {
    Name    string
    ID      string
    Age     int
    Level   int
}

type VIPUser struct {
    User                // 필드명 생략
    Level    int
    Price    int
}
```
- 필드명이 중복될 경우, 포함된 구조체명을 쓰고 접근해야 함
- `vip.User.Level`

---

### 메모리 정렬 Memory Alignment
- 레지스터 크기 4byte -> 32비트 컴퓨터
- 레지스터 크기 8byte -> 64비트 컴퓨터
- 레지스터의 크기 -> 한 번 연산에 레지스터의 크기를 연산 => 데이터가 레지스터의 크기와 같은 크기로 정렬되어 있을 때 효율적으로 데이터를 읽어 옴

#### 메모리 패딩 Memory Padding
- 메모리 정렬을 위해서 필드 사이에 공간을 띄우는 것

---

### 결합도와 응집도
#### 결합도 Coupling
- 모듈 간 상호 의존 관계를 형성해 서로 강하게 결합되어 있는 정도

#### 응집도 Cohension
- 모듈의 완성도
- 모듈 내부의 모든 기능이 단일 목적에 맞게 모여있는지를 나타내는 정도

---

- 함수: 관련 코드 블록을 묶어 응집도를 높이고 재사용성을 증가시킴
- 배열: 같은 타입의 데이터들을 묶어 응집도를 높임
- 구조체: 관련된 데이터를 묶어 응집도를 높이고 재사용성을 증가시킴