// IView.go
// Implements view related interfaces
package core

type IView interface {
	// Methods called by controller
	OnAnalysisDone(pReport *AnalysisReport)

	// Observer methods
	AddControllerAsObserver(controller IController)
}
