// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/gen/ent/team"
	"app/gen/ent/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uc *UserCreate) SetNillableName(s *string) *UserCreate {
	if s != nil {
		uc.SetName(*s)
	}
	return uc
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uc *UserCreate) SetNillableUsername(s *string) *UserCreate {
	if s != nil {
		uc.SetUsername(*s)
	}
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uc *UserCreate) SetNillableEmail(s *string) *UserCreate {
	if s != nil {
		uc.SetEmail(*s)
	}
	return uc
}

// SetToken sets the "token" field.
func (uc *UserCreate) SetToken(s string) *UserCreate {
	uc.mutation.SetToken(s)
	return uc
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (uc *UserCreate) SetNillableToken(s *string) *UserCreate {
	if s != nil {
		uc.SetToken(*s)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(s string) *UserCreate {
	uc.mutation.SetID(s)
	return uc
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (uc *UserCreate) AddTeamIDs(ids ...string) *UserCreate {
	uc.mutation.AddTeamIDs(ids...)
	return uc
}

// AddTeams adds the "teams" edges to the Team entity.
func (uc *UserCreate) AddTeams(t ...*Team) *UserCreate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uc.AddTeamIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected User.ID type: %T", _spec.ID.Value)
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	)
	_spec.OnConflict = uc.conflict
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.Token(); ok {
		_spec.SetField(user.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if nodes := uc.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TeamsTable,
			Columns: []string{user.TeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (uc *UserCreate) OnConflict(opts ...sql.ConflictOption) *UserUpsertOne {
	uc.conflict = opts
	return &UserUpsertOne{
		create: uc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uc *UserCreate) OnConflictColumns(columns ...string) *UserUpsertOne {
	uc.conflict = append(uc.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertOne{
		create: uc,
	}
}

type (
	// UserUpsertOne is the builder for "upsert"-ing
	//  one User node.
	UserUpsertOne struct {
		create *UserCreate
	}

	// UserUpsert is the "OnConflict" setter.
	UserUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *UserUpsert) SetName(v string) *UserUpsert {
	u.Set(user.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsert) UpdateName() *UserUpsert {
	u.SetExcluded(user.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *UserUpsert) ClearName() *UserUpsert {
	u.SetNull(user.FieldName)
	return u
}

// SetUsername sets the "username" field.
func (u *UserUpsert) SetUsername(v string) *UserUpsert {
	u.Set(user.FieldUsername, v)
	return u
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *UserUpsert) UpdateUsername() *UserUpsert {
	u.SetExcluded(user.FieldUsername)
	return u
}

// ClearUsername clears the value of the "username" field.
func (u *UserUpsert) ClearUsername() *UserUpsert {
	u.SetNull(user.FieldUsername)
	return u
}

// SetEmail sets the "email" field.
func (u *UserUpsert) SetEmail(v string) *UserUpsert {
	u.Set(user.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsert) UpdateEmail() *UserUpsert {
	u.SetExcluded(user.FieldEmail)
	return u
}

// ClearEmail clears the value of the "email" field.
func (u *UserUpsert) ClearEmail() *UserUpsert {
	u.SetNull(user.FieldEmail)
	return u
}

// SetToken sets the "token" field.
func (u *UserUpsert) SetToken(v string) *UserUpsert {
	u.Set(user.FieldToken, v)
	return u
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *UserUpsert) UpdateToken() *UserUpsert {
	u.SetExcluded(user.FieldToken)
	return u
}

// ClearToken clears the value of the "token" field.
func (u *UserUpsert) ClearToken() *UserUpsert {
	u.SetNull(user.FieldToken)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertOne) UpdateNewValues() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(user.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserUpsertOne) Ignore() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertOne) DoNothing() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreate.OnConflict
// documentation for more info.
func (u *UserUpsertOne) Update(set func(*UserUpsert)) *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *UserUpsertOne) SetName(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateName() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *UserUpsertOne) ClearName() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearName()
	})
}

// SetUsername sets the "username" field.
func (u *UserUpsertOne) SetUsername(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateUsername() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUsername()
	})
}

// ClearUsername clears the value of the "username" field.
func (u *UserUpsertOne) ClearUsername() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearUsername()
	})
}

// SetEmail sets the "email" field.
func (u *UserUpsertOne) SetEmail(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateEmail() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *UserUpsertOne) ClearEmail() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearEmail()
	})
}

// SetToken sets the "token" field.
func (u *UserUpsertOne) SetToken(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetToken(v)
	})
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateToken() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateToken()
	})
}

// ClearToken clears the value of the "token" field.
func (u *UserUpsertOne) ClearToken() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearToken()
	})
}

// Exec executes the query.
func (u *UserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: UserUpsertOne.ID is not supported by MySQL driver. Use UserUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
	conflict []sql.ConflictOption
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserUpsertBulk {
	ucb.conflict = opts
	return &UserUpsertBulk{
		create: ucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflictColumns(columns ...string) *UserUpsertBulk {
	ucb.conflict = append(ucb.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertBulk{
		create: ucb,
	}
}

// UserUpsertBulk is the builder for "upsert"-ing
// a bulk of User nodes.
type UserUpsertBulk struct {
	create *UserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertBulk) UpdateNewValues() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(user.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserUpsertBulk) Ignore() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertBulk) DoNothing() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreateBulk.OnConflict
// documentation for more info.
func (u *UserUpsertBulk) Update(set func(*UserUpsert)) *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *UserUpsertBulk) SetName(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateName() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *UserUpsertBulk) ClearName() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearName()
	})
}

// SetUsername sets the "username" field.
func (u *UserUpsertBulk) SetUsername(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateUsername() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUsername()
	})
}

// ClearUsername clears the value of the "username" field.
func (u *UserUpsertBulk) ClearUsername() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearUsername()
	})
}

// SetEmail sets the "email" field.
func (u *UserUpsertBulk) SetEmail(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateEmail() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *UserUpsertBulk) ClearEmail() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearEmail()
	})
}

// SetToken sets the "token" field.
func (u *UserUpsertBulk) SetToken(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetToken(v)
	})
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateToken() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateToken()
	})
}

// ClearToken clears the value of the "token" field.
func (u *UserUpsertBulk) ClearToken() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearToken()
	})
}

// Exec executes the query.
func (u *UserUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
