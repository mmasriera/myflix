# myflix

movies i've watched &amp; watchlist
5min/day project
result at [page](http://mmasriera.github.com/myflix)

```bash
  # install
  npm i
  # build
  npm run build
```

===

This version creates a bundle with all the images (100 images --> 1.1MB bundle), and serves it.
It won't scale.

===

- front
 - webpack
    - [x] babel
    - [x] compile react
    - [x] less
    - [x] images in bundle
    - [ ] eslint
    - [ ] run tests
 - libs/frameworks
    - [x] react 15.0.1
    - [x] less
    - [ ] react router ?
    - [x] ~~redux~~
    - [x] ~~bulma.io~~
 - style
   - [ ] responsive
- scripts
 - list --> json
    - [x] go lang 'not very good for scripting'
    - [x] download images bc:[imdb poster api issue](http://stackoverflow.com/questions/28676608/403-error-for-loading-image-from-http-and-not-https/28676680#28676680)
    - [ ] "deploy"
- other
  - watchlist --> file + script + front
