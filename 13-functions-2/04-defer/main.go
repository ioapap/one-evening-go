package main

import (
	"fmt"
	"time"
)

func main() {
	metrics := &Metrics{}

	err := Execute(func() error {
		fmt.Println("Executing...")
		return nil
	}, metrics)

	fmt.Println(err)
}

func Execute(f func() error, metrics *Metrics) (err error) {
	// Call StoreExecution before calling the function.
	metrics.StoreExecution()

	// Use defer to handle the err return value.
	defer func() {
		if err != nil {
			// Call StoreFailure if the function f returns with an error.
			metrics.StoreFailure()
		} else {
			// Call StoreSuccess if the function f returns without an error.
			metrics.StoreSuccess()
		}
	}()

	// Execute the function.
	err = f()
	return err
}

type Metrics struct {
	execution []time.Time
	success   []time.Time
	failure   []time.Time
}

func (m *Metrics) StoreExecution() {
	m.execution = append(m.execution, time.Now())
}

func (m *Metrics) StoreSuccess() {
	m.success = append(m.success, time.Now())
}

func (m *Metrics) StoreFailure() {
	m.failure = append(m.failure, time.Now())
}
