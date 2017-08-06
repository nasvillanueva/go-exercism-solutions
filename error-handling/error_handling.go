package erratum

const testVersion = 2

func Use(o ResourceOpener, input string) (err error) {
	resource, err := o()
	if err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(o, input)
		}
		return err
	}
	defer func() {
		r := recover()

		if frobError, ok := r.(FrobError); ok {
			resource.Defrob(frobError.defrobTag)
			err = frobError
		} else if r != nil {
			err = r.(error)
		}

		resource.Close()
	}()

	resource.Frob(input)

	return
}
