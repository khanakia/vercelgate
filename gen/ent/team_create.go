// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khanakia/vercelgate/gen/ent/team"
	"github.com/khanakia/vercelgate/gen/ent/user"
)

// TeamCreate is the builder for creating a Team entity.
type TeamCreate struct {
	config
	mutation *TeamMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (tc *TeamCreate) SetName(s string) *TeamCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tc *TeamCreate) SetNillableName(s *string) *TeamCreate {
	if s != nil {
		tc.SetName(*s)
	}
	return tc
}

// SetUserID sets the "user_id" field.
func (tc *TeamCreate) SetUserID(s string) *TeamCreate {
	tc.mutation.SetUserID(s)
	return tc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tc *TeamCreate) SetNillableUserID(s *string) *TeamCreate {
	if s != nil {
		tc.SetUserID(*s)
	}
	return tc
}

// SetSlug sets the "slug" field.
func (tc *TeamCreate) SetSlug(s string) *TeamCreate {
	tc.mutation.SetSlug(s)
	return tc
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (tc *TeamCreate) SetNillableSlug(s *string) *TeamCreate {
	if s != nil {
		tc.SetSlug(*s)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TeamCreate) SetID(s string) *TeamCreate {
	tc.mutation.SetID(s)
	return tc
}

// SetUser sets the "user" edge to the User entity.
func (tc *TeamCreate) SetUser(u *User) *TeamCreate {
	return tc.SetUserID(u.ID)
}

// Mutation returns the TeamMutation object of the builder.
func (tc *TeamCreate) Mutation() *TeamMutation {
	return tc.mutation
}

// Save creates the Team in the database.
func (tc *TeamCreate) Save(ctx context.Context) (*Team, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TeamCreate) SaveX(ctx context.Context) *Team {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TeamCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TeamCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TeamCreate) check() error {
	return nil
}

func (tc *TeamCreate) sqlSave(ctx context.Context) (*Team, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Team.ID type: %T", _spec.ID.Value)
		}
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TeamCreate) createSpec() (*Team, *sqlgraph.CreateSpec) {
	var (
		_node = &Team{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(team.Table, sqlgraph.NewFieldSpec(team.FieldID, field.TypeString))
	)
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(team.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.Slug(); ok {
		_spec.SetField(team.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if nodes := tc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   team.UserTable,
			Columns: []string{team.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Team.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TeamUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (tc *TeamCreate) OnConflict(opts ...sql.ConflictOption) *TeamUpsertOne {
	tc.conflict = opts
	return &TeamUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Team.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TeamCreate) OnConflictColumns(columns ...string) *TeamUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TeamUpsertOne{
		create: tc,
	}
}

type (
	// TeamUpsertOne is the builder for "upsert"-ing
	//  one Team node.
	TeamUpsertOne struct {
		create *TeamCreate
	}

	// TeamUpsert is the "OnConflict" setter.
	TeamUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *TeamUpsert) SetName(v string) *TeamUpsert {
	u.Set(team.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TeamUpsert) UpdateName() *TeamUpsert {
	u.SetExcluded(team.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *TeamUpsert) ClearName() *TeamUpsert {
	u.SetNull(team.FieldName)
	return u
}

// SetUserID sets the "user_id" field.
func (u *TeamUpsert) SetUserID(v string) *TeamUpsert {
	u.Set(team.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *TeamUpsert) UpdateUserID() *TeamUpsert {
	u.SetExcluded(team.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *TeamUpsert) ClearUserID() *TeamUpsert {
	u.SetNull(team.FieldUserID)
	return u
}

// SetSlug sets the "slug" field.
func (u *TeamUpsert) SetSlug(v string) *TeamUpsert {
	u.Set(team.FieldSlug, v)
	return u
}

// UpdateSlug sets the "slug" field to the value that was provided on create.
func (u *TeamUpsert) UpdateSlug() *TeamUpsert {
	u.SetExcluded(team.FieldSlug)
	return u
}

// ClearSlug clears the value of the "slug" field.
func (u *TeamUpsert) ClearSlug() *TeamUpsert {
	u.SetNull(team.FieldSlug)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Team.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(team.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TeamUpsertOne) UpdateNewValues() *TeamUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(team.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Team.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TeamUpsertOne) Ignore() *TeamUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TeamUpsertOne) DoNothing() *TeamUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TeamCreate.OnConflict
// documentation for more info.
func (u *TeamUpsertOne) Update(set func(*TeamUpsert)) *TeamUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TeamUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *TeamUpsertOne) SetName(v string) *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TeamUpsertOne) UpdateName() *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *TeamUpsertOne) ClearName() *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.ClearName()
	})
}

// SetUserID sets the "user_id" field.
func (u *TeamUpsertOne) SetUserID(v string) *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *TeamUpsertOne) UpdateUserID() *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *TeamUpsertOne) ClearUserID() *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.ClearUserID()
	})
}

