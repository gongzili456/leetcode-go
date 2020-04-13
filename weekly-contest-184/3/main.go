// 5382. HTML 实体解析器
package main

import "fmt"

func main() {
	text := "&amp; is an HTML entity but &ambassador; is not."
	res := entityParser(text)
	fmt.Println("> ", res)
}

func entityParser(text string) string {
	m := map[string]byte{
		"&quot;":  '"',
		"&apos;":  '\'',
		"&amp;":   '&',
		"&gt;":    '>',
		"&lt;":    '<',
		"&frasl;": '/',
	}

	ans := []byte{}

	var appendByte func(b, e int)
	appendByte = func(b, e int) {
		for j := b; j <= e; j++ {
			ans = append(ans, text[j])
		}
	}

	for i := 0; i < len(text); i++ {
		if text[i] == '&' {
			begin := i
			for {
				i++
				c := text[i]
				if c == ';' {
					k := text[begin : i+1]
					if v, ok := m[k]; ok {
						ans = append(ans, v)
					} else {
						appendByte(begin, i)
					}
					break
				}
				if c > 'z' || c < 'a' {
					appendByte(begin, i)
					break
				}
			}
		} else {
			ans = append(ans, text[i])
		}
	}

	return string(ans)
}
