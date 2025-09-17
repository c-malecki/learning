#include <stdio.h>

typedef struct {
  int x;
  int y;
  char *name;
} myStruct;

char name[6] = "Chris";

int main() {
  myStruct p;
  p.x = 1;
  p.y = 2;
  p.name = name;

  printf("x: %d\ny: %d\nname: %s\n", p.x, p.y, p.name);
};
