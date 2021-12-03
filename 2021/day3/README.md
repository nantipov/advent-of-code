# Day 3

## Problem

### --- Day 3: Binary Diagnostic ---
The submarine has been making some odd creaking noises, so you ask it to produce a diagnostic report just in case.

The diagnostic report (your puzzle input) consists of a list of binary numbers which, when decoded properly, can tell you many useful things about the conditions of the submarine. The first parameter to check is the power consumption.

You need to use the binary numbers in the diagnostic report to generate two new binary numbers (called the gamma rate and the epsilon rate). The power consumption can then be found by multiplying the gamma rate by the epsilon rate.

Each bit in the gamma rate can be determined by finding the most common bit in the corresponding position of all numbers in the diagnostic report. For example, given the following diagnostic report:
```
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
```

Considering only the first bit of each number, there are five 0 bits and seven 1 bits. Since the most common bit is `1`, the first bit of the gamma rate is `1`.

The most common second bit of the numbers in the diagnostic report is `0`, so the second bit of the gamma rate is `0`.

The most common value of the third, fourth, and fifth bits are `1`, `1`, and `0`, respectively, and so the final three bits of the gamma rate are `110`.

So, the gamma rate is the binary number `10110`, or `22` in decimal.

The epsilon rate is calculated in a similar way; rather than use the most common bit, the least common bit from each position is used. So, the epsilon rate is 01001, or 9 in decimal. Multiplying the gamma rate (`22`) by the epsilon rate (`9`) produces the power consumption, `198`.

Use the binary numbers in your diagnostic report to calculate the gamma rate and epsilon rate, then multiply them together. **What is the power consumption of the submarine?** (Be sure to represent your answer in decimal, not binary.)

### --- Part Two ---
Next, you should verify the life support rating, which can be determined by multiplying the oxygen generator rating by the CO2 scrubber rating.

Both the oxygen generator rating and the CO2 scrubber rating are values that can be found in your diagnostic report - finding them is the tricky part. Both values are located using a similar process that involves filtering out values until only one remains. Before searching for either rating value, start with the full list of binary numbers from your diagnostic report and consider just the first bit of those numbers. Then:

- Keep only numbers selected by the bit criteria for the type of rating value for which you are searching. Discard numbers which do not match the bit criteria.
- If you only have one number left, stop; this is the rating value for which you are searching.
- Otherwise, repeat the process, considering the next bit to the right.

The bit criteria depends on which type of rating value you want to find:

- To find **oxygen generator rating**, determine the most common value (0 or 1) in the current bit position, and keep only numbers with that bit in that position. If 0 and 1 are equally common, keep values with a 1 in the position being considered.
- To find **CO2 scrubber rating**, determine the least common value (0 or 1) in the current bit position, and keep only numbers with that bit in that position. If 0 and 1 are equally common, keep values with a 0 in the position being considered.

For example, to determine the **oxygen generator rating** value using the same example diagnostic report from above:

- Start with all 12 numbers and consider only the first bit of each number. There are more `1` bits (7) than `0` bits (5), so keep only the 7 numbers with a 1 in the first position: `11110`, `10110`, `10111`, `10101`, `11100`, `10000`, and `11001`.
- Then, consider the second bit of the 7 remaining numbers: there are more 0 bits (4) than 1 bits (3), so keep only the 4 numbers with a 0 in the second position: `10110`, `10111`, `10101`, and `10000`.
- In the third position, three of the four numbers have a `1`, so keep those three: `10110`, `10111`, and `10101`.
- In the fourth position, two of the three numbers have a `1`, so keep those two: `10110` and `10111`.
- In the fifth position, there are an equal number of `0` bits and `1` bits (one each). So, to find the oxygen generator rating, keep the number with a 1 in that position: `10111`.

As there is only one number left, stop; the **oxygen generator rating** is `10111`, or `23` in decimal.

Then, to determine the **CO2 scrubber rating** value from the same example above:

- Start again with all 12 numbers and consider only the first bit of each number. There are fewer `0` bits (5) than `1` bits (7), so keep only the 5 numbers with a `0` in the first position: `00100`, `01111`, `00111`, `00010`, and `01010`.
- Then, consider the second bit of the 5 remaining numbers: there are fewer `1` bits (2) than `0` bits (3), so keep only the 2 numbers with a `1` in the second position: `01111` and `01010`.
- In the third position, there are an equal number of `0` bits and `1` bits (one each). So, to find the CO2 scrubber rating, keep the number with a 0 in that position: 01010.
- As there is only one number left, stop; the **CO2 scrubber rating** is `01010`, or `10` in decimal.

Finally, to find the life support rating, multiply the **oxygen generator rating** (`23`) by the **CO2 scrubber rating** (`10`) to get `230`.

Use the binary numbers in your diagnostic report to calculate the oxygen generator rating and **CO2 scrubber rating**, then multiply them together. **What is the life support rating of the submarine?** (Be sure to represent your answer in decimal, not binary.)

