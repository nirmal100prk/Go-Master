Interpolation Search is a search algorithm that improves upon binary search by making more informed guesses about the position of the target element within a sorted array.

Array should be uniformly disrtibuted

Efficiency: Interpolation search can be significantly faster than binary search for uniformly distributed data.
Assumption: The algorithm assumes a uniform distribution of values. If the data is not uniformly distributed, the performance might degrade.

Worst-case: In the worst case, interpolation search has the same time complexity as linear search, O(n).
Best-case: In the best case, the time complexity is O(log log n) for uniformly distributed data.

# pos = low + [(x - arr[low]) * (high - low)] / (arr[high] - arr[low])

The formula essentially calculates the estimated position of the target element within the search space based on the relative positions of the target value compared to the values at the lowest and highest indices.

the interpolation search estimates a position based on the distribution of values in the array.