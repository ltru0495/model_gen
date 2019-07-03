package consts

const HTMLConst = `{{ define "%s"}}
{{template "header"}}
</head>
<body>
	<div class="container">
		<div class="row">
			<div class="col-md-12">%s
			</div>
		</div>
	</div>
{{template "footer"}}
</body>
{{ end }}`

const Form = `
			<form class="form" action="/%s" method="post">%s
			  	<button type="submit" class="btn btn-primary">Submit</button>
			</form>
`
const FormGroup = `
			<div class="form-group">
			    <label for="input%s">%[1]s</label>
			    <input type="text" class="form-control" id="input%[1]s" placeholder="Enter %[1]s">
			</div>
`

const TableConst = `
				<table class="table">
					<thead>
						<tr>
%s
						</tr>
					</thead>

					<tbody>
						{{ range .%ss }}
							<tr>
%s
							</tr>
						{{end}}
					</tbody>
				</table>`
