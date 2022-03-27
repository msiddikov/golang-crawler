package linkValidator

type (
	result struct {
		internals []urlInfo
		externals []urlInfo
	}
	urlInfo struct {
		url    string
		origin string
		valid  bool
		title  string
	}
)
