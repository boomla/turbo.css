package spec

import (
	"testing"
)

func TestSpec_String(t *testing.T) {
	spec := &Spec{
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
	}

	act := spec.String()
	exp := TEST_SPEC
	if act != exp {
		t.Fatalf("\nexp [%v]\n\nact [%v]", exp, act)
	}
}

const TEST_SPEC =
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


`


func TestSpec_Example(t *testing.T) {
	example := Example{
		Description: "Passing argument of type number.",
		Code: "m-1",
		Exp: ""+
			".m-1 {\n"+
			"\tmargin: 1px;\n"+
			"}",
	}

	act := example.String()
	exp := TEST_EXAMPLE
	if act != exp {
		t.Fatalf("\nexp [%v]\n\nact [%v]", exp, act)
	}
}

const TEST_EXAMPLE =
`# EXAMPLE
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
Passing argument of type number.
==================================================
m-1
--------------------------------------------------
.m-1 {
	margin: 1px;
}
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
`


func TestSpec_ExampleLibrary(t *testing.T) {
	example := ExampleLibrary{
		Description: "Description",
		Libraries: []Library{
			{
				Filename: "Filename.1",
				Source: "Source 1",
			},
			{
				Filename: "Filename.2",
				Source: "Source 2",
			},
		},
		GlobalCode: "GlobalCode",
		Code: "Code",
		Exp: "Exp",
	}

	act := example.String()
	exp := TEST_EXAMPLE_LIBRARY
	if act != exp {
		t.Fatalf("\nexp [%v]\n\nact [%v]", exp, act)
	}
}

const TEST_EXAMPLE_LIBRARY =
`# EXAMPLE LIBRARY
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
Description
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
`


func TestSpec_ExampleLibraryThatFails(t *testing.T) {
	example := ExampleLibraryThatFails{
		Description: "Description",
		Libraries: []Library{
			{
				Filename: "Filename.1",
				Source: "Source 1",
			},
			{
				Filename: "Filename.2",
				Source: "Source 2",
			},
		},
		GlobalCode: "GlobalCode",
		Code: "Code",
		ExpErr: "ExpErr",
	}

	act := example.String()
	exp := TEST_EXAMPLE_LIBRARY_THAT_FAILS
	if act != exp {
		t.Fatalf("\nexp [%v]\n\nact [%v]", exp, act)
	}
}

const TEST_EXAMPLE_LIBRARY_THAT_FAILS =
`# EXAMPLE LIBRARY THAT FAILS
<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
Description
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
ExpErr
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
`


