# update bundle
cp build/bundle.js bundle.js
# update posters
rm -rf posters
cp -r build/posters posters/
#to gh pages
git commit -am 'update'
git push
