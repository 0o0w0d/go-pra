# Pointer

## 포인터 변수 선언
```
var p *int
```

```
var a *int
var p *int

p = &a          // a의 메모리 주소를 포인터 변수 p에 대입
```

## 포인터 기본값 : nil
```
var p *int
if p != nil {
    // p가 nil이 아닌 경우: p가 유효한 메모리 주소를 가리 킴
}
```
- 포인터 변수의 값을 초기화하지 않으면 기본값은 nil
- 유효하지 않은 메모리 주소 값 -> 어떤 메모리 공간도 가리키고 있지 않음

## 구조체를 생성해 포인터 변수 초기화
- 기존 방식
    ```
    var data Data
    var p *Data = &data         // data 변수의 주소를 반환
    ```

- 구조체를 생성해 초기화하는 방식
    ```
    var p *Data = &Data{}       // Data 구조체를 만들어 주소를 반환
    ```

## 인스턴스 Instance
- 메모리에 할당된 데이터의 실체

## 인스턴스 생성
```
var data Data
var p *Data = &data
```
- 포인터 변수 p가 기존에 있던 data 인스턴스를 가리킴
- 만들어진 Data 인스턴스 개수 : 1개

```
var p *Data = &Data{}
```
- Data 인스턴스가 하나 만들어지고 포인터 변수 p가 해당 인스턴스를 가리킴

```
var p1 *Data = &Data{}
var p2 *Data = p1
var p3 *Data = p1
```
- Data 인스턴스를 하나 만들고 포인터 변수 p1, p2, p3가 가리킴
- 가리키는 포인터 변수 개수 3개 != 만들어진 Data 인스턴스 개수 1개

```
var data1 Data
var data2 Data = data1
var data3 Data = data1
```
- data1, data2, data3 -> 모두 인스턴스
- data1 값이 data2, data3에 복사된 복사본이 생성 => 값만 같을 뿐 서로 다른 인스턴스

## new() 내장 함수
```
p1 := &Data{}           // &를 사용하는 초기화
var p2 = new(Data)      // new()를 사용하는 초기화
```
- new를 이용해서 내부 필드값을 원하는 값으로 초기화할 수는 없음