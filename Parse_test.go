package spec

import (
	"testing"
	"reflect"
)

func TestParse_fail_missingHeader(t *testing.T) {
	_, err := Parse("")
	if err == nil {
		t.Fatal("expected an error")
	}
	expErr := "invalid Turbo spec, missing header [# TURBO-SPEC-FORMAT-V1\n]"
	actErr := err.Error()
	if expErr != actErr {
		t.Fatalf("\nexp [%v]\nact [%v]", expErr, actErr)
	}
}
func TestParse_ok(t *testing.T) {
	act, err := Parse(TEST_SPEC_BODY)
	if err != nil { t.Fatal(err) }

	exp := &Spec{
		Title: "Title",
		Short: "Short",
		Long: "Long\nmultiline",
		Syntaxes: []Syntax{
			Syntax{
				Class: "m-{value}",
				Properties: "margin: {value};",
			},
			Syntax{
				Class: "mx-{value}",
				Properties: "margin-left: {value};\nmargin-right: {value};",
			},
		},
		Arguments: []Argument{
			Argument{
				Name: "dense",
				Description: "This argument accepts a single type.",
				Types: []ArgumentType{
					ArgumentType{
						Type: "<string>",
						Description: "",
					},
				},
			},
			Argument{
				Name: "value",
				Description: "Argument specific description.\nmultiline",
				Types: []ArgumentType{
					ArgumentType{
						Type: "<number>",
						Description: "Type specific description\nmultiline",
					},
					ArgumentType{
						Type: "<length>",
						Description: "",
					},
				},
			},
		},
		Examples: []Example{
			Example{
				Description: "Passing argument of type number.",
				Code: "m-1",
				Exp: ""+
					".m-1 {\n"+
					"\tmargin: 1px;\n"+
					"}",
			},
			Example{
				Description: "Passing argument of type number with fraction.",
				Code: "m-1.5",
				Exp: ""+
					".m-1\\.5 {\n"+
					"\tmargin: 1.5px;\n"+
					"}",
			},
		},
		ExamplesThatFail: []ExampleThatFails{
			ExampleThatFails{
				Description: "Invalid number of arguments",
				Code: "m-1-2",
				ExpErr: "Error: utility function [m-] expects 1 argument, found 2 in [m-1-2]",
			},
			ExampleThatFails{
				Description: "Invalid argument of type <length>",
				Code: "z-1px",
				ExpErr: "Error: utility function [z-] expects argument [1] to be of type [<integer>], found [<length>] in [z-1px]",
			},
		},
		ExampleLibraries: []ExampleLibrary{
			ExampleLibrary{
				Description: "Example library",
				Libraries: []Library{
					Library{
						Filename: "Filename.1",
						Source: "Source 1",
					},
					Library{
						Filename: "Filename.2",
						Source: "Source 2",
					},
				},
				GlobalCode: "GlobalCode",
				Code: "Code",
				Exp: "Exp",
			},
			ExampleLibrary{
				Description: "Example library",
				Libraries: []Library{},
				GlobalCode: "Global library code",
				Code: "Code",
				Exp: "Exp",
			},
		},
		ExampleLibrariesThatFail: []ExampleLibraryThatFails{
			ExampleLibraryThatFails{
				Description: "Example library",
				Libraries: []Library{
					Library{
						Filename: "Filename.1",
						Source: "Source 1",
					},
					Library{
						Filename: "Filename.2",
						Source: "Source 2",
					},
				},
				GlobalCode: "GlobalCode",
				Code: "Code",
				ExpErr: "Exp error",
			},
			ExampleLibraryThatFails{
				Description: "Example library",
				Libraries: []Library{},
				GlobalCode: "Global library code",
				Code: "Code",
				ExpErr: "Exp error",
			},
		},
	}

	if ! reflect.DeepEqual(exp, act) {
		t.Fatalf("\nexp [%v]\n\nact [%v]", exp, act)
	}
}

const TEST_SPEC_BODY =
`# TURBO-SPEC-FORMAT-V1

# TITLE
Title

# SHORT
Short

# LONG
Long
multiline


# SYNTAX
m-{value}
margin: {value};

# SYNTAX
mx-{value}
margin-left: {value};
margin-right: {value};

# ARGUMENT
dense <string>
This argument accepts a single type.

# ARGUMENT NAME
value
Argument specific description.
multiline

# ARGUMENT TYPE
<number>
Type specific description
multiline

# ARGUMENT TYPE
<length>


# EXAMPLE
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
Passing argument of type number.
==================================================
m-1
--------------------------------------------------
.m-1 {
	margin: 1px;
}
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>


# EXAMPLE
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
Passing argument of type number with fraction.
==================================================
m-1.5
--------------------------------------------------
.m-1\.5 {
	margin: 1.5px;
}
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>


# EXAMPLE THAT FAILS
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
Invalid number of arguments
==================================================
m-1-2
--------------------------------------------------
Error: utility function [m-] expects 1 argument, found 2 in [m-1-2]
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>


# EXAMPLE THAT FAILS
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
Invalid argument of type <length>
==================================================
z-1px
--------------------------------------------------
Error: utility function [z-] expects argument [1] to be of type [<integer>], found [<length>] in [z-1px]
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>


# EXAMPLE LIBRARY
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
Example library
==================================================
[Filename.1]
Source 1
--------------------------------------------------
[Filename.2]
Source 2
--------------------------------------------------
GlobalCode
--------------------------------------------------
Code
--------------------------------------------------
Exp
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>


# EXAMPLE LIBRARY
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
Example library
==================================================
Global library code
--------------------------------------------------
Code
--------------------------------------------------
Exp
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>


# EXAMPLE LIBRARY THAT FAILS
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
Example library
==================================================
[Filename.1]
Source 1
--------------------------------------------------
[Filename.2]
Source 2
--------------------------------------------------
GlobalCode
--------------------------------------------------
Code
--------------------------------------------------
Exp error
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>


# EXAMPLE LIBRARY THAT FAILS
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
Example library
==================================================
Global library code
--------------------------------------------------
Code
--------------------------------------------------
Exp error
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>


`

