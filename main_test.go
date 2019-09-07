package main

import "testing"

func TestGetTickets(t *testing.T) {
	ticket := getTicket()
	if len(ticket) != 8 {
		t.Errorf("ticket should have a length of 8.")
	}
}

func TestGetFieldGenerator(t *testing.T) {
	nextField := fieldGenerator()
	field, ok := nextField()
	if !ok {
		t.Errorf("ok should be true.")
	}
	if len(field) != 6 {
		t.Errorf("ticket should have a length of 6.")
	}
}

func TestGetNumbers(t *testing.T) {
	numbers := getNumbers(10)
	if len(numbers) != 10 {
		t.Errorf("numbers should have a length of 10.")
	}
}
