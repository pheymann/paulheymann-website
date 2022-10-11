---
layout: math_post
title: Prologue - Problem 6
date: 2022-09-03
categories: article spivak calculus prologue
head_title: Prologue - Problem 6
meta_description: Spivak Calculus Prologue Problem 6
meta_keywords: spivak,calculus,prologue,solutions
---

<p>

  <strong>(a)</strong> Prove that if $0 \leq x < y$, then $x^n < y^n$, $n = 1, 2, 3, ...$.
  <br>
  <br>
  We can prove that by induction. $n = 1$:

  \[
    \begin{align*}
      x &< y  && \text{| which is obviously correct} \\
    \end{align*}
  \]

  $n = 2$:

  \[
    \begin{align*}
      x^2 &< y^2  && \text{| which is true based on 5.(ix)} \\
    \end{align*}
  \]

  $n = k$ is true, show that $n = k + 1$ is also true:

  \[
    \begin{align*}
      x^{k + 1} &< y^{k + 1} \\
      xx^{k} &< yy^{k}  && \text{| based on 5.(viii) with } a &= x, b &= y \\
      x^k &< y^k \\
    \end{align*}
  \]

</p>

<p>

  <strong>(b)</strong> Prove that if $x < y$ and $n$ is odd, then $x^n < y^n$.
  <br>
  <br>
  For $0 \leq x < y$ this is true based on (a). Now we have to show that it also holds for (A) $x < 0, y \geq 0$ and (B) $x, y < 0$.
  <br>
  <br>
  Starting with (A) using induction - $n = 1$:

  \[
    \begin{align*}
      x &< y \\
    \end{align*}
  \]

  $n = 3$:

  \[
    \begin{align*}
      x^3 &< y^3  && \text{| with } x = (-1)a, a = |x| \\
      (-1)(-1)(-1)a^3 &< y^3 \\
      (-1)|x|^3 &< y^3  && \text{| which is true because } y \geq 0 \rightarrow y^n \geq 0 \\
    \end{align*}
  \]

  $n = k$ is true with $k$ odd, show that $n = k + 2$ is also true:

  \[
    \begin{align*}
      x^{k + 2} &< y^{k + 2} \\
      (-1)^{k + 2}a^{k + 2} &< y^{k + 2}  && \text{| with } x = (-1)a, a = |x| \\
      (-1)a (-1)^{k + 1}a^{k + 1} &< yy^{k + 1}  && \text{| using 5.(viii) with } x < y \\
      (-1) (-1)^{k + 1}a^{k + 1} &< y^{k + 1} \\
      (-1)(-1) (-1)^{k}a^{k + 1} &< y^{k} \\
      (-1)(-1) x^k &< y^k \\
      x^k &< y^k \\
    \end{align*}
  \]

  (B) is identical to (a) when we just remove $-1$ on both sides and continue with $|x|$ and $|y|$.
</p>

<p>
  <strong>(c)</strong> Prove that if $x^n = y^n$ and $n$ is odd, then $x = y$.
  <br>
  <br>
  For $x, y > 0$ this statement holds true for any $n$. When $x, y < 0$ we can remove $-1$ again on both sides and deal with $|x|$ and $|y|$ which is true for any $n$.
  What is left is to show that (A) $x < 0, y > 0$ and (B) the inverse of (A) are only possible for odd $n$ when we remove the sign.

  Using induction - $n = 1$:

  \[
    \begin{align*}
      (-1)x &= y \\
      (-1)(-1)|x| &= y \\
      |x| &= y  && \text{| otherwise it is a contradiction} \\
    \end{align*}
  \]

  Using induction - $n = 3$:

  \[
    \begin{align*}
      (-1)x^3 &= y^3 \\
      (-1)(-1)(-1)(-1)|x|^3 &= y^3 \\
      |x|^3 &= y^3 \\
    \end{align*}
  \]

  $n = k$ is true with $k$ odd, show that $n = k + 2$ is also true:

  \[
    \begin{align*}
      (-1)x^{k + 2} &= y^{k + 2} \\
      (-1)(-1)(-1)(-1)^k|x|^k &= y^k \\
      (-1)(-1)^k|x|^k &= y^k \\
      (-1)x^k &= y^k \\
    \end{align*}
  \]

  This proof by indiction works exactly the same for (B).

</p>