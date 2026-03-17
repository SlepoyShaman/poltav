package main

import (
	"bufio"
	"fmt"
	"lab1/dynamic_models"
	"lab1/integrators"
	"lab1/integrators/methods"
	"lab1/vectors"
	"os"
	"strconv"
	"strings"
)

const errorReadingFloat = "Ошибка: введите корректное число"

func main() {
	fmt.Print("Задайте начальный вектор (x, y, z, Vx, Vy, Vz) через пробел: ")
	start, err := readVector()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Print("Задайте шаг интегрирования: ")
	h, err := readFloat()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	ship := dynamic_models.SpaceShip{}

	var method methods.Method
	fmt.Print("Выберете метод интегрирования: r - Рунге-Кутта, другое - Эйлера: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	m := strings.TrimSpace(input)
	if m == "r" {
		method = methods.NewRungeKutte(h, &ship)
	} else {
		method = methods.NewEuler(h, &ship)
	}

	fmt.Print("Задайте конечное время: ")
	tk, err := readFloat()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	integrator := integrators.NewIntegrator(h, method, start)
	integrator.MoveTo(tk)
}

func readVector() (vectors.StateVector, error) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Println("aboba - ", input)
	values := make([]float64, 0, 6)
	for _, part := range strings.Split(input, " ") {
		if part == "" {
			continue
		}

		num, err := strconv.ParseFloat(part, 64)
		if err != nil {
			return vectors.StateVector{}, fmt.Errorf("%s не является числом", part)
		}

		values = append(values, num)
	}

	if len(values) != 6 {
		return vectors.StateVector{}, fmt.Errorf("Необходимо ввести вектор длинной 6, получен вектор длинной %d", len(values))
	}

	return vectors.StateVector{
		X:  values[0],
		Y:  values[1],
		Z:  values[2],
		Vx: values[3],
		Vy: values[4],
		Vz: values[5],
	}, nil
}

func readFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("Ошибка: введите корректное число, введено: %s", input)
	}

	return num, nil
}
