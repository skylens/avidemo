package main

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	// 创建模板
	tmpl := template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html lang="en">
	<head>
	  <title>Avi</title>
	  <link rel="icon" href="data:image/svg+xml;base64,AAABAAEAICAAAAEAIACoEAAAFgAAACgAAAAgAAAAQAAAAAEAIAAAAAAAgBAAABMLAAATCwAAAAAAAAAAAAAAS/8GAEv/kQBL/+4AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL/+4AS/+RAEv/BgBL/5AAS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS/+JAEv/7QBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL/+0AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//1uL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wZP//9WiP//AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//9bi///8vb///r8//8GT///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8NVP///f3///H1//9WiP//AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//0B4///v9P//+Pr//4+w//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//5y5///7/P//7/T//0B4//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//97o////////7/T//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///+fv////////e6P//AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///g6f///L2///x9f//gaf//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//4aq///09///8vb//4On//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///9fj////////W4v//AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///4uv////////1+P//AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//+Rsf//+Pr///D0//8xbv//AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//0B4///v9P//+Pr//5Gx//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wxT///6+////v///7HI//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///wNP////////6+///DFP//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//6nD///9/v//9/r//xNY//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//x9h///09////f7//6nD//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///Kmn///H1///4+v//kbH//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///nLn///v8///x9f//Kmn//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///x9f////////1+P//AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL///9/v///////8jY//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//9AeP//7/T///L2//+Dp///AEv//wBL//8AS///AEv//wBL//8AS///hqr///T3///v9P//SH7//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL///e6P///////9fi//8AS///AEv//wBL//8AS///AEv//wBL///k7P///////+Ts//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//4On///y9v//7/T//zlz//8AS///AEv//wBL//8AS///SH7//+/0///09///hqr//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv///X4////////uM3//wBL//8AS///AEv//wBL///B0/////////n7//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///kbH///j6///3+f//FFn//wBL//8AS///ImP///P2///6+///l7X//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8MU///+vv///r7//+Xtf//AEv//wBL//+ivf///P3///f6//8TWP//AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//+pw////f7///n7//8AS///AEv///3+///+////scj//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//ypp///x9f//8/b//4On//+Jq///9fj///D0//8xbv//AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//8jY////////3uj//+nv////////0N7//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///SH7//+/0///v9P//7/T///D0//9Pg///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///5Oz/////////////6vD//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//+Gqv//7/T//+/0//+Jq///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//9Ifv//SH7//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv/7QBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL/+0AS/+QAEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv/iABL/wYAS/+RAEv/7QBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv//wBL//8AS///AEv/7gBL/5EAS/8GAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=">
	  <style>
		body {
		  background-color: antiquewhite;
		  text-align: center;
		  padding: 50px;
		  font-family: "Open Sans","Helvetica Neue",Helvetica,Arial,sans-serif;
		}
		#logo {
		  margin-bottom: 40px;
		}
	  </style>
	</head>
	<body>
	  <img id="logo" src="data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHJvbGU9ImltZyIgdmlld0JveD0iLTEuNDQgLTEuNDQgNzUuMDEgNDMuNzYiPjx0aXRsZT5BdmkgTmV0d29ya3MgbG9nbzwvdGl0bGU+PGRlZnM+PGNsaXBQYXRoIGlkPSJhIiB0cmFuc2Zvcm09InRyYW5zbGF0ZSgtMzYgLTcxMC42NCkiPjxwYXRoIGZpbGw9Im5vbmUiIGQ9Ik0zNiA3MTAuNjRoNzIuMTcxdjQwLjg5NkgzNnoiLz48L2NsaXBQYXRoPjwvZGVmcz48ZyBjbGlwLXBhdGg9InVybCgjYSkiPjxwYXRoIGZpbGw9IiNmMDRlMjMiIGQ9Ik00MC4zMDcgMjUuMTJhMS4xNyAxLjE3IDAgMCAwIDEuMTMxLjg0aC4wNzJhMS4xNzMgMS4xNzMgMCAwIDAgMS4xMzMtLjg0TDUzLjA4NSAxLjI0MWEuNjUzLjY1MyAwIDAgMCAuMDczLS4zNjQuOTIyLjkyMiAwIDAgMC0uOTE0LS44NzYgMS4xIDEuMSAwIDAgMC0uOTQ5LjcyOUw0MS41MSAyMy42NiAzMS43NjMuNzY3YTEuMDkgMS4wOSAwIDAgMC0uOTg2LS43NjYuOTYuOTYgMCAwIDAtLjk1LjkxMi43MjUuNzI1IDAgMCAwIC4wNzMuNHpNMTIuODUxLjg0QTEuMTY4IDEuMTY4IDAgMCAwIDExLjcyIDBoLS4wNzJhMS4xNzIgMS4xNzIgMCAwIDAtMS4xMzMuODRMLjA3MyAyNC43MTlhLjY1My42NTMgMCAwIDAtLjA3My4zNjQuOTIyLjkyMiAwIDAgMCAuOTE0Ljg3NiAxLjEgMS4xIDAgMCAwIC45NDktLjcyOUwxMS42NDggMi4zIDIxLjQgMjUuMTkzYTEuMDkgMS4wOSAwIDAgMCAuOTg2Ljc2Ni45Ni45NiAwIDAgMCAuOTUtLjkxMi43MjMuNzIzIDAgMCAwLS4wNzMtLjR6TTY2Ljk2IDBhLjk2Ny45NjcgMCAwIDAtLjk1Ljk1VjI0LjlhLjk1Ljk1IDAgMCAwIDEuOSAwVi45NWEuOTY3Ljk2NyAwIDAgMC0uOTUtLjk1Ii8+PHBhdGggZmlsbD0iIzJhMmEyZCIgZD0iTS4xNTQgMzQuNWEuMjQ2LjI0NiAwIDAgMSAuMjQyLS4yNDFILjQ3YS4zLjMgMCAwIDEgLjI0MS4xMzlsNC4zODEgNS41NnYtNS40NzZhLjIzMi4yMzIgMCAxIDEgLjQ2NCAwdjYuMTI2YS4yLjIgMCAwIDEtLjE5NS4yMDVoLS4wMzZhLjMyOS4zMjkgMCAwIDEtLjI0MS0uMTU4TC42MTggMzQuOTkydjUuNmEuMjMyLjIzMiAwIDEgMS0uNDY0IDB6TTEzLjEgNDAuODk1YTIuMzUzIDIuMzUzIDAgMCAxLTIuMzIyLTIuNDY5di0uMDE4YTIuMzM0IDIuMzM0IDAgMCAxIDIuMjQ3LTIuNDY5IDIuMjQyIDIuMjQyIDAgMCAxIDIuMTcyIDIuNDUuMjMxLjIzMSAwIDAgMS0uMjIzLjIxNGgtMy43MTFhMS44NzcgMS44NzcgMCAwIDAgMS44NTYgMS44NjUgMi4xIDIuMSAwIDAgMCAxLjUzMi0uNjQ5LjIuMiAwIDAgMSAuMTQ4LS4wNjUuMjE3LjIxNyAwIDAgMSAuMjIzLjIxMy4yMjcuMjI3IDAgMCAxLS4wNzQuMTU4IDIuNDA4IDIuNDA4IDAgMCAxLTEuODQ3Ljc3bTEuNjA1LTIuN2ExLjc1NCAxLjc1NCAwIDAgMC0xLjctMS44MzcgMS44NDMgMS44NDMgMCAwIDAtMS43NDQgMS44Mzd6bTUuNzg4IDEuMzczdi0zLjFoLS41MTFhLjIyMy4yMjMgMCAwIDEtLjIxMy0uMjEzLjIxNC4yMTQgMCAwIDEgLjIxMy0uMmguNTExdi0xLjMxOGEuMjIyLjIyMiAwIDAgMSAuMjIzLS4yMzIuMjMuMjMgMCAwIDEgLjIzMi4yMzJ2MS4zMThoMS40ODVhLjIyMi4yMjIgMCAwIDEgLjIxMy4yMTQuMjE0LjIxNCAwIDAgMS0uMjEzLjJoLTEuNDg1djMuMDU0YS44My44MyAwIDAgMCAuOTQ3LjkxOCAyLjAwNSAyLjAwNSAwIDAgMCAuNTM4LS4wOTIuMjEyLjIxMiAwIDAgMSAuMi4yLjIwNi4yMDYgMCAwIDEtLjE0OC4xOTUgMS45NjEgMS45NjEgMCAwIDEtLjY2OS4xMiAxLjIxIDEuMjEgMCAwIDEtMS4zMjctMS4zbTguNDQ1IDEuMDgxbC0xLjU1LTQuMjYxYS40MjQuNDI0IDAgMCAxLS4wMzctLjE0OC4yMzguMjM4IDAgMCAxIC4yNDItLjIyNC4yNDguMjQ4IDAgMCAxIC4yNDEuMmwxLjM3MyA0IDEuMzc0LTQuMDE4YS4yMjkuMjI5IDAgMCAxIC4yMjMtLjE3OGguMDE4YS4yMzQuMjM0IDAgMCAxIC4yMzMuMTc4bDEuMzczIDQuMDE4IDEuMzgzLTQuMDE4YS4yMzEuMjMxIDAgMCAxIC40NTUuMDM3LjQyNS40MjUgMCAwIDEtLjAzNy4xNTdsLTEuNTUgNC4yNjFhLjI3My4yNzMgMCAwIDEtLjI1MS4yaC0uMDE4YS4yNjMuMjYzIDAgMCAxLS4yNTEtLjIxNEwzMC44MSAzNi43NmwtMS4zNTUgMy44NzFhLjI2My4yNjMgMCAwIDEtLjI1MS4yMTRoLS4wMThhLjI3NC4yNzQgMCAwIDEtLjI1MS0uMm05Ljg5Mi0yLjIwOXYtLjAxOWEyLjQ0NiAyLjQ0NiAwIDAgMSAyLjQ0MS0yLjQ3OCAyLjQyIDIuNDIgMCAwIDEgMi40MjMgMi40NTl2LjAxOWEyLjQ0NyAyLjQ0NyAwIDAgMS0yLjQ0MSAyLjQ3OSAyLjQyMSAyLjQyMSAwIDAgMS0yLjQyMy0yLjQ2bTQuMzcyIDB2LS4wMTlhMS45ODUgMS45ODUgMCAwIDAtMS45NDktMi4wNTIgMS45NTUgMS45NTUgMCAwIDAtMS45MzIgMi4wMzN2LjAxOWExLjk4NSAxLjk4NSAwIDAgMCAxLjk1IDIuMDUxIDEuOTUzIDEuOTUzIDAgMCAwIDEuOTMxLTIuMDMybTUuNjEyLTIuMTkxYS4yMjguMjI4IDAgMSAxIC40NTUgMHYxLjE1MWEyLjI1MiAyLjI1MiAwIDAgMSAxLjktMS40Mi4yNDIuMjQyIDAgMCAxIC4yNDEuMjUxLjI0Ny4yNDcgMCAwIDEtLjI0MS4yNSAyLjEyOSAyLjEyOSAwIDAgMC0xLjkgMi4zNjd2MS43NDZhLjIyMS4yMjEgMCAwIDEtLjIyMy4yMzEuMjI0LjIyNCAwIDAgMS0uMjMyLS4yMzF6bTcuNDMtMi4wNDJhLjIyMi4yMjIgMCAwIDEgLjIyMy0uMjMyLjIzLjIzIDAgMCAxIC4yMzIuMjMydjQuODgzbDIuOS0zLjAwOGEuMi4yIDAgMCAxIC4xNTgtLjA2NS4yMS4yMSAwIDAgMSAuMjEzLjIxNC4xOS4xOSAwIDAgMS0uMDc0LjE1N2wtMS43MzYgMS43NjQgMS44NTcgMi4yOTNhLjIzNy4yMzcgMCAwIDEgLjA2NS4xNjcuMjA3LjIwNyAwIDAgMS0uMjIyLjIxNC4yMzIuMjMyIDAgMCAxLS4yLS4xbC0xLjgxOS0yLjI0Ny0xLjEzOCAxLjE1OXYuOTU2YS4yMTcuMjE3IDAgMCAxLS4yMjMuMjMyLjIyNC4yMjQgMCAwIDEtLjIzMi0uMjMyem04LjYxIDYuMDU3YS4yNDYuMjQ2IDAgMCAxLS4wODQtLjE3Ni4yMjcuMjI3IDAgMCAxIC4yMjMtLjIyMy4yNTguMjU4IDAgMCAxIC4xNDkuMDQ2IDIuNjUzIDIuNjUzIDAgMCAwIDEuNTc4LjUzOWMuNjQgMCAxLjE0Mi0uMzUzIDEuMTQyLS45di0uMDE4YzAtLjU1Ny0uNTk0LS43NjEtMS4yNTQtLjk0Ny0uNzctLjIyMy0xLjYyNC0uNDU1LTEuNjI0LTEuM3YtLjAyYTEuMzkxIDEuMzkxIDAgMCAxIDEuNTU5LTEuMzA4IDMuMjMyIDMuMjMyIDAgMCAxIDEuNTMyLjQyNi4yNTEuMjUxIDAgMCAxIC4xMjEuMi4yMjYuMjI2IDAgMCAxLS4yMjMuMjIyLjI0LjI0IDAgMCAxLS4xMzEtLjAzNiAyLjYxOSAyLjYxOSAwIDAgMC0xLjMxNy0uMzkxYy0uNjUgMC0xLjA3Ny4zNTQtMS4wNzcuODI3di4wMThjMCAuNTMuNjQuNzI1IDEuMzE4LjkxOS43NjEuMjE0IDEuNTU5LjQ5MyAxLjU1OSAxLjMyOHYuMDE5YTEuNDY4IDEuNDY4IDAgMCAxLTEuNjMzIDEuMzgzIDMuMjU0IDMuMjU0IDAgMCAxLTEuODM4LS42MTNNNjkuOCAzNS41NmExLjE4NSAxLjE4NSAwIDEgMSAxLjE4NSAxLjE4NEExLjE4MiAxLjE4MiAwIDAgMSA2OS44IDM1LjU2bTIuMDU3IDBhLjg3My44NzMgMCAxIDAtLjg3Mi45Ljg1Ni44NTYgMCAwIDAgLjg3Mi0uOW0tLjI4NC42NDNoLS4zMWwtLjI4NC0uNTRoLS4yMTN2LjU0SDcwLjV2LTEuMjc4aC42MTljLjMzNCAwIC41LjA5LjUuMzkxIDAgLjIzNy0uMTI1LjMzMS0uMzUuMzQ3em0tLjQ1NC0uNzMxYy4xNDQgMCAuMjQzLS4wMzEuMjQzLS4xOTFzLS4xODQtLjE1OS0uMzA1LS4xNTloLS4yOTF2LjM1eiIvPjwvZz48L3N2Zz4=" alt="Avi logo" width=300 />
	  <h1>Hello world!</h1>
	  <h3>My hostname is <font color="orange">{{ .Hostname }}</font>, and the current version is <font color="orange">v1</font></h3>
	  {{- range .IPv4s }}
	  <h3><font color="black">IPv4 is </font><font color="orange">{{ . }}</font></h3>
	  {{- end }}
	  {{- range .IPv6s }}
	  <h3><font color="black">IPv6 is </font><font color="orange">{{ . }}</font></h3>
	  {{- end }}
	  <h3><font color="black">I'm running on host </font><font color="orange"></font></h3>
	  <br />
	
	</body>
  </html>
