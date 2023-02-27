package swagger

import (
	"encoding/json"
	"github.com/getkin/kin-openapi/openapi2"
	"os"
	"strconv"
)

func Convert(docDir string) {
	input, err := os.ReadFile(docDir)
	if err != nil {
		panic(err)
	}

	var doc openapi2.T
	if err = json.Unmarshal(input, &doc); err != nil {
		panic(err)
	}

	refCnt := make(map[string]int)
	for _, path := range doc.Paths {
		for _, op := range path.Operations() {
			for _, resp := range op.Responses {
				if resp.Schema == nil || resp.Schema.Value == nil || resp.Schema.Value.AllOf == nil {
					continue
				}
				ref := resp.Schema.Value.AllOf[0].Ref
				refCnt[ref]++
				resp.Schema.Value.AllOf[0].Ref = ref + strconv.Itoa(refCnt[ref])
			}
		}
	}

	for refName, ref := range refCnt {
		refNameOri := refName[len("#/definitions/"):]
		for name, schemaRef := range doc.Definitions {
			if name == refNameOri {
				for i := 1; i <= ref; i++ {
					doc.Definitions[refNameOri+strconv.Itoa(i)] = schemaRef
				}
			}
		}
	}

	outputJSON, err := json.MarshalIndent(doc, "", "    ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(docDir, outputJSON, 0644)
	if err != nil {
		panic(err)
	}
}
