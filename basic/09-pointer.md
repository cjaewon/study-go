# 👉 포인터
그 변수의 메모리 주소를 가르키는 변수를 의미한다.

포인터를 함수 인자로 받으면 메모리 주소만 복사되지만
값을 함수 인자로 받으면 전체 속성이 다 복사된다.

```go
var a int = 5
var p *int = &a // 변수 앞에 & 를 붙여 주소를 가져올 수 있다.

fmt.Println(p, *p) // 0xc0000b4008, 5
                   // 포인터 변수 앞에 다시 *를 붙여주면 값을 가져온다. 
```

## 활용
```go
func (s *Student) InputScore(name, grade string) {
  s.score.name = name
  s.score.grade = grade
}
```