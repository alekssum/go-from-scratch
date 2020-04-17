<!DOCTYPE html>
<html>
    <head>
        <title>{{.Title}}</title>
    </head>
    <body>
        <h1>{{.H1}}</h1>
        <ul>
           {{range $value := .List}} 
            <li>{{$value}}</li>
           {{end}}
        </ul>
    </body>
</html>