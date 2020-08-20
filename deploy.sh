if [ -d deploy ]; then
	rm -rf deploy
fi
mkdir deploy
mkdir -p deploy/conf
cp -r http_server deploy/http_server
cp -r redis deploy/redis
cp -r nginx deploy/nginx
cp -r grafana deploy/grafana
cp -r http_server/conf deploy/conf/http_server
cp docker-compose.yml deploy/docker-compose.yml
cd deploy
docker-compose stop
docker-compose build
docker-compose up