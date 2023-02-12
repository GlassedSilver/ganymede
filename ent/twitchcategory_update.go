// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/zibbp/ganymede/ent/predicate"
	"github.com/zibbp/ganymede/ent/twitchcategory"
)

// TwitchCategoryUpdate is the builder for updating TwitchCategory entities.
type TwitchCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *TwitchCategoryMutation
}

// Where appends a list predicates to the TwitchCategoryUpdate builder.
func (tcu *TwitchCategoryUpdate) Where(ps ...predicate.TwitchCategory) *TwitchCategoryUpdate {
	tcu.mutation.Where(ps...)
	return tcu
}

// SetName sets the "name" field.
func (tcu *TwitchCategoryUpdate) SetName(s string) *TwitchCategoryUpdate {
	tcu.mutation.SetName(s)
	return tcu
}

// SetBoxArtURL sets the "box_art_url" field.
func (tcu *TwitchCategoryUpdate) SetBoxArtURL(s string) *TwitchCategoryUpdate {
	tcu.mutation.SetBoxArtURL(s)
	return tcu
}

// SetNillableBoxArtURL sets the "box_art_url" field if the given value is not nil.
func (tcu *TwitchCategoryUpdate) SetNillableBoxArtURL(s *string) *TwitchCategoryUpdate {
	if s != nil {
		tcu.SetBoxArtURL(*s)
	}
	return tcu
}

// ClearBoxArtURL clears the value of the "box_art_url" field.
func (tcu *TwitchCategoryUpdate) ClearBoxArtURL() *TwitchCategoryUpdate {
	tcu.mutation.ClearBoxArtURL()
	return tcu
}

// SetIgdbID sets the "igdb_id" field.
func (tcu *TwitchCategoryUpdate) SetIgdbID(s string) *TwitchCategoryUpdate {
	tcu.mutation.SetIgdbID(s)
	return tcu
}

// SetNillableIgdbID sets the "igdb_id" field if the given value is not nil.
func (tcu *TwitchCategoryUpdate) SetNillableIgdbID(s *string) *TwitchCategoryUpdate {
	if s != nil {
		tcu.SetIgdbID(*s)
	}
	return tcu
}

// ClearIgdbID clears the value of the "igdb_id" field.
func (tcu *TwitchCategoryUpdate) ClearIgdbID() *TwitchCategoryUpdate {
	tcu.mutation.ClearIgdbID()
	return tcu
}

// SetUpdatedAt sets the "updated_at" field.
func (tcu *TwitchCategoryUpdate) SetUpdatedAt(t time.Time) *TwitchCategoryUpdate {
	tcu.mutation.SetUpdatedAt(t)
	return tcu
}

