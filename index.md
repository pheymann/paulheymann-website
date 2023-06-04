---
layout: home
---

Hi there,<br/>
this is the website of Paul Heymann. If you are interested in some of my writings have a look into [articles](/articles). My open source work, including this website, is available on <a href="https://github.com/pheymann">GitHub</a>.

You are looking for a Software Engineer and think we could be a match? Feel free to contact me on <a href="https://www.linkedin.com/in/paul-h-6b4a53144">LinkedIn</a>. My [cv](/cv.pdf) is also available on this site.

<div class="row">
  <div class="col-12 post">
    <h2>Latest article: {{ site.posts.first.title }}</h2>

    <a class="title" href="{{ site.posts.first.url | relative_url }}">Find it here</a>

    <div>
      {{ site.posts.first.content }}
    </div>
  </div>
</div>
