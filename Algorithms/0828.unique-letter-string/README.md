<<<<<<< HEAD
# [828. Unique Letter String](https://leetcode.com/problems/unique-letter-string/)

## 题目

A character is unique in string S if it occurs exactly once in it.

For example, in string S = "LETTER", the only unique characters are "L" and "R".

Let's define UNIQ(S) as the number of unique characters in string S.

For example, UNIQ("LETTER") = 2.

Given a string S with only uppercases, calculate the sum of UNIQ(substring) over all non-empty substrings of S.

If there are two or more equal substrings at different positions in S, we consider them different.

Since the answer can be very large, retrun the answermodulo `10^9+7`.

Example 1:

```text
Input: "ABC"
Output: 10
Explanation: All possible substrings are: "A","B","C","AB","BC" and "ABC".
Evey substring is composed with only unique letters.
Sum of lengths of all substring is 1 + 1 + 1 + 2 + 2 + 3 = 10
```

Example 2:

```text
Input: "ABA"
Output: 8
Explanation: The same as example 1, except uni("ABA") = 1.
```

Note: 0 <= S.length <= 10000.

## 解题思路

from <https://leetcode.com/problems/unique-letter-string/discuss/128952/One-pass-O(N)-Straight-Forward>

Intuition:

Let's think about how a character can be found as a unique character.

Think about string "XAXAXXAX" and focus on making the second "A" a unique character.
We can take "XA(XAXX)AX" and between "()" is our substring.
We can see here, to make the second "A" counted as a uniq character, we need to:

insert "(" somewhere between the first and second A
insert ")" somewhere between the second and third A
For step 1 we have "A(XA" and "AX(A", 2 possibility.
For step 2 we have "A)XXA", "AX)XA" and "AXX)A", 3 possibilities.

So there are in total 2 * 3 = 6 ways to make the second A a unique character in a substring.
In other words, there are only 6 substring, in which this A contribute 1 point as unique string.

Instead of counting all unique characters and struggling with all possible substrings,
we can count for every char in S, how many ways to be found as a unique char.
We count and sum, and it will be out answer.

Explanation:

index[26][2] record last two occurrence index for every upper characters.
Initialise all values in index to -1.
Loop on string S, for every character c, update its last two occurrence index to index[c].
Count when loop. For example, if "A" appears twice at index 3, 6, 9 seperately, we need to count:
For the first "A": (6-3) * (3-(-1))"
For the second "A": (9-6) * (6-3)"
For the third "A": (N-9) * (9-6)"
Complexity:
One pass, time complexity O(N).
Space complexity O(1).
=======
# Such a SHAME on codility.com

> below content come from https://github.com/github/dmca/blob/master/2020/03/2020-03-26-Codility.md

**Are you the copyright holder or authorized to act on the copyright owner's behalf?**

Yes, I am authorized to act on the copyright owner's behalf.

**Please describe the nature of your copyright ownership or authorization to act on the owner's behalf.**

Employee of Codility

**Please provide a detailed description of the original copyrighted work that has allegedly been infringed. If possible, include a URL to where it is posted online.**

Recruitment task available only to paying customers:
https://app.codility.com/tasks/unique_characters/

**What files should be taken down? Please provide URLs for each file, or if the entire repository, the repository’s URL.**

https://github.com/aQuaYi/LeetCode-in-Go/blob/master/Algorithms/0828.unique-letter-string/README.md
https://github.com/grandyang/leetcode/issues/828

**Have you searched for any forks of the allegedly infringing files or repositories? Each fork is a distinct repository and must be identified separately if you believe it is infringing and wish to have it taken down.**

First repository has 576 forks and second repository has 689 forks. Let us know if you wish to receive a separate notice for each of them

**Is the work licensed under an open source license? If so, which open source license? Are the allegedly infringing files being used under the open source license, or are they in violation of the license?**

No, not open source.

**What would be the best solution for the alleged infringement? Are there specific changes the other person can make other than removal? Can the repository be made private?**

Removal.

**Do you have the alleged infringer’s contact information? If so, please provide it.**

No.

**I have a good faith belief that use of the copyrighted materials described above on the infringing web pages is not authorized by the copyright owner, or its agent, or the law.**

**I have taken <a href="https://www.lumendatabase.org/topics/22">fair use</a> into consideration.**

**I swear, under penalty of perjury, that the information in this notification is accurate and that I am the copyright owner, or am authorized to act on behalf of the owner, of an exclusive right that is allegedly infringed.**

**I have read and understand GitHub's <a href="https://help.github.com/articles/guide-to-submitting-a-dmca-takedown-notice/">Guide to Submitting a DMCA Takedown Notice</a>.**

**So that we can get back to you, please provide either your telephone number or physical address.**

[private]

**Please type your full legal name below to sign this request.**

[private]
>>>>>>> f33c3a477711033e1c5c5c04e72ce2c3c83f449e
