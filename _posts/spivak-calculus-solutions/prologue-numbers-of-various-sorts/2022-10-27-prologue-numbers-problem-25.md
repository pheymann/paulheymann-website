---
layout: math_post
title: Problem 27
date: 2022-10-27
categories: article spivak calculus prologue numbers
head_title: Problem 27
meta_description: Spivak Calculus Prologue Numbers Problem 27
meta_keywords: spivak,calculus,prologue,solutions
---

<p>

  <strong>(a)</strong> A set $A$ of real numbers is called inductive if
</p>

<ul>
  <li>(1) $1$ is in $A$</li>
  <li>(2) $k + 1$ is in $A$ whenever $k$ is in $A$</li>
</ul>

<p>
  Prove that:
  <strong>(a.1)</strong> $R$ is inductive.
  <br/>
  <br/>
  We know that $1 \in R$ (1) and furthermore $k + 1 \in R$ when $k \in R$ as long as we assume $R$ being well ordered. If it is and we pick some smallest element from $R$ as starting point $x$ we can apply the same proof we used for the natural numbers. But for that it is necessary to pick such a $x$ because $R$ is infinite in both directions.
  <br/>
  <br/>
  <strong>(a.2)</strong> The set of positive real numbers is inductive. This prove then is rather straight forward, because here $x = 1$ and proof (1) applies.
  <br/>
  <br/>
  <strong>(a.3)</strong> The set of positive real numbers unequal to $\frac{1}{2}$ is inductive. As shown before we start with $x = 1$ and based on the definition of induction here given we increment in $1$ steps. Following that approach, we can never reach $\frac{1}{2}$ from $x = 1$ since $x > \frac{1}{2}$. Therefore, $\frac{1}{2}$ is not part of that set.
  <br/>
  <br/>
  <strong>(a.4)</strong> The set of positive real numbers unequal to $5$ is not inductive. Again, using (2) we start at $x = 1$ but going 4 inductive steps we end up with $((((1 + 1) + 1) + 1) + 1) = 5$. That means $5 \in R^+\\{5}$ which is a contradiction.
  <br/>
  <br/>
  <strong>(a.5)</strong> If $A$ and $B$ are inductive then the set $C$ which is in both $A$ and $B$ is also inductive. If $C$ is in both sets that means $C$ defines at least some part of the intersection of $A$ and $B$. If $A$ and $B$ are inductive this intersection $I$ has to be inductive too. We can prove that by contradiction. Assume there is an element $y \in I$ with $y + 1 \notin I$ and $\exists k: k > y + 1$ and $k \in I$, in other words, there is a hole. But that would mean that while $y \in A$ and $y \in B$ one of the two sets has that hole otherwise it would be in $I$ and from that would follow that either $A$ or $B$ isn't inductive which is a contradiction.
  <br/>
  <br/>
  If $C = I$ we have our proof, but $C$ could also be $C \subset I$ and not necessarily inductive e.g. $2k \in C$ with $k \in I \land 2k \in I$. Here, $C$ is in both $A$ and $B$ but not inductive.

</p>

<p>

  <strong>(b)</strong> A real number $n$ will be called natural number if $n$ is in every inductive set.
  <br/>
  <br/>
  <strong>(b.1)</strong> Prove that $1$ is a natural number. By definition (1) every inductive set has to have $1$ included which already proves this point.
  <br/>
  <br/>
  <strong>(b.2)</strong> Prove that $k + 1$ is a natural number if $k$ is a natural numbers. If $k$ is in every inductive set of real numbers then it follows that $k + 1$ is part of each such set otherwise we would end up with a case like (a.4) which would be a contradiction. Another way to formulate it is based on (a.5) because we know that $k + 1$ needs to be part of all inductive sets since $k$ is part of the intersection and that is inductive too.
  <br/>
  <br/>
  It is also through (b.2) that we make sure that $x = 1$ is the smallest element of natural numbers. At the start $R$ is also one of the inductive steps but her $x$ could be anything even $x < 0$, but by intersecting it with $R^+$ $0$ and $R^-$ are removed and only $x = 1$ remains as smallest possible element.

</p>

