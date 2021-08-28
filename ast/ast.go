// 式は値を生成し、文は生成しない。

package ast

import "github.com/kons16/monkey/token"

type Node interface {
	TokenLiteral() string // デバッグとテスト用
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program は構文解析器が生成する全ての AST のルートノードになるもの
type Program struct {
	Statement []Statement
}

// TokenLiteral は、そのノードが関連付けられているトークンのリテラル値を返す
func (p *Program) TokenLiteral() string {
	if len(p.Statement) > 0 {
		return p.Statement[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // token.LET トークン
	Name  *Identifier // 束縛の識別子を保持する
	Value Expression  // 値を生成する式を保持する
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier 構造体は、let x = 5; における x のような束縛の識別子を保持するために使用する
type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
