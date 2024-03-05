package flags

import (
	"flag"
	"blockchain_learningtool/shared"
)

func ParseFlags() shared.Parameters {
	var params shared.Parameters

	flag.BoolVar(&params.Verbose, "verbose", false, "be verbose")
	flag.StringVar(&params.Namelist, "namelist", "mocacinno,bob,alice,james,laura,david,emily", "comma separated list of demo names")
	flag.IntVar(&params.InputValue, "value", 10000, "Starting capital from coinbase")
	flag.Parse()

	return params
}