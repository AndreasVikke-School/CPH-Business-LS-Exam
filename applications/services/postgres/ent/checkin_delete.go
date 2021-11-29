// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/postgres/ent/checkin"
	"github.com/andreasvikke/CPH-Bussines-LS-Exam/applications/services/postgres/ent/predicate"
)

// CheckInDelete is the builder for deleting a CheckIn entity.
type CheckInDelete struct {
	config
	hooks    []Hook
	mutation *CheckInMutation
}

// Where appends a list predicates to the CheckInDelete builder.
func (cid *CheckInDelete) Where(ps ...predicate.CheckIn) *CheckInDelete {
	cid.mutation.Where(ps...)
	return cid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cid *CheckInDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cid.hooks) == 0 {
		affected, err = cid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CheckInMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cid.mutation = mutation
			affected, err = cid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cid.hooks) - 1; i >= 0; i-- {
			if cid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cid *CheckInDelete) ExecX(ctx context.Context) int {
	n, err := cid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cid *CheckInDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: checkin.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: checkin.FieldID,
			},
		},
	}
	if ps := cid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cid.driver, _spec)
}

// CheckInDeleteOne is the builder for deleting a single CheckIn entity.
type CheckInDeleteOne struct {
	cid *CheckInDelete
}

// Exec executes the deletion query.
func (cido *CheckInDeleteOne) Exec(ctx context.Context) error {
	n, err := cido.cid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{checkin.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cido *CheckInDeleteOne) ExecX(ctx context.Context) {
	cido.cid.ExecX(ctx)
}
