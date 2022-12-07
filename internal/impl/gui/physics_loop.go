package gui

type PhysicsLoop struct {
	application *ApplicationImpl
}

func NewPhysicsLoop(app *ApplicationImpl) *PhysicsLoop {
	return &PhysicsLoop{app}
}

func (physicsLoop PhysicsLoop) Run() {

}
