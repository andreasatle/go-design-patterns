package hospital

// Patient contains information about the patient and what service has been done.
type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckupDone bool
	medicineDone      bool
	paymentDone       bool
}

// NewPatient creates a new patient.
func NewPatient(name string) *Patient {
	return &Patient{
		name: name,
	}
}
