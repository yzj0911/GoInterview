package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"
)

//map 定义了需要初始化
type A struct {
	a string
	b string
	c int
	d map[int]B //	a.d = make(map[int]B)
}

type B struct {
}

func main3() {
	s1 := "x我是中国人hello word!,2020 street 188#"
	var count int
	for _, v := range s1 {
		fmt.Println(string(v))
		if unicode.Is(unicode.Han, v) {
			fmt.Println("找到中文")
			count++
		}
	}
	fmt.Println(count)
	fmt.Println(IsChineseChar(s1))
}

// 或者封装函数调用
func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}

func main() {
	//DistinctFile("hsmstar.sql", "out.txt")
	// MYWORD My word
	var head string
	var tail string
	var MYWORD string
	var sep string
	var zoom float64
	flag.StringVar(&head, "head", "		", "A sentence printed on the head")
	flag.StringVar(&tail, "tail", "\t\t\t\t", "A sentence printed on the tail")
	flag.StringVar(&MYWORD, "words", "love", "The words you want to talk")
	flag.StringVar(&sep, "sep", " ", "The separator")
	flag.Float64Var(&zoom, "zoom", 1.0, "Zoom setting")
	flag.Parse()

	chars := strings.Split(MYWORD, sep)

	//time.Sleep(time.Duration(1) * time.Second)
	fmt.Println(head)
	fmt.Println()
	//time.Sleep(time.Duration(1) * time.Second)
	for _, char := range chars {
		allChar := make([]string, 0)

		for y := 12 * zoom; y > -12*zoom; y-- {

			lst := make([]string, 0)
			lstCon := ""
			for x := -30 * zoom; x < 30*zoom; x++ {
				x2 := float64(x)
				y2 := float64(y)
				formula := math.Pow(math.Pow(x2*0.04/zoom, 2)+math.Pow(y2*0.1/zoom, 2)-1, 3) - math.Pow(x2*0.04/zoom, 2)*math.Pow(y2*0.1/zoom, 3)
				if formula <= 0 {
					index := int(x) % len(char)
					if index >= 0 {
						lstCon += string(char[index])
					} else {
						lstCon += string(char[int(float64(len(char))-math.Abs(float64(index)))])
					}

				} else {
					lstCon += " "
				}
			}
			lst = append(lst, lstCon)
			allChar = append(allChar, lst...)
		}

		//for _, text := range allChar {
		//	for _, t := range text {
		//
		//		time.Sleep(time.Duration(1) * time.Microsecond)
		//		fmt.Print(string(t))
		//	}
		//	fmt.Print("\n")
		//}
		for _, text := range allChar {
			time.Sleep(time.Duration(1) * time.Microsecond)
			fmt.Print(string(text))
			fmt.Print("\n")
		}
	}
	time.Sleep(time.Duration(1) * time.Minute)
	fmt.Println("\t\t\t\t", tail)
}

//DistinctFile 为指定文件去重
func DistinctFile(file string, output string) {
	// 读取需要去重的文件内容

	f, _ := os.Open(file)
	defer func() {
		ferr := f.Close()
		if ferr != nil {
			fmt.Println(ferr.Error())
		}
	}()

	reader := bufio.NewReader(f)

	// 去重map
	var set = make(map[string]bool, 0)
	// 去重后的结果
	var result string

	for {
		a, _, c := reader.ReadLine()
		//x:=strings.ReplaceAll(string(a),"\n","")
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), "tp_") {
			continue
		}
		b := string(a)
		StrFilterNonChinese(&b)
		//var i int
		//var s []byte
		for _, v := range b {
			//i++

			//s = append(s, v)
			//if i == 3|| ai== len(a){
			lineStr := string(v)

			// key存在则跳出本次循环
			if set[lineStr] {
				continue
			}
			if lineStr == "다" {
				fmt.Print(lineStr)
			}
			fmt.Print(lineStr)
			result += fmt.Sprintf("%s", lineStr)

			set[lineStr] = true

			//s = []byte{}
			//i = 0
			//}
		}
	}

	// 写入另一个文件
	nf, err := os.Create(output)
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(nf, strings.NewReader(result))

	defer nf.Close()
}

func main1() {
	//var a A
	//var b B
	//a.d = make(map[int]B)
	//a.d[1] = b
	//fmt.Println(a)
	var a = "INSERT INTO `ac_activity` VALUES (1229397310002, 'https://mstarimg.imlatin.com/xiaS_pFqQoIHKzUw5YQvWfjgn8c=.jpg', '蛮100-10', 1223354980004, 'hs', 1, '2020-11-23 13:54:44', '2020-12-23 13:54:44', 5, NULL, 0, 0, '', 100, 1, 1, 0, 0, 0.00, 0, 0, 2, 0, 100.00, 10.00, 0.00, '', '', '', '2020-11-23 21:55:31', '2020-12-04 11:37:38');"
	var i int
	var s []byte
	output := "output.txt"
	nf, _ := os.Create(output)
	for _, v := range []byte(a) {
		fmt.Println(string(v))
		i++
		s = append(s, v)
		if i == 3 {
			fmt.Println(string(s))
			io.Copy(nf, strings.NewReader(string(s)))
			s = []byte{}
			i = 0
		}

	}
	defer nf.Close()
}

var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")

//func main() {
//	str := "问11"
//	StrFilterNonChinese(&str)
//	fmt.Println(str)
//}

func StrFilterNonChinese(src *string) {
	strn := ""
	for _, c := range *src {
		if hzRegexp.MatchString(string(c)) {
			strn += string(c)
		}
	}

	*src = strn
}
