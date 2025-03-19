// Dummy represents a dummy entity with personal information and transaction dates.
// It contains basic identification and demographic details, along with a history of transaction timestamps.
//
// Fields:
//   - Id: A unique identifier for the dummy entity.
//   - Name: The full name of the dummy entity.
//   - Sex: Gender information of the dummy entity.
//   - BirthDate: The birth date of the dummy entity as a string.
//   - TrxDates: A slice of timestamps representing transaction history.
package dummy

import "time"

type Dummy struct {
	Id        string      `json:"id"`
	Name      string      `json:"name"`
	Sex       string      `json:"sex"`
	BirthDate string      `json:"birthDate"`
	TrxDates  []time.Time `json:"trxDates"`
}
