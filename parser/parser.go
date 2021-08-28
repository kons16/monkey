package parser

import (
	"github.com/kons16/monkey/ast"
	"github.com/kons16/monkey/lexer"
	"github.com/kons16/monkey/token"
)

type Parser struct {
	l *lexer.Lexer // 字句解析器インスタンスへのポインタで、このインスタンスの NextToken() を繰り返し呼び、入力から次のトークンを取得する

	curToken  token.Token // 現在のトークン
	peekToken token.Token // 次のトークン
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// 2つのトークンを読み込む。cureToken と peekToken の両方がセットされる。
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
