func romanToInt(s string) int {
    SYMBOL_TO_INT := map[byte]int {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }


    answer := 0
    for i := 0; i < len(s) - 1; i = i + 1 {
        cur := SYMBOL_TO_INT[s[i]]
        next := SYMBOL_TO_INT[s[i + 1]]
        if cur < next {
            answer -= cur
        } else {
            answer += cur
        }
    }
    answer += SYMBOL_TO_INT[s[len(s) - 1]]
    return answer

}

/*
byte, string, rune
string is a read-only slice of byte array
rune is a code code point, which is an alias for the int32, asw well as character constant
declare byte, rune with single quote, string with double quote

https://go.dev/blog/strings
*/