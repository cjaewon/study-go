# 🔄 반복문
golang 은 다른언어와 다르게 `for` 만 존재

```go
	// for [전처리문]; [조건문]; [후처리문] {}
	for i := 0; i < 10; i++ {
    // i 는 스코프 안에서만 접근 가능
	}

	// for {} == while {}
	for {

  }
  
  // break 로 for 문 종료
  // continue 로 반복문 안에서 다음 이동
```

## 예시
### 구구단
```go
  for i := 0; i < 10; i++ {
    for j := 0; j < 10; j++ {
      fmt.Printf("%d * %d = %d\n", i, i, i*j)
    }
    fmt.Println()
  }
```

### 별 찍기
```go
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
```

### 리버스 별 찍기
```go
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
```