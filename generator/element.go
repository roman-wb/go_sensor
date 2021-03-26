package generator

var Fields = []string{}
var Operators = []string{}

type Element struct {
	Tag      string
	Attr     map[string]string
	Children []Element
}

// shared methods (id name)
// specific methods (value optional length...)

func (el Element) IsTemplates() bool {
	return el.Tag == "templates"
}

func (el Element) IsTemplate() bool {
	return el.Tag == "template"
}

func (el Element) IsField() bool {
	return contains(Fields, el.Tag)
}

func (el Element) IsOperator() bool {
	return contains(Operators, el.Tag)
}

func contains(slice []string, search string) bool {
	for _, item := range slice {
		if item == search {
			return true
		}
	}
	return false
}
