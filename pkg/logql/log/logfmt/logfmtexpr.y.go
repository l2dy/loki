// Code generated by goyacc -p LogfmtExpr -o logfmtexpr.y.go logfmtexpr.y. DO NOT EDIT.

//line logfmtexpr.y:4
package logfmt

import __yyfmt__ "fmt"

//line logfmtexpr.y:4

func setScannerData(lex interface{}, data []interface{}) {
	lex.(*Scanner).data = data
}

//line logfmtexpr.y:12
type LogfmtExprSymType struct {
	yys  int
	str  string
	key  string
	list []interface{}
}

const STRING = 57346
const KEY = 57347

var LogfmtExprToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"STRING",
	"KEY",
}

var LogfmtExprStatenames = [...]string{}

const LogfmtExprEofCode = 1
const LogfmtExprErrCode = 2
const LogfmtExprInitialStackSize = 16

//line yacctab:1
var LogfmtExprExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const LogfmtExprPrivate = 57344

const LogfmtExprLast = 8

var LogfmtExprAct = [...]int8{
	6, 5, 6, 4, 1, 2, 7, 3,
}

var LogfmtExprPact = [...]int16{
	-4, -1000, -2, -1000, -1000, -1000, -1000, -1000,
}

var LogfmtExprPgo = [...]int8{
	0, 7, 3, 5, 4,
}

var LogfmtExprR1 = [...]int8{
	0, 4, 3, 3, 3, 1, 2,
}

var LogfmtExprR2 = [...]int8{
	0, 1, 1, 1, 2, 1, 1,
}

var LogfmtExprChk = [...]int16{
	-1000, -4, -3, -1, -2, 5, 4, -2,
}

var LogfmtExprDef = [...]int8{
	0, -2, 1, 2, 3, 5, 6, 4,
}

var LogfmtExprTok1 = [...]int8{
	1,
}

var LogfmtExprTok2 = [...]int8{
	2, 3, 4, 5,
}

var LogfmtExprTok3 = [...]int8{
	0,
}

var LogfmtExprErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	LogfmtExprDebug        = 0
	LogfmtExprErrorVerbose = false
)

type LogfmtExprLexer interface {
	Lex(lval *LogfmtExprSymType) int
	Error(s string)
}

type LogfmtExprParser interface {
	Parse(LogfmtExprLexer) int
	Lookahead() int
}

type LogfmtExprParserImpl struct {
	lval  LogfmtExprSymType
	stack [LogfmtExprInitialStackSize]LogfmtExprSymType
	char  int
}

func (p *LogfmtExprParserImpl) Lookahead() int {
	return p.char
}

func LogfmtExprNewParser() LogfmtExprParser {
	return &LogfmtExprParserImpl{}
}

const LogfmtExprFlag = -1000

