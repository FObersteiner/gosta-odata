package godata

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestFilterTokenizerGeoFunc(t *testing.T) {
	tokenizer := GlobalFilterTokenizer
	input := "st_within(location, geography'LINESTRING(7.5 51.5, 7.5 53.5)')"
	expect := []*Token{
		{Value: "st_within", Type: FilterTokenFunc},
		{Value: "(", Type: FilterTokenOpenParen},
		{Value: "location", Type: FilterTokenLiteral},
		{Value: ",", Type: FilterTokenComma},
		{Value: "geography", Type: FilterTokenGeography},
		{Value: "'LINESTRING(7.5 51.5, 7.5 53.5)'", Type: FilterTokenString},
		{Value: ")", Type: FilterTokenCloseParen},
	}
	output, err := tokenizer.Tokenize(input)
	if err != nil {
		t.Error(err)
	}

	result, err := CompareTokens(expect, output)
	if !result {
		t.Error(err)
	}
}

func TestLogicalOperator(t *testing.T) {
	tokenizer := GlobalFilterTokenizer
	input := "properties/gemeentecode eq '518'"
	expect := []*Token{
		{Value: "properties", Type: FilterTokenLiteral},
		{Value: "/", Type: FilterTokenNav},
		{Value: "gemeentecode", Type: FilterTokenLiteral},
		{Value: "eq", Type: FilterTokenLogical},
		{Value: "'518'", Type: FilterTokenString},
	}
	output, err := tokenizer.Tokenize(input)
	if err != nil {
		t.Error(err)
	}

	result, err := CompareTokens(expect, output)
	if !result {
		t.Error(err)
	}
}

func TestDateTimeWithOffset(t *testing.T) {
	tokenizer := FilterTokenizer()
	input := "time lt 2015-10-14T23:30:00.104+02:00"
	expect := []*Token{
		{Value: "time", Type: FilterTokenLiteral},
		{Value: "lt", Type: FilterTokenLogical},
		{Value: "2015-10-14T23:30:00.104+02:00", Type: FilterTokenDateTime},
	}
	output, err := tokenizer.Tokenize(input)
	if err != nil {
		t.Error(err)
	}

	result, err := CompareTokens(expect, output)
	if !result {
		t.Error(err)
	}
}

// TestFilterLiteralContainingFunctionName checks if "time" is not seen as an function
func TestFilterLiteralContainingFunctionName(t *testing.T) {
	tokenizer := FilterTokenizer()
	input := "time lt '2015-10-14T23:30:00.104+02:00'"
	expect := []*Token{
		{Value: "time", Type: FilterTokenLiteral},
		{Value: "lt", Type: FilterTokenLogical},
		{Value: "'2015-10-14T23:30:00.104+02:00'", Type: FilterTokenString},
	}
	output, err := tokenizer.Tokenize(input)
	if err != nil {
		t.Error(err)
	}

	result, err := CompareTokens(expect, output)
	if !result {
		t.Error(err)
	}
}

func TestFilterAny(t *testing.T) {
	tokenizer := FilterTokenizer()
	input := "Tags/any(d:d/Key eq 'Site' and d/Value lt 10)"
	expect := []*Token{
		{Value: "Tags", Type: FilterTokenLiteral},
		{Value: "/", Type: FilterTokenNav},
		{Value: "any", Type: FilterTokenLambda},
		{Value: "(", Type: FilterTokenOpenParen},
		{Value: "d", Type: FilterTokenLiteral},
		{Value: ":", Type: FilterTokenColon},
		{Value: "d", Type: FilterTokenLiteral},
		{Value: "/", Type: FilterTokenNav},
		{Value: "Key", Type: FilterTokenLiteral},
		{Value: "eq", Type: FilterTokenLogical},
		{Value: "'Site'", Type: FilterTokenString},
		{Value: "and", Type: FilterTokenLogical},
		{Value: "d", Type: FilterTokenLiteral},
		{Value: "/", Type: FilterTokenNav},
		{Value: "Value", Type: FilterTokenLiteral},
		{Value: "lt", Type: FilterTokenLogical},
		{Value: "10", Type: FilterTokenInteger},
		{Value: ")", Type: FilterTokenCloseParen},
	}
	output, err := tokenizer.Tokenize(input)
	if err != nil {
		t.Error(err)
	}

	result, err := CompareTokens(expect, output)
	if !result {
		t.Error(err)
	}
}

func TestFilterAll(t *testing.T) {
	tokenizer := FilterTokenizer()
	input := "Tags/all(d:d/Key eq 'Site')"
	expect := []*Token{
		{Value: "Tags", Type: FilterTokenLiteral},
		{Value: "/", Type: FilterTokenNav},
		{Value: "all", Type: FilterTokenLambda},
		{Value: "(", Type: FilterTokenOpenParen},
		{Value: "d", Type: FilterTokenLiteral},
		{Value: ":", Type: FilterTokenColon},
		{Value: "d", Type: FilterTokenLiteral},
		{Value: "/", Type: FilterTokenNav},
		{Value: "Key", Type: FilterTokenLiteral},
		{Value: "eq", Type: FilterTokenLogical},
		{Value: "'Site'", Type: FilterTokenString},
		{Value: ")", Type: FilterTokenCloseParen},
	}
	output, err := tokenizer.Tokenize(input)
	if err != nil {
		t.Error(err)
	}

	result, err := CompareTokens(expect, output)
	if !result {
		t.Error(err)
	}
}

