package selected

type SelectedField struct {
	Name     string
	Args     map[string]interface{}
	Selected []SelectedField
}
