Resister finder is a small utility to generate a list of all possible resistance values that can be created from a given set of fixed value resistors and the number of resistors you are willing to combine together in any combination of series or parallel.

Set the list of resistor values in the resistor_values array near the top of resistor.go.

Set the number of resistors to combine by changing the value of the DEPTH constant in the main function within resistor.go.

The program generates the output on the standard-output with two columns:
* the resistance value obtained by the combination
* the arrangement and values of the combination of resistors. '+' indicates a series connection, '|' indicates a parallel connection.

The output is sorted in order of resistance value.

No attempt is made to eliminate duplicate combinations or values.
