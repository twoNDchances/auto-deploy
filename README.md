
# Auto-Deploy

Công cụ tự động hóa đơn giản cho Linux


## Installation

```bash
  git clone https://github.com/twoNDchances/auto-deploy.git
  cd auto-deploy
```

## Setup

Tự động tạo các tệp cần thiết:
```bash
  sh init.txt
```
Trong tệp `creds.txt`, điền thông tin của máy chủ:
```bash
  USERNAME=root
  HOSTNAME=
  PORT=
  PASSWORD=
  SITE_FILE=sitepath.txt
  COMMANDS=comms.txt
```
Trong tệp `sitepath.txt`, là danh sách các đường dẫn tuyệt đối:
```bash
  /home/user
  /tmp
  /etc/passwd
```
Trong tệp `comms.txt`, là danh sách các lệnh cần chạy tại các đường dẫn trong `sitepath.txt`:
```bash
  ls -la
  cat /etc/passwd
  touch test.txt && cat test.txt
```
Mặc định, các đường dẫn trong `sitepath.txt` đều sẽ được `cd` tự động đến và thực hiện các lệnh ở `comms.txt`, nhưng nếu cần `cd` sang đường dẫn khác rồi lại `cd` về, Auto-Deploy có cung cấp biến `$site` để làm được điều đó:
```bash
  cd $site && echo 123
```

# Usage
Sau khi đã hoàn thành các bước thiết lập trên:
```bash
  go run .
```

## Feedback

Liên hệ caoxuanthien13122002@gmail.com

