#include <stdio.h>
#include "person.h"

int main(int argc, char** argv) {
    fprintf(stderr, "C MAIN START\n");
    APerson * of;
    of = get_person("tim", "tim hughes");
    printf("GO BINDING>        Name: %s | Long Name: %s | \n", of->name, of->long_name);
    fprintf(stderr, "C MAIN END\n");
    return 0;
}
