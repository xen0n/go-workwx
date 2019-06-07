package main

import (
	"errors"
	"strings"

	"gopkg.in/russross/blackfriday.v2"
)

var (
	errOnlyOneTopicAllowed = errors.New("only one topic allowed per document")
	errToplevelTopicNotH1  = errors.New("the top-level topic must be h1")
	errTopicChildNotH2     = errors.New("the children sections of a topic must be h2")
	errUnknownTopicChild   = errors.New("unknown child section of topic")
	errModelDefNotH3       = errors.New("model definition header must be h3")

	errMultipleModelTables    = errors.New("only one table allowed per model")
	errUnknownFieldTableTitle = errors.New("unknown column title of field table")
)

func analyzeDocument(doc *mdTocNode) (hir, error) {
	empty := hir{}

	// currently only 1 topic is allowed per md file
	if len(doc.TocChildren) != 1 {
		return empty, errOnlyOneTopicAllowed
	}

	result := hir{}
	for _, n := range doc.TocChildren {
		topic, err := analyzeH1(n)
		if err != nil {
			return empty, err
		}
		result.topics = append(result.topics, topic)
	}

	return result, nil
}

func analyzeH1(doc *mdTocNode) (topic, error) {
	empty := topic{}

	if doc.Level != 1 {
		return empty, errToplevelTopicNotH1
	}

	result := topic{}
	for _, n := range doc.TocChildren {
		// fmt.Printf("H2: %s\n", n.ThisText())
		title := n.ThisInnerText()
		title = strings.TrimSpace(title)

		switch title {
		case "Models":
			models, err := analyzeH2Models(n)
			if err != nil {
				return empty, err
			}
			result.models = models

		default:
			return empty, errUnknownTopicChild
		}
	}
	return result, nil
}

func analyzeH2Models(doc *mdTocNode) ([]apiModel, error) {
	if doc.Level != 2 {
		return nil, errTopicChildNotH2
	}

	result := make([]apiModel, 0)
	for _, n := range doc.TocChildren {
		model, err := analyzeH3Model(n)
		if err != nil {
			return nil, err
		}
		result = append(result, model)
	}
	return result, nil
}

func analyzeH3Model(doc *mdTocNode) (apiModel, error) {
	empty := apiModel{}

	if doc.Level != 3 {
		return empty, errModelDefNotH3
	}

	result := apiModel{
		vis: visibilityPublic,
	}

	// model metadata
	var docSb strings.Builder
	for _, n := range doc.ThisContent {
		switch n.ThisType() {
		case blackfriday.Code:
			// ident
			result.ident = n.ThisLit()

		case blackfriday.Text:
			docSb.WriteString(n.ThisLit())

		default:
			// ignore
		}
	}
	result.doc = strings.TrimSpace(docSb.String())

	// model fields
	// currently only one table is allowed
	seenTable := false
	for _, n := range doc.Content {
		switch n.ThisType() {
		case blackfriday.Table:
			if seenTable {
				return empty, errMultipleModelTables
			}
			seenTable = true

			fields, err := analyzeModelFieldTable(n)
			if err != nil {
				return empty, err
			}
			result.fields = fields

		default:
			// ignore for now
			// TODO: allow collecting paragraphs into doc comments
		}
	}
	return result, nil
}

func analyzeModelFieldTable(tbl *mdContentNode) ([]apiModelField, error) {
	// initially mark the columns as non-existent
	var idxIdent int = -1
	var idxType int = -1
	var idxDesc int = -1
	var idxTagJson int = -1

	result := make([]apiModelField, 0)

	// TODO: disallow multiple header rows
	for _, n := range tbl.ThisContent {
		switch n.ThisType() {
		case blackfriday.TableHead:
			// only look at the first row
			tr := n.ThisContent[0]

			// parse out the column titles
			for i, td := range tr.ThisContent {
				colTitle := strings.ToLower(td.ThisInnerText())
				switch colTitle {
				case "name":
					idxIdent = i
				case "type":
					idxType = i
				case "doc":
					idxDesc = i
				case "json":
					idxTagJson = i
				default:
					return nil, errUnknownFieldTableTitle
				}
			}

		case blackfriday.TableBody:
			// parse the fields
			for _, tr := range n.ThisContent {
				field := apiModelField{
					vis:  visibilityPublic,
					tags: make(map[string]string),
				}
				for i, td := range tr.ThisContent {
					if i == idxIdent {
						for _, n2 := range td.ThisContent {
							switch n2.ThisType() {
							case blackfriday.Code:
								field.ident = n2.ThisLit()

							default:
								// ignored
							}
						}
					}

					if i == idxType {
						for _, n2 := range td.ThisContent {
							switch n2.ThisType() {
							case blackfriday.Code:
								field.typ = n2.ThisLit()

							default:
								// ignored
							}
						}
					}

					if i == idxDesc {
						// I'm too lazy
						field.doc = td.ThisInnerText()
					}

					if i == idxTagJson {
						for _, n2 := range td.ThisContent {
							switch n2.ThisType() {
							case blackfriday.Code:
								field.tags["json"] = n2.ThisLit()

							default:
								// ignored
							}
						}
					}
				}

				result = append(result, field)
			}
		}
	}

	return result, nil
}