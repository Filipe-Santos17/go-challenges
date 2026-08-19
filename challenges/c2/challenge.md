# Health Calculator

You are building a health app. There is a user registry with name, weight (kg) and height (m). When measuring a person again, the new values arrive separately and you need to update the existing record — not create a new one.

Implement a function that, given a user and new weight and height values, updates the original struct. Then, calculate the BMI (weight/height²) and classify: underweight (< 18.5), normal (18.5–24.9), overweight (25–29.9), obesity (≥ 30).

Hint: What guarantees that the original struct will be modified? Think about how to pass mutable data to functions in Go.
