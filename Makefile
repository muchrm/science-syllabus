repo = muchrm/science-syllabus
commit = latest
docker:
	docker build -f Dockerfile -t $(repo):$(commit) .
run:
	docker run --link mongo:mongo --rm --name krud -v $(PWD)/sheet_2.xlsx:/app/sheet_2.xlsx -v $(PWD)/sheet.xlsx:/app/sheet.xlsx $(repo)