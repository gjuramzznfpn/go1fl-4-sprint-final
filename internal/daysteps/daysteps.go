package daysteps

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	tokens := strings.Split(data, ",")

	if len(tokens) != 2 {
		return 0, 0, errors.New("Строка должна содержать два параметра разделенныйх запятой")
	}

	steps, err := strconv.Atoi(tokens[0])

	if err != nil {
		return 0, 0, err
	}

	if steps <= 0 {
		return 0, 0, errors.New("Колличество шагов должно быть больше 0")
	}

	duration, err := time.ParseDuration(tokens[1])

	if err != nil {
		return 0, 0, err
	}

	if duration <= 0 {
		return 0, 0, errors.New("Продолжительность должна быть больше 0")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	return ""
}
