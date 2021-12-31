if [[ "$OSTYPE" == "linux-gnu" ]]; then
 
elif [[ "$OSTYPE" == "darwin" ]]; then
    cp configs/grafana_dashboards.json /usr/local/opt/grafana/share/grafana/conf/provisioning/dashboards
    cp configs/grafana_datasource.yml /usr/local/opt/grafana/share/grafana/conf/provisioning/datasources
    brew services restart grafana
fi