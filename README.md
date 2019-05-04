# Sorted Segments

Reference implementation of the sorted segment lookup tree.

```
// soseg.Tree describes a list of weights sorted by unique keys.
// The tree also keeps track of the running total/sum of weights preceding each entry.
// Effectively, it's a specialized segment tree whose range entries all touch but don't overlap.
// The length of each range is equal to the entry weight.
```

Implements the Nimiq Validator List Selection Algorithm in sublinear time:
Algorithm 1 of https://katallassos.com/papers/Albatross.pdf
