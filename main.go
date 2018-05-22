// TODO: Convert everything to a more "OOP" style
package main

import (
  // "os"
  "fmt"
  "math/rand"
  "time"
  "strconv"
  "github.com/fatih/color"
)

// Available moves:
// - straight cut
// - turn over stack
// - cull (Hofsinzer style)

func main() {

  fmt.Println("Stack Explorer v0.0.1")

  const stack_len = 6

  // Initialising the Seed for randomness
  rand.Seed(time.Now().Unix())

  fmt.Print(rand.Intn(stack_len))
  fmt.Print("\n")

  correct_stack := get_ordered_stack(stack_len)
  shuffled_stack := get_shuffled_stack(stack_len)

  print_stack(correct_stack)
  fmt.Print("\n")
  print_stack(shuffled_stack)

  // Available colors (7)
  // red
  // green
  // yellow
  // blue
  // magenta
  // cyan
  // white

  for tryNb := 0 ; !is_eq(shuffled_stack, correct_stack); tryNb++ {
    shuffled_stack = cull(shuffled_stack, gen_asc_uniq_numbers(3, stack_len))
    color.Blue("TRY #%s", strconv.Itoa(tryNb))
    fmt.Println("Current state : ", shuffled_stack)
    fmt.Println("Wanted state  : ", correct_stack)
    fmt.Println("")
  }

  // stack = cull(shuffled_stack, []int {2, 3, 12})
  // print_stack(stack)

  // stack = straight_cut(stack, 5)
  // print_stack(stack)

  // stack = turn_over(stack)
  // fmt.Println(stack)

  // shuffled_stack := get_shuffled_stack(13)
}
