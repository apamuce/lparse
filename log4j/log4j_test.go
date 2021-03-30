package log4j

import (
	"github.com/apamuce/lparse"
	"testing"
	"time"
)

const reg_exp = `([0-9A-Za-z,-\\/]+\s[0-9A-Za-z-:]+),[0-9]+\s([0-9]+)\s+([\w]+)\s+(\[.+\])\s(\(.+\))(.*)`

func TestLog4j_VerifyParserType(t *testing.T) {
	//ARRANGE
	p := NewLog4jParser(reg_exp)

	//ACT
	pt := p.GetParserType()

	//ASSERT
	if pt != lparse.Log4j {
		t.Errorf("Invalid parser type %v. Expected: %v", pt, lparse.Log4j)
	}
}

func TestLog4j_ParseLine(t *testing.T) {
	//ARRANGE
	p := NewLog4jParser(reg_exp)
	log_line := "2020-10-30 17:36:28,000 73679146 INFO  [org.springframework.beans.factory.annotation.AutowiredAnnotationBeanPostProcessor] (sharedTaskSchedulerFactoryBean_QuartzSchedulerThread:) JSR-330 'javax.inject.Inject' annotation found and supported for autowiring"
	date := time.Date(2020, 10, 30, 17, 36, 28, 0, time.UTC)
	sev := lparse.INFO
	src := "[org.springframework.beans.factory.annotation.AutowiredAnnotationBeanPostProcessor]"
	th := "(sharedTaskSchedulerFactoryBean_QuartzSchedulerThread:)"
	cont := " JSR-330 'javax.inject.Inject' annotation found and supported for autowiring"

	//ACT
	l, err := p.Parse(log_line)

	//ASSERT
	if err != nil {
		t.Errorf("Failed to parse data, received error %v", err)
	} else {
		if l.Date != date {
			t.Errorf("Expected Date `%v`, received `%v`", date, l.Date)
		}
		if l.Severity != sev {
			t.Errorf("Expected Severity `%v`, received `%v`", sev, l.Severity)
		}
		if l.SrcFile != src {
			t.Errorf("Expected SrcFile `%v`, received `%v`", src, l.SrcFile)
		}
		if l.Thread != th {
			t.Errorf("Expected ThreadID `%v`, received `%v`", th, l.Thread)
		}
		if l.Content != cont {
			t.Errorf("Expected Content `%v`, received `%v`", cont, l.Content)
		}
	}
}

func TestLog4j_ParseContentLine(t *testing.T) {
	//ARRANGE
	p := NewLog4jParser(reg_exp)
	logLine := " Info about 'DistributedBlockingQueueTemplate.execute' "
	expectedContent := "\n" + logLine

	//ACT
	l, err := p.Parse(logLine)

	//ASSERT
	if err != nil {
		t.Errorf("Failed to parse data, received error %v", err)
	} else {
		if l.Content != expectedContent {
			t.Errorf("Expected content '%v', received '%v'", expectedContent, l.Content)
		}
	}
}

func TestEmptyLogLineParse(t *testing.T) {
	//ARRANGE
	p := NewLog4jParser(reg_exp)

	//ACT
	_, err := p.Parse("")

	//ASSERT
	if err == nil {
		t.Errorf("Error expected, received nil")
	}
}
