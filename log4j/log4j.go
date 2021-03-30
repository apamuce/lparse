package log4j

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/apamuce/lparse"
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

type Log4j struct {
	regexp     string
	parserType lparse.ParserType
}

func (l Log4j) GetParserType() lparse.ParserType {
	return l.parserType
}

func (l Log4j) Parse(line string) (data *lparse.LogEntry, err error) {
	r := *regexp.MustCompile(l.regexp)

	if len(line) < 1 {
		return nil, fmt.Errorf("empty string to parse")
	}

	if r.MatchString(line) {
		res := r.FindAllStringSubmatch(line, -1)
		date, _ := time.Parse(timeLayout, res[0][1])
		data = &lparse.LogEntry{
			Date:     date,
			Severity: stringToSeverity(strings.TrimSpace(res[0][3])),
			SrcFile:  strings.TrimSpace(res[0][4]),
			Thread:   strings.TrimSpace(res[0][5]),
			Content:  res[0][6],
		}
	} else {
		data = &lparse.LogEntry{
			Content: fmt.Sprintf("\n%s", line),
		}
	}

	return data, nil
}

func (l Log4j) ParseBulk(lines []string) (data []*lparse.LogEntry, err error) {
	for _, line := range lines {
		d, err := l.Parse(line)
		if err != nil {
			data = append(data, d)
		} else {
			return nil, err
		}
	}

	return data, nil
}

func NewLog4jParser(regexp string) *Log4j {
	return &Log4j{
		regexp:     regexp,
		parserType: lparse.Log4j,
	}
}

func stringToSeverity(severity string) lparse.Severity {
	switch severity {
	case "INFO":
		return lparse.INFO
	case "WARNING":
		return lparse.WARNING
	case "ERROR":
		return lparse.ERROR
	case "DEBUG":
		return lparse.DEBUG
	default:
		return lparse.UNKNOWN
	}
}
