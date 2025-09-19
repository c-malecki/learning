# C

notes from https://www.learn-c.org/en/Hello%2C_World%21

## Integers

- char
- int
- short
- long
- long long

can also be unsigned

- unsigned char
- unsigned int
- unsigned short
- unsigned long
- unsigned long long

size of data type is dependent on architecture (32bit vs 64bit etc...)

```c
    printf("The size of int: %d bytes\n", sizeof(int));
    printf("The size of char: %d bytes\n", sizeof(char));
    printf("The size of float: %d bytes\n", sizeof(float));
    printf("The size of double: %d bytes", sizeof(double));
```

there is no `boolean` data primitive in C

```c
/* example definition of boolean */
#define BOOL char
#define FALSE 0
#define TRUE 1
```

## Arrays

```c
/* defines an array of 10 integers */
int numbers[10];
/* populate the array */
numbers[0] = 10;
numbers[1] = 20;

/* multi-dimensional */
char vowels[1][5] = {
    {'a', 'e', 'i', 'o', 'u'}
};
/* vowels is an array with capacity of 1
the element in the capacity 1 array is an array with a capacity of 5
the elements in the capaicty 5 are the characters */

/* array with 3 rows and each row has 4 columns */
int a[3][4] = {
   {0, 1, 2, 3} ,   /*  initializers for row indexed by 0 */
   {4, 5, 6, 7} ,   /*  initializers for row indexed by 1 */
   {8, 9, 10, 11}   /*  initializers for row indexed by 2 */
};

/* can omit the first capacity declaration and the compiler will still work */
int a[][4] = {
   {0, 1, 2, 3} ,
   {4, 5, 6, 7} ,
   {8, 9, 10, 11}
};
/* can also omit the inside braces which initializes to the same thing as above */
int a[][4] = {0,1,2,3,4,5,6,7,8,9,10,11};
```

## Strings

Strings in C are arrays of characters.

Pointer to character array to define simple string which is READ ONLY

```c
char * name = "John Smith";
```

For a string that can be manipulated, define it as local character array

```c
char name[] = "John Smith";
```

Empty brackets notation tells the compiler to calculate the size of the array automatically.

```c
char name[] = "John Smith";
/* is the same as */
char name[11] = "John Smith";
```

The reason we allocate 11 instead of 10 (the actual length of "John Smith") is a special character equal to 0 terminates the string.

```c
char * name = "John Smith";
int age = 27;

printf("%s is %d years old.\n", name, age);
/* John Smith is 27 years old. */
```

### String functions:

`strlen` length of string

`strncmp` compares strings, return 0 if equal or different number if not

```c
char * name = "John";
/* args - strings to compare and maximum comparison length */
if (strncmp(name, "John", 4) == 0) {
    printf("Hello, John!\n");
} else {
    printf("You are not John. Go away.\n");
}
```

## Loops

### for

```c
int array[10] = { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 };
int sum = 0;
int i;

for (i = 0; i < 10; i++) {
    sum += array[i];
}

printf("Sum of the array is %d\n", sum);
```

### while

```c
int n = 0;

while (n < 10) {
    n++;
}
```

### infinite

```c
while (1) {
   /* do something */
   break
}
```

Uses typical directives like `break` and `continue`

## Functions

Functions must be first defined before they are used in the code. They can be either declared first and then implemented later on using a header file or in the beginning of the C file.

```c
/* function declaration */
int foo(int bar);

int main() {
    /* calling foo from main */
    printf("The value of foo is %d", foo(1));
}

int foo(int bar) {
    return bar + 1;
}
```

void

```c
void moo() {
    /* do something and don't return a value */
}

int main() {
    moo();
}
```

## Scope - static, global

By default, variables are local to the scope they are defined. Using `static` increases their scope up to the file containing them. These variables can be accessed anywhere inside a file.

```c
#include <stdio.h>

int notStatic()
{
  int count = 0;
  count++;
  return count;
}

int isStatic()
{
  static int count = 0;
  count++;
  return count;
}

int main()
{
  printf("%d ", notStatic());
  printf("%d ", notStatic());
  /* 1 1 */

  printf("%d ", isStatic());
  printf("%d ", isStatic());
  /* 1 2 */
  return 0;
}
```

