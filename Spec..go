package spec

import ()

type Spec struct {
	Title string
	Short string
	Long string
	Syntaxes []Syntax
	Arguments []Argument
	Examples []Example
	ExamplesThatFail []ExampleThatFails
	ExampleLibraries []ExampleLibrary
	ExampleLibrariesThatFail []ExampleLibraryThatFails
}

type Syntax struct {
	Class string
	Properties string
}

type Argument struct {
	Name string
	Description string
	Types []ArgumentType
}

type ArgumentType struct {
	Type string
	Description string
}

type Example struct {
	Description string
	Code string
	Exp string
}

type ExampleThatFails struct {
	Description string
	Code string
	ExpErr string
}

type Library struct {
	Filename string
	Source string
}

type ExampleLibrary struct {
	Description string
	Libraries []Library
	GlobalCode string
	Code string
	Exp string
}

type ExampleLibraryThatFails struct {
	Description string
	Libraries []Library
	GlobalCode string
	Code string
	ExpErr string
}

