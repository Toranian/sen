package token

type TokenType string

type TokenPosition struct {
	Row int
	Col int
}

type Token struct {
	Type    TokenType     // Type of the value (int)
	Literal string        // The actual value (64)
	Pos     TokenPosition // Track where each token is for better error handling
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	// Comparisons
	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="

	// Logical operators
	LAND = "&&"
	LOR  = "||"

	// Delimiters
	COMMA    = ","
	NEWLINE  = "\n"
	CARRIAGE = "\r"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	LBRACKET = "["
	RBRACKET = "]"

	// Additional Types
	STRING = "STRING"
	COLON  = ":"
)

// Keywords
const (
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	CLASS    = "CLASS"
	MOD      = "MOD"
	PUB      = "PUB"
	WHILE    = "WHILE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,

	"class": CLASS,
	"mod":   MOD,
	"pub":   PUB,
	"while": WHILE,
}

// Determine if the identifier is a keyword or not
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
