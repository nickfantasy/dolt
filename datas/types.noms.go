// This file was generated by nomdl/codegen.

package datas

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __datasPackageInFile_types_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.Type{
		types.MakeStructType("Commit",
			[]types.Field{
				types.Field{"value", types.MakePrimitiveType(types.ValueKind), false},
				types.Field{"parents", types.MakeCompoundType(types.SetKind, types.MakeCompoundType(types.RefKind, types.MakeType(ref.Ref{}, 0))), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	__datasPackageInFile_types_CachedRef = types.RegisterPackage(&p)
}

// Commit

type Commit struct {
	_value   types.Value
	_parents SetOfRefOfCommit

	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewCommit(cs chunks.ChunkStore) Commit {
	return Commit{
		_value:   types.Bool(false),
		_parents: NewSetOfRefOfCommit(cs),

		cs:  cs,
		ref: &ref.Ref{},
	}
}

type CommitDef struct {
	Value   types.Value
	Parents SetOfRefOfCommitDef
}

func (def CommitDef) New(cs chunks.ChunkStore) Commit {
	return Commit{
		_value:   def.Value,
		_parents: def.Parents.New(cs),
		cs:       cs,
		ref:      &ref.Ref{},
	}
}

func (s Commit) Def() (d CommitDef) {
	d.Value = s._value
	d.Parents = s._parents.Def()
	return
}

var __typeForCommit types.Type

func (m Commit) Type() types.Type {
	return __typeForCommit
}

func init() {
	__typeForCommit = types.MakeType(__datasPackageInFile_types_CachedRef, 0)
	types.RegisterStruct(__typeForCommit, builderForCommit, readerForCommit)
}

func builderForCommit(cs chunks.ChunkStore, values []types.Value) types.Value {
	i := 0
	s := Commit{ref: &ref.Ref{}, cs: cs}
	s._value = values[i]
	i++
	s._parents = values[i].(SetOfRefOfCommit)
	i++
	return s
}

func readerForCommit(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(Commit)
	values = append(values, s._value)
	values = append(values, s._parents)
	return values
}

func (s Commit) Equals(other types.Value) bool {
	return other != nil && __typeForCommit.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s Commit) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Commit) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeForCommit.Chunks()...)
	chunks = append(chunks, s._value.Chunks()...)
	chunks = append(chunks, s._parents.Chunks()...)
	return
}

func (s Commit) ChildValues() (ret []types.Value) {
	ret = append(ret, s._value)
	ret = append(ret, s._parents)
	return
}

func (s Commit) Value() types.Value {
	return s._value
}

func (s Commit) SetValue(val types.Value) Commit {
	s._value = val
	s.ref = &ref.Ref{}
	return s
}

func (s Commit) Parents() SetOfRefOfCommit {
	return s._parents
}

func (s Commit) SetParents(val SetOfRefOfCommit) Commit {
	s._parents = val
	s.ref = &ref.Ref{}
	return s
}

// MapOfStringToRefOfCommit

type MapOfStringToRefOfCommit struct {
	m   types.Map
	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewMapOfStringToRefOfCommit(cs chunks.ChunkStore) MapOfStringToRefOfCommit {
	return MapOfStringToRefOfCommit{types.NewTypedMap(cs, __typeForMapOfStringToRefOfCommit), cs, &ref.Ref{}}
}

type MapOfStringToRefOfCommitDef map[string]ref.Ref

func (def MapOfStringToRefOfCommitDef) New(cs chunks.ChunkStore) MapOfStringToRefOfCommit {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), NewRefOfCommit(v))
	}
	return MapOfStringToRefOfCommit{types.NewTypedMap(cs, __typeForMapOfStringToRefOfCommit, kv...), cs, &ref.Ref{}}
}

func (m MapOfStringToRefOfCommit) Def() MapOfStringToRefOfCommitDef {
	def := make(map[string]ref.Ref)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v.(RefOfCommit).TargetRef()
		return false
	})
	return def
}

func (m MapOfStringToRefOfCommit) Equals(other types.Value) bool {
	return other != nil && __typeForMapOfStringToRefOfCommit.Equals(other.Type()) && m.Ref() == other.Ref()
}

func (m MapOfStringToRefOfCommit) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfStringToRefOfCommit) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, m.Type().Chunks()...)
	chunks = append(chunks, m.m.Chunks()...)
	return
}

func (m MapOfStringToRefOfCommit) ChildValues() []types.Value {
	return append([]types.Value{}, m.m.ChildValues()...)
}

// A Noms Value that describes MapOfStringToRefOfCommit.
var __typeForMapOfStringToRefOfCommit types.Type

func (m MapOfStringToRefOfCommit) Type() types.Type {
	return __typeForMapOfStringToRefOfCommit
}

