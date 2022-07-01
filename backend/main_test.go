package main

import (
	"github.com/ssst0n3/source-analysis-system/test/test_config"
	"testing"
)

func Test_main(t *testing.T) {
	test_config.Init()
	main()
}
