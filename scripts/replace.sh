for filename in `find . -type f -name 'book*'`; do mv -v "$filename" "${filename//book/[replacement]}"; done
find ./ -type f \( ! -iname "replace.sh" \) -exec sed -i -e 's/book/[replacement]/g' {} \;
find ./ -type f \( ! -iname "replace.sh" \) -exec sed -i -e 's/Book/[Replacement]/g' {} \;
find ./ -type f \( ! -iname "replace.sh" \) -exec sed -i -e 's/"bookshop"/"[project]"/g' {} \;
find ./ -type f \( ! -iname "replace.sh" \) -exec sed -i -e 's/POSTGRES_DB: bookshop/POSTGRES_DB: [project]/g' {} \;
