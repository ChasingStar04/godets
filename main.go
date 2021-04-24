package main // Java에서 package와 같은 느낌, Python에서 폴더 이름과 같은 느낌

// 사용되지 않은 module은 에러를 내뿜는다
// fmt(Formatting의 약자) -> #include <stdio.h>
import (
    "encoding/json"
    "fmt"
)

func main() { // 어떤 파일에서 main 함수를 먼저 무조건 실행한다
    // fmt라는 모듈에서 불러오는 함수잖아요
    // Println이라는게
    // 그래서 어떤 모듈에서 export되는, import가 가능한 함수들은 무조건 대문자로 시작을 합니다
    fmt.Println("Hello, World!") // Go언어는 세미콜론이 필요없는 언어

    test := map[string]int{"a":10, "b":11, "c":12} // map이라는 파이썬에서 딕셔너리처럼 기능하는 변수를 만든건데 string이 key의 type이고 int가 value의 type
    value, err := json.Marshal(test) // json으로 변환하는 함수인 Marshal에 넣어주는데

    // 모든 error를 프로그래머가 직접 제어해줘야 한다
    // error에 취약하지만 error를 잘 관리만 해주면 자유자재로 컨트롤 가능!
    if err != nil { // () 묶지 않는다. 
        // 오류가 나면 어찌저찌 한다
    }

    fmt.Println(value)
}

// Go 언어에 대한 설교
// 2009년 구글에서 개발된 언어
// 개발진: C++ 개발자, ~~~~~~~~

// Stackoverflow 2020 - Most Wanted 3위
// " - Most Loved 5위
// " - 평균 연봉 세계 3위($74k = $74000 = 82,658,000)
// 2009, 2016 올해의 언어
// 언어 순위 11위
// Rust 50, Typescript 20, Kotlin 50, dart 30
