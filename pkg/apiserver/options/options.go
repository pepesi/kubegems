package options

type Option interface {
	Apply(map[string]interface{})
}

type FieldEqual struct {
	field string
	value interface{}
}

func (fe *FieldEqual) Apply(condtion map[string]interface{}) {
	condtion[fe.field] = fe.value
}

func Equal(field string, value interface{}) *FieldEqual {
	return &FieldEqual{
		field: field,
		value: value,
	}
}

type FieldIn struct {
	field string
	value interface{}
}

func (fi *FieldIn) Apply(condtion map[string]interface{}) {
	condtion[fi.field+" in"] = fi.value
}

func In(field string, value interface{}) *FieldIn {
	return &FieldIn{
		field: field,
		value: value,
	}
}
