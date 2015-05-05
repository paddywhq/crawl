package myregex

import(
	"fmt"
	"regexp"
	"strings"
)

const preposition = "a,abaft,abeam,aboard,about,above,absent,across,afore,after,against,along,alongside,amid,amidst,among,amongst,an,anenst,apropos,apud,around,as,astride,at,athwart,atop,barring,before,behind,below,beneath,beside,besides,between,beyond,but,by,chez,circa,concerning,despite,down,during,except,excluding,failing,following,for,forenenst,from,given,in,including,inside,into,like,mid,midst,minus,modulo,near,next,notwithstanding,o',of,off,on,onto,opposite,out,outside,over,pace,past,per,plus,pro,qua,regarding,round,sans,save,since,than,through,thru,throughout,thruout,till,times,to,toward,towards,under,underneath,unlike,until,unto,up,upon,versus,via,vice,vis-à-vis,with,within,without,worth"

func Parse(s string, regex string) [][]string{
	//fmt.Printf( "str:%s, reg:%s \n", s, regex )
	r, err := regexp.Compile( regex )
	if err!=nil {
		fmt.Printf( "regex error: %v\n", err )
	}
	return r.FindAllStringSubmatch( s, -1 )
}


/**
 * 这个方法是干这个事儿的
 * 输入一个字符串和一个正则表达式
 * 返回一个数据，包含用该正则表达式打散的段落（每个段落头部包含该正则表达式匹配的字符串）
 */
func FindMatchBodyMap(arr []string, regex string) []string{
	r, _ := regexp.Compile(regex)

	res := make([]string, 0, 1000)
	tmpMatch := ""
	tmpBody := make( []string, 0, 10000 )

	for _, line := range arr{
		if r.Match( []byte( line ) ){
			matchArr := Parse( line, regex )
			if len(matchArr) >0 && len(matchArr[0]) >= 2 { //match
				match := matchArr[0][1]

				if tmpMatch == "" {
					tmpMatch = match
				}else{
					tmpS := "";
					for _, h := range tmpBody{
						tmpS += h + "\n"
					}

					res = append( res, tmpS)
					tmpMatch = match
					tmpBody = tmpBody[:0]
				}
			}
		}

		tmpBody = append( tmpBody, line )
	}

	tmpS := "";
	for _, h := range tmpBody{
		tmpS += h + "\n"
	}
	res = append( res, tmpS )

	return res
}

func Exists(s string, regex string) bool{
	r, _ := regexp.MatchString( regex, s )
	return r
}

func RemoveHTMLTag(input string) string{
	re := regexp.MustCompile("<[^>]*>")
	s := re.ReplaceAllString(input, "")
	s = strings.Replace( s, "&amp;", "&", -1 )
	s = strings.Replace( s, "&nbsp;", " ", -1 )
	s = strings.Replace( s, "&lt;", "<", -1 )
	s = strings.Replace( s, "&gt;", ">", -1 )
	s = strings.Replace( s, "&quot;", "\"", -1 )
	s = strings.Replace( s, "&apos;", "'", -1 )
	s = strings.Replace( s, "&cent;", "￠", -1 )
	s = strings.Replace( s, "&pound;", "£", -1 )
	s = strings.Replace( s, "&yen;", "¥", -1 )
	s = strings.Replace( s, "&euro;", "€", -1 )
	s = strings.Replace( s, "&sect;", "§", -1 )
	s = strings.Replace( s, "&copy;", "©", -1 )
	s = strings.Replace( s, "&reg;", "®", -1 )
	s = strings.Replace( s, "&trade;", "™", -1 )
	s = strings.Replace( s, "&times;", "×", -1 )
	s = strings.Replace( s, "&divide;", "÷", -1 )
	return s
}

// func FindContentInBracket(input string) string{
// 	re := regexp.MustCompile( "[^\\(]*\\(([^\\)]*)\\)" )
// 	result := re.FindAllStringSubmatch( input, -1 )
// 	if  len(result)==1 && len(result[0]) == 2{
// 		return result[0][1]
// 	}
// 	return ""
// }

func RemoveSymbol(input string, repl string) string{
	symbolRegex := "[-!$^&*\\(\\)_+|~={}:\";'<>?,.（）～]"
	re := regexp.MustCompile(symbolRegex)
	return string(re.ReplaceAll( []byte(input), []byte(repl) ))
}

func FindEnglishAndDigit( input string ) string{
	re := regexp.MustCompile( "[^a-zA-Z0-9]*([a-zA-Z0-9 ]*)[^a-zA-Z0-9]*" )
	result := re.FindAllStringSubmatch( input, -1 )
	if  len(result)==1 && len(result[0]) >= 2{
		return result[0][1]
	}

	return ""
}

func UppercaseWord(input string) string{
	arr := strings.Split(input, "")
	newStr := ""
	for i, s:=range arr{
		if i==0{
			newStr += strings.ToUpper(s)
		}else{
			newStr += strings.ToLower(s)
		}
	}
	return newStr
}

func IsPreposition(input string) bool{
	prepArr := strings.Split( preposition, "," )
	for _, p := range prepArr {
		if p == input{
			return true
		}
	} 
	return false
}

func ReplaceAllSpace(input string, to string) string{
	re := regexp.MustCompile( "[\\s]+" )
	return string(re.ReplaceAll( []byte(input), []byte(to) ))
}

func Camel(input string) string{
	input = strings.TrimSpace(input)
	arr := strings.Split( input, " " )
	newStr := ""
	for i, s := range arr{
		if !IsPreposition( strings.ToLower(s) ) {
			newStr += UppercaseWord(s)
		}else{
			newStr += strings.ToLower(s)
		}
		
		if i!=len(arr)-1 {
			newStr += " "
		}
	}
	return newStr
}