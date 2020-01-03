package pattern

type RegularExpr struct {
	matched bool   // 是否匹配
	pattern []rune // 正则表达式
	plen    int    // 正则表表达式长度
}

func (re *RegularExpr) Match(text []rune, tlen int) bool {
	re.matched = false
	re.rmatch(0, 0, text, tlen)
	return re.matched
}

func (re *RegularExpr) rmatch(ti, pj int, text []rune, tlen int) {
	if re.matched {
		return
	}
	if pj == re.plen {
		if ti == tlen {
			re.matched = true
		}
		return
	}
	if re.pattern[pj] == '*' {
		for k := 0; k <= tlen-ti; k++ {
			re.rmatch(ti+k, pj+1, text, tlen)
		}
	} else if re.pattern[pj] == '?' {
		re.rmatch(ti, pj+1, text, tlen)
		re.rmatch(ti+1, pj+1, text, tlen)
	} else if ti < tlen && re.pattern[pj] == text[ti] {
		re.rmatch(ti+1, pj+1, text, tlen)
	}
}
