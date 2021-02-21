// Exercise: Loops and Functions

// As a way to play with functions and loops, let's implement a square root funciton: five a numebr x, we want to find the number z for which z^2 is most early x.

// Computers typically compute the square root of x using a loop Starting with some guess z, we can adjust z based on how close z^2 is to x, producing a better guess:

//  z -= (z*z - x) / (2*z)