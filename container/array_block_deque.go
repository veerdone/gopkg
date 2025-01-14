//go:build go1.18

/*
 * Copyright 2023 veerdone
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package container

import "sync"

type ArrayBlockDeque[T any] struct {
	element                   []*node[T]
	notFull, notEmpty         *sync.Cond
	takeIndex, putIndex, size int
}

func NewArrayBlockDeque[T any](num int) *ArrayBlockDeque[T] {
	if num <= 0 {
		num = 16
	}

	return &ArrayBlockDeque[T]{
		element:  make([]*node[T], num),
		notFull:  sync.NewCond(&sync.Mutex{}),
		notEmpty: sync.NewCond(&sync.Mutex{}),
	}
}

func (a *ArrayBlockDeque[T]) Put(value T) {
	a.notFull.L.Lock()
	defer a.notFull.L.Unlock()
	if a.size == len(a.element) {
		a.notFull.Wait()
	}
	a.element[a.putIndex] = &node[T]{value}
	a.putIndex = (a.putIndex+1) % len(a.element)
	a.size++
	a.notEmpty.Signal()
}

func (a *ArrayBlockDeque[T]) Take() T {
	a.notEmpty.L.Lock()
	defer a.notEmpty.L.Unlock()
	if a.size == 0 {
		a.notEmpty.Wait()
	}
	n := a.element[a.takeIndex]
	a.element[a.takeIndex] = nil
	a.takeIndex = (a.takeIndex+1) % len(a.element)
	a.size--
	a.notFull.Signal()

	return n.Val()
}
