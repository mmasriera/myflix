# myflix

[page] (http://mmasriera.github.com/myflix)

movies i've watched

TODO : 
 - [x] search
 - [ ] watchlist
 - [ ] series
 - [ ] personal recommendations/ranking

list of movies ``` lists/titles.txt```

```bash
  # install deps
  npm i
  # download data
  npm run list2json
  # build
  npm run build
  # open
  open build/index.html
```

## branches/versions
- script **lists/list2json**: 
  - **list2json-seq** fist version of the script that retrieves the data, sequential :exclamation:sloooooow
  - **conc-waitgroup** up to 10 concurrent downloads
- images
  - **images-in-bundle** images in bundle.js via webpack, :exclamation:bundle.js size
  - **images-lazyload** images saved in a specific directory, and rendered as the page loads
- front
  - **sidebar-grid** ![v1](https://gyazo.com/3d286e1238e2a520353341cb4d9a6de6.png)
