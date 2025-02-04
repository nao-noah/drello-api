// Code generated by entc, DO NOT EDIT.

package workspace

const (
	// Label holds the string label denoting the workspace type in the database.
	Label = "workspace"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// Table holds the table name of the workspace in the database.
	Table = "workspaces"
)

// Columns holds all SQL columns for workspace fields.
var Columns = []string{
	FieldID,
	FieldTitle,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
