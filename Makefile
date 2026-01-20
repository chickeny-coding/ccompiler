ccompiler.exe : ccompiler.go clexer.go cparser.go creplacer.go canalyzer.go
	go build ccompiler.go clexer.go cparser.go creplacer.go canalyzer.go
