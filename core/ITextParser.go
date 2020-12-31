// Implements the basic interface of the text parser
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package core

type ITextParser interface {
	ParseLine(line string) error
	GetKeyFromLine(line string) (string, error)
}
