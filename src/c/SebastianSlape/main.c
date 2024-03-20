#pragma GCC optimize("O3")
#pragma GCC optimize("unroll-loops")

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
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

		// If I squish the result to an int16 by clamping, lineTotal will always be equal to 65,535 as lineTotal never goes below 30,000,000 let alone below 65,535.
		//
		// I looked at your implementation and when you "squish" the sum to an int16 using the line "sum := uint16(rawSum)" it actually trunicates the rest of the digits
		// That means that a number like 33033787 will be trunicated to 3643.
		//
		//  Decimal                Binary
		// 33033787: 00000001 11111000 00001110 00111011
		//     3643:                   00001110 00111011
		//
		// I ended up doing what you did by using the "and" operator on 65535 (11111111 11111111). If this is correct and by "squish" you meant trunicate the binary then please make the 
		// infomation page clearer about this as this might be the most important part of the challenge.
		lineTotal = lineTotal & 65535;

		int64_t lineFinal = (int64_t)(lineTotal * squareRoot(lineTotal));

		if (max < lineFinal) {
			max = lineFinal;
		}
		if (min > lineFinal) {
			min = lineFinal;
		}

		total += lineFinal;

		count++;
	}

	double mean = (int64_t)(total / count);

	printf("MIN: %f\n",min);
	printf("MAX: %f\n",max);
	printf("MEAN: %f",mean);

	// Close the file
	fclose(fptr);
}
