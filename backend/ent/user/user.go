// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUSEREMAIL holds the string denoting the user_email field in the database.
	FieldUSEREMAIL = "user_email"
	// FieldUSERNAME holds the string denoting the user_name field in the database.
	FieldUSERNAME = "user_name"

	// EdgeBooking holds the string denoting the booking edge name in mutations.
	EdgeBooking = "booking"
	// EdgeRoleplay holds the string denoting the roleplay edge name in mutations.
	EdgeRoleplay = "roleplay"

	// Table holds the table name of the user in the database.
	Table = "users"
	// BookingTable is the table the holds the booking relation/edge.
	BookingTable = "bookings"
	// BookingInverseTable is the table name for the Booking entity.
	// It exists in this package in order to avoid circular dependency with the "booking" package.
	BookingInverseTable = "bookings"
	// BookingColumn is the table column denoting the booking relation/edge.
	BookingColumn = "USER_ID"
	// RoleplayTable is the table the holds the roleplay relation/edge.
	RoleplayTable = "users"
	// RoleplayInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RoleplayInverseTable = "roles"
	// RoleplayColumn is the table column denoting the roleplay relation/edge.
	RoleplayColumn = "ROLE_ID"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUSEREMAIL,
	FieldUSERNAME,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the User type.
var ForeignKeys = []string{
	"ROLE_ID",
}

var (
	// USEREMAILValidator is a validator for the "USER_EMAIL" field. It is called by the builders before save.
	USEREMAILValidator func(string) error
	// USERNAMEValidator is a validator for the "USER_NAME" field. It is called by the builders before save.
	USERNAMEValidator func(string) error
)
