---
layout: home
---

Hi there,<br/>
this is my personal website and blog. Here, I publish my [writings](/articles) or share some of my [photography](/photos). My open source work, including this website, is available on <a href="https://github.com/pheymann">GitHub</a>.

I am also maintaing <a href="https://wernigerode-in-zahlen.de">wernigerode-in-zahlen.de</a> a website hosting a visual description of my city's budget.

You want to get into contact? Feel free to contact me on <a href="https://www.linkedin.com/in/paul-heymann-6b4a53144/">LinkedIn</a>. My [cv](/cv.pdf) is also available on this site.

<div class="row">
  <div class="col-12 post">
    <h2>Latest article: {{ site.posts.first.title }}</h2>

    <a class="title" href="{{ site.posts.first.url | relative_url }}">Find it here</a>

    <div>
      {{ site.posts.first.content }}
    </div>
  </div>
</div>
