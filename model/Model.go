package model

import (
	"fmt"

	"github.com/allep/cpp_dependency_analyzer/core"
)

// model

type Model struct {
	path_projects             []core.PathProjectPair
	excluded_extensions       []string
	abstract_report_enable    bool
	instability_report_enable bool

	// controllers - observer
	controller_list []core.IController
}

//--------------------------------------------------------------------
// Methods called by Controllers
func (m *Model) OnPathProjectList(list []core.PathProjectPair) {
	fmt.Println("Model: OnPathProjectList called")
}

func (m *Model) OnExcludedFileExtList(extensions []string) {
	fmt.Println("Model: OnExcludedFileExtList called")
}

func (m *Model) OnAbstractReportEnable(enable bool) {
	fmt.Println("Model: OnAbstractReportEnable called with enable =", enable)
	m.abstract_report_enable = enable
}

func (m *Model) OnInstabilityReportEnable(enable bool) {
	fmt.Println("Model: OnInstabilityReportEnable called with enable =", enable)
	m.instability_report_enable = enable
}

func (m *Model) OnStart() bool {
	fmt.Println("Model: OnStart called")
	return true
}

func (m *Model) OnStop() bool {
	fmt.Println("Model: OnStop called")
	return true
}

//--------------------------------------------------------------------
// Observer methods
func (m *Model) AddControllerAsObserver(controller core.IController) {
	fmt.Println("Model: adding a controller as observer")
	m.controller_list = append(m.controller_list, controller)
	fmt.Println("Model: current length of controller_list:", len(m.controller_list))
}
