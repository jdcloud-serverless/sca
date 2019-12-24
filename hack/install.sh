mkdir ~/.sca
wget -P ~/.sca/ https://github.com/jdcloud-serverless/sca/releases/download/v0.0.1/sca

sudo chmod +777 ~/.sca/sca

echo "export PATH=$PATH:$HOME/.sca" >> $HOME/.bash_profile
source $HOME/.bash_profile