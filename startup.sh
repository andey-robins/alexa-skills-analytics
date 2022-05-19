# tear down any old containers
docker kill $(docker ps -qa)
docker rm $(docker ps -qa)
docker network rm traefik_default

# startup services
cd /root/alexa-skills-analytics/traefik
docker-compose up -d

cd /root/alexa-skills-analytics/server
docker-compose up -d

# temporarily disabled until we get this working
# cd /root/alexa-skills-analytics/web
# docker-compose up -d

cd /root/alexa-skills-analytics/admin
docker-compose up -d