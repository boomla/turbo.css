package spec

import ()

func (spec *Spec) String() (s string) {
	s += SPEC_HEADER + "\n"

	s += SECTION_TITLE
	s += spec.Title + "\n\n"

	s += SECTION_SHORT
	s += spec.Short + "\n\n"

	s += SECTION_LONG
	s += spec.Long + "\n\n"

	for _, syntax := range spec.Syntaxes {
		s += SECTION_SYNTAX
		s += syntax.Class + "\n"
		s += syntax.Properties + "\n\n"
	}

	for _, arg := range spec.Arguments {
		if len(arg.Types) == 1 {
			// Use dense format when there is only 1 argument
			s += SECTION_ARGUMENT
			s += arg.Name + " " + arg.Types[0].Type + "\n"
			if arg.Description != "" {
				s += arg.Description + "\n"
			}
			s += "\n"
		} else {
			// Use long format when there are multiple arguments
			s += SECTION_ARGUMENT_NAME
			s += arg.Name + "\n"
			if arg.Description != "" {
				s += arg.Description + "\n"
			}
			s += "\n"

			for _, typ := range arg.Types {
				s += SECTION_ARGUMENT_TYPE
				s += typ.Type + "\n"
				if typ.Description != "" {
					s += typ.Description + "\n"
				}
				s += "\n"
			}
		}
	}

	for _, example := range spec.Examples {
		s += example.String() + "\n\n"
	}
	for _, exampleThatFails := range spec.ExamplesThatFail {
		s += exampleThatFails.String() + "\n\n"
	}
	for _, exampleLib := range spec.ExampleLibraries {
		s += exampleLib.String() + "\n\n"
	}
	for _, exampleLibThatFails := range spec.ExampleLibrariesThatFail {
		s += exampleLibThatFails.String() + "\n\n"
	}

	return s
}

func (example Example) String() (s string) {
	s += SECTION_EXAMPLE
	s += SEPARATOR_START
	s += example.Description + "\n"
	s += SEPARATOR_CODE
	s += example.Code + "\n"
	s += SEPARATOR
	s += example.Exp + "\n"
	s += SEPARATOR_END
	return s
}

func (exampleThatFails ExampleThatFails) String() (s string) {
	s += SECTION_EXAMPLE_THAT_FAILS
	s += SEPARATOR_START
	s += exampleThatFails.Description + "\n"
	s += SEPARATOR_CODE
	s += exampleThatFails.Code + "\n"
	s += SEPARATOR
	s += exampleThatFails.ExpErr + "\n"
	s += SEPARATOR_END
	return s
}

func (exampleLib ExampleLibrary) String() (s string) {
	s += SECTION_EXAMPLE_LIBRARY
	s += SEPARATOR_START
	s += exampleLib.Description + "\n"
	s += SEPARATOR_CODE
	for _, lib := range exampleLib.Libraries {
		s += "[" + lib.Filename + "]\n"
		s += lib.Source + "\n"
		s += SEPARATOR
	}
	if exampleLib.GlobalCode != "" {
		s += exampleLib.GlobalCode + "\n"
		s += SEPARATOR
	}
	s += exampleLib.Code + "\n"
	s += SEPARATOR
	s += exampleLib.Exp + "\n"
	s += SEPARATOR_END
	return s
}

func (exampleLib ExampleLibraryThatFails) String() (s string) {
	s += SECTION_EXAMPLE_LIBRARY_THAT_FAILS
	s += SEPARATOR_START
	s += exampleLib.Description + "\n"
	s += SEPARATOR_CODE
	for _, lib := range exampleLib.Libraries {
		s += "[" + lib.Filename + "]\n"
		s += lib.Source + "\n"
		s += SEPARATOR
	}
	if exampleLib.GlobalCode != "" {
		s += exampleLib.GlobalCode + "\n"
		s += SEPARATOR
	}
	s += exampleLib.Code + "\n"
	s += SEPARATOR
	s += exampleLib.ExpErr + "\n"
	s += SEPARATOR_END
	return s
}

