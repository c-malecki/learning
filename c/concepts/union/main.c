#include <stdio.h>

union intParts {
  int theInt;
  char bytes[sizeof(int)];
};

enum NumberKind { FLOAT, INT };
struct Number {
  enum NumberKind kind;
  union {
    int i;
    float f;
  };
};

void output_number(struct Number *n) {
  switch (n->kind) {
  case INT:
    printf("The integer %d\n", n->i);
    break;
  case FLOAT:
    printf("The float %f\n", n->f);
    break;
  }
}

int main() {
  struct Number three = {.kind = INT, {.i = 3}};
  struct Number two_point_five = {.kind = FLOAT, {.f = 2.5}};

  output_number(&three);
  output_number(&two_point_five);

  return 0;
}