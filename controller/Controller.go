package controller

import (
	"fmt"

	"github.com/allep/cpp_dependency_analyzer/core"
)

//--------------------------------------------------------------------
// controller
type Controller struct {
	path_projects             []core.PathProjectPair
	excluded_extensions       []string
	abstract_report_enable    bool
	instability_report_enable bool

	// views and models - observers
	view_list  []core.IView
	model_list []core.IModel
}

//--------------------------------------------------------------------
// Methods called by Views
func (c *Controller) OnPathProjectList(list []core.PathProjectPair) {
	fmt.Println("Controller: OnPathProjectList called")
}

func (c *Controller) OnExcludedFileExtList(extensions []string) {
	fmt.Println("Controller: OnExcludedFileExtList called")
}

func (c *Controller) OnAbstractReportEnable(enable bool) {
	fmt.Println("Controller: OnAbstractReportEnable called with enable =", enable)
	c.abstract_report_enable = enable
}

func (c *Controller) OnInstabilityReportEnable(enable bool) {
	fmt.Println("Controller: OnInstabilityReportEnable called with enable =", enable)
	c.instability_report_enable = enable
}

func (c *Controller) OnStart() bool {
	fmt.Println("Controller: OnStart called")
	return true
}

func (c *Controller) OnStop() bool {
	fmt.Println("Controller: OnStop called")
	return true
}

//--------------------------------------------------------------------
// Methods called by the Model
func (c *Controller) OnAnalysisDone(pReport *core.AnalysisReport) {
	fmt.Println("Controller: OnAnalysisDone called")
}

//--------------------------------------------------------------------
// Observer methods
func (c *Controller) AddViewAsObserver(view core.IView) {
	fmt.Println("Controller: adding a view as observer")
	c.view_list = append(c.view_list, view)
	fmt.Println("Controller: current length of view_list:", len(c.view_list))
}

func (c *Controller) AddModelAsObserver(model core.IModel) {
	fmt.Println("Controller: adding a model as observer")
	c.model_list = append(c.model_list, model)
	fmt.Println("Controller: current length of model_list:", len(c.model_list))
}
