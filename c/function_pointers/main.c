#include <stdio.h>

// Function pointer type definition
typedef int (*operation)(int, int);

// Simple math functions
int add(int a, int b) { return a + b; }
int subtract(int a, int b) { return a - b; }
int multiply(int a, int b) { return a * b; }

// Function that takes a function pointer as parameter
int calculate(int x, int y, operation op) { return op(x, y); }

int main() {
  int a = 10, b = 5;

  // Direct function pointer usage
  operation func = add;
  printf("%d + %d = %d\n", a, b, func(a, b));

  // Passing functions as arguments
  printf("%d - %d = %d\n", a, b, calculate(a, b, subtract));
  printf("%d * %d = %d\n", a, b, calculate(a, b, multiply));

  // Array of function pointers
  operation ops[] = {add, subtract, multiply};
  char symbols[] = {'+', '-', '*'};

  for (int i = 0; i < 3; i++) {
    printf("%d %c %d = %d\n", a, symbols[i], b, ops[i](a, b));
  }
}