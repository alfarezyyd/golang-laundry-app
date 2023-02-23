package exception

type ProgramError struct {
	Error string
}

func NewProgramError(error string) *ProgramError {
	return &ProgramError{Error: error}
}

func ResponseIfProgramError(err error) {
	if err != nil {
		panic(NewProgramError(err.Error()))
	}
}
