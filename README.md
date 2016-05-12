# myflix

movies i've watched &amp; watchlist

5min/day project

[page](http://mmasriera.github.com/myflix)

```bash
  # install
  npm i
  # download data
  npm run list2json
  # build
  npm run build
```
===
- front
 - webpack
    - [x] babel
    - [x] compile react
    - [x] less
    - [x] ~~images in bundle~~ lazy load
    - [ ] eslint
    - [ ] run tests
 - libs/frameworks
    - [x] react 15.0.1
    - [x] less
    - [x] react-lazy-load
    - [x] ~~redux~~
    - [x] ~~bulma.io~~
 - style
   - [ ] responsive
 - [ ] highlight movies
- scripts
 - list --> json
    - [x] go lang 'not very good for scripting'
    - [x] download images bc:[imdb poster api issue](http://stackoverflow.com/questions/28676608/403-error-for-loading-image-from-http-and-not-https/28676680#28676680)
    - [ ] "deploy"
    - versions:
      - [sequential](https://github.com/mmasriera/myflix/blob/list2json-seq/lists/list2json.go)
      - **current** [concurrent (waitgroup+mutex)](https://github.com/mmasriera/myflix/blob/conc-waitgroup/lists/list2json.go)
- other
  - [ ] watchlist --> file + script + front
