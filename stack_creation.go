/*
** STACK CREATION FUNCTIONS
*/

package main

import (
  "math/rand"
)


func get_ordered_stack(stack_size int) []int {
  stack := make([]int, stack_size)

  for i := 0; i < stack_size ; i++ {
    // For simplification purposes, we simply consider cards as numbers ranging from 1 to infinity.
    stack[i] = i + 1
  }
  return stack
}

func get_shuffled_stack(stack_size int) []int {
  // Get shuffle from 0 to stack_size
  list := rand.Perm(stack_size)

  // Increment every value by 1 to weed out the zero.
  // WITH MAGIC (range keyword plus blank identifier aka underscore)
  // for i, _ := range list {
  //   list[i]++
  // }
  // or WITHOUT MAGIC
  for i := 0; i < stack_size ; i++ {
    list[i]++
  }
  return list
}
