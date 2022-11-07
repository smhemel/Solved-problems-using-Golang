func checkVowel(c rune) bool {
    switch string(c) {
    case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U" :
            return true
        default:
            return false;
    }
    return false;
}

func reverseVowels(s string) string {
    str := []rune(s);
    
    for i, j := 0, len(s)-1; i<j; {
        if checkVowel(str[i]) && checkVowel(str[j]) {
            str[i], str[j] = str[j], str[i];
            i++; j--;
        } else if !checkVowel(str[i]) {
            i++;
        } else if !checkVowel(str[j]) {
            j--;
        }
    }
    
    return string(str);
}