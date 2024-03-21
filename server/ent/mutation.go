// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"melp-back/ent/pnu"
	"melp-back/ent/predicate"
	"melp-back/ent/schema"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypePnu = "Pnu"
)

// PnuMutation represents an operation that mutates the Pnu nodes in the graph.
type PnuMutation struct {
	config
	op            Op
	typ           string
	id            *string
	jibun         *string
	bchk          *string
	pnu           *string
	col_adm_se    *string
	geometry      **schema.Geometry
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Pnu, error)
	predicates    []predicate.Pnu
}

var _ ent.Mutation = (*PnuMutation)(nil)

// pnuOption allows management of the mutation configuration using functional options.
type pnuOption func(*PnuMutation)

// newPnuMutation creates new mutation for the Pnu entity.
func newPnuMutation(c config, op Op, opts ...pnuOption) *PnuMutation {
	m := &PnuMutation{
		config:        c,
		op:            op,
		typ:           TypePnu,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPnuID sets the ID field of the mutation.
func withPnuID(id string) pnuOption {
	return func(m *PnuMutation) {
		var (
			err   error
			once  sync.Once
			value *Pnu
		)
		m.oldValue = func(ctx context.Context) (*Pnu, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Pnu.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPnu sets the old Pnu of the mutation.
func withPnu(node *Pnu) pnuOption {
	return func(m *PnuMutation) {
		m.oldValue = func(context.Context) (*Pnu, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PnuMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PnuMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Pnu entities.
func (m *PnuMutation) SetID(id string) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PnuMutation) ID() (id string, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PnuMutation) IDs(ctx context.Context) ([]string, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []string{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Pnu.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetJibun sets the "jibun" field.
func (m *PnuMutation) SetJibun(s string) {
	m.jibun = &s
}

// Jibun returns the value of the "jibun" field in the mutation.
func (m *PnuMutation) Jibun() (r string, exists bool) {
	v := m.jibun
	if v == nil {
		return
	}
	return *v, true
}

// OldJibun returns the old "jibun" field's value of the Pnu entity.
// If the Pnu object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PnuMutation) OldJibun(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldJibun is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldJibun requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldJibun: %w", err)
	}
	return oldValue.Jibun, nil
}

// ResetJibun resets all changes to the "jibun" field.
func (m *PnuMutation) ResetJibun() {
	m.jibun = nil
}

// SetBchk sets the "bchk" field.
func (m *PnuMutation) SetBchk(s string) {
	m.bchk = &s
}

// Bchk returns the value of the "bchk" field in the mutation.
func (m *PnuMutation) Bchk() (r string, exists bool) {
	v := m.bchk
	if v == nil {
		return
	}
	return *v, true
}

// OldBchk returns the old "bchk" field's value of the Pnu entity.
// If the Pnu object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PnuMutation) OldBchk(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBchk is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBchk requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBchk: %w", err)
	}
	return oldValue.Bchk, nil
}

// ResetBchk resets all changes to the "bchk" field.
func (m *PnuMutation) ResetBchk() {
	m.bchk = nil
}

// SetPnu sets the "pnu" field.
func (m *PnuMutation) SetPnu(s string) {
	m.pnu = &s
}

// Pnu returns the value of the "pnu" field in the mutation.
func (m *PnuMutation) Pnu() (r string, exists bool) {
	v := m.pnu
	if v == nil {
		return
	}
	return *v, true
}

// OldPnu returns the old "pnu" field's value of the Pnu entity.
// If the Pnu object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PnuMutation) OldPnu(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPnu is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPnu requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPnu: %w", err)
	}
	return oldValue.Pnu, nil
}

// ResetPnu resets all changes to the "pnu" field.
func (m *PnuMutation) ResetPnu() {
	m.pnu = nil
}

// SetColAdmSe sets the "col_adm_se" field.
func (m *PnuMutation) SetColAdmSe(s string) {
	m.col_adm_se = &s
}

// ColAdmSe returns the value of the "col_adm_se" field in the mutation.
func (m *PnuMutation) ColAdmSe() (r string, exists bool) {
	v := m.col_adm_se
	if v == nil {
		return
	}
	return *v, true
}

// OldColAdmSe returns the old "col_adm_se" field's value of the Pnu entity.
// If the Pnu object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PnuMutation) OldColAdmSe(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldColAdmSe is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldColAdmSe requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldColAdmSe: %w", err)
	}
	return oldValue.ColAdmSe, nil
}

// ResetColAdmSe resets all changes to the "col_adm_se" field.
func (m *PnuMutation) ResetColAdmSe() {
	m.col_adm_se = nil
}

// SetGeometry sets the "geometry" field.
func (m *PnuMutation) SetGeometry(s *schema.Geometry) {
	m.geometry = &s
}

// Geometry returns the value of the "geometry" field in the mutation.
func (m *PnuMutation) Geometry() (r *schema.Geometry, exists bool) {
	v := m.geometry
	if v == nil {
		return
	}
	return *v, true
}

// OldGeometry returns the old "geometry" field's value of the Pnu entity.
// If the Pnu object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PnuMutation) OldGeometry(ctx context.Context) (v *schema.Geometry, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGeometry is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGeometry requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGeometry: %w", err)
	}
	return oldValue.Geometry, nil
}

