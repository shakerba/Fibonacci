counter=1
while [ $counter -le 1005 ]
do
#echo $counter
curl 127.0.0.1:8080/current &
curl 127.0.0.1:8080/next &
curl 127.0.0.1:8080/previous &
((counter++))
done
echo All done
