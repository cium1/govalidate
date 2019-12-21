package govalidate

// Error struct
type Error struct {
	field        string
	fieldData    interface{}
	fieldAlias   string
	rule         string
	ruleArgs     interface{}
	errorMessage string
}

// GetField get field
func (e *Error) GetField() string {
	return e.field
}

// GetFieldData get field data
func (e *Error) GetFieldData() interface{} {
	return e.fieldData
}

// GetFieldAlias get field alias
func (e *Error) GetFieldAlias() string {
	return e.fieldAlias
}

// GetRule get rule
func (e *Error) GetRule() string {
	return e.rule
}

// GetErrorMessage get error message
func (e *Error) GetErrorMessage() string {
	return e.errorMessage
}

// GetRuleArg get rule arg
func (e *Error) GetRuleArg() interface{} {
	return e.ruleArgs
}
