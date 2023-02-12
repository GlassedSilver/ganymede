// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/zibbp/ganymede/ent/channel"
	"github.com/zibbp/ganymede/ent/predicate"
)

// ChannelDelete is the builder for deleting a Channel entity.
type ChannelDelete struct {
	config
	hooks    []Hook
	mutation *ChannelMutation
}

// Where appends a list predicates to the ChannelDelete builder.
func (cd *ChannelDelete) Where(ps ...predicate.Channel) *ChannelDelete {
	cd.mutation.Where(ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *ChannelDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ChannelMutation](ctx, cd.sqlExec, cd.mutation, cd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *ChannelDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *ChannelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(channel.Table, sqlgraph.NewFieldSpec(channel.FieldID, field.TypeUUID))
	if ps := cd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cd.mutation.done = true
	return affected, err
}

// ChannelDeleteOne is the builder for deleting a single Channel entity.
type ChannelDeleteOne struct {
	cd *ChannelDelete
}

// Where appends a list predicates to the ChannelDelete builder.
func (cdo *ChannelDeleteOne) Where(ps ...predicate.Channel) *ChannelDeleteOne {
	cdo.cd.mutation.Where(ps...)
	return cdo
}

// Exec executes the deletion query.
func (cdo *ChannelDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{channel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *ChannelDeleteOne) ExecX(ctx context.Context) {
	if err := cdo.Exec(ctx); err != nil {
		panic(err)
	}
}
