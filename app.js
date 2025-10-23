const http = require('http');

const server = http.createServer((req, res) => {
  res.writeHead(200, {'Content-Type': 'text/html; charset=utf-8'});
  res.end('<h1>🎉 Jenkins + Docker 자동 배포 성공!</h1><p>현재 시간: ' + new Date().toLocaleString('ko-KR') + '</p>');
});

server.listen(3000, () => {
  console.log('서버 실행 중: 3000번 포트');
});
