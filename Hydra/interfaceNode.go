package main

import (
	"errors"
)

var ErrInvalidNode = errors.New("Node is not valid")

type Node interface {
	SetValue(v int) error
	GetValue() int
}

type SLLNode struct {
	next         *SLLNode
	value        int
	SNodeMessage string
}

func (SNode *SLLNode) SetValue(v int) error {
	if SNode == nil {
		return ErrInvalidNode
	}
	SNode.value = v
	return nil
}

func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

func NewSLLNode() *SLLNode {
	return &SLLNode{SNodeMessage: "This is a message from the normal Node"}
}

type PowerNode struct {
	next         *PowerNode
	value        int
	PNodeMessage string
}

func (sNode *PowerNode) SetValue(v int) error {
	if sNode == nil {
		return ErrInvalidNode
	}
	sNode.value = v * 10
	return nil
}

func (sNode *PowerNode) GetValue() int {
	return sNode.value
}

func NewPowerNode() *PowerNode {
	return &PowerNode{
		PNodeMessage: "This is a message from the power Node.",
	}
}
