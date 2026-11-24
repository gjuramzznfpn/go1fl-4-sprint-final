package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	tokens := strings.Split(data, ",")

	if len(tokens) != 3 {
		return 0, "", 0, errors.New("Строка должна содержать 3 параметра разделенныйх запятой")
	}

	steps, err := strconv.Atoi(tokens[0])

	if err != nil {
		return 0, "", 0, err
	}

	if steps <= 0 {
		return 0, "", 0, errors.New("Колличество шагов должно быть больше 0")
	}

	duration, err := time.ParseDuration(tokens[2])

	if err != nil {
		return 0, "", 0, err
	}

	if duration <= 0 {
		return 0, "", 0, errors.New("Продолжительность должна быть больше 0")
	}

	return steps, tokens[1], duration, nil
}

func distance(steps int, height float64) float64 {
	return stepLengthCoefficient * height * float64(steps) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	return distance(steps, height) / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)

	if err != nil {
		log.Println(err)
		return "", err
	}

	var calories float64

	switch activity {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	default:
		err := errors.New("неизвестный тип тренировки")
		log.Println(err)
		return "", err
	}

	output := fmt.Sprintf("Тип тренировки: %v\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activity, duration.Hours(), distance(steps, height), meanSpeed(steps, height, duration), calories)

	return output, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("Колличество шагов должно быть больше 0")
	}

	if weight <= 0 {
		return 0, errors.New("Вес должен быть больше 0")
	}

	if height <= 0 {
		return 0, errors.New("Рост должен быть больше 0")
	}

	if duration <= 0 {
		return 0, errors.New("Продолжительность должна быть больше 0")
	}

	return (weight * meanSpeed(steps, height, duration) * duration.Minutes()) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	output, err := RunningSpentCalories(steps, weight, height, duration)

	return output * float64(walkingCaloriesCoefficient), err
}
