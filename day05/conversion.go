package day05

import (
	"errors"
	"strconv"
	"strings"
)

type Conversion struct {
	maps []rangeMap
}

func ParseConversion(lines []string) (Conversion, error) {
	rangeMaps := make([]rangeMap, 0)
	for _, mapLine := range lines[1:] {
		split := strings.Fields(mapLine)
		converted := make([]int, 0)
		for _, intString := range split {
			conv, err := strconv.Atoi(intString)
			if err != nil {
				return Conversion{}, err
			}
			converted = append(converted, conv)
		}

		if len(converted) != 3 {
			return Conversion{}, errors.New("expected range map line of length 3: " + mapLine)
		}
		rangeMaps = append(rangeMaps, rangeMap{
			destination: converted[0],
			source:      converted[1],
			rangeLength: converted[2]})
	}
	return Conversion{maps: rangeMaps}, nil
}

func (conv Conversion) Convert(source int) int {
	rm, ok := conv.FindRange(source)
	if ok {
		return source + rm.Offset()
	}
	return source
}

func (conv Conversion) FindRange(source int) (rangeMap, bool) {
	for _, rm := range conv.maps {
		if rm.InRange(source) {
			return rm, true
		}
	}
	return rangeMap{}, false
}

type rangeMap struct {
	source      int
	destination int
	rangeLength int
}

func (rm rangeMap) InRange(n int) bool {
	return n >= rm.source && n <= rm.source+rm.rangeLength
}

func (rm rangeMap) Offset() int {
	return rm.destination - rm.source
}
