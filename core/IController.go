// Implements controller related interfaces
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package core

// Structure used to tie together a path and the project it refers to.
type PathProjectPair struct {
	Path    string
	Project string
}

// The main analysis structure, holding inside it
type AnalysisReport struct {
	NumClasses         uint32
	NumAbstractClasses uint32
	NumFiles           uint32
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
	OnAnalysisDone(pReport *AnalysisReport)

	// Observer methods
	AddViewAsObserver(view IView)
	AddModelAsObserver(model IModel)
}
