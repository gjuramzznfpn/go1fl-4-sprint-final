package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
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
	steps, duration, err := parsePackage(data)

	if steps <= 0 {
		return ""
	}

	distance := (float64(steps) * stepLength) / float64(mInKm)

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	if err != nil {
		return ""
	}

	output := fmt.Sprintf("Количество шагов: %v.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)

	return output
}
