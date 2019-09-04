package wkt

import (
	"errors"
	"fmt"
	"github.com/Tony-Kwan/go-geo/geo"
	lex "github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
	"reflect"
	"strconv"
	"strings"
)

const (
	// Literal

	// Token
	tokenType       = "TYPE"
	tokenEmpty      = "EMPTY"
	tokenNumber     = "NUMBER"
	tokenLeftParen  = "L_PAREN"
	tokenRightParen = "R_PAREN"
	tokenComma      = "COMMA"

	// Supported wkt type
	point      = "POINT"
	lineString = "LINESTRING"
	polygon    = "POLYGON"
	//TODO: support other type
)

var (
	lexer      *lex.Lexer
	tokens     []string
	tokenMapId map[string]int

	ErrorUnexpectedEOF = errors.New("unexpected EOF")
)

func init() {
	tokens = []string{
		tokenType,
		tokenEmpty,
		tokenNumber,
		tokenLeftParen,
		tokenRightParen,
		tokenComma,
	}
	tokenMapId = make(map[string]int)
	for i, tk := range tokens {
		tokenMapId[tk] = i
	}

	lexer = lex.NewLexer()
	lexer.Add([]byte("EMPTY"), token(tokenEmpty))
	lexer.Add([]byte(`[a-zA-Z]+`), token(tokenType))
	lexer.Add([]byte(`\(`), token(tokenLeftParen))
	lexer.Add([]byte(`\)`), token(tokenRightParen))
	lexer.Add([]byte(`[-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?`), token(tokenNumber))
	lexer.Add([]byte(`,`), token(tokenComma))
	lexer.Add([]byte(`( |\t|\n|\r)+`), skip)

	if err := lexer.Compile(); err != nil {
		panic(err)
	}
}

type WktReader struct {
	scanner *lex.Scanner
}

func MustPolygon(s interface{}, err error) geo.Polygon {
	if err != nil {
		panic(err)
	}
	ret, ok := s.(geo.Polygon)
	if !ok {
		panic(fmt.Errorf("%v can't cast to *geo.Polygon", reflect.TypeOf(s)))
	}
	return ret
}

func (r WktReader) Read(wkt string) (interface{}, error) {
	scanner, err := lexer.Scanner([]byte(wkt))
	if err != nil {
		return nil, err
	}
	r.scanner = scanner

	wktType, err := r.getNextExpectedToken(tokenType)
	if err != nil {
		return nil, err
	}

	var geometry interface{}
	switch strings.ToUpper(wktType) {
	case point:
		geometry, err = r.readPoint()
	case lineString:
		geometry, err = r.readLineString()
	case polygon:
		geometry, err = r.readPolygon()
	default:
		return nil, fmt.Errorf("unsupported wktType: %s", wktType)
	}
	if err != nil {
		return nil, err
	}
	return geometry, r.getNextEOF()
}

func skip(*lex.Scanner, *machines.Match) (interface{}, error) {
	return nil, nil
}

func token(name string) lex.Action {
	return func(s *lex.Scanner, m *machines.Match) (interface{}, error) {
		return s.Token(tokenMapId[name], string(m.Bytes), m), nil
	}
}

func unexpectedToken(expect, found, value string) error {
	return fmt.Errorf("unexpected token: expect=%s, found=%s, value=%s", expect, found, value)
}

// For Debug
func printRestToken(scanner *lex.Scanner) {
	for tk, err, eof := scanner.Next(); !eof; tk, err, eof = scanner.Next() {
		if err != nil {
			fmt.Println("[ERROR]", err)
			return
		}
		token := tk.(*lex.Token)
		fmt.Printf("%10v | %-10v\n", tokens[token.Type], string(token.Lexeme))
	}
}

func (r *WktReader) readPoint() (geo.Point, error) {
	tkt, err := r.getNextEmptyOrLeftParen()
	if err != nil {
		return geo.Point{}, err
	}
	if strings.Compare(tkt, tokenEmpty) == 0 {
		return geo.NewPoint(0, 0, nil), nil
	}

	x, err := r.getNextNumber()
	if err != nil {
		return geo.Point{}, err
	}
	y, err := r.getNextNumber()
	if err != nil {
		return geo.Point{}, err
	}
	_, err = r.getNextExpectedToken(tokenRightParen)
	if err != nil {
		return geo.Point{}, err
	}
	return geo.NewPoint(x, y, nil), nil
}

