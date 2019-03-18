
antlr4 = java -jar /usr/local/lib/antlr-4.7.1-complete.jar

version_parser:
	cd parser && $(antlr4) -Dlanguage=Go -package parser Version.g4 && gofmt -w .