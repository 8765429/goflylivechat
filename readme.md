## GOFLY LIVE CHAT
Open-source live chat support system, built for modern customer service

Real-time messaging - Instant connection between customers and support teams

Lightning-fast performance - Powered by Golang for high-concurrency handling

Multilingual support - Chinese / English switchable via URL parameter `?lang=cn` or `?lang=en`

### Technical Architecture

A modern stack built for performance and scalability

- Backend: `gin`, `jwt-go`, `websocket`, `go.uuid`, `gorm`, `cobra`
- Frontend: `VueJS`, `ElementUI`
- Database: `MySQL`

---

### Quick Start with Docker (Recommended)

#### 1. Docker Compose Deployment

```bash
git clone https://github.com/8765429/goflylivechat.git
cd goflylivechat

# Create config from template
cp config/mysql.example.json config/mysql.json

# Edit mysql.json to match your database credentials
# For docker-compose deployment, set Server to "mysql"
vi config/mysql.json

# Build and start
docker-compose up -d --build

# Initialize database (first time only)
docker-compose exec gofly ./gochat install
```

Once running, the service listens on port 8081. Access via http://[your-ip]:8081.

#### 2. Export Docker Image for Deployment Elsewhere

```bash
# Build the image
docker-compose build

# Export image to tar file
docker save gofly-livechat:latest -o gofly-livechat.tar

# Copy to another server and load
scp gofly-livechat.tar user@server:/path/
ssh user@server "docker load -i /path/gofly-livechat.tar"
```

On the target server, create a `docker-compose.yml` with:

```yaml
version: "3"
services:
  mysql:
    image: mysql:5.7
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: goflychat@root
      MYSQL_DATABASE: goflychat
      MYSQL_USER: goflychat
      MYSQL_PASSWORD: goflychat
    volumes:
      - ./data/mysql:/var/lib/mysql
    networks:
      - gofly_net

  gofly:
    image: gofly-livechat:latest
    restart: always
    ports:
      - "8081:8081"
    volumes:
      - ./config/mysql.json:/app/config/mysql.json
      - ./data/upload:/app/static/upload
      - ./data/logs:/app/logs
    depends_on:
      - mysql
    networks:
      - gofly_net

networks:
  gofly_net:
    driver: bridge
```

Then run:
```bash
docker-compose up -d
docker-compose exec gofly ./gochat install
```

---

### Manual Installation

#### 1. Set Up MySQL Database
- Install and run MySQL (version >= 5.5).
- Create a database:
```sql
CREATE DATABASE goflychat CHARSET utf8mb4;
```

#### 2. Configure Database Connection
Edit `config/mysql.json`:
```json
{
    "Server":"127.0.0.1",
    "Port":"3306",
    "Database":"goflychat",
    "Username":"goflychat",
    "Password":"goflychat"
}
```

#### 3. Install and Configure Golang
```bash
wget https://studygolang.com/dl/golang/go1.21.0.linux-amd64.tar.gz
tar -C /usr/local -xvf go1.21.0.linux-amd64.tar.gz
echo "PATH=\$PATH:/usr/local/go/bin" >> /etc/profile
source /etc/profile
go version
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

#### 4. Download the Source Code
```bash
git clone https://github.com/8765429/goflylivechat.git
cd goflylivechat
```

#### 5. Initialize the Database
```bash
go run main.go install
```

#### 6. Run the Application
```bash
go run main.go server
```

#### 7. Build executable
```bash
go build -o gochat
./gochat server -p 8081
```

---

### Multilingual Support

The system supports Chinese (cn) and English (en) languages.

- Add `?lang=cn` or `?lang=en` to any URL to switch language
- Example: `http://your-ip:8081/livechat?user_id=agent&lang=en`
- Default language is Chinese (cn)

Language translations are defined in `static/js/chat-lang.js`.

---

### Customer Service Integration

Chat Link:
```
http://127.0.0.1:8081/livechat?user_id=agent
```

Popup Integration:
```html
<script>
    (function(global, document, scriptUrl, callback) {
        const head = document.getElementsByTagName('head')[0];
        const script = document.createElement('script');
        script.type = 'text/javascript';
        script.src = scriptUrl + "/static/js/chat-widget.js";
        script.onload = script.onreadystatechange = function () {
            if (!this.readyState || this.readyState === "loaded" || this.readyState === "complete") {
                callback(scriptUrl);
            }
        };
        head.appendChild(script);
    })(window, document, "http://127.0.0.1:8081", function(baseUrl) {
        CHAT_WIDGET.initialize({
            API_URL: baseUrl,
            AGENT_ID: "agent",
            LANG: "cn"
        });
    });
</script>
```

For domain access, configure a reverse proxy to port 8081 to hide the port number.

### Important Notice
The use of this project for illegal or non-compliant purposes, including but not limited to viruses, trojans, pornography, gambling, fraud, prohibited items, counterfeit products, false information, cryptocurrencies, and financial violations, is strictly prohibited.

This project is intended solely for personal learning and testing purposes. Any commercial use or illegal activities are explicitly forbidden!!!

### Copyright Notice
This project provides full-featured code but is intended only for personal demonstration and testing. Commercial use is strictly prohibited.

By using this software, you agree to comply with all applicable local laws and regulations. You are solely responsible for any legal consequences arising from misuse.
