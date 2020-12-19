// IController.go
// Implements controller related interfaces
package core

type PathProjectPair struct {
	Path    string
	Project string
}

type IController interface {
	// Methods toward views
	OnPathProjectList(list []PathProjectPair)
	OnExcludedFileExtList(extensions []string)
	OnAbstractReportEnable(enable bool)
	OnInstabilityReportEnable(enable bool)
	OnStart() bool
	OnStop() bool

	// Methods toward models
	// TODO
}
