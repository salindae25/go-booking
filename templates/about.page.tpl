{{template "base" .}}
 {{define "content"}}

<h1>About Page</h1>
<div>This was in the template data: {{index .StringMap "Name"}}</div>
<p>ip :{{index .StringMap "remote_ip"}}</p>
<p>
  {{if eq (index .StringMap "remote_ip") ""}} visit <a href="/">Home </a> page
  and comeback {{else}}
  <span>extra content for people who visited home</span>

  {{end}}
</p>
{{end}}
