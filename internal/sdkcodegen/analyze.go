package main

func analyzeDocument(doc *mdTocNode) {
	for _, n := range doc.TocChildren {
		_ = n
	}

}

func analyzeH1(doc *mdTocNode) {}
