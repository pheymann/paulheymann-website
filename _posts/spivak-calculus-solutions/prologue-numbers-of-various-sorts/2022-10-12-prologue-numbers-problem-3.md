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

<p>

  <strong>(b)</strong> Prove that $\binom{n}{k}$ is always a natural number.
  <br>
  <br>
  Using induction we start with $n = 0$ and $k = 0$:
  $$\binom{0}{0} = 1 \in N$$
  In general, $k = 0$ for any $n$:
  $$\binom{n}{0} = 1 \in N$$
  And $k = n$:
  $$\binom{n}{n} = 1 \in N$$
  Next, $n = i$ for all $k$ are given now prove $n = i + 1$:

  \[
    \begin{align*}
      \binom{i + 1}{j} &= \binom{n}{k - 1} + \binom{n}{k}  && \text{| based on (i)} \\
    \end{align*}
  \]

  Both $\binom{n}{k} \in N$ and $\binom{n}{k - 1} \in N$ by definition and $N$ being closed under $+$.

</p>

<p>

  <strong>(c)</strong> Prove that the binomial coefficient is a natural number by showing that it is the number of $k$ sets one can
  select from $1, 2, ..., n$.
  <br>
  <br>
  Since $\binom{n}{k}$ is supposed to be the number of sets it has to be $N$ by definition, under the assumption that
  we exclue $\emptyset$. A proof that doesn't realy on the definition itself is to show that we can define an
  expression that defines the number of sets $k$ and how that connects to the coefficient.
  <br>
  <br>
  If we select the sets of size $k$ we select one of the $n$ elements. Next we select one from $n - 1$:

  \[
    \begin{align*}
      n(n - 1)(n - 2) ... (n - k + 1) &= \frac{n(n - 1) ... 1}{(n - k)(n - k - 1)...1} \\
      &= \frac{n!}{(n - k)!} \\
    \end{align*}
  \]

  Now, we have sets in there which have the same elements but in different order. To exclude those we have to divide by the number of possible combination of sets of $k$ elements. $k(k - 1)(k - 2) ... 1 = k!$

  \[
    \begin{align*}
      \frac{n!}{k!(n - k)!} = \binom{n}{k} \\
    \end{align*}
  \]

</p>

<p>

  <strong>(d)</strong> Prove the binomial theore $(a + b)^n = \sum_{i = 0}^n \binom{n}{i}a^(n - i)b^i$ for any $a, b$ and $n \in N$.
  <br>
  <br>
  Using induction we start with $n = 1$:

  \[
    \begin{align*}
      \sum_{j = 0}^1 \binom{n}{j}a^{n - j}b^j &= \binom{1}{0}(a^1b^0) + \binom{1}{1}(a^0b^1) \\
      &= a + b \\
      &= (a + b)^1 \\
    \end{align*}
  \]

  And follow with the induction step from $n = k$ to $n = k + 1$:

  \[
    \begin{align*}
      (a + b)^{k + 1} &= (a + b)^k(a + b) \\
      &= (\sum_{j = 0}^{k} \binom{k}{j}a^{n - j}b^j)(a + b) \\
      &= \binom{k}{0}a^ka + \binom{k}{1}a^{k - 1}ab ... \binom{k}{k - 1}aab^{k - 1} \binom{k}{k}b^ka \\
      &+ \binom{k}{0}a^kb + \binom{k}{1}a^{k - 1}bb ... \binom{k}{k - 1}ab^{k - 1}b \binom{k}{k}b^kb \\
      &= \binom{k}{0}a^{k + 1} + \binom{k}{1}a^{k}b ... \binom{k}{k - 1}a^2b^{k - 1} \binom{k}{k}ab \\
      &+ \binom{k}{0}a^kb + \binom{k}{1}a^{k - 1}b^2 ... \binom{k}{k - 1}ab^{k} \binom{k}{k}b^{k + 1} \\
      &= a^{k + 1} + (\binom{k}{0} + \binom{k}{1})a^{k}b + ... + (\binom{k}{1} + \binom{k}{0})ab^k + b^{k + 1} \\
      &= a^{k + 1} + \binom{k + 1}{1}a^{k}b + ... + \binom{k + 1}{1}ab^k + b^{k + 1} \\
      &= \sum_{j = 0}^{k + 1} \binom{k}{j}a^{n - j}b^j \\
    \end{align*}
  \]

</p>
