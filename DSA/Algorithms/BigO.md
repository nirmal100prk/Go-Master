# Big O Notation

provides a way to measure how the runtime or memory usage of an algorithm grows as the input size increases

## Why is it important?
Predicting performance: It helps you estimate how an algorithm will perform with larger datasets.   
Comparing algorithms: You can evaluate different approaches to a problem and choose the most efficient one.
Optimizing code: It identifies potential bottlenecks and areas for improvement in your code.

Big O notation typically represents the worst-case scenario

Big O notations

O(1) Constant time - the runtime is independent of the input size

O(log n) Logarithmic time - the runtime grows logarithmically with the input size 
O(n) Linear time - the runtime grows directly with the input size
O(n log n) Linearithmic time - the runtime grows linearly with the logarithm of
the input size
O(n^2) Quadratic time - the runtime grows quadratically with the input size
O(2^n) Exponential time - the runtime grows exponentially with the input size
O(n!) Factorial time - the runtime grows factorially with the input size

## Logarithm time complexity explained
Logarithmic time complexity is a time complexity class that describes an algorithm whose running time grows logarithm
with the size of the input. This means that as the input size increases, the running time of the algorithm increases, but at a much slower rate.

## Common Examples of O(log n) Algorithms
Binary search: Efficiently finding a specific value within a sorted array by repeatedly dividing the search interval in half.
Calculating the height of a balanced binary tree: The height of a balanced binary tree is logarithmic in relation to the number of nodes.

Exponents: When we raise a number to a power, we multiply the number by itself repeatedly. This leads to rapid growth. For example, 2^10 is 1024.
Logarithms: A logarithm is the inverse of an exponent. It answers the question: "To what power must we raise a given number (the base) to get a specific value?" For instance, log₂(1024) = 10.

## John Napier introduced logarithm

Big O notation is used to classify algorithms according to how their run time or space requirements grow as the input size grows. It provides a way to talk about the efficiency of algorithms in a way that abstracts away constant factors and lower-order terms, focusing on the dominant term that describes the growth rate.

