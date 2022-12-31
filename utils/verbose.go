package utils

var _verbose bool = false

func SetVerbose(verbose bool) {
	_verbose = verbose
}

func Verbose() bool {
	return _verbose
}
