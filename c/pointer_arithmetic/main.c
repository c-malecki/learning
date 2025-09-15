#include <stdio.h>

int main() {
  int arr[] = {10, 20, 30, 40, 50};
  int *ptr = arr; // Point to first element

  // Traditional array indexing
  for (int i = 0; i < 5; i++) {
    printf("arr[%d] = %d\n", i, arr[i]);
  }

  printf("\n");

  // Pointer arithmetic - equivalent to above
  for (int i = 0; i < 5; i++) {
    printf("*(ptr + %d) = %d\n", i, *(ptr + i));
    // ptr + i moves i positions forward
    // *(ptr + i) dereferences to get the value
  }

  printf("\n");

  // Moving the pointer itself
  for (int i = 0; i < 5; i++) {
    printf("*ptr = %d\n", *ptr);
    ptr++; // Move pointer to next element
  }

  return 0;
}