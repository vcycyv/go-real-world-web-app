for filename in `find . -type f -name 'post*'`; do mv -v "$filename" "${filename//post/[replacement]}"; done
find ./ -type f \( ! -iname "replace.sh" \) -exec sed -i -e 's/post/[replacement]/g' {} \;
find ./ -type f \( ! -iname "replace.sh" \) -exec sed -i -e 's/Post/[Replacement]/g' {} \;
find ./ -type f \( ! -iname "replace.sh" \) -exec sed -i -e 's/Method[Replacement]/MethodPost/g' {} \;
find ./ -type f \( ! -iname "replace.sh" \) -exec sed -i -e 's/[replacement]gres/postgres/g' {} \;
find ./ -type f \( ! -iname "replace.sh" \) -exec sed -i -e 's/"bookshop"/"[project]"/g' {} \;
find ./ -type f \( ! -iname "replace.sh" \) -exec sed -i -e 's/POSTGRES_DB: bookshop/POSTGRES_DB: [project]/g' {} \;
