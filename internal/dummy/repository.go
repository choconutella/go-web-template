// GetDummyData retrieves dummy data from the database based on the provided ID.
// It executes a SQL query with the given ID parameter and scans the result into a Dummy struct.
//
// Parameters:
//   - ctx: The context for the database query, which can be used for cancellation or timeout.
//   - id: A string representing the unique identifier for the dummy data.
//
// Returns:
//   - *Dummy: A pointer to the Dummy struct containing the retrieved data.
//   - error: An error if the query fails or if no data is found for the given ID.
//
// The function will return an error if:
//   - The database query fails.
//   - No rows are found for the given ID.
//   - Row scanning fails.
package dummy

import (
	"context"
	"database/sql"
	"fmt"
)

type CummulativeRepository struct {
	Db *sql.DB
}

func (cr *CummulativeRepository) GetDummyData(ctx context.Context, id string) (*Dummy, error) {
	query := `your query here where id = :id`
	row := cr.Db.QueryRowContext(ctx, query, sql.Named("id", id))
	var dummy Dummy
	if err := row.Scan(&dummy.Name, &dummy.Sex, &dummy.BirthDate); err != nil {
		return nil, fmt.Errorf("fail to get dummy info for id %s: %v", id, err)
	}
	return &dummy, nil
}
