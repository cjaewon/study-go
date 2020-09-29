# 🗂 배열과 문자열
배열은 같은 타입의 변수들로 이루어진 유한집합을 의미한다

## 배열
```go
// var [이름] [[개수]][타입] 
// 사이즈 : 10 * 4 = 40byte
var a [10]int32
b := [10]int32{} 

for i := 0; i < 10; i+= { // 10 을 len(a) 라고 해도 됨
  fmt.Println(a[i]) // a[i] => 0
}
```

## 문자열
```go
hw := "Hello World"

for i := 0; i < len(hw); i++ {
  fmt.Print(hw[i], string(hw[i])) // => 아스키 코드, 문자 
}

hangul := "한글" // 한글은 영어와 달리 1byte 가 아님
fmt.Println(len(hangul)) // 4byte
// 따라서 위에 for 문으로 순회하면 이상한 글자가 나오는데
// 그 이유는 for문이 바이트로 반복하기 때문이다.

// 이를 방지하기 위해 []rune 타입으로 선언하면 된다.
```

```go
s := "Hello 월드"
s2 := []rune(s)

for i := 0; i < len(s); i++ {
  fmt.Print(string(s2[i])) // rune 을 사용해야 정상적으로 출력
}

for t := range s {
  fmt.Print(string(t)) // range 사용 시 rune 을 안 사용하고도 정상적으로 출력
}
```