func TestFilterTokenizer(t *testing.T) {
	tokenizer := FilterTokenizer()
	input := "Name eq 'Milk' and Price lt 2.55"
	expect := []*Token{
		{Value: "Name", Type: FilterTokenLiteral},
		{Value: "eq", Type: FilterTokenLogical},
		{Value: "'Milk'", Type: FilterTokenString},
		{Value: "and", Type: FilterTokenLogical},
		{Value: "Price", Type: FilterTokenLiteral},
		{Value: "lt", Type: FilterTokenLogical},
		{Value: "2.55", Type: FilterTokenFloat},
	}
	output, err := tokenizer.Tokenize(input)
	if err != nil {
		t.Error(err)
	}

	result, err := CompareTokens(expect, output)
	if !result {
		t.Error(err)
	}
}

func TestFilterTokenizerFunc(t *testing.T) {
	tokenizer := FilterTokenizer()
	input := "not endswith(Name,'ilk')"
	expect := []*Token{
		{Value: "not", Type: FilterTokenLogical},
		{Value: "endswith", Type: FilterTokenFunc},
		{Value: "(", Type: FilterTokenOpenParen},
		{Value: "Name", Type: FilterTokenLiteral},
		{Value: ",", Type: FilterTokenComma},
		{Value: "'ilk'", Type: FilterTokenString},
		{Value: ")", Type: FilterTokenCloseParen},
	}
	output, err := tokenizer.Tokenize(input)
	if err != nil {
		t.Error(err)
	}

	result, err := CompareTokens(expect, output)
	if !result {
		t.Error(err)
	}
}

func BenchmarkFilterTokenizer(b *testing.B) {
	t := FilterTokenizer()
	for i := 0; i < b.N; i++ {
		input := "Name eq 'Milk' and Price lt 2.55"
		t.Tokenize(input)
	}
}

// Check if two slices of tokens are the same.
func CompareTokens(a, b []*Token) (bool, error) {
	if len(a) != len(b) {
		res := ""
		for _, t := range b {
			res += fmt.Sprintf("| %s ", t.Value)
		}
		return false, fmt.Errorf("Different lengths, result: %s", res)
	}
	for i := range a {
		if a[i].Value != b[i].Value || a[i].Type != b[i].Type {
			return false, errors.New("Different at index " + strconv.Itoa(i) + " " +
				a[i].Value + " != " + b[i].Value +
				fmt.Sprintf(" or types are different. Type expected: %v Type result: %v", a[i].Type, b[i].Type))
		}
	}
	return true, nil
}

func TestFilterParserTree(t *testing.T) {
	input := "not (A eq B)"

	tokens, err := GlobalFilterTokenizer.Tokenize(input)
	if err != nil {
		t.Error(err)
		return
	}
	output, err := GlobalFilterParser.InfixToPostfix(tokens)
	if err != nil {
		t.Error(err)
		return
	}

	tree, err := GlobalFilterParser.PostfixToTree(output)
	if err != nil {
		t.Error(err)
		return
	}

	if tree.Token.Value != "not" {
		t.Error("Root is '" + tree.Token.Value + "' not 'not'")
	}
	if tree.Children[0].Token.Value != "eq" {
		t.Error("First child is '" + tree.Children[1].Token.Value + "' not 'eq'")
	}
}

func printTree(n *ParseNode, level int) {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}
	fmt.Printf("%s %-10s %-10d\n", indent, n.Token.Value, n.Token.Type)
	for _, v := range n.Children {
		printTree(v, level+1)
	}
}

func TestNestedPath(t *testing.T) {
	input := "Address/City eq 'Redmond'"
	tokens, err := GlobalFilterTokenizer.Tokenize(input)
	if err != nil {
		t.Error(err)
		return
	}
	output, err := GlobalFilterParser.InfixToPostfix(tokens)
	if err != nil {
		t.Error(err)
		return
	}

	tree, err := GlobalFilterParser.PostfixToTree(output)
	if err != nil {
		t.Error(err)
		return
	}
	// printTree(tree, 0)
	if tree.Token.Value != "eq" {
		t.Error("Root is '" + tree.Token.Value + "' not 'eq'")
	}
	if tree.Children[0].Token.Value != "/" {
		t.Error("First child is \"" + tree.Children[0].Token.Value + "\", not '/'")
	}
	if tree.Children[1].Token.Value != "'Redmond'" {
		t.Error("First child is \"" + tree.Children[1].Token.Value + "\", not 'Redmond'")
	}
}

func TestLambda(t *testing.T) {
	input := "Tags/any(var:var/Key eq 'Site' and var/Value eq 'London')"
	tokens, err := GlobalFilterTokenizer.Tokenize(input)
	if err != nil {
		t.Error(err)
		return
	}
	output, err := GlobalFilterParser.InfixToPostfix(tokens)
	if err != nil {
		t.Error(err)
		return
	}

	tree, err := GlobalFilterParser.PostfixToTree(output)
	if err != nil {
		t.Error(err)
		return
	}
	// printTree(tree, 0)

	if tree.Token.Value != "/" {
		t.Error("Root is '" + tree.Token.Value + "' not '/'")
	}
	if tree.Children[0].Token.Value != "Tags" {
		t.Error("First child is '" + tree.Children[0].Token.Value + "' not 'Tags'")
	}
	if tree.Children[1].Token.Value != "any" {
		t.Error("First child is '" + tree.Children[1].Token.Value + "' not 'any'")
	}
	if tree.Children[1].Children[0].Token.Value != ":" {
		t.Error("First child is '" + tree.Children[1].Children[0].Token.Value + "' not ':'")
	}
}
