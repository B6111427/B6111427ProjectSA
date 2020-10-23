// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/B6111427/app/ent/booking"
	"github.com/B6111427/app/ent/bookingtype"
	"github.com/B6111427/app/ent/cliententity"
	"github.com/B6111427/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// BookingCreate is the builder for creating a Booking entity.
type BookingCreate struct {
	config
	mutation *BookingMutation
	hooks    []Hook
}

// SetBOOKINGDATE sets the BOOKING_DATE field.
func (bc *BookingCreate) SetBOOKINGDATE(t time.Time) *BookingCreate {
	bc.mutation.SetBOOKINGDATE(t)
	return bc
}

// SetNillableBOOKINGDATE sets the BOOKING_DATE field if the given value is not nil.
func (bc *BookingCreate) SetNillableBOOKINGDATE(t *time.Time) *BookingCreate {
	if t != nil {
		bc.SetBOOKINGDATE(*t)
	}
	return bc
}

// SetTIMELEFT sets the TIME_LEFT field.
func (bc *BookingCreate) SetTIMELEFT(t time.Time) *BookingCreate {
	bc.mutation.SetTIMELEFT(t)
	return bc
}

// SetUsedbyID sets the usedby edge to User by id.
func (bc *BookingCreate) SetUsedbyID(id int) *BookingCreate {
	bc.mutation.SetUsedbyID(id)
	return bc
}

// SetNillableUsedbyID sets the usedby edge to User by id if the given value is not nil.
func (bc *BookingCreate) SetNillableUsedbyID(id *int) *BookingCreate {
	if id != nil {
		bc = bc.SetUsedbyID(*id)
	}
	return bc
}

// SetUsedby sets the usedby edge to User.
func (bc *BookingCreate) SetUsedby(u *User) *BookingCreate {
	return bc.SetUsedbyID(u.ID)
}

// SetBookID sets the book edge to Bookingtype by id.
func (bc *BookingCreate) SetBookID(id int) *BookingCreate {
	bc.mutation.SetBookID(id)
	return bc
}

// SetNillableBookID sets the book edge to Bookingtype by id if the given value is not nil.
func (bc *BookingCreate) SetNillableBookID(id *int) *BookingCreate {
	if id != nil {
		bc = bc.SetBookID(*id)
	}
	return bc
}

// SetBook sets the book edge to Bookingtype.
func (bc *BookingCreate) SetBook(b *Bookingtype) *BookingCreate {
	return bc.SetBookID(b.ID)
}

// SetUsingID sets the using edge to ClientEntity by id.
func (bc *BookingCreate) SetUsingID(id int) *BookingCreate {
	bc.mutation.SetUsingID(id)
	return bc
}

// SetNillableUsingID sets the using edge to ClientEntity by id if the given value is not nil.
func (bc *BookingCreate) SetNillableUsingID(id *int) *BookingCreate {
	if id != nil {
		bc = bc.SetUsingID(*id)
	}
	return bc
}

// SetUsing sets the using edge to ClientEntity.
func (bc *BookingCreate) SetUsing(c *ClientEntity) *BookingCreate {
	return bc.SetUsingID(c.ID)
}

// Mutation returns the BookingMutation object of the builder.
func (bc *BookingCreate) Mutation() *BookingMutation {
	return bc.mutation
}

// Save creates the Booking in the database.
func (bc *BookingCreate) Save(ctx context.Context) (*Booking, error) {
	if _, ok := bc.mutation.BOOKINGDATE(); !ok {
		v := booking.DefaultBOOKINGDATE()
		bc.mutation.SetBOOKINGDATE(v)
	}
	if _, ok := bc.mutation.TIMELEFT(); !ok {
		return nil, &ValidationError{Name: "TIME_LEFT", err: errors.New("ent: missing required field \"TIME_LEFT\"")}
	}
	var (
		err  error
		node *Booking
	)
	if len(bc.hooks) == 0 {
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BookingCreate) SaveX(ctx context.Context) *Booking {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bc *BookingCreate) sqlSave(ctx context.Context) (*Booking, error) {
	b, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	b.ID = int(id)
	return b, nil
}

func (bc *BookingCreate) createSpec() (*Booking, *sqlgraph.CreateSpec) {
	var (
		b     = &Booking{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: booking.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: booking.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.BOOKINGDATE(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: booking.FieldBOOKINGDATE,
		})
		b.BOOKINGDATE = value
	}
	if value, ok := bc.mutation.TIMELEFT(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: booking.FieldTIMELEFT,
		})
		b.TIMELEFT = value
	}
	if nodes := bc.mutation.UsedbyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.UsedbyTable,
			Columns: []string{booking.UsedbyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.BookIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.BookTable,
			Columns: []string{booking.BookColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bookingtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.UsingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.UsingTable,
			Columns: []string{booking.UsingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: cliententity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return b, _spec
}
