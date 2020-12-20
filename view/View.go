package view

import (
	"fmt"

	"github.com/allep/cpp_dependency_analyzer/core"
)

type CLIView struct {
	path_project_pair_list    []core.PathProjectPair
	excluded_extensions       []string
	abstract_report_enable    bool
	instability_report_enable bool

	// controllers - observer
	controller_list []core.IController
}

//--------------------------------------------------------------------
// API
func (v *CLIView) SetPathProjectPairList(list []core.PathProjectPair) {
	copy(v.path_project_pair_list, list)

	for _, c := range v.controller_list {
		c.OnPathProjectList(v.path_project_pair_list)
	}
}

func (v *CLIView) SetExcludedFileExtList(extensions []string) {
	copy(v.excluded_extensions, extensions)

	for _, c := range v.controller_list {
		c.OnExcludedFileExtList(v.excluded_extensions)
	}
}

func (v *CLIView) SetAbstractReportEnable(enable bool) {
	v.abstract_report_enable = enable

	for _, c := range v.controller_list {
		c.OnAbstractReportEnable(v.abstract_report_enable)
	}
}

func (v *CLIView) SetInstabilityReportEnable(enable bool) {
	v.instability_report_enable = enable

	for _, c := range v.controller_list {
		c.OnInstabilityReportEnable(v.instability_report_enable)
	}
}

func (v *CLIView) Start() bool {
	var ret bool = true
	for _, c := range v.controller_list {
		ret = ret && c.OnStart()
	}
	return ret
}

func (v *CLIView) Stop() bool {
	var ret bool = true
	for _, c := range v.controller_list {
		ret = ret && c.OnStop()
	}
	return ret
}

//--------------------------------------------------------------------
// Methods called by Controllers
func (v *CLIView) OnAnalysisDone(pReport *core.AnalysisReport) {
	fmt.Println("CLIView: OnAnalysisDone called")
}

//--------------------------------------------------------------------
// Observer methods
func (v *CLIView) AddControllerAsObserver(controller core.IController) {
	fmt.Println("CLIView: adding a controller as observer")
	v.controller_list = append(v.controller_list, controller)
	fmt.Println("CLIView: current length of controller_list:", len(v.controller_list))
}
