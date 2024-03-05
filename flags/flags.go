package flags

import (
	"flag"
	"blockchain_learningtool/shared"
)

func ParseFlags()  {
	var params shared.Parameters

	flag.BoolVar(&params.Verbose, "verbose", false, "be verbose")
	flag.StringVar(&params.Namelist, "namelist", "mocacinno,bob,alice,james,laura,david,emily", "comma separated list of demo names")
	flag.IntVar(&params.InputValue, "value", 10000, "Starting capital from coinbase")
	flag.IntVar(&params.NumberOfBlocks, "nbblocks", 10, "number of blocks to generate")
	flag.Parse()

	shared.Myparameters = params
}