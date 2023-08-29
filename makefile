pr:
	gh pr create --web -B master

push:
	git push origin staging

add:
	git add .

ppr:
	git push origin staging
	gh pr create --web -B master

