package main

type employee interface {
	PrintName() string
	PrintAddress(id int)
	PrintSalary(b int, t int) float64
}
