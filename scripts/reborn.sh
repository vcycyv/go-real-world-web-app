sed -i 's/\[replacement\]/new-text/g' replace.sh
sed -i 's/\[Replacement\]/new-text/g' replace.sh
sed -i 's/\[project\]/new-text/g' replace.sh
chmod +x replace.sh
mv replace.sh ..
cd ..
./replace.sh
mv replace.sh scripts/
cd scripts