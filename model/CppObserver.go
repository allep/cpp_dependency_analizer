// Implements Cpp observer interfaces
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package model

type IncludeObserver interface {
	UpdateIncludeList(list []string)
}

type TypedefObserver interface {
	UpdateTypedefList(list []string)
}

type EnumObserver interface {
	UpdateEnumList(list []string)
}

type ClassObserver interface {
	UpdateClassList(list []string)
}