By default, functions are global.

`static` will declare a function where the scope is reduced to the file containing it.

```c
static void fun(void) {
   printf("I am a static function.");
}
```

### typedef

using `typedef` will declare a type to be reused

```c
#include <stdio.h>

typedef struct
{
  int x;
  int y;
  char *name;
} myStruct;

char name[6] = "Chris";

int main()
{
  myStruct p;
  p.x = 1;
  p.y = 2;
  p.name = name;

  printf("x: %dy: %dname: %s", p.x, p.y, p.name);
  /* x: 1 y: 2 name: Chris */
};
```

## Structs

```c
struct point {
    int x;
    int y;
};

p.x = 10;
p.y = 5;
```

## Pointers

The following does 3 things:

1. allocates a local stack variable called `name` which is a pointer to a single character
2. causes the string "John" to appear somewhere in memory (after compilation and execution)
3. initializes `name` arg to point to where "J" character is in memory and the rest of the characters follow contiguously

```c
char * name = "John";
```

uses typical `*` and `&` syntax

```c
int a = 1;
int * pointer_to_a = &a;

a += 1;
printf("%d\n", a);
/* 2 */

*pointer_to_a += 1;
printf("%d\n", a);
/* 3 */
```

functions pass arguments by value/copies like Go

```c
void addone(int n) {
    // n is local variable which only exists within the function scope
    n++; // therefore incrementing it has no effect
}

int n;
printf("Before: %d\n", n);
addone(n);
printf("After: %d\n", n);
```

## Dynamic Allocation

Allocating memory dynamically helps us to store data without initially knowing the size of the data in the time we wrote the program.

To allocate memory dynamically, we need a pointer to store the location of the allocated memory.
We use that pointer to free the memory when we're done using it.

```c
typedef struct {
    char * name;
    int age;
} person;

/* allocate a new person in the myperson argument */
person * myperson = (person *) malloc(sizeof(person));

myperson->name = "John";
myperson->age = 27;

/* deallocate memory */
free(myperson);
```

This tells the compiler to dynamically allocation just enough memory to hold the `person` struct and returns a pointer of type `person`.

We write `(person *)` before `malloc()` because `malloc()` returns a `void` pointer.

`(person *)` in front is a typecast so that the return type is `person`

The typecast is not necessary because C will implicitly convert the type of the returned pointer to that of the poitner it was assigned to.

```c
person *myperson = malloc(sizeof(person));
```

`free` does not delete the `myperson` variable but releases the data it points to

`myperson` will still point to somewhere in memory, but we cannot access it until we allocate new data using it.

### Arrays and Pointers

```c
char vowels[] = {'A', 'E', 'I', 'O', 'U'};
char *pvowels = vowels;
int i;

// Print the addresses
for (i = 0; i < 5; i++) {
    printf("&vowels[%d]: %p, pvowels + %d: %p, vowels + %d: %p\n", i, &vowels[i], i, pvowels + i, i, vowels + i);
}

/*
&vowels[0]: 0x7ffee146da17, pvowels + 0: 0x7ffee146da17, vowels + 0: 0x7ffee146da17

&vowels[1]: 0x7ffee146da18, pvowels + 1: 0x7ffee146da18, vowels + 1: 0x7ffee146da18
...
*/

// Print the values
for (i = 0; i < 5; i++) {
    printf("vowels[%d]: %c, *(pvowels + %d): %c, *(vowels + %d): %c\n", i, vowels[i], i, *(pvowels + i), i, *(vowels + i));
}

/*
vowels[0]: A, *(pvowels + 0): A, *(vowels + 0): A

vowels[1]: E, *(pvowels + 1): E, *(vowels + 1): E
...
*/
```

`&vowels[i]` and `pvowels + i` and `vowels + i` are equivalent and return the address of the ith element

`vowels[1]` and `*(pvowels + i)` and `*(vowels + i)` all return the ith element

