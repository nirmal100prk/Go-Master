Binary Search Algorithm is a searching algorithm used in a sorted array by repeatedly dividing the search interval in half. The idea of binary search is to use the information that the array is sorted and reduce the time complexity to O(log N). 

worst case O(log N)

The Binary Search Algorithm can be implemented in the following two ways

Iterative Binary Search Algorithm
Recursive Binary Search Algorithm



Initial Problem:

You have a sorted array of n elements, and you want to search for a specific element.
Binary Search Process:

Binary search compares the target element with the middle element of the array.
If the target element is equal to the middle element, the search is complete.
If the target is smaller than the middle element, you only need to search the left half of the array.
If the target is larger, you search the right half.
In each step, the array size is effectively halved.
Dividing Search Space:

In each iteration of binary search, the search space is reduced by half. If the array has n elements, after the first iteration, you're left with n/2 elements to search, then n/4, n/8, and so on.
This halving process continues until either the element is found or there are no more elements to search.

Number of Steps:
The maximum number of times you can halve the array is the number of times n can be divided by 2 until you reach 1 element. This is equivalent to the number of times you need to divide n by 2 to get 1, which is log2n.
Hence, the number of steps needed to reduce the search space to a single element is 
log2n.

Worst Case:
In the worst case, the element you're searching for may be the last element you check (i.e., you may have to halve the search space 

log2n times and find the element in the last possible place).
Therefore, in the worst case, binary search takes log2n steps.
Why O(logn)?
The time complexity is expressed as O(logn) because in algorithm analysis, we ignore constant factors and base of logarithms. The base of the logarithm is not significant in Big-O notation, so whether it's log2n, log10n, or logen, it simplifies to O(logn)