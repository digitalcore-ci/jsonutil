# jsonutil

jsonutil is a small Go package that provides utility types for working with JSON data following the convention used at Digital Core CI.
It includes types for representing JSON errors and responses, as well as constant error codes for common error cases.

## Example of our JSON convention

JSON responses should be in format `{"name": "object"}` where name is the name of the object being returned and object is the actual object.

JSON response examples:

```
{
    "users": [
        {
            "id": 1,
            "name": "Jhon"
        },
        {
            "id": 2,
            "name": "Doe"
        }
    ]
}
```

```
{
    "media": {
        "id": "YmVob3ViYWtvdWFtZW1hbmFzc2U=",
        "contentType": "image/png"
    }
}
```

JSON errors should be in the format `{"errorCode": "code", "message": "message value"}`.

JSON errors examples:

```
{
    "errorCode": "INVALID_INPUT",
    "message": "invalid field input"
}
```

##