The name of an array itself is a constant pointer to the first element of the array, so they all point to the same location.

### Dynamic memory allocation for arrays

```c
// Allocate memory to store five characters
int n = 5;
char *pvowels = (char *) malloc(n * sizeof(char));
int i;

pvowels[0] = 'A';
pvowels[1] = 'E';
*(pvowels + 2) = 'I';
pvowels[3] = 'O';
*(pvowels + 4) = 'U';

for (i = 0; i < n; i++) {
    printf("%c ", pvowels[i]);
}
// A E I O U

printf("\n");

free(pvowels);
```

Allocate 5 contiguous bytes of memory to store 5 characters

Still use array notation to traverse the blocks of memory even though `pvowels` is a pointer

Useful because when declaring an array, the number of elements it will contain must be known beforehand.

Must remember to call `free()` when relevant to prevent memory leaks

For a two-dimensional array, we use a pointer to a pointer:

```c
int nrows = 2;
int ncols = 5;
int i, j;

// Allocate memory for nrows pointers
char **pvowels = (char **) malloc(nrows * sizeof(char *));

// For each row, allocate memory for ncols elements
pvowels[0] = (char *) malloc(ncols * sizeof(char));
pvowels[1] = (char *) malloc(ncols * sizeof(char));

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

for (i = 0; i < nrows; i++) {
    for(j = 0; j < ncols; j++) {
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
```

### From Claude

`char **` is a pointer to a pointer to char. It's used to create 2D arrays dynamically.

**Breaking it down:**

- `char *` = pointer to a char (can point to a string)
- `char **` = pointer to a `char *` (can point to an array of strings)

**Here's what happens:**

1. `char **pvowels` allocates memory for an array of `char *` pointers (the "rows")
2. Each `pvowels[i]` then gets memory allocated for actual characters (the "columns")

```
pvowels → [ptr] → "actual string data"
          [ptr] → "actual string data"
          [ptr] → "actual string data"
```

So `pvowels` points to an array of pointers, and each of those pointers points to a string.

The "array-like" behavior comes from the **pointer arithmetic** and memory layout:

1. `char **pvowels`: `pvowels` points to a block of memory containing `char *` pointers
2. `pvowels[0]`, `pvowels[1]`: these use array notation, but it's really pointer arithmetic: `pvowels[i]` is equivalent to `*(pvowels + i)`
3. `pvowels[0] = malloc(cols * sizeof(char))`: this allocates a contiguous block of cols characters
4. `pvowels[0][2]` - again, array notation but really `*(pvowels[0] + 2)`

So the "array" behavior is because:

- You allocated contiguous memory blocks
- C lets you use `[]` notation on pointers
- `ptr[i]` is just syntactic sugar for `*(ptr + i)`

A true array would be declared like char arr[5][10] - that's a fixed-size 2D array. Your code creates a dynamic structure that behaves like an array but is really just pointers to memory blocks.

## Recursion

```c
#include <stdio.h>

unsigned int multiply(unsigned int x, unsigned int y)
{
    if (x == 1)
    {
        /* Terminating case */
        return y;
    }
    else if (x > 1)
    {
        /* Recursive step */
        return y + multiply(x-1, y);
    }

    /* Catch scenario when x is zero */
    return 0;
}

int main() {
    printf("3 times 5 is %d", multiply(3, 5));
    return 0;
}
```

## Linked Lists

```c
// defining struct in recursive manner
typedef struct node {
    int val;
    struct node * next;
} node_t;

// declare head as NULL pointer
node_t * head = NULL;

// allocate memory to head
head = (node_t *) malloc(sizeof(node_t));

// make sure malloc did not return a NULL pointer
if (head == NULL) {
    return 1;
}

// set value for head
head->val = 1;

// allocate next node
head->next = (node_t *) malloc(sizeof(node_t));

// set value for node
head->next->val = 2;
head->next->next = NULL;
```

**add item to end:**

