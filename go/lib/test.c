#include "test.h"

char *test(char *s){
    if(*s == '\0'){
        return "c";
    }else{
        return s;
    }
}