### Route Register

Route registration library

Payload:
```
{
    "service": string //service name
    "routes": {
        "name": string //route name
        "path": string //route
        "method": string //Route verb
        "isPublic": boolean //if is public or not
    }[]
}
```