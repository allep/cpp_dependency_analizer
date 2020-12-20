// IModel.go
// Implements the model related interfaces
package core

type IModel interface {
	// Methods toward controller
	OnPathProjectList(list []PathProjectPair)
	OnExcludedFileExtList(extensions []string)
	OnAbstractReportEnable(enable bool)
	OnInstabilityReportEnable(enable bool)
	OnStart() bool
	OnStop() bool

	// Observer methods
	AddControllerAsObserver(controller IController)
}
