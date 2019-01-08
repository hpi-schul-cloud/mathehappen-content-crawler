package crawler

import (
	"github.com/schul-cloud/mathehappen-content-crawler/mathehappen"
)

// Crawl downloads the content, transforms it,
// and posts it to the target server.
func Crawl() {
	dataIn := downloadContent()
	dataOut := mathehappen.Transform(dataIn)
	postContent(dataOut)
}
