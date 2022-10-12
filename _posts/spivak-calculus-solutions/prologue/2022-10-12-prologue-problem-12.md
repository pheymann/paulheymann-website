---
layout: math_post
title: Prologue - Problem 12
date: 2022-10-12
categories: article spivak calculus prologue
head_title: Prologue - Problem 12
meta_description: Spivak Calculus Prologue Problem 12
meta_keywords: spivak,calculus,prologue,solutions
---

<p>

  Prove the following:

</p>

<p>

  <strong>(i)</strong> $|xy| = |x| \cdot |y|$

  \[
    \begin{align*}
      (|xy|)^2 &= (xy)^2 \\
      &= x^2y^2 \\
      &= (|x|)^2 \cdot (|y|)^2 \\
      &= (|x| \cdot |y|)^2  && \text{| exponential can be eliminated} \\
    \end{align*}
  \]

</p>

<p>

  <strong>(ii)</strong> $|\frac{1}{x}| = \frac{1}{|x|}$, if $x \neq 0$.

  \[
    \begin{align*}
      |\frac{1}{x}| &= |1x^{-1}| \\
      \rightarrow (|1x^{-1}|) &= (|1| \cdot |x^{-1}|)  && \text{| based on (i)} \\
      &= |1| \cdot (\sqrt{(x^{-1})^2}) \\
      &= |1| \cdot (\sqrt{\frac{xx}{x}})  && \text{| no inherent order to exponentials here} \\
      &= |1| \cdot (\sqrt{(x^2)^{-1}}) \\
      &= |1| \cdot ((x^2)^{-1})^{\frac{1}{2}}) \\
      &= |1| \cdot (((x^2)^{\frac{1}{2}})^{-1}) \\
      &= |1| \cdot ((\sqrt{x^2})^{-1}) \\
      &= |1| \cdot \frac{1}{|x|} \\
    \end{align*}
  \]

</p>

<p>

  <strong>(iii)*</strong>  $\frac{|x|}{|y|} = |\frac{x}{y}|$, if $y \neq 0$

  \[
    \begin{align*}
      \frac{|x|}{|y|} &= |x| \cdot |y|^{-1} \\
      &= |xy^{-1}|  && \text{| based on (ii)} \\
      &= |\frac{x}{y}| \\
    \end{align*}
  \]

</p>

<p>

  <strong>(iv)</strong> $|x - y| \leq |x| + |y|$
  <br>
  <br>
  When $z = -y$ we get $|x + z| \leq |x| + |z|$ and that is already proven to be true.

</p>

<p>

  <strong>(v)</strong> $|x| - |y| \leq |x - y|$

  \[
    \begin{align*}
      (x - y)^2 &= x^2 - 2xy - y^2 \\
      &\geq x^2 - 2|x| \cdot |y| - y^2 \\
      &= |x|^2 - 2|x| \cdot |y| - |y|^2 \\
      &= (|x| - |y|)^2  && \text{| and then } a^2 \geq b^2 \rightarrow a \geq b \\
    \end{align*}
  \]

</p>

<p>

  <strong>(vi)</strong> $|(|x| - |y|)| \leq |x - y|$

  \[
    \begin{align*}
      |x - y| &= ||x - y|| \\
      &\geq |(|x| - |y|)| \\
    \end{align*}
  \]

</p>

<p>

  <strong>(vii)</strong> $|x + y + z| \leq |x| + |y| + |z|$. Indicate when equality holds, and prove your statement.

  \[
    \begin{align*}
      |x + y + z| &= |x + a| \\
      &\leq |x| + |a| \\
      \\
      |y + z| &\leq |y| + |z| \\
      \\
      |x + y + z| &\leq |x| + |y| + |z| \\
    \end{align*}
  \]

  For this expression to be equal $|y + z| = |y| + |z|$ and based on that $|x + a| = |x| + |a|$ which is only the case when $x, y, z \geq 0$. See:

  \[
    \begin{align*}
      (|a| + |b|)^2 &= |a|^2 + 2|a||b| + |b|^2 \\
      &= a^2 + 2|a||b| + b^2 \\
      &= a^2 + 2ab + b^2  && \text{| with } a, b \geq 0 \\
      &= (|a + b|)^2 \\
    \end{align*}
  \]

</p>
