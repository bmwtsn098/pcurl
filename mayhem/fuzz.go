package fuzz

import "strconv"
import "github.com/antlabs/pcurl"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 1 {
        num, _ = strconv.Atoi(string(bytes[0]))

        switch num {
    
        case 0:
            content := string(bytes)
            pcurl.GetArgsToken(content)
            return 0

        case 1:
            content := string(bytes)
            pcurl.ParseString(content)
            return 0

        default:
            var strArr = make([]string, len(bytes))
            for i, byte := range bytes {

                strArr[i] = string(byte)
            }
            pcurl.ParseSlice(strArr)
            return 0

        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}