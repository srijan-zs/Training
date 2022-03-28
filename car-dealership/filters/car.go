package filters

// Car struct is defined as a filter to extract cars of specific brand and whether engine is included or not
type Car struct {
	IncludeEngine bool
	Brand         string
}
