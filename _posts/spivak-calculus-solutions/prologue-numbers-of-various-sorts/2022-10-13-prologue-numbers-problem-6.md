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

  We can even derive $\sum k$ using the following equality:

  \[
    \begin{align*}
      (k + 1)^2 &= k^2 + 2k + 1 \\
      (k + 1)^2 - k^2 &= 2k + 1 \\
    \end{align*}
  \]

  With that we get:

  \[
    \begin{align*}
      2^2 - 1^2 &= 2 \cdot 1 + 1 \\
      3^2 - 2^2 &= 2 \cdot 2 + 1 \\
      &. \\
      &. \\
      &. \\
      (n + 1)^2 - n^2 &= 2 \cdot n + 1 \\
      ------ \\
      (n + 1)^2 - 1 &= 2 \cdot [1 + ... + n] + n \\
      \\
      \sum_{i = 1}^n i &= \frac{(n + 1)^2 - n - 1}{2} \\
      &= \frac{n^2 + 2n + 1 - n - 1}{2} \\
      &= \frac{n^2 + n}{2} \\
      &= \frac{n(n + 1)}{2} \\
    \end{align*}
  \]

  Using that result for $\sum k^2$ we get:

  \[
    \begin{align*}
      (n + 1)^3 - 1 &= 3 \cdot [1^2 + ... + n^2] + 3 \cdot [1 + ... + n] + n \\
      \sum_{i = 1}^n i^2 &= \frac{(n + 1)^3 - 1 - 3 \sum_{i = 1}^n i - n}{3} \\
      &= \frac{n^3 + 3n^2 + 3n + 1 - 1 - 3 \sum_{i = 1}^n i - n}{3} \\
      &= \frac{n^3 + 3n^2 + 2n - 3 \frac{n(n + 1)}{2}}{3} \\
      &= \frac{n^3 + 3n^2 + 2n - \frac{3n^2}{2} - \frac{3n}{2}}{3} \\
      &= \frac{n^3 + \frac{3n^2}{2} + \frac{n}{2}}{3} \\
    \end{align*}
  \]

</p>


<p>
  <strong>(i)</strong> $1^3 + ... + n^3$

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
      \sum_{i = 1}^n i^3 &= \frac{(n + 1)^4 - 1 - 6 \cdot \sum_{i = 1}^n i^2 - 4 \cdot \sum_{i = 1}^n i - n}{4}
    \end{align*}
  \]

  We know $\sum_{i = 1}^n i^2$ and $\sum_{i = 1}^n i$ already.

</p>

<p>
  <strong>(ii)</strong> $1^4 + ... + n^4$

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
      \sum_{i = 1}^n i^4 &= \frac{(n + 1)^5 - 1 - 10 \cdot \sum_{i = 1}^n i^3 - 10 \cdot \sum_{i = 1}^n i^2 - 5 \cdot \sum_{i = 1}^n i - n}{5}
    \end{align*}
  \]
</p>

<p>

  <strong>(iii)</strong> $\frac{1}{1 \cdot 2} + ... + \frac{1}{n \cdot (n + 1)}$

  \[
    \begin{align*}
      \frac{1}{k(k + 1)} &= \frac{1}{k} - \frac{1}{k + 1} \\
    \end{align*}
  \]

  Using that we can apply $k = 1, ..., n$:

  \[
    \begin{align*}
      \frac{1}{1 \cdot 2} &= \frac{1}{1} - \frac{1}{2} \\
      &. \\
      &. \\
      &. \\
      \frac{1}{n \cdot (n + 1)} &= \frac{1}{n} - \frac{1}{n + 1} \\
      ------ \\
      [\frac{1}{1 \cdot 2} + ... + \frac{1}{n \cdot (n + 1)}] &= [\frac{1}{1} + ... + \frac{1}{n}] - [\frac{1}{2} + ... + \frac{1}{n + 1}] \\
      \sum_{i = 1}^n \frac{1}{n(n + 1)} &= \sum_{i = 1}^n \frac{1}{i} - \sum_{i = 1}^n \frac{1}{i + 1} \\
      &= \frac{1}{1} - \frac{1}{n + 1} \\
      &= \frac{n + 1 - 1}{n + 1} \\
      &= \frac{n}{n + 1} \\
    \end{align*}
  \]
</p>

<p>
  <strong>(iv)</strong> $\frac{3}{1^2 \cdot 2^2} + ... + \frac{2n + 1}{n^2 \cdot (n + 1)^2}$

  \[
    \begin{align*}
      \frac{2k + 1}{k^2(k + 1)^2} &= \frac{1}{k^2} - \frac{1}{(k + 1)^2} \\
    \end{align*}
  \]

  Using that we can apply $k = 1, ..., n$:

  \[
    \begin{align*}
      \frac{2 + 1}{1^2 2^2} &= \frac{1}{1^2} - \frac{1}{2^2} \\
      &. \\
      &. \\
      &. \\
      \frac{2n + 1}{n^2(n + 1)^2} &= \frac{1}{n^2} - \frac{1}{(n + 1)^2} \\
      ------ \\
      [\frac{3}{1^2 2^2} + ... + \frac{2n + 1}{n^2(n + 1)^2}] &= [\frac{1}{1^2} + ... + \frac{1}{n^2}] - [\frac{1}{2^2} + ... + \frac{1}{(n + 1)^2}] \\
      \sum_{i = 1}^n \frac{2i + 1}{i^2(i + 1)^2} &= \sum_{i = 1}^n \frac{1}{i^2} - \sum_{i = 1}^n \frac{1}{(i + 1)^2} \\
      &= \frac{1}{1} - \frac{1}{(n + 1)^2} \\
      &= \frac{(n + 1)^2 - 1}{(n + 1)^2} \\
      &= \frac{n^2 + 2n + 1 - 1}{(n + 1)^2} \\
      &= \frac{n(n + 2)}{(n + 1)^2} \\
    \end{align*}
  \]

</p>
