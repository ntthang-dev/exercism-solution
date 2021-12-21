/* C program to find Armstrong number */

#include "armstrong_numbers.h"

// Function to check whether the given number is
// Armstrong number or not
bool is_armstrong_number(int number)
{
    int exponent;
    int final = number;
    int sum = 0;   

    // Find total digits in number
    exponent = floor(log10(number) + 1);

    // Calculate sum of power of digits
    while (number > 10) {
        int digit = number % 10;
        sum += pow(digit, exponent);
        number /= 10;
    }
    // Check for Armstrong number
    return sum == final; 
}