// SetSlug sets the "slug" field.
func (u *TeamUpsertOne) SetSlug(v string) *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.SetSlug(v)
	})
}

// UpdateSlug sets the "slug" field to the value that was provided on create.
func (u *TeamUpsertOne) UpdateSlug() *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.UpdateSlug()
	})
}

// ClearSlug clears the value of the "slug" field.
func (u *TeamUpsertOne) ClearSlug() *TeamUpsertOne {
	return u.Update(func(s *TeamUpsert) {
		s.ClearSlug()
	})
}

// Exec executes the query.
func (u *TeamUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TeamCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TeamUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TeamUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TeamUpsertOne.ID is not supported by MySQL driver. Use TeamUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TeamUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TeamCreateBulk is the builder for creating many Team entities in bulk.
type TeamCreateBulk struct {
	config
	err      error
	builders []*TeamCreate
	conflict []sql.ConflictOption
}

// Save creates the Team entities in the database.
func (tcb *TeamCreateBulk) Save(ctx context.Context) ([]*Team, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Team, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TeamMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TeamCreateBulk) SaveX(ctx context.Context) []*Team {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TeamCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TeamCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Team.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TeamUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (tcb *TeamCreateBulk) OnConflict(opts ...sql.ConflictOption) *TeamUpsertBulk {
	tcb.conflict = opts
	return &TeamUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Team.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TeamCreateBulk) OnConflictColumns(columns ...string) *TeamUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TeamUpsertBulk{
		create: tcb,
	}
}

// TeamUpsertBulk is the builder for "upsert"-ing
// a bulk of Team nodes.
type TeamUpsertBulk struct {
	create *TeamCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Team.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(team.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TeamUpsertBulk) UpdateNewValues() *TeamUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(team.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Team.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TeamUpsertBulk) Ignore() *TeamUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TeamUpsertBulk) DoNothing() *TeamUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TeamCreateBulk.OnConflict
// documentation for more info.
func (u *TeamUpsertBulk) Update(set func(*TeamUpsert)) *TeamUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TeamUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *TeamUpsertBulk) SetName(v string) *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *TeamUpsertBulk) UpdateName() *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *TeamUpsertBulk) ClearName() *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.ClearName()
	})
}

// SetUserID sets the "user_id" field.
func (u *TeamUpsertBulk) SetUserID(v string) *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *TeamUpsertBulk) UpdateUserID() *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *TeamUpsertBulk) ClearUserID() *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.ClearUserID()
	})
}

// SetSlug sets the "slug" field.
func (u *TeamUpsertBulk) SetSlug(v string) *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.SetSlug(v)
	})
}

// UpdateSlug sets the "slug" field to the value that was provided on create.
func (u *TeamUpsertBulk) UpdateSlug() *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.UpdateSlug()
	})
}

// ClearSlug clears the value of the "slug" field.
func (u *TeamUpsertBulk) ClearSlug() *TeamUpsertBulk {
	return u.Update(func(s *TeamUpsert) {
		s.ClearSlug()
	})
}

// Exec executes the query.
func (u *TeamUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TeamCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TeamCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TeamUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
