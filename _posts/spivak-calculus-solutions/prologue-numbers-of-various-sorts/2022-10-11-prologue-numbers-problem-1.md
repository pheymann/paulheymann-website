---
layout: math_post
title: Problem 1
date: 2022-10-11
categories: article spivak calculus prologue numbers
head_title: Problem 1
meta_description: Spivak Calculus Prologue Numbers Problem 
meta_keywords: spivak,calculus,prologue,solutions
---


<p>

  Prove the following formulas by induction.
  <br/>
  <br/>
  <strong>(i)</strong> $1^2 + ... + n^2 = \frac{n(n + 1)(2n + 1)}{6}$
  <br/>
  <br/>
  We start with $n = 1$:

  \[
    \begin{align*}
      \frac{1(1 + 1)(2 + 1)}{6} &= \frac{2 \cdot 3}{6} \\
      &= 1 \\
    \end{align*}
  \]

  Next, we do the induction step from a given $n = k$ to $n = k + 1$:

  \[
    \begin{align*}
      \frac{k(k + 1)(2k + 1)}{6} + (k + 1)^2 &= \frac{k(k + 1)(2k + 1) + 6(k + 1)^2}{6} \\
      &= \frac{k(k + 1)(2k + 1) + 6k^2 + 12k + 6}{6} \\
      &= \frac{k(2k^2 + k + 2k + 1) + 6k^2 + 12k + 6}{6} \\
      &= \frac{2k^3 + 3k^2 + k + 6k^2 + 12k + 6}{6} \\
      &= \frac{(2k^3 + 7k^2 + 6k) + (2k^2 + 7k + 6)}{6} \\
      &= \frac{k(2k^2 + 7k + 6) + (2k^2 + 7k + 6)}{6} \\
      &= \frac{k(k + 2)(2k + 3) + (k + 2)(2k + 3)}{6} \\
      &= \frac{k((k + 1) + 1)(2(k + 1) + 1) + ((k + 1) + 1)(2(k + 1) + 1)}{6} \\
      &= \frac{(k + 1)(k + 1) + 1)(2(k + 1) + 1)}{6} \\
    \end{align*}
  \]

</p>

<p>

  <strong>(ii)</strong> $1^3 + ... + n^3 = (1 + ... + n)^2$
  <br/>
  <br/>
  We again start with $n = 1$:

  \[
    \begin{align*}
      1^2 &= 1 &= 1^3 \\
    \end{align*}
  \]

  Next, we do the induction step from a given $n = k$ to $n = k + 1$:

  \[
    \begin{align*}
      (1 + ... + (k + 1))^2 &= 1^2 + ... + k^2 + (k + 1)^2 + 2 \cdot (1 \cdot 2) + ... + 2 \cdot (1 \cdot (k + 1)) + ... + 2 \cdot (k \cdot (k + 1)) \\
      &= 1^2 + ... + k^2 + 2 \cdot (1 \cdot 2) + ... + 2 \cdot (1 \cdot k) + ... + 2 \cdot ((k - 1) \cdot k) + (k + 1)^2 + 2 \cdot (1 \cdot (k + 1)) + ... + 2 \cdot (k \cdot (k + 1)) \\
      &= (1 + ... + k)^2 + (k + 1)^2 + 2 \cdot (1 + ... + k)(k + 1) \\
      &= (1 + ... + k)^2 + (k + 1)^2 + 2 \frac{k(k + 1)}{2}(k + 1) \\
      &= (1 + ... + k)^2 + k^2 + 2k + 1 + 2 \frac{(k^2 + k)(k + 1)}{2} \\
      &= (1 + ... + k)^2 + k^2 + 2k + 1 + k^3 + k^2 + k^2 + k \\
      &= (1 + ... + k)^2 + k^3 + 3k^2 + 3k + 1 \\
      &= (1 + ... + k)^2 + (k + 1)^3 \\
    \end{align*}
  \]

</p>
