package sqlbuild

import (
	"fmt"
	"strings"
)

type (
	PlaceholderFormat string

	Placeholder interface {
		Next() string
		Reset(to int)
		Get() int
		Fill(count int) []string
	}

	placeholder struct {
		count  int
		format PlaceholderFormat
	}
)

const (
	FormatPostgres PlaceholderFormat = "$%d"
	FormatMysql    PlaceholderFormat = "?"
)

// NewPlaceholder init new placeholder
func NewPlaceholder(options ...PlaceholderOption) Placeholder {
	var args = &PlaceholderOptions{
		format: FormatPostgres,
		offset: 0,
	}

	var opt PlaceholderOption
	for _, opt = range options {
		opt(args)
	}

	return &placeholder{
		format: args.format,
		count:  args.offset,
	}
}

// Next get next placeholder, started with 1
func (p *placeholder) Next() string {
	p.count++

	if !strings.Contains(p.format.String(), "%d") {
		return fmt.Sprintf(p.format.String())
	}

	return fmt.Sprintf(p.format.String(), p.count)
}

// Reset set cursor with new value
func (p *placeholder) Reset(to int) {
	p.count = to
}

// Get get current cursor value
func (p *placeholder) Get() int {
	return p.count
}

// Fill generate placeholders slice
func (p *placeholder) Fill(count int) []string {
	var ret = make([]string, count)
	var i int
	for i = 0; i < count; i++ {
		ret[i] = p.Next()
	}
	return ret
}

// String stringer for PlaceholderFormat
func (p PlaceholderFormat) String() string {
	return string(p)
}
