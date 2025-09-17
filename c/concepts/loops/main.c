#include <stdio.h>

int main() {
  int n = 0;
  int arr[10];

  while (n < 10) {
    n++;
    printf("while: %d\n", n);
  }

  int i;

  for (i = n; i > 0; i--) {
    n--;
    printf("for: %d\n", n);
  }
};