package ast

import (
	"monkey/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
func (p *Program) String() string {
	var sb strings.Builder
	for _, s := range p.Statements {
		sb.WriteString(s.String())
	}
	return sb.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
func (ls *LetStatement) String() string {
	var sb strings.Builder
	sb.WriteString(ls.TokenLiteral() + " ")
	sb.WriteString(ls.Name.String())
	sb.WriteString(" = ")
	if ls.Value != nil {
		sb.WriteString(ls.Value.String())
	}
	sb.WriteString(";")
	return sb.String()
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var sb strings.Builder
	sb.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		sb.WriteString(rs.ReturnValue.String())
	}
	sb.WriteString(";")
	return sb.String()
}

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(pe.Operator)
	sb.WriteString(pe.Right.String())
	sb.WriteString(")")
	return sb.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(ie.Left.String())
	sb.WriteString(" " + ie.Operator + " ")
	if ie.Right != nil {
		sb.WriteString(ie.Right.String())
	}
	sb.WriteString(")")
	return sb.String()
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var sb strings.Builder
	sb.WriteString("if")
	sb.WriteString(ie.Condition.String())
	sb.WriteString(" ")
	sb.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {

		sb.WriteString("else ")
		sb.WriteString(ie.Alternative.String())
	}
	return sb.String()
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var sb strings.Builder
	for _, s := range bs.Statements {
		sb.WriteString(s.String())
	}
	return sb.String()
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var sb strings.Builder
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	sb.WriteString(fl.TokenLiteral())
	sb.WriteString("(")
	sb.WriteString(strings.Join(params, ", "))
	sb.WriteString(") ")
	sb.WriteString(fl.Body.String())
	return sb.String()
}

type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var sb strings.Builder
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	sb.WriteString(ce.Function.String())
	sb.WriteString("(")
	sb.WriteString(strings.Join(args, ", "))
	sb.WriteString(")")
	return sb.String()
}

