
# update bundle
cp build/bundle.js bundle.js
# set current date
#currentdate=$(date +%d-%m-%Y)
#sed -i '' "s/LASTUPDATE/$currentdate/" public/bundle.js
# update posters
rm -rf posters
cp -r build/posters posters/
