---
layout: math_post
title: Problem 5
date: 2022-09-02
categories: article spivak calculus prologue basics
head_title: Problem 5
meta_description: Spivak Calculus Prologue Basics Problem 5
meta_keywords: spivak,calculus,prologue,solutions
---

Prove the following:

<p>
  <strong>(i)</strong> If $a < b$ and $c < d$, then $a + c < b + d$. With $b - a \in P$ and $d - c \in P$ we get:

  \[
    \begin{align*}
      b - a + d - c &> 0 \\
      b + d - a - c &> 0 \\
      b + d &> a + c
    \end{align*}
  \]
</p>

<p>
  <strong>(ii)</strong> If $a < b$, then $-b < -a$.

  \[
    \begin{align*}
      a &< b \\
      0 &< b - a \\
      -b &< -a
    \end{align*}
  \]
</p>

<p>
  <strong>(iii)</strong> If $a < b$ and $c > d$, then $a - c < b - d$.

  \[
    \begin{align*}
      b - a + c - d &> 0 \\
      b - d + c - a &> 0 \\
      b - d &> a - c
    \end{align*}
  \]
</p>

<p>
  <strong>(iv)</strong> If $a < b$ and $c > 0$, then $ac < bc$.

  \[
    \begin{align*}
      b - a &> 0 \\
      bc - ac &> 0 \\
      bc &> ac
    \end{align*}
  \]
</p>

<p>
  <strong>(v)</strong> If $a < b$ and $c < 0$, then $ac > bc$.

  \[
    \begin{align*}
      a(-c) &< b(-c)  && \text{| based on (iv)} \\
      bc &< ac \\
    \end{align*}
  \]
</p>

<p>

  <strong>(vi)</strong> If $a > 1$, then $a^2 > a$.

  \[
    \begin{align*}
      a^2 &> a \\
      a^2 - a &> 0 \\
      a - 1 &> 0 \\
      a &> 1 \\
    \end{align*}
  \]

</p>

<p>

  <strong>(vii)</strong> If $0 < a < 1$, then $a^2 < a$.

  \[
    \begin{align*}
      a^2 &< a \\
      a - a^2 &> 0 \\
      1 - a &> 0 \\
      a &< 1 \\
    \end{align*}
  \]

  When $a = 0$ then $a^2 = a$ which means we have to avoid that case: $0 < a$. What happens when $a < 0$?

  \[
    \begin{align*}
      a &< 0 \\
      a^2 &> 0 \\
      a &< a^2  && \text{| for any } a \\
    \end{align*}
  \]

</p>

<p>

  <strong>(viii)</strong> If $0 \leq a < b$ and $0 \leq c < d$, then $ac < bd$.

  \[
    \begin{align*}
      b - a &> 0 \\
      bc - ac &> 0 \\
      bc &> ac \\
    \end{align*}
  \]

  Then:

  \[
    \begin{align*}
      d - c &> 0 \\
      bd - bc &> 0 \\
      bd &> bc \\
    \end{align*}
  \]

  Both can now be combined through $bc$:

  \[
    \begin{align*}
      bd &> ac \\
      ac &< bd \\
    \end{align*}
  \]

</p>

<p>

  <strong>(ix)</strong> If $0 \leq a < b$, then $a^2 < b^2$. (Use (viii))
  <br>
  <br>
  From (viii) we have $ac < bd$ if $a < b$ and $c < d$. If $c = a$ and $d = b$ then we get $aa < bb = a^2 < b^2$.

</p>

<p>
  <strong>(x)</strong> If $a, b \geq 0$ and $a^2 < b^2$, then $a < b$. (User (ix) backwards.)
  <br>
  <br>
  $a^2 < b^2 = aa < bb$ and based on (viii) $a = c$ and $b = d$ and with that $c < d$ or $a < b$.

</p>
