repo = muchrm/science-syllabus
commit = latest
docker:
	docker build -f Dockerfile -t $(repo):$(commit) .
run:
	docker run --link mongo:mongo --rm -it -d --name krud -v $(PWD):/go/src/github.com/muchrm/science-syllabus $(repo)