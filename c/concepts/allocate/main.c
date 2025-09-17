#include <stdio.h>
#include <stdlib.h>

typedef struct {
  char *name;
  int age;
} person;

int main() {
  person *myperson = malloc(sizeof(person));

  myperson->name = "John";
  myperson->age = 27;

  printf("name: %s age: %d\n", myperson->name, myperson->age);

  free(myperson);

  myperson = malloc(sizeof(person));

  myperson->name = "Jim";
  myperson->age = 40;

  printf("name: %s age: %d\n", myperson->name, myperson->age);

  free(myperson);

  int rows = 2;
  int cols = 5;
  int i, j;

  // Allocate memory for rows pointers
  char **pvowels = (char **)malloc(rows * sizeof(char *));

  // For each row, allocate memory for cols elements
  pvowels[0] = (char *)malloc(cols * sizeof(char));
  pvowels[1] = (char *)malloc(cols * sizeof(char));

  pvowels[0][0] = 'A';
  pvowels[0][1] = 'E';
  pvowels[0][2] = 'I';
  pvowels[0][3] = 'O';
  pvowels[0][4] = 'U';

  pvowels[1][0] = 'a';
  pvowels[1][1] = 'e';
  pvowels[1][2] = 'i';
  pvowels[1][3] = 'o';
  pvowels[1][4] = 'u';

  for (i = 0; i < rows; i++) {
    for (j = 0; j < cols; j++) {
      printf("%c ", pvowels[i][j]);
    }

    printf("\n");
  }

  /*
    A E I O U
    a e i o u
  */

  // Free individual rows
  free(pvowels[0]);
  free(pvowels[1]);

  // Free the top-level pointer
  free(pvowels);
};