---
layout: math_post
title: Problem 2
date: 2022-08-02
categories: article spivak calculus prologue basics
head_title: Problem 2
meta_description: Spivak Calculus Prologue Basics Problem 2
meta_keywords: spivak,calculus,prologue,solutions
---

<p>
  We have to show what is wrong with the following proof. Let $x = y$. Then:

  \[
    \begin{align*}
      x^2 &= xy && \text{1} \\
      x^2 - y^2 &= xy - y^2 && \text{2} \\
      (x + y)(x - y) &= y(x - y) && \text{3} \\
      x + y &= y && \text{4} \\
      2y &= y && \text{5} \\
      2 &= 1 && \text{6}
    \end{align*}
  \]

In line (3) to (4) we are getting rid of $(x - y)$ by the multiplicative inverse, but that change is illegal, because it would mean a division by $0$.
</p>