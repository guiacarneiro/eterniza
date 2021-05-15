package producao

// UnityType ...
type UnityType int

const (
	Quantity UnityType = 0
	Size     UnityType = 1
	Weight   UnityType = 2
)

func (nodeType UnityType) String() string {
	names := [...]string{
		"Quantity",
		"Size",
		"Weight"}
	if nodeType < Quantity || nodeType > Weight {
		return "Unknown"
	}
	return names[nodeType]
}

type MateriaPrima struct {
	ID    string    `json:"id"`
	Label string    `json:"label"`
	Unity UnityType `json:"unity"`
	Value float64	`json:"value"`
}
