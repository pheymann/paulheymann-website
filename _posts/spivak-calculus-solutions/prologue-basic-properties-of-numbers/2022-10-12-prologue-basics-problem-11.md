---
layout: math_post
title: Problem 11
date: 2022-10-12
categories: article spivak calculus prologue basics
head_title: Problem 11
meta_description: Spivak Calculus Prologue Basics Problem 11
meta_keywords: spivak,calculus,prologue,solutions
---

<p>

  Find all numbers $x$ for which:

  <strong>(i)</strong>

  \[
    \begin{align*}
      |x - 3| &= 8 \\
      \sqrt{(x - 3)^2} &= 8 \\
      \pm(x - 3) &= 8 \\
      \\
      -x_1 + 3 &= 8 \\
      x_1 &= -5 \\
      \\
      x_2 - 3 &= 8 \\
      x_2 &= 11 \\
    \end{align*}
  \]

</p>

<p>

  <strong>(ii)</strong>

  \[
    \begin{align*}
      |x - 3| &< 8 \\
      \sqrt{(x - 3)^2} &< 8 \\
      \pm(x - 3) &< 8 \\
      \\
      -x_1 + 3 &< 8 \\
      x_1 &> -5 \\
      \\
      x_2 - 3 &< 8 \\
      x_2 &< 11 \\
    \end{align*}
  \]

</p>

<p>

  <strong>(iii)*</strong>

  \[
    \begin{align*}
      |x + 4| &< 2 \\
      \pm(x + 4) &< 2 \\
      \\
      -x_1 - 4 &< 2 \\
      x_1 &> -6 \\
      \\
      x_2 + 4 &< 2 \\
      x_2 &< -2 \\
    \end{align*}
  \]

</p>

<p>

  <strong>(iv)</strong>

  \[
    \begin{align*}
      |x - 1| + |x + 1| &< 2 \\
      \pm(x - 1) + \pm(x + 1) &< 2 \\
      \\
      -x_1 + 1 - x_1 + 2 &> 1 \\
      -2x_1 + 3 &> 1 \\
      2x_1 &< 2 \\
      x_1 &< 1 \\
      \\
      -x_2 + 1 + x_2 - 2 &> 1 \\
      -1 &> 1  && \text{| contradiction} \\
      \\
      x_3 - 1 - x_3 + 2 &> 1 \\
      1 &> 1  && \text{| contradiction} \\
      \\
      x_4 - 1 + x_4 - 2 &> 1 \\
      2x_4 - 3 &> 1 \\
      x_4 &> 2 \\
    \end{align*}
  \]

  $x_2$ and $x_3$ are an impossibility and therefore only $x_1$ and $x_4$ are left.

</p>

<p>

  <strong>(v)</strong>

  \[
    \begin{align*}
      |x - 1| + |x + 1| &< 2 \\
      \pm(x - 1) + \pm(x + 1) &< 2 \\
      \\
      -x_1 + 1 - x_1 - 1 &< 2 \\
      x_1 &> -1 \\
      \\
      x_2 - 1 + x_2 + 1 &< 2 \\
      x_2 &< 1 \\
    \end{align*}
  \]

</p>

<p>

  <strong>(vi)</strong>

  \[
    \begin{align*}
      |x - 1| + |x + 1| &< 1 \\
      \pm(x - 1) + \pm(x + 1) &< 1 \\
      \\
      -x_1 + 1 - x_1 - 1 &< 1 \\
      x_1 &> -\frac{1}{2} \\
      \\
      x_2 - 1 + x_2 + 1 &< 1 \\
      x_2 &< \frac{1}{2} \\
    \end{align*}
  \]

</p>

<p>

  <strong>(vii)</strong>

  \[
    \begin{align*}
      |x - 1| \cdot |x + 1| &= 0 \\
      \pm(x - 1) \cdot \pm(x + 1) &= 0 \\
      (x - 1)(x + 1) &= 0  && \text{| the signs can be removed by multiplication with } -1 \\
      \\
      x_1 - 1 &= 0 \\
      x_1 &= 1 \\
      \\
      x_2 &= -1 \\
    \end{align*}
  \]

</p>

<p>

  <strong>(viii)</strong>

  \[
    \begin{align*}
      |x - 1| \cdot |x + 2| &= 3 \\
      \pm(x - 1) \cdot \pm(x + 2) &= 3 \\
      \\
      (-1)(x_1 - 1) \cdot (-1)(x_1 + 2) &= 3 \\
      (x_1 - 1)(x_1 + 2) &= 3 \\
      x_1^2 + x_1 - 2 &= 3 \\
      x_1^2 + x_1 - 5 &= 0  && \text{| quadratic equation} \\
      x_{1_{12}} &= \frac{-1 \pm\sqrt{1 - 4 \cdot 1 \cdot -5}}{2} \\
      x_{1_1} &= \frac{-1 -\sqrt{21}}{2} \\
      x_{1_2} &= \frac{-1 +\sqrt{21}}{2} \\
      \\
      (-1)(x_2 - 1) \cdot (x_2 + 2) &= 3 \\
      (-x_2 + 1)(x_2 + 2) &= 3 \\
      -x_2^2 - x_2 + 2 &= 3 \\
      -x_2^2 - x_2 - 1 &= 0  && \text{| quadratic equation} \\
      x_{2_1} &= \frac{-1 + \sqrt{-3}}{2} \\
      x_{2_2} &= \frac{-1 - \sqrt{-3}}{2} \\
      \\
      (x_3 - 1) \cdot (-1)(x_3 + 2) &= 3 \\
      (x_3 - 1)(-x_3 - 2) &= 3 \\
      -x_3^2 - x_3 + 2 &= 3  && \text{| identical to } x_2 \\
      \\
      (x_4 - 1) \cdot (x_4 + 2) &= 3 \\
      x_4^2 + x_4 - 2 &= 3  && \text{| identical to } x_1 \\
    \end{align*}
  \]

</p>