```c
void push(node_t * head, int val) {
    node_t * current = head;
    while (current->next != NULL) {
        current = current->next;
    }

    /* now we can add a new variable */
    current->next = (node_t *) malloc(sizeof(node_t));
    current->next->val = val;
    current->next->next = NULL;
}
```

**add item to beginning:**

1. Create a new item and set its value
2. Link the new item to point to the head of the list
3. Set the head of the list to be our new item

pass a pointer to the pointer variable (a double pointer) so we will be able to modify the pointer itself.

```c
void push(node_t ** head, int val) {
    node_t * new_node;
    new_node = (node_t *) malloc(sizeof(node_t));

    new_node->val = val;
    new_node->next = *head;
    *head = new_node;
}
```

**remove first item from list:**

1. Take the next item that the head points to and save it
2. Free the head item
3. Set the head to be the next item that we've stored on the side

```c
int pop(node_t ** head) {
    int retval = -1;
    node_t * next_node = NULL;

    if (*head == NULL) {
        return -1;
    }

    next_node = (*head)->next;
    retval = (*head)->val;
    free(*head);
    *head = next_node;

    return retval;
}
```

**remove last item from list:**

have to look two items ahead and see if the next item is the last one in the list

```c
int remove_last(node_t * head) {
    int retval = 0;
    /* if there is only one item in the list, remove it */
    if (head->next == NULL) {
        retval = head->val;
        free(head);
        return retval;
    }

    /* get to the second to last node in the list */
    node_t * current = head;
    while (current->next->next != NULL) {
        current = current->next;
    }

    /* now current points to the second to last item of the list, so let's remove current->next */
    retval = current->next->val;
    free(current->next);
    current->next = NULL;
    return retval;
}
```

**remove specific item:**

1. Iterate to the node before the node we wish to delete
2. Save the node we wish to delete in a temporary pointer
3. Set the previous node's next pointer to point to the node after the node we wish to delete
4. Delete the node using the temporary pointer

```c
int remove_by_index(node_t ** head, int n) {
    int i = 0;
    int retval = -1;
    node_t * current = *head;
    node_t * temp_node = NULL;

    if (n == 0) {
        return pop(head);
    }

    for (i = 0; i < n-1; i++) {
        if (current->next == NULL) {
            return -1;
        }
        current = current->next;
    }

    if (current->next == NULL) {
        return -1;
    }

    temp_node = current->next;
    retval = temp_node->val;
    current->next = temp_node->next;
    free(temp_node);

    return retval;
}
```

## Binary Trees

https://www.learn-c.org/en/Binary_trees

## Unions

https://www.learn-c.org/en/Unions

C Unions are essentially the same as C Structures, except that instead of containing multiple variables each with their own memory a Union allows for multiple names to the same variable.

These names can treat the memory as different types (and the size of the union will be the size of the largest type, + any padding the compiler might decide to give it)

So if you wanted to be able to read a variable's memory in different ways, for example read an integer one byte at a time, you could have something like this:

```c
union intParts {
  int theInt;
  char bytes[sizeof(int)];
};
```

Allowing you to look at each byte individually without casting a pointer and using pointer arithmetic:

```c
union intParts parts;
parts.theInt = 5968145; // arbitrary number > 255 (1 byte)

printf("The int is %i\nThe bytes are [%i, %i, %i, %i]\n",
parts.theInt, parts.bytes[0], parts.bytes[1], parts.bytes[2], parts.bytes[3]);

// vs

int theInt = parts.theInt;
printf("The int is %i\nThe bytes are [%i, %i, %i, %i]\n",
theInt, *((char*)&theInt+0), *((char*)&theInt+1), *((char*)&theInt+2), *((char*)&theInt+3));

// or with array syntax which can be a tiny bit nicer sometimes

printf("The int is %i\nThe bytes are [%i, %i, %i, %i]\n",
    theInt, ((char*)&theInt)[0], ((char*)&theInt)[1], ((char*)&theInt)[2], ((char*)&theInt)[3]);

/*
  The int is 5968145
  The bytes are [17, 17, 91, 0]
*/
```

