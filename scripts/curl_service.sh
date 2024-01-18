curl --location --request GET 'localhost:3000/getUserById' \
--header 'Content-Type: application/json' \
--data '{
    "userId":1
}'