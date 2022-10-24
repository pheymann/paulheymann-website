---
layout: math_post
title: Problem 11
date: 2022-10-24
categories: article spivak calculus prologue numbers
head_title: Problem 11
meta_description: Spivak Calculus Prologue Numbers Problem 11
meta_keywords: spivak,calculus,prologue,solutions
---

<p>

  Prove the principle of complete induction from the ordinary principle of induction.
  <br>
  <br>
  Complete induction can be expressed as: Set $A$ that contains $1$ and $n + 1$ when it also contains $1, ..., n$. To show that this is true using ordinary induction, we create a new set $B$ with $1, ..., k \in B$ for every $k \in A$. First, we start with the base assumption $k \leq 1; k \in B$, which is simply $1$. Now, we show $n + 1$ which means we construct a set $B$ with $k \leq n + 1; k \in B$ under the assumption $n \in B$.
</p>

<ul>
  <li>If $k < n + 1$ then we only look at $k = n$ which is automatically true because of our induction assumption $n \in B$.</li>
  <li>If $k = n + 1$ we already know $k \leq n \in B$ and that $n + 1 \in A$ from which follows that $n + 1 \in B$.</li>
</ul>

<p>
  But that means we constructed a set $B$ that contains $1 \leq k \leq n$ with $k \in A$ for any $n \in N$.
</p>
