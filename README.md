# go-real-world-web-app

for filename in `find . -type f -name 'post*'`; do mv -v "$filename" "${filename//post/[replacement]}"; done
find ./ -type f -exec sed -i -e 's/post/[replacement]/g' {} \;
find ./ -type f -exec sed -i -e 's/Post/[Replacement]/g' {} \;
find ./ -type f -exec sed -i -e 's/Methed[Replacement]/MethodPost/g' {} \;
find ./ -type f -exec sed -i -e 's/"blog"/"[replacement]"/g' {} \;
find ./ -type f -exec sed -i -e 's/[replacement]gres/postgres/g' {} \;

