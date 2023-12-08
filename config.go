package unimorph

import (
	"os"
	"strconv"
)

var Config = struct {
	SortRadixInput bool
	CompressTrees  bool
	DrawTrees      bool
}{}

func init() {
	Config.SortRadixInput = parseEnvBool("SORT_RADIX_INPUT", false)
	Config.CompressTrees = parseEnvBool("COMPRESS_TREES", false)
	Config.DrawTrees = parseEnvBool("DRAW_TREES", false)
}

func parseEnvBool(key string, def bool) bool {
	val, ok := os.LookupEnv(key)

	if !ok {
		return def
	}

	b, err := strconv.ParseBool(val)

	if err != nil {
		return def
	}

	return b
}
