---
layout: math_post
title: Problem 23
date: 2022-10-27
categories: article spivak calculus prologue numbers
head_title: Problem 23
meta_description: Spivak Calculus Prologue Numbers Problem 23
meta_keywords: spivak,calculus,prologue,solutions
---

<p>

  The following is a recursive defintion of $a^n$:

  \[
    \begin{align*}
      a^1 &= a \\
      a^{n + 1} &= a^n \cdot a \\
    \end{align*}
  \]

  Prove, by induction, that:

  \[
    \begin{align*}
      a^{n + m} &= a^n \cdot a^m \\
      (a^n)^m &= a^{nm} \\
    \end{align*}
  \]

  We start by proving the first proposition for $n = 1$:

  \[
    \begin{align*}
      a^1 \cdot a^m &= a^m \cdot a \\
      &= a^{1 + m} \\
    \end{align*}
  \]

  Now, follows the induction step frpm $n = k$ to $n = k + 1$:

  \[
    \begin{align*}
      a^1 \cdot a^{k + m} &= a^1 \cdot a^x \\
      &= a^{1 + x} \\
      &= a^{(k + 1) + m} \\
    \end{align*}
  \]

  Next, we prove the second proposition and start again witg $m = 1$:

  \[
    \begin{align*}
      (a^n)^1 &= a^n \\
      &= a^{n \cdot 1} \\
    \end{align*}
  \]

  Now, follows the induction step frpm $m = k$ to $m = k + 1$:

  \[
    \begin{align*}
      (a^n)^{m + 1} &= (a^n)^m \cdot a^n \\
      &= [a^n \cdot ... \cdot a^n] \cdot a^n \\
      &= a^{[n + ... + n]} \cdot a^n \\
      &= a^{nm} \cdot a^n \\
      &= a^{nm + n} \\
      &= a^{n(m + 1)} \\
    \end{align*}
  \]

</p>
