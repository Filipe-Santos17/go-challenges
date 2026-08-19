# Guardian of Prime Numbers

In a simplified cryptography system, security depends on the choice of prime numbers. You are responsible for implementing the function that validates whether a positive integer received from the user is prime.

The program must:

- Receive an integer N (1 ≤ N ≤ 10⁹) via standard input or as a command-line argument.
- Implement a function `ehPrimo(n int) bool` that returns `true` if `n` is prime and `false` otherwise.
- For very large numbers, the function must be efficient — it is not enough to test all divisors up to N−1.
- The program must display a message: "O número X é primo" or "O número X não é primo".
- If the number is less than 2, treat it as a special case (not prime).