## Calculations
```shell
$ go build day3_a.go
$ go build day3_b.go

$ ./day3_a input.txt
> 3320834

$ ./day3_b input.txt
> [current ]id= 1 , q= 1000 , q0= 499 , q1= 501 , pos= 1 , value= 010100010011
> OXY follow to 1 (majority or equal) 
> [next    ]id= 2 , q= 501 , q0= 252 , q1= 249 , pos= 2 , value= 101111100010
> [current ]id= 2 , q= 501 , q0= 252 , q1= 249 , pos= 2 , value= 101111100010
> OXY follow to 0 (majority) 
> [next    ]id= 3 , q= 252 , q0= 120 , q1= 132 , pos= 3 , value= 101111100010
> [current ]id= 3 , q= 252 , q0= 120 , q1= 132 , pos= 3 , value= 101111100010
> OXY follow to 1 (majority or equal) 
> [next    ]id= 4 , q= 132 , q0= 68 , q1= 64 , pos= 4 , value= 101111100010
> [current ]id= 4 , q= 132 , q0= 68 , q1= 64 , pos= 4 , value= 101111100010
> OXY follow to 0 (majority) 
> [next    ]id= 5 , q= 68 , q0= 28 , q1= 40 , pos= 5 , value= 101000001101
> [current ]id= 5 , q= 68 , q0= 28 , q1= 40 , pos= 5 , value= 101000001101
> OXY follow to 1 (majority or equal) 
> [next    ]id= 6 , q= 40 , q0= 20 , q1= 20 , pos= 6 , value= 101010010110
> [current ]id= 6 , q= 40 , q0= 20 , q1= 20 , pos= 6 , value= 101010010110
> OXY follow to 1 (majority or equal) 
> [next    ]id= 7 , q= 20 , q0= 7 , q1= 13 , pos= 7 , value= 101011000100
> [current ]id= 7 , q= 20 , q0= 7 , q1= 13 , pos= 7 , value= 101011000100
> OXY follow to 1 (majority or equal) 
> [next    ]id= 8 , q= 13 , q0= 7 , q1= 6 , pos= 8 , value= 101011100011
> [current ]id= 8 , q= 13 , q0= 7 , q1= 6 , pos= 8 , value= 101011100011
> OXY follow to 0 (majority) 
> [next    ]id= 302 , q= 7 , q0= 3 , q1= 4 , pos= 9 , value= 101011100011
> [current ]id= 302 , q= 7 , q0= 3 , q1= 4 , pos= 9 , value= 101011100011
> OXY follow to 1 (majority or equal) 
> [next    ]id= 1204 , q= 4 , q0= 1 , q1= 3 , pos= 10 , value= 101011101100
> [current ]id= 1204 , q= 4 , q0= 1 , q1= 3 , pos= 10 , value= 101011101100
> OXY follow to 1 (majority or equal) 
> [next    ]id= 1700 , q= 3 , q0= 1 , q1= 2 , pos= 11 , value= 101011101100
> [current ]id= 1700 , q= 3 , q0= 1 , q1= 2 , pos= 11 , value= 101011101100
> OXY follow to 1 (majority or equal) 
> [next    ]id= 1701 , q= 2 , q0= 1 , q1= 1 , pos= 12 , value= 101011101111
> [current ]id= 1701 , q= 2 , q0= 1 , q1= 1 , pos= 12 , value= 101011101111
> OXY follow to 1 (majority or equal) 
> [next    ]
> [current ]id= 1 , q= 1000 , q0= 499 , q1= 501 , pos= 1 , value= 010100010011
> CO2 follow to 0 (minority or equal) 
> [next    ]id= 28 , q= 499 , q0= 252 , q1= 247 , pos= 2 , value= 010100010011
> [current ]id= 28 , q= 499 , q0= 252 , q1= 247 , pos= 2 , value= 010100010011
> CO2 follow to 1 (minority) 
> [next    ]id= 29 , q= 247 , q0= 128 , q1= 119 , pos= 3 , value= 010100010011
> [current ]id= 29 , q= 247 , q0= 128 , q1= 119 , pos= 3 , value= 010100010011
> CO2 follow to 1 (minority) 
> [next    ]id= 227 , q= 119 , q0= 53 , q1= 66 , pos= 4 , value= 011010100110
> [current ]id= 227 , q= 119 , q0= 53 , q1= 66 , pos= 4 , value= 011010100110
> CO2 follow to 0 (minority or equal) 
> [next    ]id= 228 , q= 53 , q0= 24 , q1= 29 , pos= 5 , value= 011010100110
> [current ]id= 228 , q= 53 , q0= 24 , q1= 29 , pos= 5 , value= 011010100110
> CO2 follow to 0 (minority or equal) 
> [next    ]id= 587 , q= 24 , q0= 13 , q1= 11 , pos= 6 , value= 011000011010
> [current ]id= 587 , q= 24 , q0= 13 , q1= 11 , pos= 6 , value= 011000011010
> CO2 follow to 1 (minority) 
> [next    ]id= 588 , q= 11 , q0= 5 , q1= 6 , pos= 7 , value= 011001110000
> [current ]id= 588 , q= 11 , q0= 5 , q1= 6 , pos= 7 , value= 011001110000
> CO2 follow to 0 (minority or equal) 
> [next    ]id= 589 , q= 5 , q0= 2 , q1= 3 , pos= 8 , value= 011001010010
> [current ]id= 589 , q= 5 , q0= 2 , q1= 3 , pos= 8 , value= 011001010010
> CO2 follow to 0 (minority or equal) 
> [next    ]id= 590 , q= 2 , q0= 1 , q1= 1 , pos= 9 , value= 011001000001
> [current ]id= 590 , q= 2 , q0= 1 , q1= 1 , pos= 9 , value= 011001000001
> CO2 follow to 0 (minority or equal) 
> [next    ]id= 917 , q= 1 , q0= 1 , q1= 0 , pos= 10 , value= 011001000001
> oxygen  101011101111   2799
> co2     011001000001   1601
> 4481199
```
