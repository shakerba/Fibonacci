


echo "Scanning HTTP port 8080"
nc -w 60 -z localhost 8080
if [ $? != "0" ]; then
  echo "HTTP Server not running"
  exit 1
fi

echo "All servers running"
exit 0
