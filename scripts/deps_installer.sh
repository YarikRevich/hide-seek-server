if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    apt update
    apt install golang=1.17.5
    apt install prometheus
elif [[ "$OSTYPE" == "darwin"* ]]; then
   

    if ! command -v brew &> /dev/null; then
        /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    fi 

    if ! command -v go &> /dev/null; then
        brew install go@1.17
    fi 

    if ! command -v docker &> /dev/null; then
        brew install docker >& /dev/null;
        brew install --cask docker >& /dev/null;
        open -a /Applications/Docker.app;
    fi 

    # if ! command -v grafana-server &> /dev/null; then

    # fi 
    # brew install go@1.17.5;
    # brew install prometheus@2.32.1;
    # brew install grafana@8.3.3;
fi