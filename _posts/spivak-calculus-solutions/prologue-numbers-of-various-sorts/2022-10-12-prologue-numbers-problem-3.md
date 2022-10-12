---
layout: math_post
title: Problem 3
date: 2022-10-12
categories: article spivak calculus prologue numbers
head_title: Problem 3
meta_description: Spivak Calculus Prologue Numbers Problem 3
meta_keywords: spivak,calculus,prologue,solutions
---

<p>

  If $0 \leq k \leq n$ and using the "binomial coefficient" $\binom{n}{k}$:

</p>

<p>

  <strong>(a)</strong> Prove that:
  $$\binom{n + 1}{k} = \binom{n}{k - 1} + \binom{n}{k}$$

  \[
    \begin{align*}
      \binom{n}{k - 1} + \binom{n}{k} &= \frac{n!}{(k - 1)!(n - k + 1)!} + \frac{n!}{k!(n - k)!} \\
      &= \frac{n!}{\frac{k!}{k}((n - k)!(n - k + 1))} + \frac{n!}{k!(n - k)!} \\
      &= \frac{n!k}{k!(n - k)!(n - k + 1)} + \frac{n!(n - k + 1)}{k!(n - k)!(n - k + 1)} \\
      &= \frac{n!k + n!n - n!k + n!}{k!(n - k)!(n - k + 1)} \\
      &= \frac{n!n + n!}{k!(n - k)!(n - k + 1)} \\
      &= \frac{n!(n + 1)}{k!(n - k)!(n - k + 1)} \\
      &= \frac{(n + 1)!}{k!(n - k)!(n - k + 1)} \\
      &= \frac{(n + 1)!}{k!(n + 1 - k)!} \\
      &= \binom{n + 1}{k} \\
    \end{align*}
  \]

</p>