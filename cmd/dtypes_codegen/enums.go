package main

import (
	"fmt"
	"github.com/gomlx/exceptions"
	"github.com/janpfeifer/must"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

const (
	DTypeEnumGoFileName = "gen_dtype_enum.go"
)

var (
	reEnums = regexp.MustCompile(
		`(?m)(typedef enum \{\n([^}]+)}` + // Enum definition
			`\s+(PJRT_Buffer_Type)\s*;)`) // Enum type name
	reEnumComment    = regexp.MustCompile(`^\s*(//.*)$`)
	reEnumDefinition = regexp.MustCompile(`^\s*PJRT_Buffer_Type_(\w+)(\s*=\s*(\w+))?\s*,?$`)

	enumsFromCTemplate = template.Must(template.New(DTypeEnumGoFileName).Parse(`
package dtypes

/***** File generated by ./cmd/dtypes_codegen, don't edit it directly. *****/

// DType is an enum represents the data type of a buffer or a scalar.
// These are all the types supported by XLA.
//
// The names come from the C/C++ constants, so they are not Go idiomatic. 
// The package provides some aliases.
type DType int32
const ({{range .}}
	// {{.Name}} is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h (as PJRT_Buffer_Type_{{.Name}}). {{range .Comments}}
	{{.}}{{end}}
	{{.Name}} DType = {{.Value}}
{{end}})
`))
)

type enumValue struct {
	Name     string
	Comments []string
	Value    int
}

func generateEnums(contents string) {
	var allValues []*enumValue
	var enumV *enumValue
	matches := reEnums.FindStringSubmatch(contents)
	if len(matches) == 0 {
		exceptions.Panicf("failed to match PJRT_Buffer_Types enum from pjrt_c_api.h")
	}
	for _, line := range strings.Split(matches[2], "\n") {
		if line == "" {
			continue
		}
		if enumV == nil {
			enumV = &enumValue{}
		}
		if subMatches := reEnumComment.FindStringSubmatch(line); len(subMatches) > 0 {
			enumV.Comments = append(enumV.Comments, subMatches[1])
			continue
		}
		subMatches := reEnumDefinition.FindStringSubmatch(line)
		if len(subMatches) == 0 {
			continue
		}
		enumV.Name = subMatches[1]
		if v := subMatches[3]; v != "" {
			enumV.Value = int(must.M1(strconv.ParseInt(subMatches[3], 10, 64)))
		} else {
			if len(allValues) == 0 {
				enumV.Value = 0
			} else {
				enumV.Value = allValues[len(allValues)-1].Value + 1
			}
		}
		allValues = append(allValues, enumV)
		enumV = nil
	}

	f := must.M1(os.Create(DTypeEnumGoFileName))
	must.M(enumsFromCTemplate.Execute(f, allValues))
	must.M(exec.Command("gofmt", "-w", DTypeEnumGoFileName).Run())
	fmt.Printf("Generated %q based on pjrt_c_api.h\n", DTypeEnumGoFileName)
}
