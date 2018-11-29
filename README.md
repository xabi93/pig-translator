Pig translator is a service to translate from english to Latin pig

# Requierements
- Docker

# Usage
Just:
```bash
docker-compose up
```
And the service will start in port `8080`

Then make a `POST` request with `application/json` header to `http://localhost:8080` with the following structure:
```json
{
    "lang": "pig",
    "text": "{what you want to translate}"
}
```
test