Combining this with a structure allows you to create a "tagged" union which can be used to store multiple different types, one at a time.

For example, you might have a "number" struct, but you don't want to use something like this:

```c
struct operator {
    int intNum;
    float floatNum;
    int type;
    double doubleNum;
};
```

Because your program has a lot of them and it takes a bit too much memory for all of the variables, so you could use this:

```c
struct operator {
    int type;
    union {
      int intNum;
      float floatNum;
      double doubleNum;
    } types;
};
```

Like this the size of the struct is just the size of the int type + the size of the largest type in the union (the double). Not a huge gain, only 8 or 16 bytes, but the concept can be applied to similar structs.

use:

```c
operator op;
op.type = 0; // int, probably better as an enum or macro constant
op.types.intNum = 352;
```

Also, if you don't give the union a name then it's members are accessed directly from the struct:

```c
struct operator {
    int type;
    union {
        int intNum;
        float floatNum;
        double doubleNum;
    }; // no name!
};

operator op;
op.type = 0; // int
// intNum is part of the union, but since it's not named you access it directly off the struct itself
op.intNum = 352;
```

Another, perhaps more useful feature, is when you always have multiple variables of the same type, and you want to be able to use both names (for readability) and indexes (for ease of iteration), in that case you can do something like this:

```c
union Coins {
    struct {
        int quarter;
        int dime;
        int nickel;
        int penny;
    }; // anonymous struct acts the same way as an anonymous union, members are on the outer container
    int coins[4];
};
```

In that example you can see that there is a struct which contains the four (common) coins in the United States.

since the union makes the variables share the same memory the coins array matches with each int in the struct (in order):

```c
union Coins change;
for(int i = 0; i < sizeof(change) / sizeof(int); ++i)
{
    scanf("%i", change.coins + i); // BAD code! input is always suspect!
}
printf("There are %i quarters, %i dimes, %i nickels, and %i pennies\n",
    change.quarter, change.dime, change.nickel, change.penny);
```

## Pointer Arithmetic

https://www.learn-c.org/en/Pointer_Arithmetics

### From Claude:

```c
#include <stdio.h>

int main() {
    int arr[] = {10, 20, 30, 40, 50};
    int *ptr = arr;  // Point to first element

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
        ptr++;  // Move pointer to next element
    }

    return 0;
}
```

Key points:

- `ptr + 1` doesn't add 1 byte: it adds `sizeof(int)` bytes (usually 4)
- `arr[i]` is exactly equivalent to `*(ptr + i)`
- `ptr++` moves to the next element of that type

Where/why to use:

- String processing: Walking through characters
- Dynamic arrays: When you don't know size at compile time
- Performance: Sometimes faster than array indexing
- C library functions: Many expect pointers (like `strcpy`, `memcpy`)

Pointer arithmetic is fundamental to how arrays work in C: `arr[i]` is just syntactic sugar for pointer arithmetic.

## Function Pointers

https://www.learn-c.org/en/Function_Pointers

### Function Pointer Syntax

`void (*pf)(int);`

- `*pf` is the function pointer
- `void` is the return type
- `int` is the argument type

```c
#include <stdio.h>

void someFunction(int arg)
{
    printf("This is someFunction being called and arg is: %d\n", arg);
    printf("Whoops leaving the function now!\n");
}

main()
{
    void (*pf)(int);
    pf = &someFunction;
    printf("We're about to call someFunction() using a pointer!\n");
    (pf)(5);
    printf("Wow that was cool. Back to main now!\n\n");
}
```

```c
#include <stdio.h>
#include <stdlib.h> //for qsort()

int compare(const void* left, const void* right)
{
    return (*(int*)right - *(int*)left);
    // go back to ref if this seems complicated: https://cplusplus.com/reference/cstdlib/qsort/
}
main()
{
    int (*cmp) (const void* , const void*);
    cmp = &compare;

    int iarray[] = {1,2,3,4,5,6,7,8,9};
    qsort(iarray, sizeof(iarray)/sizeof(*iarray), sizeof(*iarray), cmp);

    int c = 0;
    while (c < sizeof(iarray)/sizeof(*iarray))
    {
        printf("%d \t", iarray[c]);
        c++;
    }
}
```

