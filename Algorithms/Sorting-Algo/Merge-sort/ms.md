Merge sort is a sorting algorithm that follows the <divide-and-conquer/> approach. It works by recursively dividing the input array into smaller subarrays and sorting those subarrays then merging them back together to obtain the sorted array.

Function Definitions:

mergeSort: The function that splits the array and sorts the subarrays.
merge: The function that merges two sorted subarrays.
Algorithm:

Divide: Recursively split the array into halves until you have subarrays of length 1 or 0.
Conquer: Sort the subarrays (this happens recursively).
Combine: Merge the sorted subarrays to produce the sorted result.