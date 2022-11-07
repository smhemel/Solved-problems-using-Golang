import "strconv";

func replaceAt(s string, i int, c rune) string {
    r := []rune(s)
    r[i] = c
    return string(r)
}

func maximum69Number (num int) int {
    str := strconv.Itoa(num);
    
    for i := 0; i < len(str); i++ {
        if string(str[i]) == "6" {
            str = replaceAt(str, i, '9');
            break;
        }
    }
    ans, err := strconv.Atoi(str);
    
    if (err != nil) {
        return num;
    }
    
    return ans;
}