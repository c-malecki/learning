#include <stdio.h>

typedef struct {
  int *y;
} nest;

typedef struct {
  int x;
  nest *n;
} point;

void addNoPointer(int i) { i++; };

void addWithPointer(int *i) { (*i)++; };

void incrementStruct(point *p) {
  // short hand dereference for (*p).x++;
  p->x++;
  // dereference pointer inside struct pointer
  (*p->n->y)++;
};

int i;

int main() {
  printf("i: %d\n", i);
  /* i: 0 */

  addNoPointer(i);

  printf("i: %d\n", i);
  /* i: 0 */

  addWithPointer(&i);

  printf("i: %d\n", i);
  /* i: 1 */

  nest n;
  n.y = &i;

  point p;
  p.x = i;
  p.n = &n;

  addWithPointer(p.n->y);

  incrementStruct(&p);

  printf("p.x: %d\np.n.y: %d\n", p.x, *p.n->y);
  /*
  p.x: 2
  p.n.y: 3
  */
};