func LogfmtExprTokname(c int) string {
	if c >= 1 && c-1 < len(LogfmtExprToknames) {
		if LogfmtExprToknames[c-1] != "" {
			return LogfmtExprToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func LogfmtExprStatname(s int) string {
	if s >= 0 && s < len(LogfmtExprStatenames) {
		if LogfmtExprStatenames[s] != "" {
			return LogfmtExprStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func LogfmtExprErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !LogfmtExprErrorVerbose {
		return "syntax error"
	}

	for _, e := range LogfmtExprErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + LogfmtExprTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(LogfmtExprPact[state])
	for tok := TOKSTART; tok-1 < len(LogfmtExprToknames); tok++ {
		if n := base + tok; n >= 0 && n < LogfmtExprLast && int(LogfmtExprChk[int(LogfmtExprAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if LogfmtExprDef[state] == -2 {
		i := 0
		for LogfmtExprExca[i] != -1 || int(LogfmtExprExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; LogfmtExprExca[i] >= 0; i += 2 {
			tok := int(LogfmtExprExca[i])
			if tok < TOKSTART || LogfmtExprExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if LogfmtExprExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += LogfmtExprTokname(tok)
	}
	return res
}

func LogfmtExprlex1(lex LogfmtExprLexer, lval *LogfmtExprSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(LogfmtExprTok1[0])
		goto out
	}
	if char < len(LogfmtExprTok1) {
		token = int(LogfmtExprTok1[char])
		goto out
	}
	if char >= LogfmtExprPrivate {
		if char < LogfmtExprPrivate+len(LogfmtExprTok2) {
			token = int(LogfmtExprTok2[char-LogfmtExprPrivate])
			goto out
		}
	}
	for i := 0; i < len(LogfmtExprTok3); i += 2 {
		token = int(LogfmtExprTok3[i+0])
		if token == char {
			token = int(LogfmtExprTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(LogfmtExprTok2[1]) /* unknown char */
	}
	if LogfmtExprDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", LogfmtExprTokname(token), uint(char))
	}
	return char, token
}

func LogfmtExprParse(LogfmtExprlex LogfmtExprLexer) int {
	return LogfmtExprNewParser().Parse(LogfmtExprlex)
}

func (LogfmtExprrcvr *LogfmtExprParserImpl) Parse(LogfmtExprlex LogfmtExprLexer) int {
	var LogfmtExprn int
	var LogfmtExprVAL LogfmtExprSymType
	var LogfmtExprDollar []LogfmtExprSymType
	_ = LogfmtExprDollar // silence set and not used
	LogfmtExprS := LogfmtExprrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	LogfmtExprstate := 0
	LogfmtExprrcvr.char = -1
	LogfmtExprtoken := -1 // LogfmtExprrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		LogfmtExprstate = -1
		LogfmtExprrcvr.char = -1
		LogfmtExprtoken = -1
	}()
	LogfmtExprp := -1
	goto LogfmtExprstack

ret0:
	return 0

ret1:
	return 1

LogfmtExprstack:
	/* put a state and value onto the stack */
	if LogfmtExprDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", LogfmtExprTokname(LogfmtExprtoken), LogfmtExprStatname(LogfmtExprstate))
	}

	LogfmtExprp++
	if LogfmtExprp >= len(LogfmtExprS) {
		nyys := make([]LogfmtExprSymType, len(LogfmtExprS)*2)
		copy(nyys, LogfmtExprS)
		LogfmtExprS = nyys
	}
	LogfmtExprS[LogfmtExprp] = LogfmtExprVAL
	LogfmtExprS[LogfmtExprp].yys = LogfmtExprstate

LogfmtExprnewstate:
	LogfmtExprn = int(LogfmtExprPact[LogfmtExprstate])
	if LogfmtExprn <= LogfmtExprFlag {
		goto LogfmtExprdefault /* simple state */
	}
	if LogfmtExprrcvr.char < 0 {
		LogfmtExprrcvr.char, LogfmtExprtoken = LogfmtExprlex1(LogfmtExprlex, &LogfmtExprrcvr.lval)
	}
	LogfmtExprn += LogfmtExprtoken
	if LogfmtExprn < 0 || LogfmtExprn >= LogfmtExprLast {
		goto LogfmtExprdefault
	}
	LogfmtExprn = int(LogfmtExprAct[LogfmtExprn])
	if int(LogfmtExprChk[LogfmtExprn]) == LogfmtExprtoken { /* valid shift */
		LogfmtExprrcvr.char = -1
		LogfmtExprtoken = -1
		LogfmtExprVAL = LogfmtExprrcvr.lval
		LogfmtExprstate = LogfmtExprn
		if Errflag > 0 {
			Errflag--
		}
		goto LogfmtExprstack
	}

LogfmtExprdefault:
	/* default state action */
	LogfmtExprn = int(LogfmtExprDef[LogfmtExprstate])
	if LogfmtExprn == -2 {
		if LogfmtExprrcvr.char < 0 {
			LogfmtExprrcvr.char, LogfmtExprtoken = LogfmtExprlex1(LogfmtExprlex, &LogfmtExprrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if LogfmtExprExca[xi+0] == -1 && int(LogfmtExprExca[xi+1]) == LogfmtExprstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			LogfmtExprn = int(LogfmtExprExca[xi+0])
			if LogfmtExprn < 0 || LogfmtExprn == LogfmtExprtoken {
				break
			}
		}
		LogfmtExprn = int(LogfmtExprExca[xi+1])
		if LogfmtExprn < 0 {
			goto ret0
		}
	}
	if LogfmtExprn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			LogfmtExprlex.Error(LogfmtExprErrorMessage(LogfmtExprstate, LogfmtExprtoken))
			Nerrs++
			if LogfmtExprDebug >= 1 {
				__yyfmt__.Printf("%s", LogfmtExprStatname(LogfmtExprstate))
				__yyfmt__.Printf(" saw %s\n", LogfmtExprTokname(LogfmtExprtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for LogfmtExprp >= 0 {
				LogfmtExprn = int(LogfmtExprPact[LogfmtExprS[LogfmtExprp].yys]) + LogfmtExprErrCode
				if LogfmtExprn >= 0 && LogfmtExprn < LogfmtExprLast {
					LogfmtExprstate = int(LogfmtExprAct[LogfmtExprn]) /* simulate a shift of "error" */
					if int(LogfmtExprChk[LogfmtExprstate]) == LogfmtExprErrCode {
						goto LogfmtExprstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if LogfmtExprDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", LogfmtExprS[LogfmtExprp].yys)
				}
				LogfmtExprp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if LogfmtExprDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", LogfmtExprTokname(LogfmtExprtoken))
			}
			if LogfmtExprtoken == LogfmtExprEofCode {
				goto ret1
			}
			LogfmtExprrcvr.char = -1
			LogfmtExprtoken = -1
			goto LogfmtExprnewstate /* try again in the same state */
		}
	}

	/* reduction by production LogfmtExprn */
	if LogfmtExprDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", LogfmtExprn, LogfmtExprStatname(LogfmtExprstate))
	}

	LogfmtExprnt := LogfmtExprn
	LogfmtExprpt := LogfmtExprp
	_ = LogfmtExprpt // guard against "declared and not used"

	LogfmtExprp -= int(LogfmtExprR2[LogfmtExprn])
	// LogfmtExprp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if LogfmtExprp+1 >= len(LogfmtExprS) {
		nyys := make([]LogfmtExprSymType, len(LogfmtExprS)*2)
		copy(nyys, LogfmtExprS)
		LogfmtExprS = nyys
	}
	LogfmtExprVAL = LogfmtExprS[LogfmtExprp+1]

	/* consult goto table to find next state */
	LogfmtExprn = int(LogfmtExprR1[LogfmtExprn])
	LogfmtExprg := int(LogfmtExprPgo[LogfmtExprn])
	LogfmtExprj := LogfmtExprg + LogfmtExprS[LogfmtExprp].yys + 1

	if LogfmtExprj >= LogfmtExprLast {
		LogfmtExprstate = int(LogfmtExprAct[LogfmtExprg])
	} else {
		LogfmtExprstate = int(LogfmtExprAct[LogfmtExprj])
		if int(LogfmtExprChk[LogfmtExprstate]) != -LogfmtExprn {
			LogfmtExprstate = int(LogfmtExprAct[LogfmtExprg])
		}
	}
	// dummy call; replaced with literal code
	switch LogfmtExprnt {

	case 1:
		LogfmtExprDollar = LogfmtExprS[LogfmtExprpt-1 : LogfmtExprpt+1]
//line logfmtexpr.y:27
		{
			setScannerData(LogfmtExprlex, LogfmtExprDollar[1].list)
		}
	case 2:
		LogfmtExprDollar = LogfmtExprS[LogfmtExprpt-1 : LogfmtExprpt+1]
//line logfmtexpr.y:30
		{
			LogfmtExprVAL.list = []interface{}{LogfmtExprDollar[1].str}
		}
	case 3:
		LogfmtExprDollar = LogfmtExprS[LogfmtExprpt-1 : LogfmtExprpt+1]
//line logfmtexpr.y:31
		{
			LogfmtExprVAL.list = []interface{}{LogfmtExprDollar[1].str}
		}
	case 4:
		LogfmtExprDollar = LogfmtExprS[LogfmtExprpt-2 : LogfmtExprpt+1]
//line logfmtexpr.y:32
		{
			LogfmtExprVAL.list = append(LogfmtExprDollar[1].list, LogfmtExprDollar[2].str)
		}
	case 5:
		LogfmtExprDollar = LogfmtExprS[LogfmtExprpt-1 : LogfmtExprpt+1]
//line logfmtexpr.y:36
		{
			LogfmtExprVAL.str = LogfmtExprDollar[1].key
		}
	case 6:
		LogfmtExprDollar = LogfmtExprS[LogfmtExprpt-1 : LogfmtExprpt+1]
//line logfmtexpr.y:39
		{
			LogfmtExprVAL.str = LogfmtExprDollar[1].str
		}
	}
	goto LogfmtExprstack /* stack new state and value */
}