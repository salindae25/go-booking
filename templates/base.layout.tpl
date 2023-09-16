{{define "base"}}
<!DOCTYPE html>
<html lang="en">
  <head>
    <title></title>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    {{block "css" .}} {{end}}
  </head>
  <body>
    <nav>
      <li><a href="/">Home</a></li>
      <li><a href="/about">About</a></li>
      
    </nav>
    {{block "content" .}} {{end}} {{block "javascript" .}} {{end}}
    
  </body>
</html>
{{end}}
