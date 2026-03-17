package math

type Vector2 struct {
	X float32
	Y float32
}

func NewVector2(x float32, y float32) Vector2 {
	return Vector2{
		X: x,
		Y: y,
	}
}

func (a Vector2) Inverse() (Vector2) {
	return Vector2{
		X: -a.X,
		Y: -a.Y,
	}
}

func (a Vector2) Sub(b Vector2) (Vector2) {
	return Vector2{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func (a Vector2) Add(b Vector2) (Vector2) {
	return Vector2{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (a Vector2) Mul(v float32) (Vector2) {
	return Vector2{
		X: a.X * v,
		Y: a.Y * v,
	}
}

func (a Vector2) Cross(b Vector2) (float32) {
	return a.X * b.Y - a.Y * b.X
}

func (a Vector2) Div(v float32) (Vector2) {
	return Vector2{
		X: a.X / v,
		Y: a.Y / v,
	}
}

func (a Vector2) Dot(b Vector2) (float32) {
	return a.X * b.X + a.Y * b.Y
}