// Mutation returns the TwitchCategoryMutation object of the builder.
func (tcu *TwitchCategoryUpdate) Mutation() *TwitchCategoryMutation {
	return tcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tcu *TwitchCategoryUpdate) Save(ctx context.Context) (int, error) {
	tcu.defaults()
	return withHooks[int, TwitchCategoryMutation](ctx, tcu.sqlSave, tcu.mutation, tcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tcu *TwitchCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := tcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tcu *TwitchCategoryUpdate) Exec(ctx context.Context) error {
	_, err := tcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcu *TwitchCategoryUpdate) ExecX(ctx context.Context) {
	if err := tcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcu *TwitchCategoryUpdate) defaults() {
	if _, ok := tcu.mutation.UpdatedAt(); !ok {
		v := twitchcategory.UpdateDefaultUpdatedAt()
		tcu.mutation.SetUpdatedAt(v)
	}
}

func (tcu *TwitchCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(twitchcategory.Table, twitchcategory.Columns, sqlgraph.NewFieldSpec(twitchcategory.FieldID, field.TypeString))
	if ps := tcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcu.mutation.Name(); ok {
		_spec.SetField(twitchcategory.FieldName, field.TypeString, value)
	}
	if value, ok := tcu.mutation.BoxArtURL(); ok {
		_spec.SetField(twitchcategory.FieldBoxArtURL, field.TypeString, value)
	}
	if tcu.mutation.BoxArtURLCleared() {
		_spec.ClearField(twitchcategory.FieldBoxArtURL, field.TypeString)
	}
	if value, ok := tcu.mutation.IgdbID(); ok {
		_spec.SetField(twitchcategory.FieldIgdbID, field.TypeString, value)
	}
	if tcu.mutation.IgdbIDCleared() {
		_spec.ClearField(twitchcategory.FieldIgdbID, field.TypeString)
	}
	if value, ok := tcu.mutation.UpdatedAt(); ok {
		_spec.SetField(twitchcategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{twitchcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tcu.mutation.done = true
	return n, nil
}

// TwitchCategoryUpdateOne is the builder for updating a single TwitchCategory entity.
type TwitchCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TwitchCategoryMutation
}

// SetName sets the "name" field.
func (tcuo *TwitchCategoryUpdateOne) SetName(s string) *TwitchCategoryUpdateOne {
	tcuo.mutation.SetName(s)
	return tcuo
}

// SetBoxArtURL sets the "box_art_url" field.
func (tcuo *TwitchCategoryUpdateOne) SetBoxArtURL(s string) *TwitchCategoryUpdateOne {
	tcuo.mutation.SetBoxArtURL(s)
	return tcuo
}

// SetNillableBoxArtURL sets the "box_art_url" field if the given value is not nil.
func (tcuo *TwitchCategoryUpdateOne) SetNillableBoxArtURL(s *string) *TwitchCategoryUpdateOne {
	if s != nil {
		tcuo.SetBoxArtURL(*s)
	}
	return tcuo
}

// ClearBoxArtURL clears the value of the "box_art_url" field.
func (tcuo *TwitchCategoryUpdateOne) ClearBoxArtURL() *TwitchCategoryUpdateOne {
	tcuo.mutation.ClearBoxArtURL()
	return tcuo
}

// SetIgdbID sets the "igdb_id" field.
func (tcuo *TwitchCategoryUpdateOne) SetIgdbID(s string) *TwitchCategoryUpdateOne {
	tcuo.mutation.SetIgdbID(s)
	return tcuo
}

// SetNillableIgdbID sets the "igdb_id" field if the given value is not nil.
func (tcuo *TwitchCategoryUpdateOne) SetNillableIgdbID(s *string) *TwitchCategoryUpdateOne {
	if s != nil {
		tcuo.SetIgdbID(*s)
	}
	return tcuo
}

// ClearIgdbID clears the value of the "igdb_id" field.
func (tcuo *TwitchCategoryUpdateOne) ClearIgdbID() *TwitchCategoryUpdateOne {
	tcuo.mutation.ClearIgdbID()
	return tcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tcuo *TwitchCategoryUpdateOne) SetUpdatedAt(t time.Time) *TwitchCategoryUpdateOne {
	tcuo.mutation.SetUpdatedAt(t)
	return tcuo
}

// Mutation returns the TwitchCategoryMutation object of the builder.
func (tcuo *TwitchCategoryUpdateOne) Mutation() *TwitchCategoryMutation {
	return tcuo.mutation
}

// Where appends a list predicates to the TwitchCategoryUpdate builder.
func (tcuo *TwitchCategoryUpdateOne) Where(ps ...predicate.TwitchCategory) *TwitchCategoryUpdateOne {
	tcuo.mutation.Where(ps...)
	return tcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tcuo *TwitchCategoryUpdateOne) Select(field string, fields ...string) *TwitchCategoryUpdateOne {
	tcuo.fields = append([]string{field}, fields...)
	return tcuo
}

// Save executes the query and returns the updated TwitchCategory entity.
func (tcuo *TwitchCategoryUpdateOne) Save(ctx context.Context) (*TwitchCategory, error) {
	tcuo.defaults()
	return withHooks[*TwitchCategory, TwitchCategoryMutation](ctx, tcuo.sqlSave, tcuo.mutation, tcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tcuo *TwitchCategoryUpdateOne) SaveX(ctx context.Context) *TwitchCategory {
	node, err := tcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tcuo *TwitchCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := tcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcuo *TwitchCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := tcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcuo *TwitchCategoryUpdateOne) defaults() {
	if _, ok := tcuo.mutation.UpdatedAt(); !ok {
		v := twitchcategory.UpdateDefaultUpdatedAt()
		tcuo.mutation.SetUpdatedAt(v)
	}
}

func (tcuo *TwitchCategoryUpdateOne) sqlSave(ctx context.Context) (_node *TwitchCategory, err error) {
	_spec := sqlgraph.NewUpdateSpec(twitchcategory.Table, twitchcategory.Columns, sqlgraph.NewFieldSpec(twitchcategory.FieldID, field.TypeString))
	id, ok := tcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TwitchCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, twitchcategory.FieldID)
		for _, f := range fields {
			if !twitchcategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != twitchcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcuo.mutation.Name(); ok {
		_spec.SetField(twitchcategory.FieldName, field.TypeString, value)
	}
	if value, ok := tcuo.mutation.BoxArtURL(); ok {
		_spec.SetField(twitchcategory.FieldBoxArtURL, field.TypeString, value)
	}
	if tcuo.mutation.BoxArtURLCleared() {
		_spec.ClearField(twitchcategory.FieldBoxArtURL, field.TypeString)
	}
	if value, ok := tcuo.mutation.IgdbID(); ok {
		_spec.SetField(twitchcategory.FieldIgdbID, field.TypeString, value)
	}
	if tcuo.mutation.IgdbIDCleared() {
		_spec.ClearField(twitchcategory.FieldIgdbID, field.TypeString)
	}
	if value, ok := tcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(twitchcategory.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &TwitchCategory{config: tcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{twitchcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tcuo.mutation.done = true
	return _node, nil
}
