package main

type Doc struct {
	ID       int
	Keywords []string
}

func BuildInvertIndex(docs []*Doc) map[string][]int {
	idxMap := make(map[string][]int)
	for _, doc := range docs {
		for _, keyword := range doc.Keywords {
			idxMap[keyword] = append(idxMap[keyword], doc.ID)
		}
	}

	return idxMap
}
