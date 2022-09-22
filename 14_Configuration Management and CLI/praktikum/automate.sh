if [ $# -eq 0 ]; then
  echo "Usage: $0 <file>"
  exit 1
fi

# shellcheck disable=SC1007
folderName="$1 at $(date)"

# shellcheck disable=SC2154
mkdir -p "$folderName"/about_me/personal
mkdir "$folderName"/about_me/professional
mkdir "$folderName"/about_me/my_friends
mkdir "$folderName"/about_me/my_system_info

facebookAccount="https://www.facebook.com/$2"
echo "$facebookAccount" > "$folderName"/about_me/personal/facebook.txt

linkedinAccount="https://www.linkedin.com/in/$2"
echo "$linkedinAccount" > "$folderName"/about_me/professional/linkedin.txt

curl -s https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt > "$folderName"/about_me/my_friends/list_of_my_friends.txt

echo "My username: $(whoami)" >> "$folderName"/about_me/my_system_info/about_this_laptop.txt "\n" "With host:" uname -a "$folderName"/about_me/my_system_info/about_this_laptop.txt


ping google.com -c 3 > "$folderName"/about_me/my_system_info/internet_connection.txt