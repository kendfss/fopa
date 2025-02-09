package internal

const SourceFilePath = "testdata/src.htm"

func Map[Have, Want any](fn func(Have) Want, args ...Have) []Want {
	out := make([]Want, len(args))
	for i, elem := range args {
		out[i] = fn(elem)
	}
	return out
}
