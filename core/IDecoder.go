// Implements the basic decoder interface
// Author: Alessandro Paganelli (alessandro.paganelli@gmail.com)

package core

type IDecoder interface {
	DecodeLine(line string) (bool, error)
	GetDecoderDescription() string
}
