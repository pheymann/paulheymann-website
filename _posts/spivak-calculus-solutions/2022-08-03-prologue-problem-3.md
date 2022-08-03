---
layout: math_post
title: Prologue - Problem 3
date: 2022-08-02
categories: article spivak calculus
head_title: Prologue - Problem 3
meta_description: Spivak Calculus Prologue Problem 3
meta_keywords: spivak,calculus,prologue,solutions
---

We have to prove the following assertions.

<p>
  <strong>(i)</strong> $\frac{a}{b} = \frac{ac}{bc}$, if $b,c \neq 0$.

  \[
    \begin{align*}
      \frac{a}{b} &= ab^{-1} \\
      &= ab^{-1}cc^{-1} \\
      &= acb^{-1}c^{-1} \\
      &= ac(bc)^{-1} && \text{| see (iii)} \\
      &= \frac{ac}{bc}
    \end{align*}
  \]
</p>

<p>
  <strong>(ii)</strong> $\frac{a}{b} + \frac{c}{d} = \frac{ad + bc}{bd}$.

  \[
    \begin{align*}
      \frac{a}{b} + \frac{c}{d} &= ab^{-1} + cd^{-1} \\
      &= ab^{-1}dd^{-1} + cd^{-1}bb^{-1} \\
      &= adb^{-1}d^{-1} + cbd^{-1}b^{-1} \\
      &= ad(bd)^{-1} + bc(bd)^{-1} && \text{| see (iii)} \\
      &= (ad + bc)(bd)^{-1} \\
      &= \frac{ad + bc}{bd}
    \end{align*}
  \]
</p>

<p>
  <strong>(iii)</strong> $(ab)^{-1} = a^{-1}b^{-1}$, if $a,b \neq 0$.

  \[
    \begin{align*}
      aa^{-1}bb^{-1} &= 1 \\
      aba^{-1}b^{-1} &= 1 \\
      (ab)(a^{-1}b^{-1}) &= 1 \\
      (ab)(a^{-1}b^{-1}) &= (ab)(ab)^{-1} && \text{| inverse } (ab) \\
      a^{-1}b^{-1} &= (ab)^{-1}
    \end{align*}
  \]
</p>

<p>
  <strong>(iv)</strong> $\frac{a}{b} \cdot \frac{c}{d} = \frac{ac}{db}$, if $b,d \neq 0$.

  \[
    \begin{align*}
      \frac{a}{b} \cdot \frac{c}{d} &= ab^{-1}cd^{-1} \\
      &= ac(bd)^{-1} \\
      &= \frac{ac}{bd} = \frac{ac}{db}
    \end{align*}
  \]
</p>

<p>
  <strong>(v)</strong> $\frac{\frac{a}{b}}{\frac{c}{d}} = \frac{ad}{bc}$, if $b,c \neq 0$.

  \[
    \begin{align*}
      ab^{-1}(cd^{-1})^{-1} &= ab^{-1}c^{-1}d && \text{| } (a^{-1})^{-1} = a \\
      &= ad(bc)^{-1} \\
      &= \frac{ad}{bc}
    \end{align*}
  \]
</p>

<p>
  <strong>(vi)</strong> If $b,d \neq 0$, then $\frac{a}{b} = \frac{c}{d}$ if and only if $ad = bc$. Also determine when $\frac{a}{b} = \frac{b}{a}$.

  \[
    \begin{align*}
      \frac{a}{b} &= \frac{c}{d} \\
      ab^{-1} &= cd^{-1} \\
      ab^{-1}bd &= cd^{-1}bd \\
      ad &= cb \\
      ad &= bc
    \end{align*}
  \]

  And:

\[
    \begin{align*}
      ad &= bc \\
      aa &= bb \\
      a &= b
    \end{align*}
  \]
</p>