package lparse

import (
	"time"
)

type ParserType = int
type Severity = int

const (
	Log4j ParserType = iota
)

const (
	INFO Severity = iota
	WARNING
	ERROR
	DEBUG
	UNKNOWN
)

type LogEntry struct {
	Date     time.Time
	Severity Severity
	SrcFile  string
	Thread   string
	Content  string
}

type LogParser interface {
	GetParserType() ParserType
	Parse(line string) (*LogEntry, error)
	ParseBulk(lines []string) ([]*LogEntry, error)
}

//type Parser struct {
//	parserType ParserType
//	LogParser LogParser
//}
//
//func NewParser(parserType ParserType, regexp string) *Parser {
//	switch parserType {
//	case Log4j:
//		return &Parser{
//			parserType: parserType,
//			LogParser:  log4j.NewLog4j(regexp),
//		}
//	default:
//		panic(fmt.Sprintf("Unimplemented parser type: %v", parserType))
//	}
//}
