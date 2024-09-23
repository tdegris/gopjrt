package main

import (
	"fmt"
	"github.com/gomlx/gopjrt/protos/xla_data"
	"github.com/janpfeifer/must"
	"github.com/pkg/errors"
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

// panicf panics with formatted description.
//
// It is only used for "bugs in the code" -- when parameters don't follow the specifications.
// In principle, it should never happen -- the same way nil-pointer panics should never happen.
func panicf(format string, args ...any) {
	panic(errors.Errorf(format, args...))
}

var aliases = map[string]string{
	"INVALID": "InvalidDType",
	"PRED":    "Bool",
	"S8":      "Int8",
	"S16":     "Int16",
	"S32":     "Int32",
	"S64":     "Int64",
	"U8":      "Uint8",
	"U16":     "Uint16",
	"U32":     "Uint32",
	"U64":     "Uint64",
	"F16":     "Float16",
	"BF16":    "BFloat16",
	"F32":     "Float32",
	"F64":     "Float64",
	"C64":     "Complex64",
	"C128":    "Complex128",
}

var (
	reEnums = regexp.MustCompile(
		`(?m)(typedef enum \{\n([^}]+)}` + // Enum definition
			`\s+(PJRT_Buffer_Type)\s*;)`) // Enum type name
	reEnumComment    = regexp.MustCompile(`^\s*(//.*)$`)
	reEnumDefinition = regexp.MustCompile(`^\s*PJRT_Buffer_Type_(\w+)(\s*=\s*(\w+))?\s*,?$`)

	enumsFromCTemplate = template.Must(template.New(DTypeEnumGoFileName).Parse(`
package dtypes

/***** File generated by ./cmd/dtypes_codegen, don't edit it directly. *****/

import "github.com/gomlx/gopjrt/protos/xla_data"

// DType is an enum represents the data type of a buffer or a scalar.
// These are all the types supported by XLA/PJRT.
//
// The names come from the C/C++ constants, so they are not Go idiomatic. 
// The package provides some aliases.
//
// It is unfortunate, but the data types enums used in XLA/PJRT (which DType is modeled after) 
// and in C++ XlaBuilder (and other parts of XLA) don't match. 
// The gopjrt project uses the PJRT enum everywhere, and makes the conversions when needed to call C++ code (see
// DType.PrimitiveType and FromPrimitiveType for conversion).
type DType int32
const ({{range .}}
	// {{.Name}} is a 1:1 mapping of the corresponding C enum value defined in pjrt_c_api.h (as PJRT_Buffer_Type_{{.Original}}). {{range .Comments}}
	{{.}}{{end}}
	{{.Name}} DType = {{.Value}}
{{end}})

// Aliases from PJRT C API.
const ({{range .}}{{if .HasAlias}}
	// {{.Original}} (or PJRT_Buffer_Type_{{.Original}}) is the C enum name for {{.Name}}.
	{{.Original}} = {{.Name}}
{{end}}{{end}})

// MapOfNames to their dtypes. It includes also aliases to the various dtypes. 
// It is also later initialized to include the lower-case version of the names.
var MapOfNames = map[string]DType{
{{range .}}	"{{.Name}}": {{.Name}},
{{if .HasAlias}}	"{{.Original}}": {{.Name}},
{{end}}{{end}}}

// PrimitiveType returns the DType equivalent used in C++ XlaBuilder. 
// For internal use only.
//
// It is unfortunate, but the data types enums used in PJRT (which DType is modeled after) 
// and C++ XlaBuilder (and other parts of XLA) don't match.
func (dtype DType) PrimitiveType() xla_data.PrimitiveType {
	switch (dtype) {
	case InvalidDType:
		return xla_data.PrimitiveType_PRIMITIVE_TYPE_INVALID{{range .}}{{if .HasPrimitiveType}}
	case {{.Name}}:
		return xla_data.PrimitiveType_{{.Original}}{{end}}{{end}}
	default:
		return xla_data.PrimitiveType_PRIMITIVE_TYPE_INVALID
	}
}

// FromPrimitiveType returns the equivalent DType.
// For internal use only.
//
// It is unfortunate, but the data types enums used in PJRT (which DType is modeled after) 
// and C++ XlaBuilder (and other parts of XLA) don't match.
func FromPrimitiveType(primitiveType xla_data.PrimitiveType) DType {
	switch (primitiveType) {
	case xla_data.PrimitiveType_PRIMITIVE_TYPE_INVALID:
		return InvalidDType{{range .}}{{if .HasPrimitiveType}}
	case xla_data.PrimitiveType_{{.Original}}:
		return {{.Name}}{{end}}{{end}}
	default:
		return InvalidDType
	}
}
`))
)

type enumValue struct {
	Name             string
	Comments         []string
	Value            int
	HasAlias         bool
	Original         string
	HasPrimitiveType bool
}

func generateEnums(contents string) {
	var allValues []*enumValue
	var enumV *enumValue
	matches := reEnums.FindStringSubmatch(contents)
	if len(matches) == 0 {
		panicf("failed to match PJRT_Buffer_Types enum from pjrt_c_api.h")
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
		enumV.Original = subMatches[1]
		if v := subMatches[3]; v != "" {
			enumV.Value = int(must.M1(strconv.ParseInt(subMatches[3], 10, 64)))
		} else {
			if len(allValues) == 0 {
				enumV.Value = 0
			} else {
				enumV.Value = allValues[len(allValues)-1].Value + 1
			}
		}

		// Find preferred alias.
		enumV.Name, enumV.HasAlias = aliases[enumV.Original]
		if !enumV.HasAlias {
			enumV.Name = enumV.Original
		}

		// Find PrimitiveType:
		_, enumV.HasPrimitiveType = xla_data.PrimitiveType_value[enumV.Original]

		allValues = append(allValues, enumV)
		enumV = nil
	}

	f := must.M1(os.Create(DTypeEnumGoFileName))
	must.M(enumsFromCTemplate.Execute(f, allValues))
	must.M(exec.Command("gofmt", "-w", DTypeEnumGoFileName).Run())
	fmt.Printf("Generated %q based on pjrt_c_api.h\n", DTypeEnumGoFileName)
	must.M(exec.Command("enumer", "-type=DType", "-yaml", "-json", "-text", "-values", DTypeEnumGoFileName).Run())
}
