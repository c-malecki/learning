#include <stdio.h>

int notStatic() {
  int count = 0;
  count++;
  return count;
}

int isStatic() {
  static int count = 0;
  count++;
  return count;
}

int main() {
  printf("%d ", notStatic());
  printf("%d ", notStatic());
  /* prints 1 1 */

  printf("%d ", isStatic());
  printf("%d ", isStatic());
  /* prints 1 2 */
}