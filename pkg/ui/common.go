package ui

import "fmt"

type Size interface {
	Size() string
}

type size struct {
	value string
}

func (s *size) Size() string {
	return s.value
}

func Px(value int) Size {
	return &size{
		value: fmt.Sprintf("%dpx", value),
	}
}

func Fr(value int) Size {
	return &size{
		value: fmt.Sprintf("%dfr", value),
	}
}

func Rem(value float64) Size {
	return &size{
		value: fmt.Sprintf("%.2frem", value),
	}
}

func Em(value float64) Size {
	return &size{
		value: fmt.Sprintf("%.2fem", value),
	}
}

func Percent(value float64) Size {
	return &size{
		value: fmt.Sprintf("%.2f%%", value),
	}
}

func Vh(value float64) Size {
	return &size{
		value: fmt.Sprintf("%.2fvh", value),
	}
}

func Vw(value float64) Size {
	return &size{
		value: fmt.Sprintf("%.2fvw", value),
	}
}
