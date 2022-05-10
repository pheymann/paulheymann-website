# Blog
## Development setup
[Install RVM](https://rvm.io/). `gpg2` isn't required it seems.

```sh
curl -sSL https://get.rvm.io | bash -s stable
```

Install Ruby:

```sh
rvm install 2.6.8
```

Install [Jekyll](https://jekyllrb.com/docs/):

```sh
gem install bundler jekyll
```

## Build blog

First install all dependencies:

```sh
bundle install
```

Then build the blog:

```sh
bundle exec jekyll serve
```