dependencies:
	pip freeze > requirements.txt

environment:
	python3 -m venv rss

venv:
	cd rss/bin && source ./activate

install:
	pip install -r requirements.txt 