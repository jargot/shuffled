/*
** SHUFFLE FUNCTIONS
*/

package main

// import (
//   "fmt"
// )

// var _ = fmt.Println // For debugging only

/*
**
**
*/

func turn_over(stack []int) (new_stack []int) {
  stack_len := len(stack) // Can't make it constant since constants are evaluated at compile time
  new_stack = make([]int, stack_len)

  for i := 0; i < stack_len ; i++ {
    new_stack[i] = stack[stack_len - i - 1]
  }
  return
}

func straight_cut(stack []int, depth int) (new_stack []int) {
  first_half := stack[:depth]
  second_half := stack[depth:]
  new_stack = append(new_stack, second_half...)
  new_stack = append(new_stack, first_half...)
  return
}

// Positions must be > 1 and be in ascending order
func cull(stack []int, asc_positions []int) (new_stack []int) {
  culled_stack := make([]int, len(asc_positions))
  new_stack = make([]int, len(stack))

  culled_stack_index := 0;
  asc_positions_index := 0;
  for i := 0; i < len(stack) ; i++ {
    if (asc_positions_index < len(asc_positions) && i == asc_positions[asc_positions_index]) {
      culled_stack[culled_stack_index] = stack[i]
      stack[i] = 0;
      asc_positions_index++
      culled_stack_index++
    }
  }
  new_stack_index := 0;
  for j := 0; j < len(stack) ; j++ {
    if stack[j] != 0 {
      new_stack[new_stack_index] = stack[j]
      new_stack_index++
    }
  }
  new_stack = new_stack[:len(stack) - len(asc_positions)]
  new_stack = append(new_stack, culled_stack...)
  return
}
