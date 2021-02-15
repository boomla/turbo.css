package spec

import (
	"strings"
	"fmt"
)

func Parse(specBody string) (*Spec, error) {
	s := strings.ReplaceAll(specBody, "\r", "")

	// HEADER
	if ! strings.HasPrefix(s, SPEC_HEADER) {
		return nil, fmt.Errorf("invalid Turbo spec, missing header ["+SPEC_HEADER+"]")
	}
	s = strings.TrimPrefix(s, SPEC_HEADER)

	spec := Spec{}

	sections := strings.Split(s, "\n# ")
	for _, section := range sections {
		section = strings.TrimSpace(section)
		if section == "" {
			continue
		}

		line, sectionBody, _ := eatNonEmptyLine(section)
		sectionType := "# " + line + "\n"
		switch sectionType {
			case SECTION_TITLE:   spec.Title = sectionBody
			case SECTION_SHORT:   spec.Short = sectionBody
			case SECTION_LONG:    spec.Long = sectionBody
			case SECTION_SYNTAX: {
				pos := strings.Index(sectionBody, "\n")
				if pos == -1 {
					return nil, fmt.Errorf("invalid Turbo spec syntax section ["+sectionBody+"]")
				}

				spec.Syntaxes = append(spec.Syntaxes, Syntax{
					Class: sectionBody[:pos],
					Properties: sectionBody[pos+1:],
				})
			}
			case SECTION_ARGUMENT: {
				// Separate name+type and description
				nameType, description, _ := eatNonEmptyLine(sectionBody)

				// Separate name and type
				pos := strings.Index(nameType, " ")
				if pos < 0 {
					return nil, fmt.Errorf("invalid Turbo spec syntax section ["+SECTION_ARGUMENT+"], missing space between argument name and type")
				}
				name := nameType[:pos]
				typ := nameType[pos+1:]
				
				spec.Arguments = append(spec.Arguments, Argument{
					Name: name,
					Description: description,
					Types: []ArgumentType{
						ArgumentType{
							Type: typ,
							Description: "",
						},
					},
				})
			}
			case SECTION_ARGUMENT_NAME: {
				// Separate name and description
				name, description, _ := eatNonEmptyLine(sectionBody)

				spec.Arguments = append(spec.Arguments, Argument{
					Name: name,
					Description: description,
					Types: []ArgumentType{},
				})
			}
			case SECTION_ARGUMENT_TYPE: {
				// Separate type and description
				typ, description, _ := eatNonEmptyLine(sectionBody)

				index := len(spec.Arguments) - 1
				if index < 0 {
					return nil, fmt.Errorf("invalid Turbo spec syntax, section ["+SECTION_ARGUMENT_TYPE+"] must come after a section ["+SECTION_ARGUMENT_NAME+"]")
				}

				spec.Arguments[index].Types = append(spec.Arguments[index].Types, ArgumentType{
					Type: typ,
					Description: description,
				})
			}
			case SECTION_EXAMPLE: {
				parts := splitExampleSection(sectionBody)
				expLen := 3
				if len(parts) != expLen {
					return nil, fmt.Errorf("invalid Turbo spec syntax, section ["+sectionType+"] contains unexpected number of sub-sections [%v], expected [%v]", len(parts), expLen)
				}
				spec.Examples = append(spec.Examples, Example{
					Description: parts[0],
					Code: parts[1],
					Exp: parts[2],
				})
			}
			case SECTION_EXAMPLE_THAT_FAILS: {
				parts := splitExampleSection(sectionBody)
				expLen := 3
				if len(parts) != expLen {
					return nil, fmt.Errorf("invalid Turbo spec syntax, section ["+sectionType+"] contains unexpected number of sub-sections [%v], expected [%v]", len(parts), expLen)
				}
				spec.ExamplesThatFail = append(spec.ExamplesThatFail, ExampleThatFails{
					Description: parts[0],
					Code: parts[1],
					ExpErr: parts[2],
				})
			}
			case SECTION_EXAMPLE_LIBRARY: {
				parts := splitExampleSection(sectionBody)
				expMinLen := 4
				if len(parts) < expMinLen {
					return nil, fmt.Errorf("invalid Turbo spec syntax, section ["+sectionType+"] contains unexpected number of sub-sections [%v], expected minimum [%v]", len(parts), expMinLen)
				}

				end := len(parts)

				libSubSegments := parts[1:end-3]
				libraries := []Library{}
				for _, libCode := range libSubSegments {
					libraries = append(libraries, parseLib(libCode))
				}

				spec.ExampleLibraries = append(spec.ExampleLibraries, ExampleLibrary{
					Description: parts[0],
					Libraries: libraries,
					GlobalCode: parts[end-3],
					Code: parts[end-2],
					Exp: parts[end-1],
				})
			}
			case SECTION_EXAMPLE_LIBRARY_THAT_FAILS: {
				parts := splitExampleSection(sectionBody)
				expMinLen := 4
				if len(parts) < expMinLen {
					return nil, fmt.Errorf("invalid Turbo spec syntax, section ["+sectionType+"] contains unexpected number of sub-sections [%v], expected minimum [%v]", len(parts), expMinLen)
				}

				end := len(parts)

				libSubSegments := parts[1:end-3]
				libraries := []Library{}
				for _, libCode := range libSubSegments {
					libraries = append(libraries, parseLib(libCode))
				}

				spec.ExampleLibrariesThatFail = append(spec.ExampleLibrariesThatFail, ExampleLibraryThatFails{
					Description: parts[0],
					Libraries: libraries,
					GlobalCode: parts[end-3],
					Code: parts[end-2],
					ExpErr: parts[end-1],
				})
			}
			default: {
				return nil, fmt.Errorf("unexpected section type ["+sectionType+"]")
			}
		}
	}

	return &spec, nil
}

func eatNonEmptyLine(s string) (line, remainder string, end bool) {
	s = strings.TrimSpace(s)

	pos := strings.Index(s, "\n")
	if pos == -1 {
		return s, "", true
	}

	return s[:pos], s[pos+1:], false
}

func splitExampleSection(sectionBody string) ([]string) {
	sectionBody += "\n"
	sectionBody = strings.Replace(sectionBody, SEPARATOR_START, SEPARATOR, -1)
	sectionBody = strings.Replace(sectionBody, SEPARATOR_CODE, SEPARATOR, -1)
	sectionBody = strings.Replace(sectionBody, SEPARATOR_END, SEPARATOR, -1)

	parts := strings.Split(sectionBody, SEPARATOR)
	parts = parts[1:len(parts)-1]
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}

	return parts
}

func parseLib(libCode string) (Library) {
	if ! strings.HasPrefix(libCode, "[") {
		return Library{
			Filename: "",
			Source: libCode,
		}
	}
	pos := strings.Index(libCode, "]\n")
	if pos < 0 {
		return Library{
			Filename: "",
			Source: libCode,
		}
	}

	return Library{
		Filename: libCode[1:pos],
		Source: libCode[pos+2:],
	}
}

