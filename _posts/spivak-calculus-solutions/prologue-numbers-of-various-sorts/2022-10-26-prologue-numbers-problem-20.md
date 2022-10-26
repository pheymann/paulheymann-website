---
layout: math_post
title: Problem 20
date: 2022-10-26
categories: article spivak calculus prologue numbers
head_title: Problem 20
meta_description: Spivak Calculus Prologue Numbers Problem 20
meta_keywords: spivak,calculus,prologue,solutions
---

<p>

  Prove that the Fibonacci sequence can be represented as:
  $$a_n = \frac{(\frac{1 + \sqrt{5}}{2})^n - (\frac{1 - \sqrt{5}}{2})^n}{\sqrt{5}}$$
  Using complete induction(*) we start with $n = 1$:

  \[
    \begin{align*}
      a_1 &= \frac{\frac{1 + \sqrt{5}}{2} - \frac{1 - \sqrt{5}}{2}}{\sqrt{5}} \\
      &= \frac{\sqrt{5}}{\sqrt{5}} \\
      &= 1 \\
    \end{align*}
  \]

  Now, also have to show that $a_2 = 1$ (based on the Fibonacci definition):

  \[
    \begin{align*}
      a_2 &= \frac{(\frac{1 + \sqrt{5}}{2})^2 - (\frac{1 - \sqrt{5}}{2})^2}{\sqrt{5}} \\
      &= \frac{\frac{1 + 2\sqrt{5} + 5}{4} - \frac{1 - 2\sqrt{5} + 5}{4}}{\sqrt{5}} \\
      &= \frac{\sqrt{5}}{\sqrt{5}} \\
      &= 1 \\
    \end{align*}
  \]

  In the induction step, we assume that $1, ..., n - 1$ is correct and now show that this equation holds for $n$:

  \[
    \begin{align*}
      a_n &= a_{n - 1} + a_{n - 2} \\
      &= \frac{(\frac{1 + \sqrt{5}}{2})^{n - 1} - (\frac{1 - \sqrt{5}}{2})^{n - 1}}{\sqrt{5}} \\
      &+ \frac{(\frac{1 + \sqrt{5}}{2})^{n - 2} - (\frac{1 - \sqrt{5}}{2})^{n - 2}}{\sqrt{5}} \\
      &= \frac{(\frac{1 + \sqrt{5}}{2})^{n}(\frac{1 + \sqrt{5}}{2})^{-1} - (\frac{1 - \sqrt{5}}{2})^{n}(\frac{1 - \sqrt{5}}{2})^{-1}}{\sqrt{5}} \\
      &+ \frac{(\frac{1 + \sqrt{5}}{2})^{n}(\frac{1 + \sqrt{5}}{2})^{-2} - (\frac{1 - \sqrt{5}}{2})^{n}(\frac{1 - \sqrt{5}}{2})^{-2}}{\sqrt{5}} \\
      &= \frac{(\frac{1 + \sqrt{5}}{2})^{n}(\frac{1 + \sqrt{5}}{2})^{-1} + (\frac{1 + \sqrt{5}}{2})^{n}(\frac{1 + \sqrt{5}}{2})^{-2}}{\sqrt{5}} \\
      &- \frac{(\frac{1 - \sqrt{5}}{2})^{n}(\frac{1 - \sqrt{5}}{2})^{-1} + (\frac{1 - \sqrt{5}}{2})^{n}(\frac{1 - \sqrt{5}}{2})^{-2}}{\sqrt{5}} \\
      &= \frac{1}{\sqrt{5}} \cdot ((\frac{1 + \sqrt{5}}{2})^{n}(\frac{1 + \sqrt{5}}{2})^{-1} + (\frac{1 + \sqrt{5}}{2})^{n}(\frac{1 + \sqrt{5}}{2})^{-2} \\
      &- (\frac{1 - \sqrt{5}}{2})^{n}(\frac{1 - \sqrt{5}}{2})^{-1} + (\frac{1 - \sqrt{5}}{2})^{n}(\frac{1 - \sqrt{5}}{2})^{-2}) \\
      &= \frac{1}{\sqrt{5}} \cdot ((\frac{1 + \sqrt{5}}{2})^{n} \cdot ((\frac{1 + \sqrt{5}}{2})^{-1} + (\frac{1 + \sqrt{5}}{2})^{-2}) \\
      &- (\frac{1 - \sqrt{5}}{2})^{n} \cdot ((\frac{1 - \sqrt{5}}{2})^{-1} + (\frac{1 - \sqrt{5}}{2})^{-2})) \\
    \end{align*}
  \]

  For simplification we focus on 2 sub-sections of that equation now. First:

  \[
    \begin{align*}
      (\frac{1 + \sqrt{5}}{2})^{-1} + (\frac{1 + \sqrt{5}}{2})^{-2} &= \frac{2}{1 + \sqrt{5}} + \frac{4}{(1 + \sqrt{5})^2} \\
      &= \frac{2}{1 + \sqrt{5}} \cdot \frac{1 + \sqrt{5}}{1 + \sqrt{5}} + \frac{4}{(1 + \sqrt{5})^2} \\
      &= \frac{2 + 2\sqrt{5}}{(1 + \sqrt{5})^2} + \frac{4}{(1 + \sqrt{5})^2} \\
      &= \frac{2 + 2\sqrt{5} + 4}{(1 + \sqrt{5})^2} \\
      &= \frac{1 + 2\sqrt{5} + 5}{(1 + \sqrt{5})^2} \\
      &= \frac{(1 + \sqrt{5})^2}{(1 + \sqrt{5})^2} \\
      &= 1
    \end{align*}
  \]

  Second:

  \[
    \begin{align*}
      (\frac{1 - \sqrt{5}}{2})^{-1} + (\frac{1 - \sqrt{5}}{2})^{-2} &= \frac{2}{1 - \sqrt{5}} + \frac{4}{(1 - \sqrt{5})^2} \\
      &= \frac{2 - 2\sqrt{5}}{(1 - \sqrt{5})^2} + \frac{4}{(1 - \sqrt{5})^2} \\
      &= \frac{1 - 2\sqrt{5} + 5}{(1 - \sqrt{5})^2} \\
      &= \frac{(1 - \sqrt{5})^2}{(1 - \sqrt{5})^2} \\
      &= 1
    \end{align*}
  \]

  And with that the initial equation becomes:

  \[
    \begin{align*}
      a_n &= \frac{1}{\sqrt{5}} \cdot ((\frac{1 + \sqrt{5}}{2})^{n} \cdot 1 - (\frac{1 - \sqrt{5}}{2})^{n} \cdot 1) \\
      &= \frac{(\frac{1 + \sqrt{5}}{2})^n - (\frac{1 - \sqrt{5}}{2})^n}{\sqrt{5}}
    \end{align*}
  \]

</p>
