# 39. Combination Sum

Given an array of distinct integers `candidates` and a target integer `target`, return a list of all __*unique combinations*__ of `candidates` where the chosen numbers sum to `target`. You may return the combinations in any order.

The __*same*__ number may be chosen from `candidates` an __unlimited number of times__. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

It is __*guaranteed*__ that the number of unique combinations that sum up to `target` is less than `150` combinations for the given input.


__Example 1:__

```plain
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
```

__Exmaple 2:__

```
Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]
```

__Example 3:__

```
Input: candidates = [2], target = 1
Output: []
```

__Constraints:__

- `1 <= candidates.length <= 30`
- `1 <= candidates[i] <= 200`
- All elements of `candidates` are **distinct**.
- `1 <= target <= 500`
