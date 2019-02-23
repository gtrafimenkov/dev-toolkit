// SPDX-License-Identifier: MIT
// Copyright (c) 2019 Gennady Trafimenkov

package main

import "fmt"

type queue struct {
	incoming []int
	outgoing []int
}

func (q *queue) push(value int) {
	q.incoming = append(q.incoming, value)
}

func (q *queue) pop() (value int, ok bool) {
	outgoingSize := len(q.outgoing)
	if outgoingSize > 0 {
		value = q.outgoing[outgoingSize-1]
		q.outgoing = q.outgoing[0 : outgoingSize-1]
		return value, true
	}

	incomingSize := len(q.incoming)
	if incomingSize > 0 {

		// move items from incoming to outgoing reversing the order in the process
		for i := incomingSize - 1; i >= 0; i-- {
			q.outgoing = append(q.outgoing, q.incoming[i])
		}
		q.incoming = q.incoming[0:0]

		return q.pop()
	}
	return 0, false
}

func main() {
	q := new(queue)
	q.push(1)
	q.push(2)
	q.push(3)

	value, _ := q.pop()
	fmt.Println(value)

	q.push(4)
	q.push(5)
	for {
		value, ok := q.pop()
		if !ok {
			break
		}
		fmt.Println(value)
	}
}
