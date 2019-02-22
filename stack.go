package main

import (
	"errors"
)

type Stack struct {
	memory []uint16
	index uint
	capacity uint
}


func newStack(size uint) *Stack {
	stack := new(Stack)
	stack.memory = make([]uint16, size)
	return stack
}


func (stack *Stack) push(item uint16) error {
	if index == capacity {
		return errors.New("Stack is full")
	}

	stack.memory[stack.index] = item
	stack.index++

	return nil
}


func (stack *Stack) pop(item uint16) (uint16, error) {
	if index == 0 {
		return 0, errors.New("Stack is empty")
	}

	stack.index--
	item := stack.memory[stack.index]
	stack.memory[stack.index] = 0

	return item, nil
}


func (stack *Stack) peek() (uint16, error) {
	if index == 0 {
		return 0, errors.New("Stack is empty")
	}

	return stack.memory[stack.index - 1], nil
}


func (stack *Stack) size() uint {
	return len(stack.memory)
}

func (stack *Stack) capacity() uint {
	return stack.capacity
}