func init() {
	__typeForMapOfStringToRefOfCommit = types.MakeCompoundType(types.MapKind, types.MakePrimitiveType(types.StringKind), types.MakeCompoundType(types.RefKind, types.MakeType(__datasPackageInFile_types_CachedRef, 0)))
	types.RegisterValue(__typeForMapOfStringToRefOfCommit, builderForMapOfStringToRefOfCommit, readerForMapOfStringToRefOfCommit)
}

func builderForMapOfStringToRefOfCommit(cs chunks.ChunkStore, v types.Value) types.Value {
	return MapOfStringToRefOfCommit{v.(types.Map), cs, &ref.Ref{}}
}

func readerForMapOfStringToRefOfCommit(v types.Value) types.Value {
	return v.(MapOfStringToRefOfCommit).m
}

func (m MapOfStringToRefOfCommit) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToRefOfCommit) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToRefOfCommit) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToRefOfCommit) Get(p string) RefOfCommit {
	return m.m.Get(types.NewString(p)).(RefOfCommit)
}

func (m MapOfStringToRefOfCommit) MaybeGet(p string) (RefOfCommit, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return NewRefOfCommit(ref.Ref{}), false
	}
	return v.(RefOfCommit), ok
}

func (m MapOfStringToRefOfCommit) Set(k string, v RefOfCommit) MapOfStringToRefOfCommit {
	return MapOfStringToRefOfCommit{m.m.Set(types.NewString(k), v), m.cs, &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfStringToRefOfCommit) Remove(p string) MapOfStringToRefOfCommit {
	return MapOfStringToRefOfCommit{m.m.Remove(types.NewString(p)), m.cs, &ref.Ref{}}
}

type MapOfStringToRefOfCommitIterCallback func(k string, v RefOfCommit) (stop bool)

func (m MapOfStringToRefOfCommit) Iter(cb MapOfStringToRefOfCommitIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(RefOfCommit))
	})
}

type MapOfStringToRefOfCommitIterAllCallback func(k string, v RefOfCommit)

func (m MapOfStringToRefOfCommit) IterAll(cb MapOfStringToRefOfCommitIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v.(RefOfCommit))
	})
}

func (m MapOfStringToRefOfCommit) IterAllP(concurrency int, cb MapOfStringToRefOfCommitIterAllCallback) {
	m.m.IterAllP(concurrency, func(k, v types.Value) {
		cb(k.(types.String).String(), v.(RefOfCommit))
	})
}

type MapOfStringToRefOfCommitFilterCallback func(k string, v RefOfCommit) (keep bool)

func (m MapOfStringToRefOfCommit) Filter(cb MapOfStringToRefOfCommitFilterCallback) MapOfStringToRefOfCommit {
	out := m.m.Filter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(RefOfCommit))
	})
	return MapOfStringToRefOfCommit{out, m.cs, &ref.Ref{}}
}

// SetOfRefOfCommit

type SetOfRefOfCommit struct {
	s   types.Set
	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewSetOfRefOfCommit(cs chunks.ChunkStore) SetOfRefOfCommit {
	return SetOfRefOfCommit{types.NewTypedSet(cs, __typeForSetOfRefOfCommit), cs, &ref.Ref{}}
}

type SetOfRefOfCommitDef map[ref.Ref]bool

func (def SetOfRefOfCommitDef) New(cs chunks.ChunkStore) SetOfRefOfCommit {
	l := make([]types.Value, len(def))
	i := 0
	for d, _ := range def {
		l[i] = NewRefOfCommit(d)
		i++
	}
	return SetOfRefOfCommit{types.NewTypedSet(cs, __typeForSetOfRefOfCommit, l...), cs, &ref.Ref{}}
}

func (s SetOfRefOfCommit) Def() SetOfRefOfCommitDef {
	def := make(map[ref.Ref]bool, s.Len())
	s.s.Iter(func(v types.Value) bool {
		def[v.(RefOfCommit).TargetRef()] = true
		return false
	})
	return def
}

func (s SetOfRefOfCommit) Equals(other types.Value) bool {
	return other != nil && __typeForSetOfRefOfCommit.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s SetOfRefOfCommit) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s SetOfRefOfCommit) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, s.Type().Chunks()...)
	chunks = append(chunks, s.s.Chunks()...)
	return
}

func (s SetOfRefOfCommit) ChildValues() []types.Value {
	return append([]types.Value{}, s.s.ChildValues()...)
}

// A Noms Value that describes SetOfRefOfCommit.
var __typeForSetOfRefOfCommit types.Type

func (m SetOfRefOfCommit) Type() types.Type {
	return __typeForSetOfRefOfCommit
}

