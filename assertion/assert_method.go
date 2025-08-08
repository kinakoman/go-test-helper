package assertion

func (a *Assert) IsErrNil(err error, message string) {
	if err != nil {
		a.t.Errorf("Expected no error, but got: %s. %s", err.Error(), message)
	}
}

func (a *Assert) IsTrue(condition bool, message string) {
	if !condition {
		a.t.Errorf("Expected condition to be true, but it was false. %s", message)
	}
}

func (a *Assert) AreEqual(expected, actual interface{}, message string) {
	if expected != actual {
		a.t.Errorf("Expected %v, but got %v. %s", expected, actual, message)
	}
}

func (a *Assert) IsNotNil(object interface{}, message string) {
	if object == nil {
		a.t.Errorf("Expected object to be not nil, but it was nil. %s", message)
	}
}

func (a *Assert) NotEmpty(s string, message string) {
	if s == "" {
		a.t.Errorf("Expected string to be not empty, but it was empty. %s", message)
	}
}
