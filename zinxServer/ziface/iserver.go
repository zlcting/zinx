package ziface

//IServer is interface
type IServer interface {
	Start()
	Stop()
	Serve()
}
