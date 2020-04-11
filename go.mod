module github.com/yanyi/generate-multifields

go 1.13

require (
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v0.0.7
	github.com/stretchr/testify v1.5.1
)

// Use local fork with relative on-disk location.
replace github.com/yanyi/generate-multifields => ./
