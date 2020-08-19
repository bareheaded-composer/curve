if [ -d deploy ]; then
	rm -rf deploy
fi
mkdir deploy
mkdir -p deploy/conf
cp -r go deploy/go
cp -r go/conf deploy/conf/go
cp docker-compose.yml deploy/docker-compose.yml
cd deploy
docker-compose stop
docker-compose build
docker-compose up