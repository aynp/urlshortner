curl --request POST \
  --url http://localhost:8080/v1/tiny/create \
  --header 'content-type: application/json' \
  --data '{
  "original_url": "http://www.google.com",
  "url_params": [
    {
      "source_param": "x",
      "target_param": "y",
      "is_mandatory": true
    }
  ],
  "auto_gen_params": [
    {
      "type": "uuid",
      "target_key": "uid"
    }
  ]
}'
