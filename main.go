package main

import (
  "fmt"
  "math/rand"
)

var debug_mode = true

func main() {

  fmt.Println("Stack Explorer v0.0.1")
  const stack_len = 13

  // Initialising the Seed for randomness
  rand.Seed(123)

  stack := get_ordered_stack(13)
  print_stack(stack)


  stack = cull(stack, []int {2, 3, 12})
  print_stack(stack)

  // stack = straight_cut(stack, 5)
  // print_stack(stack)

  // stack = turn_over(stack)
  // fmt.Println(stack)

  // shuffled_stack := get_shuffled_stack(13)
}
