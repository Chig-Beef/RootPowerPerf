#pragma GCC optimize("O3")
#pragma GCC optimize("unroll-loops")

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>
#include <stdint.h>

// Newton–Raphson method
double squareRoot(double value) {
	double guess = value;
	double root = 0;

	for (int a = 0; a < 20; a++) { 
        root = 0.5 * (guess + (value / guess));

        guess = root;
	}

	return root;
}

int main() {
	FILE* fptr;

	fptr = fopen("data.txt", "r");

	char line[6000];
	int count = 0;

	double total = 0;
	double max = 0;
	double min = 999999999999;

	while (fgets(line, 6000, fptr)) {
		char *p = strtok(line, ",");
		int i = 0;
		
		int64_t lineTotal = 0;

		while (p != NULL) {
			int convertedNumber = atoi(p);
			lineTotal += convertedNumber;
			p = strtok(NULL, ",");
		}

		lineTotal = lineTotal & 65535;

		double lineFinal = trunc(lineTotal * squareRoot(lineTotal));

		if (max < lineFinal) {
			max = lineFinal;
		}
		if (min > lineFinal) {
			min = lineFinal;
		}

		total += lineFinal;

		count++;
	}

	double mean = trunc(total / count);

	printf("MIN: %f\n",min);
	printf("MAX: %f\n",max);
	printf("MEAN: %f",mean);

	// Close the file
	fclose(fptr);
}
