package types

// Assignment is a struct that holds the name of an assignment and the name of the profile it assigns.
type Assignment struct {
	// Name is the name given to the Assignment resource itself.
	Name string
	// ProfileName is the name of the assigned profile.
	ProfileName string
}
