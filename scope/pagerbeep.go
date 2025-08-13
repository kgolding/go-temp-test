package scope

type PagerBeep byte

func (pb PagerBeep) String() string {
	return string(pb)
}

func (pb PagerBeep) IsValid() bool {
	return pb >= 'A' && pb <= 'D'
}