func (r *WktReader) readLineString() (geo.LineString, error) {
	tkt, err := r.getNextEmptyOrLeftParen()
	if err != nil {
		return nil, err
	}
	pts := make(geo.LineString, 0)
	if strings.Compare(tkt, tokenEmpty) == 0 {
		return pts, nil
	}

	var x, y float64
	for {
		x, err = r.getNextNumber()
		if err != nil {
			return nil, err
		}
		y, err = r.getNextNumber()
		if err != nil {
			return nil, err
		}
		pts = append(pts, geo.NewPoint(x, y, nil))

		tkt, err = r.getNextCommaOfRightParen()
		if err != nil {
			return nil, err
		}
		if strings.Compare(tkt, tokenRightParen) == 0 {
			break
		}
	}
	return pts, nil
}

func (r *WktReader) readPolygon() (geo.Polygon, error) {
	tkt, err := r.getNextEmptyOrLeftParen()
	if err != nil {
		return geo.Polygon{}, err
	}
	if strings.Compare(tkt, tokenEmpty) == 0 {
		return geo.NewPolygon(make(geo.LinearRing, 0)), nil
	}

	rings := make([]geo.LinearRing, 0)
	for {
		pts, err := r.readLineString()
		if err != nil {
			return geo.Polygon{}, err
		}
		rings = append(rings, geo.LinearRing(pts))
		if tkt, err = r.getNextCommaOfRightParen(); err != nil {
			return geo.Polygon{}, err
		} else if strings.Compare(tkt, tokenRightParen) == 0 {
			break
		}
	}
	if len(rings) == 0 {
		return geo.NewPolygon(make(geo.LinearRing, 0)), nil
	}

	if len(rings) == 1 {
		return geo.NewPolygon(rings[0]), nil
	} else {
		return geo.NewPolygon(rings[0], rings[1:]...), nil
	}
}

func (r *WktReader) getNextEmptyOrLeftParen() (string, error) {
	tkt, tkv, err := r.getNextToken()
	if err != nil {
		return "", err
	}
	if strings.Compare(tkt, tokenEmpty) != 0 && strings.Compare(tkt, tokenLeftParen) != 0 {
		return "", unexpectedToken(tokenEmpty+"/"+tokenLeftParen, tkt, tkv)
	}
	return tkt, nil
}

func (r *WktReader) getNextCommaOfRightParen() (string, error) {
	tkt, tkv, err := r.getNextToken()
	if err != nil {
		return "", err
	}
	if strings.Compare(tkt, tokenComma) != 0 && strings.Compare(tkt, tokenRightParen) != 0 {
		return "", unexpectedToken(tokenComma+"/"+tokenRightParen, tkt, tkv)
	}
	return tkt, nil
}

func (r *WktReader) getNextRightParenAndEOF() error {
	_, err := r.getNextExpectedToken(tokenRightParen)
	if err != nil {
		return err
	}
	return r.getNextEOF()
}

func (r *WktReader) getNextNumber() (float64, error) {
	tkv, err := r.getNextExpectedToken(tokenNumber)
	if err != nil {
		return 0, err
	}
	f, err := strconv.ParseFloat(tkv, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

func (r *WktReader) getNextExpectedToken(expect string) (string, error) {
	token, value, err := r.getNextToken()
	if err != nil {
		return "", err
	}
	if strings.Compare(token, expect) != 0 {
		return "", unexpectedToken(expect, token, value)
	}
	return value, nil
}

func (r *WktReader) getNextToken() (string, string, error) {
	tk, err, eof := r.scanner.Next()
	if eof {
		return "", "", ErrorUnexpectedEOF
	}
	if err != nil {
		return "", "", err
	}
	tok := tk.(*lex.Token)
	return tokens[tok.Type], string(tok.Lexeme), nil
}

func (r *WktReader) getNextEOF() error {
	tk, err, eof := r.scanner.Next()
	if !eof {
		tok := tk.(*lex.Token)
		return unexpectedToken("EOF", tokens[tok.Type], string(tok.Lexeme))
	}
	if err != nil {
		return err
	}
	return nil
}
