# traffic light
Golang application to show your Jenkins view status by usb-connected traffic light powered by arduino
- **app** directory contains desktop-side code
- **sketch** directory contains code for arduino

## Usage
go build -o traffic_light app/main.go
./traffic_light -jenkins=JENKINS_URL -view=VIEW_NAME -tty=TTY