### From Claude:

This example shows how `qsort()` uses function pointers for custom sorting.

**The compare function:**

```c
cint compare(const void* left, const void* right)
{
    return (*(int*)right - *(int*)left);
}
```

- Takes `void*` pointers (generic pointers that can point to any type)
- Casts them to `int*` then dereferences: `*(int*)left` gets the actual int value
- Returns `right - left` (descending order):
  - Positive: left > right, swap needed
  - Negative: left < right, keep order
  - Zero: equal

**Function pointer setup:**

```c
cint (*cmp) (const void* , const void*);
cmp = &compare;
```

- Declares a function pointer matching `qsort()`'s expected signature
- Points it to your compare function

**qsort() call:**

```c
cqsort(iarray, sizeof(iarray)/sizeof(*iarray), sizeof(*iarray), cmp);
```

- `iarray`: array to sort
- `sizeof(iarray)/sizeof(*iarray)`: number of elements (9)
- `sizeof(*iarray)`: size of each element (4 bytes for int)
- `cmp`: your comparison function

Result: Array gets sorted in descending order: `9 8 7 6 5 4 3 2 1`

**The key insight:** `qsort()` doesn't know how to compare your data types, so you provide a function pointer that tells it how to compare two elements.

### Claude example usage

```c
#include <stdio.h>

// Function pointer type definition
typedef int (*operation)(int, int);

// Simple math functions
int add(int a, int b) { return a + b; }
int subtract(int a, int b) { return a - b; }
int multiply(int a, int b) { return a * b; }

// Function that takes a function pointer as parameter
int calculate(int x, int y, operation op) {
    return op(x, y);
}

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
```

Why use function pointers?

- Runtime selection: Choose which function to call based on user input or conditions
- Callbacks: Pass functions to other functions (like event handlers)
- Polymorphism: Different functions with same signature can be swapped easily
- Cleaner code: Avoid long switch statements

Function pointers let you treat functions as data - store them, pass them around, and call them dynamically.

## Bitmasks

https://www.learn-c.org/en/Bitmasks

Bit masking is simply the process of storing data truly as bits, as opposed to storing it as chars/ints/floats. It is incredibly useful for storing certain types of data compactly and efficiently.

Bit masking is based on boolean logic (logic gates)

- NOT a - the final value is the opposite of the input value (1 -> 0, 0 -> 1)
- a AND b - if both values are 1, the final value is 1, otherwise the final value is 0
- a OR b - if either value is 1, the final value is 1, otherwise the final value is 0
- a XOR b - if one value is 1 and the other value is 0, the final value is 1, otherwise the final value is 0

C guarantees that certain primitives are at least some number of bytes in size.

https://en.wikipedia.org/wiki/C_data_types#Basic_types

Bit masks are often used when setting flags. Flags are values that can be in two states, such as 'on/off' and 'moving/stationary'.

**Setting bit n**

Setting bit `n` is as simple as ORing the value of the storage variable with the value `2^n`.

`storage |= 1 << n;`

As an example, here is the setting of bit 3 where storage is a char (8 bits):

`01000010 OR 00001000 == 01001010`

The `2^n` logic places the '1' value at the proper bit in the mask itself, allowing access to that same bit in the storage variable.

**Clearing bit n**

Clearing bit `n` is the result of ANDing the value of the storage variable with the inverse (NOT) of the value 2^n:

`storage &= ~(1 << n);`

Here's the example again:

`01001010 AND 11110111 == 01000010`

**Flipping bit n**

Flipping bit `n` is the result of XORing the value of the storage variable with `2^n`:

`storage ^= 1 << n;`

`01000010 01001010 XOR XOR 00001000 00001000 == == 01001010 01000010`

**Checking bit n**

Checking a bit is ANDing the value of `2^n` with the bit storage:

`bit = storage & (1 << n);`

`01000010 01001010 AND AND 00001000 00001000 == == 00000000 00001000`
