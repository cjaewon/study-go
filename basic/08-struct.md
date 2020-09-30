# 구조체
구조화 된 데이터를 처리할 때 struct를 사용하는데 이를 구조체라고 한다

```go
// 예시
type Person struct {
  name string
  age int
}
```

## 메소드
go 는 c언어랑 다르게 구조체가 속성 뿐만 아니라 기능 또한 포함 할 수 있다.  

```go
type Student struct {
  name string
  class int

  score Score
}

type Score struct {
  name int
  grade int
}

func (s Student) ViewScore() {
  fmt.Println(s.score)
}

func (s Student) InputScore(name, grade string) {
  s.score.name = name
  s.score.grade = grade

  // s.score 의 name, grade 는 변하지 않는다.
  // go 는 기본적으로 복사로 이루어지기 때문에 복사된 s 에 score 를 대입한 것이다.
}

```