func init() {
	__typeForSetOfRefOfCommit = types.MakeCompoundType(types.SetKind, types.MakeCompoundType(types.RefKind, types.MakeType(__datasPackageInFile_types_CachedRef, 0)))
	types.RegisterValue(__typeForSetOfRefOfCommit, builderForSetOfRefOfCommit, readerForSetOfRefOfCommit)
}

func builderForSetOfRefOfCommit(cs chunks.ChunkStore, v types.Value) types.Value {
	return SetOfRefOfCommit{v.(types.Set), cs, &ref.Ref{}}
}

func readerForSetOfRefOfCommit(v types.Value) types.Value {
	return v.(SetOfRefOfCommit).s
}

func (s SetOfRefOfCommit) Empty() bool {
	return s.s.Empty()
}

func (s SetOfRefOfCommit) Len() uint64 {
	return s.s.Len()
}

func (s SetOfRefOfCommit) Has(p RefOfCommit) bool {
	return s.s.Has(p)
}

type SetOfRefOfCommitIterCallback func(p RefOfCommit) (stop bool)

func (s SetOfRefOfCommit) Iter(cb SetOfRefOfCommitIterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(v.(RefOfCommit))
	})
}

type SetOfRefOfCommitIterAllCallback func(p RefOfCommit)

func (s SetOfRefOfCommit) IterAll(cb SetOfRefOfCommitIterAllCallback) {
	s.s.IterAll(func(v types.Value) {
		cb(v.(RefOfCommit))
	})
}

func (s SetOfRefOfCommit) IterAllP(concurrency int, cb SetOfRefOfCommitIterAllCallback) {
	s.s.IterAllP(concurrency, func(v types.Value) {
		cb(v.(RefOfCommit))
	})
}

type SetOfRefOfCommitFilterCallback func(p RefOfCommit) (keep bool)

func (s SetOfRefOfCommit) Filter(cb SetOfRefOfCommitFilterCallback) SetOfRefOfCommit {
	out := s.s.Filter(func(v types.Value) bool {
		return cb(v.(RefOfCommit))
	})
	return SetOfRefOfCommit{out, s.cs, &ref.Ref{}}
}

func (s SetOfRefOfCommit) Insert(p ...RefOfCommit) SetOfRefOfCommit {
	return SetOfRefOfCommit{s.s.Insert(s.fromElemSlice(p)...), s.cs, &ref.Ref{}}
}

func (s SetOfRefOfCommit) Remove(p ...RefOfCommit) SetOfRefOfCommit {
	return SetOfRefOfCommit{s.s.Remove(s.fromElemSlice(p)...), s.cs, &ref.Ref{}}
}

func (s SetOfRefOfCommit) Union(others ...SetOfRefOfCommit) SetOfRefOfCommit {
	return SetOfRefOfCommit{s.s.Union(s.fromStructSlice(others)...), s.cs, &ref.Ref{}}
}

func (s SetOfRefOfCommit) First() RefOfCommit {
	return s.s.First().(RefOfCommit)
}

func (s SetOfRefOfCommit) fromStructSlice(p []SetOfRefOfCommit) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfRefOfCommit) fromElemSlice(p []RefOfCommit) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

// RefOfCommit

type RefOfCommit struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfCommit(target ref.Ref) RefOfCommit {
	return RefOfCommit{target, &ref.Ref{}}
}

func (r RefOfCommit) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfCommit) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfCommit) Equals(other types.Value) bool {
	return other != nil && __typeForRefOfCommit.Equals(other.Type()) && r.Ref() == other.Ref()
}

func (r RefOfCommit) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.Type().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

func (r RefOfCommit) ChildValues() []types.Value {
	return nil
}

// A Noms Value that describes RefOfCommit.
var __typeForRefOfCommit types.Type

func (r RefOfCommit) Type() types.Type {
	return __typeForRefOfCommit
}

func (r RefOfCommit) Less(other types.OrderedValue) bool {
	return r.TargetRef().Less(other.(types.RefBase).TargetRef())
}

func init() {
	__typeForRefOfCommit = types.MakeCompoundType(types.RefKind, types.MakeType(__datasPackageInFile_types_CachedRef, 0))
	types.RegisterRef(__typeForRefOfCommit, builderForRefOfCommit)
}

func builderForRefOfCommit(r ref.Ref) types.Value {
	return NewRefOfCommit(r)
}

func (r RefOfCommit) TargetValue(cs chunks.ChunkStore) Commit {
	return types.ReadValue(r.target, cs).(Commit)
}

func (r RefOfCommit) SetTargetValue(val Commit, cs chunks.ChunkSink) RefOfCommit {
	return NewRefOfCommit(types.WriteValue(val, cs))
}
