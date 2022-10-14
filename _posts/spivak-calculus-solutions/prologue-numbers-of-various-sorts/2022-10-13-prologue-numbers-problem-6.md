---
layout: math_post
title: Problem 6
date: 2022-10-13
categories: article spivak calculus prologue numbers
head_title: Problem 6
meta_description: Spivak Calculus Prologue Numbers Problem 6
meta_keywords: spivak,calculus,prologue,solutions
---

<p>

  The forumla for $1^2 + ... + n^2$ may be derived as follows. We begin with the formula:
  $$(k + 1)^3 - k^3 = 3k^2 + 3k + 1$$
  (For myself):

  \[
    \begin{align*}
      (k + 1)^3 - k^3 &= (k^2 + 2k + 1)(k + 1) - k^3 \\
      &=k^3 + k^2 + 2k^2 + 2k + k + 1 - k^3 \\
      &= 3k^2 + 3k + 1 \\
    \end{align*}
  \]

  Writing this formula for $k = 1, ..., n$ and adding, we obtain:

  \[
    \begin{align*}
      2^3 - 1^3 &= 3 + 3 + 1 \\
      3^3 - 2^3 &= 3 \cdot 2^2 + 3 \cdot 2 + 1 \\
      &. \\
      &. \\
      &. \\
      (n + 1)^3 - n^3 &= 3 \cdot n^2 + 3 \cdot n + 1 \\
      ------\\
      (n + 1)^3 - 1 &= 3 \cdot [1^2 + ... + n^2] + 3 \cdot [1 + ... + n] + n \\
    \end{align*}
  \]

</p>

<p>

  \[
    \begin{align*}
      (k + 1)^4 - k^4 &= (k^2 + 2k + 1)(k^2 + 2k + 1) - k^4 \\
      &= k^4 + 2k^3 + k^2 + 2k^3 + 4k^2 + 2k + k^2 + 2k + 1 - k^4 \\
      &= 4k^3 + 6k^2 + 4k + 1 \\
    \end{align*}
  \]

  Now we can write this formula for $k = 1, ..., n$ again:

  \[
    \begin{align*}
      2^4 - 1^4 &= 4 + 6 + 4 + 1 \\
      3^4 - 2^4 &= 4 \cdot 2^3 + 6 \cdot 2^2 + 4 \cdot 2 + 1 \\
      &. \\
      &. \\
      &. \\
      (n + 1)^4 - n^4 &= 4 \cdot n^3 + 6 \cdot n^2 + 4 \cdot n + 1 \\
      ------\\
      (n + 1)^4 - 1 &= 4 \cdot [1^3 + ... + n^3] + 6 \cdot [1^2 + ... + n^2] + 4 \cdot [1 + ... + n] + n \\
    \end{align*}
  \]

  We know $\sum_{k = 1}^n k^2$ already as long as we know $\sum_{k = 1}^n k$.

</p>

<p>

  \[
    \begin{align*}
      (k + 1)^5 - k^5 &= (k + 1)^5 - k^5 \\
      &= 5k^4 + 10k^3 + 10k^2 + 5k + 1 \\
    \end{align*}
  \]

  Now we can write this formula for $k = 1, ..., n$ again:

  \[
    \begin{align*}
      (n + 1)^5 - 1 &= 5 \cdot [1^4 + ... + n^4] + 10 \cdot [1^3 + ... + n^3] + 10 \cdot [1^2 + ... + n^2] + 5 \cdot [1 + ... + n] + n \\
    \end{align*}
  \]

</p>
