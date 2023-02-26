echo 'Building server'
cd server
GOOS=linux go build
cd ..

echo 'Building web page'
cd web
npm install
npm run build
rm -rf node_modules
cd ..

echo 'Building admin page'
cd admin
npm install
npm run build
rm -rf node_modules
cd ..

echo 'Zipping up project'
cd ..
zip -r alexa-deployment.zip alexa-skills-analytics/

echo 'Deploying to web server...'
scp alexa-deployment.zip ubuntu@uwalexaskills.org:/home/ubuntu
cd alexa-skills-analytics
ssh ubuntu@uwalexaskills.org 'cd /home/ubuntu && unzip alexa-deployment.zip'
echo 'Connecting to server. Run the startup.sh script to start all new services'
ssh ubuntu@uwalexaskills.org