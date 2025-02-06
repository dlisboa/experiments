#include <stdio.h>
#include <unistd.h>

int main(void) {
  printf("\033[2J"); // clear screen
  printf("\033[%d;%dH", 0, 0); // go to point (0,0)
  fflush(stdout);
  printf("sleeping 5 seconds...");
  fflush(stdout);
  for (int i=0; i<5; i++) {
    sleep(1);
  }
}
