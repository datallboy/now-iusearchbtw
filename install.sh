# /bin/bash

if ! [ $(id -u) = 0 ]; then
   echo "The script need to be run as root." >&2
   exit 1
fi

if [ $SUDO_USER ]; then
    real_user=$SUDO_USER
else
    real_user=$(whoami)
fi

## Build project
sudo -u $real_user make build

## Create systemd service
SERVICE_NAME="now-iusearchbtw"
SECRET=$(openssl rand -base64 32)

IS_ACTIVE=$(systemctl is-active $SERVICE_NAME)
if [ "$IS_ACTIVE" = "active" ]; then
  # restart the service
  echo "Service is running"
  echo "Restarting service"
  systemctl restart $SERVICE_NAME
  echo "Service restarted"
else 
  # create service file
  echo "Creating service file"
  cat > /etc/systemd/system/${SERVICE_NAME}.service << EOF
[Unit]
Description=Now you can say you use Arch btw. 
After=network.target

[Service]
ExecStart=$PWD/bin/${SERVICE_NAME}
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOF
  ## Restart daemon, enable, and start service
  systemctl daemon-reload
  systemctl enable ${SERVICE_NAME}
  systemctl start ${SERVICE_NAME}
  echo "Service Started"
fi

exit 0