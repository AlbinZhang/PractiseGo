//go run issuesreport.go repo:golang/go is:open json decoder
//go run issuesreport.go repo:golang/go commenter:gopherbot json encoder
package main

import (
	"html/template"
	"log"
	"os"
	"time"

	"./github"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}---------------------------------
Number:	{{.Number}}
User: 	{{.User.Login}}
Title:	{{.Title | printf "%.64s"}}
Age:	{{.CreateAt | daysAgo}} days
{{end}} `

/*
var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align:left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
    <td>{{.State}}</td>
	<td><ahref='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><ahref='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))
*/

var report = template.Must(template.New("report").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	/*
		if err := report.Execute(os.Stdout, result); err != nil {
			log.Fatal(err)
		}
	*/
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

}

func daysAgo(t time.Time) int {
	//fmt.Println(t)
	return int(time.Since(t).Hours() / 24)
}

var issueList = template.Must(template.New("issuelist").Parse(`
 <h1>{{.TotalCount}} issues</h1>
 <table>
 <tr style='text-align: left'>
   <th>#</th>
   <th>State</th>
   <th>User</th>
   <th>Title</th>
  </tr>
 {{range .Items}}
 <tr>
   <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
   <td>{{.State}}</td>
   <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
   <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
 </tr>
 {{end}}
 </table>
 `))
