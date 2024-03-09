package tokenizer

type TokenType int

const (
	eof TokenType = iota
	h1
	h2
	h3
	h4
	h5
	h6
	p
	bold
	italic
	link
	image
	mono
	code
	ol
	ul
	escaped
	hr
)
