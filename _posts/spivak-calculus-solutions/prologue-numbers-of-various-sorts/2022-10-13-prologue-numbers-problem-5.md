---
layout: math_post
title: Problem 5
date: 2022-10-13
categories: article spivak calculus prologue numbers
head_title: Problem 5
meta_description: Spivak Calculus Prologue Numbers Problem 5
meta_keywords: spivak,calculus,prologue,solutions
---

<p>

  <strong>(a)</strong> Prove by induction on $n$ that:
  $$1 + r + r^2 + ... + r^n = \frac{1 - r^{n + 1}}{1 - r}$$
  If $r \neq 1$.
  <br>
  <br>
  Starting with $n = 0$:

  \[
    \begin{align*}
      \frac{1 - r^1}{1 - r} &= 1 &= r^0 \\
    \end{align*}
  \]

  Now, assuming $n = k$ show $n = k + 1$:

  \[
    \begin{align*}
      1 + r + ... + r^k + r^{k + 1} &= \frac{1 - r^{k + 1}}{1 - r} + r^{k + 1} \\
      &= \frac{1 - r^{k + 1}}{1 - r} + \frac{r^{k + 1}(1 - r)}{1 - r} \\
      &= \frac{1 - r^{k + 1}}{1 - r} + \frac{r^{k + 1} - r^{k + 2})}{1 - r} \\
      &= \frac{1 - r^{k + 1} + r^{k + 1} - r^{k + 2}}{1 - r} \\
      &= \frac{1 - r^{k + 2}}{1 - r} \\
      &= \frac{1 - r^{(k + 1) + 1}}{1 - r} \\
    \end{align*}
  \]

</p>

<p>

  <strong>(b)</strong> Derive this result by setting $S = 1 + r + ... + r^n$, multiplying this equation by $r$, and solving the two equations for $S$.

  \[
    \begin{align*}
      S &= 1 + r + ... + r^n \\
      rS &= r + r^2 + ... + r^{n + 1} \\
      \\
      S(1 - r) &= S - rS \\
      &= 1 - r^{n + 1} \\
      \rightarrow S &= \frac{1 - r^{n + 1}}{1 - r} \\
    \end{align*}
  \]

</p>
