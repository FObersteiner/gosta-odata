package godata

const (
	SearchTokenLiteral int = iota
	SearchTokenOpenParen
	SearchTokenCloseParen
	SearchTokenOp
	SearchTokenWhitespace
)

var (
	GlobalSearchTokenizer = SearchTokenizer()
	GlobalSearchParser    = SearchParser()
)

// Convert an input string from the $filter part of the URL into a parse
// tree that can be used by providers to create a response.
func ParseSearchString(filter string) (*GoDataSearchQuery, error) {
	tokens, err := GlobalSearchTokenizer.Tokenize(filter)
	if err != nil {
		return nil, err
	}
	postfix, err := GlobalSearchParser.InfixToPostfix(tokens)
	if err != nil {
		return nil, err
	}
	tree, err := GlobalSearchParser.PostfixToTree(postfix)
	if err != nil {
		return nil, err
	}
	return &GoDataSearchQuery{tree}, nil
}

// Create a tokenizer capable of tokenizing filter statements
func SearchTokenizer() *Tokenizer {
	t := Tokenizer{}
	t.Add("^\\\"[^\\\"]+\\\"", SearchTokenLiteral)
	t.Add("^\\(", SearchTokenOpenParen)
	t.Add("^\\)", SearchTokenCloseParen)
	t.Add("^(OR|AND|NOT)", SearchTokenOp)
	t.Add("^[\\w]+", SearchTokenLiteral)
	t.Ignore("^ ", SearchTokenWhitespace)

	return &t
}

func SearchParser() *Parser {
	parser := EmptyParser()
	parser.DefineOperator("NOT", 1, OpAssociationNone, 3)
	parser.DefineOperator("AND", 2, OpAssociationLeft, 2)
	parser.DefineOperator("OR", 2, OpAssociationLeft, 1)
	return parser
}
