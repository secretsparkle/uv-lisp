package structs

type Control struct {
	Args []string
	Body []string
}

type Function struct {
	Name     string
	Args     []string
	Bindings map[string]string
	Body     SExpression
}

type SExpression struct {
	Data        bool
	SExpression []interface{}
}
