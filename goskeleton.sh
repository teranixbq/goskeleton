directory="$PWD"

echo -n "Folder Name: "
read folder_name

mkdir "$folder_name"

echo "Pilih pattern:"
echo "1. Default"

read -p "Pilihan Anda: " choice

case $choice in
    1)
        default() {
            git clone --depth=1 $(curl -s https://api.github.com/repos/teranixbq/goskeleton/contents/default | grep -o 'git://[^"]*' | head -n 1) "$folder_name"
        }
        default
        ;;
    *)
        echo "Pilihan tidak valid"
        ;;
esac
