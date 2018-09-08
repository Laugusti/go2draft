package pair

// Pair is a parametric type with two type parameters and contains two values of those type
type Pair(type T1, T2) struct {
	l T1
	r T2
}

// New returns a new Pair using the parameters
func New(type Pair(T1, T2))(l T1, r T2) *Pair(T1, T2) {
	return &Pair(T1, T2)(l, r)
}

// Left returns the left value of the Pair
func (p *Pair(T1, T2)) Left() T1 {
	return p.l
}

// Right returns the right value of the Pair
func (p *Pair(T1, T2)) Right() T1 {
	return p.r
}
