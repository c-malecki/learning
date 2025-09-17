# Cheatsheet

```c
int compare(const void* left, const void* right)
{
    return (*(int*)right - *(int*)left);
}
```

- `const` prevents the function from modifying the data
- `void*` is a typecast for a generic pointer type (can point to any type)
- `*(int*)left`: `(int*)` is casting `left` to type `int` and `*(typecast)left` is dereferencing the pointer to get the value cast as an `int`

Example of what `const` prevents:

```c
int bad_compare(const void* left, const void* right) {
    *(int*)left = 999;  // Compiler error! Can't modify const data
    return *(int*)left - *(int*)right;
}
```

## Dynamic data

### `sizeof`

```c
int iarray[] = {1,2,3,4,5,6,7,8,9};
// int is 4 bytes
sizeof(iarray)/sizeof(*iarray) // 9
// iarray has 9 elements
```

Get the size of variable `iarray` in bytes and divide it by the size of the elements in `iarray`

### Pointer Arithmetic

Useful for:

- string processing and walking through characters
- iterating dynamic arrays when size isn't known at compile time
- c lib functions that expect pointers

```c
#include <stdio.h>

int main() {
    int arr[] = {10, 20, 30, 40, 50};
    int *ptr = arr;  // Point to first element

    // Traditional array indexing
    for (int i = 0; i < 5; i++) {
        printf("%d\n", arr[i]);
    }
    // 10 20 30 40 50

    printf("\n");

    // Pointer arithmetic - equivalent to above
    for (int i = 0; i < 5; i++) {
        printf("*%d\n", *(ptr + i));
        // ptr + i moves i positions forward
        // *(ptr + i) dereferences to get the value
    }
    // 10 20 30 40 50

    printf("\n");

    // Moving the pointer itself
    ptr = arr;  // Reset to start
    for (int i = 0; i < 5; i++) {
        printf("*ptr = %d\n", *ptr);
        ptr++;  // Move pointer to next element
    }

    return 0;
}
```

## Data types

Integer types:

- `char`: 1 byte, stores single characters or small integers (-128 to 127)
- `int`: typically 4 bytes, general-purpose integers
- `short`: typically 2 bytes, smaller range integers
- `long`: typically 8 bytes, larger range integers
- `long long`: at least 8 bytes, very large integers
- `void`: no type

Floating-point types:

- `float`: 4 bytes, single-precision decimal numbers (6-7 digits precision)
- `double`: 8 bytes, double-precision decimal numbers (15-17 digits precision)
- `long double`: extended precision (varies by system)

Modifiers:

- `unsigned`: makes integer types only positive, doubles positive range
- `signed`: explicitly signed (default for most types)

## Format specifiers

https://en.wikipedia.org/wiki/C_data_types#Basic_types

- char, signed char, unsigned char: `%c`
- short, short int, signed short, signed short int: `%hi` or `%hd`
- unsigned short, unsigned short int: `%hu`
- int, signed, signed int: `%i` or `%d`
- unsigned, unsigned int: `%u`
- long, long int, signed long, signed long int: `%li` or `%ld`
- unsigned long, unsigned long int: `%lu`
- long long, long long int, signed long long, signed long long int: `%lli` or `%lld`
- unsigned long long, unsigned long long int: `%llu`
- float: `%f` `%g` `%e` `%a` or `%F` `%G` `%E` `%A
- double: `%lf` `%lg` `%le` `%la` or `%lF` `%lG` `%lE` `%lA`