// ResetGeometry resets all changes to the "geometry" field.
func (m *PnuMutation) ResetGeometry() {
	m.geometry = nil
}

// Where appends a list predicates to the PnuMutation builder.
func (m *PnuMutation) Where(ps ...predicate.Pnu) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the PnuMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *PnuMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Pnu, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *PnuMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *PnuMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Pnu).
func (m *PnuMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PnuMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.jibun != nil {
		fields = append(fields, pnu.FieldJibun)
	}
	if m.bchk != nil {
		fields = append(fields, pnu.FieldBchk)
	}
	if m.pnu != nil {
		fields = append(fields, pnu.FieldPnu)
	}
	if m.col_adm_se != nil {
		fields = append(fields, pnu.FieldColAdmSe)
	}
	if m.geometry != nil {
		fields = append(fields, pnu.FieldGeometry)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PnuMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case pnu.FieldJibun:
		return m.Jibun()
	case pnu.FieldBchk:
		return m.Bchk()
	case pnu.FieldPnu:
		return m.Pnu()
	case pnu.FieldColAdmSe:
		return m.ColAdmSe()
	case pnu.FieldGeometry:
		return m.Geometry()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PnuMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case pnu.FieldJibun:
		return m.OldJibun(ctx)
	case pnu.FieldBchk:
		return m.OldBchk(ctx)
	case pnu.FieldPnu:
		return m.OldPnu(ctx)
	case pnu.FieldColAdmSe:
		return m.OldColAdmSe(ctx)
	case pnu.FieldGeometry:
		return m.OldGeometry(ctx)
	}
	return nil, fmt.Errorf("unknown Pnu field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PnuMutation) SetField(name string, value ent.Value) error {
	switch name {
	case pnu.FieldJibun:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetJibun(v)
		return nil
	case pnu.FieldBchk:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBchk(v)
		return nil
	case pnu.FieldPnu:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPnu(v)
		return nil
	case pnu.FieldColAdmSe:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetColAdmSe(v)
		return nil
	case pnu.FieldGeometry:
		v, ok := value.(*schema.Geometry)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGeometry(v)
		return nil
	}
	return fmt.Errorf("unknown Pnu field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PnuMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PnuMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PnuMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Pnu numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PnuMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PnuMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PnuMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Pnu nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PnuMutation) ResetField(name string) error {
	switch name {
	case pnu.FieldJibun:
		m.ResetJibun()
		return nil
	case pnu.FieldBchk:
		m.ResetBchk()
		return nil
	case pnu.FieldPnu:
		m.ResetPnu()
		return nil
	case pnu.FieldColAdmSe:
		m.ResetColAdmSe()
		return nil
	case pnu.FieldGeometry:
		m.ResetGeometry()
		return nil
	}
	return fmt.Errorf("unknown Pnu field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PnuMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PnuMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PnuMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PnuMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PnuMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PnuMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PnuMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Pnu unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PnuMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Pnu edge %s", name)
}