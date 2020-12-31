// Implements view related interfaces
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package core

type IView interface {
	// Methods called by controller
	OnAnalysisDone(pReport *AnalysisReport)

	// Observer methods
	AddControllerAsObserver(controller IController)
}
