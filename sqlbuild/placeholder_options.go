package sqlbuild

type (
	PlaceholderOptions struct {
		format PlaceholderFormat
		offset int
	}

	PlaceholderOption func(*PlaceholderOptions)
)

// WithOffset init Placeholder with offset
func WithOffset(offset int) PlaceholderOption {
	return func(args *PlaceholderOptions) {
		args.offset = offset
	}
}

// WithFormatPostgres init Placeholder with postgres format
func WithFormatPostgres() PlaceholderOption {
	return func(args *PlaceholderOptions) {
		args.format = FormatPostgres
	}
}

// WithFormatMysql init Placeholder with mysql format
func WithFormatMysql() PlaceholderOption {
	return func(args *PlaceholderOptions) {
		args.format = FormatMysql
	}
}

// WithFormat init Placeholder with custom format
func WithFormat(format string) PlaceholderOption {
	return func(args *PlaceholderOptions) {
		args.format = PlaceholderFormat(format)
	}
}
