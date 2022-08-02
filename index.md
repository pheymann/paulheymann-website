---
layout: home
---

Hi there,<br/>
if you are interested in some of my writings have a look into [articles](/articles). My open source work, including this website, is available on <a href="https://github.com/pheymann">GitHub</a>.

You are recruiter? Feel free to contact me on <a href="https://www.linkedin.com/in/paul-h-6b4a53144">LinkedIn</a>. My [cv](/cv.pdf) is also available on this site.

{% assign thought_posts = site.posts | where_exp:"post","post.categories contains 'thought'" %}
<div class="row">
  <div class="col-12 post">
    <h2>Latest article: {{ thought_posts.first.title }}</h2>

    <a class="title" href="{{ thought_posts.first.url | relative_url }}">Find it here</a>

    <div>
      {{ thought_posts.first.content }}
    </div>
  </div>
</div>