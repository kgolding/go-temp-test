package scope

type PagerType byte

func (pt PagerType) String() string {
	switch pt {
	case 'A':
		return "alphanumeric"
	case 'N':
		return "numeric"
	}
	return string(pt)
}

func (pt PagerType) IsNumeric() bool {
	return pt == 'N'
}

func (pt PagerType) IsAlphanumeric() bool {
	return pt == 'A'
}