`))

	// 创建 http.Handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 获取服务端信息
		hostname, _ := os.Hostname()
		ips := getIPs()

		// 区分 IPv4 和 IPv6 地址
		ipv4s := make([]string, 0)
		ipv6s := make([]string, 0)
		for _, ip := range ips {
			if ip == "127.0.0.1" || ip == "::1" {
				continue
			}
			if strings.Contains(ip, ":") {
				ipv6s = append(ipv6s, ip)
			} else {
				ipv4s = append(ipv4s, ip)
			}
		}

		// 获取客户端信息
		userAgent := r.UserAgent()
		remoteAddr := r.RemoteAddr

		// 渲染模板
		tmpl.Execute(w, struct {
			Hostname   string
			IPv4s      []string
			IPv6s      []string
			UserAgent  string
			RemoteAddr string
		}{
			Hostname:   hostname,
			IPv4s:      ipv4s,
			IPv6s:      ipv6s,
			UserAgent:  userAgent,
			RemoteAddr: remoteAddr,
		})

		// 记录日志
		logRequest(r)
	})

	mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "v1")
		// 记录日志
		logRequest(r)
	})

	// 创建 http.Server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// 监听和服务
	fmt.Println("Listening on port 8080...")
	server.ListenAndServe()
}

// 获取服务端 IP地址
func getIPs() []string {
	ifs, err := net.Interfaces()
	if err != nil {
		return nil
	}

	ips := make([]string, 0)
	for _, i := range ifs {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}

		for _, a := range addrs {
			ip := a.(*net.IPNet).IP.String()
			if ip != "127.0.0.1" {
				ips = append(ips, ip)
			}
		}
	}

	return ips
}

// 记录日志
func logRequest(r *http.Request) {
	// 这里可以使用更复杂的日志记录库，例如 zap
	fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.RequestURI)
}
