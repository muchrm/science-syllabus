repo = muchrm/science-syllabus
commit = latest
docker:
	docker build -f Dockerfile -t $(repo):$(commit) .