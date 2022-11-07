#include "hello.h"
#include <stdio.h>

int hello(const char *name, int year, char *out) {
    int n;
    n = sprintf(out, "Greetings, %s from %d! We come in peace :)", name, year);
    return n;
}

char* hello_alloc(const char *name, int year) {
    size_t needed = snprintf(NULL, 0, "Greetings, %s from %d! We come in peace :)", name, year) + 1;
    char *str = malloc(needed);
    sprintf(str, "Greetings, %s from %d! We come in peace :)", name, year);
    return str;
}

typedef struct {
    char * name;
    int age;
} user;

user new_user(char *name, int age) {
    return (user){
        .name = name,
        .age = age,
    };
}

void process_user(user * u) {
    u->name = "";
    u->age = 0;
}