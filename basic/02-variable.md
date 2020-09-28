# 🎲 변수
값이 변하거나 저장할 수 있는 메모리상의 공간을 의미한다.

```go
var a int
// var [이름] [타입]
var b int = 10
// 기본적 0으로 초기화됨
// 10을 대입했기 때문에 10으로 초기화 됨

var c = 3.14 // 타입을 생략해도 자동으로 타입을 유추해서 변수를 생성한다.

d := 10
// 또는 [이름] := [값] 을 하면
// 자동으로 타입을 유추한 후 변수를 생성한다.
```

변수를 생성하면 그에 맞는 메모리가 활당되고 
해당 변수를 읽으면 해당 주소로 부터 타입 사이즈에 맞게 값들을 읽어오게된다.

## 타입
```go
var a int // 32bit -> 4byte, 64bit -> 8byte
var b int8 // 1byte
var c int16 // 2byte
var d int32 // 4byte
var e int64 // 8byte

var ue uint8 // 0 ~ 255, 256 이 아닌 이유는 음수, 양수 구분을 위해서이다.

// int, float + [숫자] 에서 숫자는 bit 라고 보면 된다.
// 따라서 int8 은 8bit 이므로 = 1byte

var f float32 // 4byte, 소수점 이하 7개 까지 가능
var g float64 // 8byte, 소수점 이하 15개 까지 가능

var h bool // true / false, 아마도 1byte 추정

var i string // 사이즈가 정해져있지 않고 가변적임, 글자 수 * 1 ~ 3byte
var j rune // 1 ~ 3byte, string -> []rune 으로 char 과 비슷함 
```
