package mega

// Mega represents the advanced Mega system.
type Mega struct {
	a *ComponentA
	b *ComponentB
}

// New returns a new Mega system instance.
func New() *Mega {
	return &Mega{
		a: &ComponentA{},
		b: &ComponentB{},
	}
}

// Run executes Step A then Step B.
func (m *Mega) Run() string {
	return m.a.Execute() + " then " + m.b.Execute()
}
