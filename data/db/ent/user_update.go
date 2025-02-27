// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-template/data/db/ent/predicate"
	"go-template/data/db/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// SetPasswd sets the "passwd" field.
func (uu *UserUpdate) SetPasswd(s string) *UserUpdate {
	uu.mutation.SetPasswd(s)
	return uu
}

// SetNillablePasswd sets the "passwd" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePasswd(s *string) *UserUpdate {
	if s != nil {
		uu.SetPasswd(*s)
	}
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetPhone sets the "phone" field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.mutation.SetPhone(s)
	return uu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhone(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhone(*s)
	}
	return uu
}

// ClearPhone clears the value of the "phone" field.
func (uu *UserUpdate) ClearPhone() *UserUpdate {
	uu.mutation.ClearPhone()
	return uu
}

// SetRole sets the "role" field.
func (uu *UserUpdate) SetRole(s string) *UserUpdate {
	uu.mutation.SetRole(s)
	return uu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRole(s *string) *UserUpdate {
	if s != nil {
		uu.SetRole(*s)
	}
	return uu
}

// SetIsVip sets the "is_vip" field.
func (uu *UserUpdate) SetIsVip(b bool) *UserUpdate {
	uu.mutation.SetIsVip(b)
	return uu
}

// SetNillableIsVip sets the "is_vip" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsVip(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsVip(*b)
	}
	return uu
}

// SetBalance sets the "balance" field.
func (uu *UserUpdate) SetBalance(f float64) *UserUpdate {
	uu.mutation.ResetBalance()
	uu.mutation.SetBalance(f)
	return uu
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBalance(f *float64) *UserUpdate {
	if f != nil {
		uu.SetBalance(*f)
	}
	return uu
}

// AddBalance adds f to the "balance" field.
func (uu *UserUpdate) AddBalance(f float64) *UserUpdate {
	uu.mutation.AddBalance(f)
	return uu
}

// SetCreateTime sets the "create_time" field.
func (uu *UserUpdate) SetCreateTime(s string) *UserUpdate {
	uu.mutation.SetCreateTime(s)
	return uu
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreateTime(s *string) *UserUpdate {
	if s != nil {
		uu.SetCreateTime(*s)
	}
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Passwd(); ok {
		_spec.SetField(user.FieldPasswd, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if uu.mutation.PhoneCleared() {
		_spec.ClearField(user.FieldPhone, field.TypeString)
	}
	if value, ok := uu.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeString, value)
	}
	if value, ok := uu.mutation.IsVip(); ok {
		_spec.SetField(user.FieldIsVip, field.TypeBool, value)
	}
	if value, ok := uu.mutation.Balance(); ok {
		_spec.SetField(user.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := uu.mutation.AddedBalance(); ok {
		_spec.AddField(user.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := uu.mutation.CreateTime(); ok {
		_spec.SetField(user.FieldCreateTime, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// SetPasswd sets the "passwd" field.
func (uuo *UserUpdateOne) SetPasswd(s string) *UserUpdateOne {
	uuo.mutation.SetPasswd(s)
	return uuo
}

// SetNillablePasswd sets the "passwd" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePasswd(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPasswd(*s)
	}
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetPhone sets the "phone" field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.mutation.SetPhone(s)
	return uuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhone(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhone(*s)
	}
	return uuo
}

// ClearPhone clears the value of the "phone" field.
func (uuo *UserUpdateOne) ClearPhone() *UserUpdateOne {
	uuo.mutation.ClearPhone()
	return uuo
}

// SetRole sets the "role" field.
func (uuo *UserUpdateOne) SetRole(s string) *UserUpdateOne {
	uuo.mutation.SetRole(s)
	return uuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRole(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetRole(*s)
	}
	return uuo
}

// SetIsVip sets the "is_vip" field.
func (uuo *UserUpdateOne) SetIsVip(b bool) *UserUpdateOne {
	uuo.mutation.SetIsVip(b)
	return uuo
}

// SetNillableIsVip sets the "is_vip" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsVip(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsVip(*b)
	}
	return uuo
}

// SetBalance sets the "balance" field.
func (uuo *UserUpdateOne) SetBalance(f float64) *UserUpdateOne {
	uuo.mutation.ResetBalance()
	uuo.mutation.SetBalance(f)
	return uuo
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBalance(f *float64) *UserUpdateOne {
	if f != nil {
		uuo.SetBalance(*f)
	}
	return uuo
}

// AddBalance adds f to the "balance" field.
func (uuo *UserUpdateOne) AddBalance(f float64) *UserUpdateOne {
	uuo.mutation.AddBalance(f)
	return uuo
}

// SetCreateTime sets the "create_time" field.
func (uuo *UserUpdateOne) SetCreateTime(s string) *UserUpdateOne {
	uuo.mutation.SetCreateTime(s)
	return uuo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreateTime(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetCreateTime(*s)
	}
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Passwd(); ok {
		_spec.SetField(user.FieldPasswd, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if uuo.mutation.PhoneCleared() {
		_spec.ClearField(user.FieldPhone, field.TypeString)
	}
	if value, ok := uuo.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeString, value)
	}
	if value, ok := uuo.mutation.IsVip(); ok {
		_spec.SetField(user.FieldIsVip, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.Balance(); ok {
		_spec.SetField(user.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := uuo.mutation.AddedBalance(); ok {
		_spec.AddField(user.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := uuo.mutation.CreateTime(); ok {
		_spec.SetField(user.FieldCreateTime, field.TypeString, value)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
