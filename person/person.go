package person

type Person struct {
	FirstName string
	LastName  string
}

func NewPerson(firstName, lastName string) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
	}
}
