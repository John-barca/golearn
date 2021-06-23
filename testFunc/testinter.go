package test

// 当接口名和方法名都大写时，表示该方法公开
// 可以被接口所在包以外的地方被访问
type Tank interface {
	Walk() 		// move 
	Fire()		// attack
}

type Plane interface {
	Fly()
}

type PlaneTank interface {
	Tank
	Plane
}

type Person struct {
	Name string 
	Birth string
	ID int64
}

