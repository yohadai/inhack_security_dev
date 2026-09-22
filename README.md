> ⚠️ **이 코드는 교육용으로 일부러 취약하게 만든 코드입니다. 실제 서비스에 절대 쓰면 안 됩니다.**
> 개발 파이프라인 보안(시크릿 관리, CI, 컨테이너 빌드, 의존성 공급망, 서명 키, 서비스 간 인증)을
> 직접 고쳐 보며 배우기 위한 스터디용 템플릿입니다.

### inhack-pipeline-template

작은 웹 애플리케이션을 예제로 삼아, 10주 동안 매주 보안 설정을 하나씩 붙여 나가는 템플릿입니다.

### 서비스 구조

세 개의 서비스로 이루어져 있습니다. 데이터베이스는 따로 두지 않고 각 서비스의 메모리에 저장합니다.

| 서비스 | 스택 | 역할 | 포트 |
| --- | --- | --- | --- |
| `auth-service` | Go + Echo | 회원가입(`POST /signup`), 로그인(`POST /login`). 로그인에 성공하면 JWT 발급 | 8081 |
| `api-service` | Go + Echo | 내 정보 조회(`GET /me`, JWT 검증), 상태 확인(`GET /health`). `/health` 에서 auth-service 를 HTTP 로 호출 | 8082 |
| `frontend` | React + Vite | 회원가입 / 로그인 / 내 정보 보기 세 화면 | 5173 |

```
브라우저 ──▶ frontend (5173)
                │
                ├──▶ auth-service (8081)  회원가입 · 로그인 · 토큰 발급
                └──▶ api-service  (8082)  토큰 검증 · 내 정보
                          │
                          └──▶ auth-service (8081)  /health 확인
```

### 실행 방법

#### 1) docker compose 로 한 번에 실행

```bash
docker compose up --build
```

- frontend: http://localhost:5173
- auth-service: http://localhost:8081/health
- api-service: http://localhost:8082/health

#### 2) 직접 실행

터미널 세 개를 띄웁니다.

```bash
# 터미널 1 - auth-service
cd auth-service
go run .

# 터미널 2 - api-service
cd api-service
go run .

# 터미널 3 - frontend
cd frontend
npm install
npm run dev
```

### 동작 확인 (curl)

```bash
# 회원가입
curl -X POST http://localhost:8081/signup \
  -H 'Content-Type: application/json' \
  -d '{"username":"alice","password":"pw1234"}'

# 로그인 → 토큰 받기
curl -X POST http://localhost:8081/login \
  -H 'Content-Type: application/json' \
  -d '{"username":"alice","password":"pw1234"}'

# 받은 토큰으로 내 정보 조회
curl http://localhost:8082/me \
  -H 'Authorization: Bearer <위에서 받은 토큰>'
```
---

##### 1주차 — 시크릿 스캔 (git 히스토리)

##### 2주차 — .gitignore / pre-commit / 브랜치 보호

##### 3주차 — CI 파이프라인 (빌드 + 시크릿 검사)

##### 4주차 — 컨테이너 이미지 보안 (Dockerfile / Trivy)

##### 5주차 — GitHub Actions 워크플로우 보안

##### 6주차 — 의존성 공급망 (Dependabot / audit)

##### 7주차 — JWT 서명 키 취약점 이해

##### 8주차 — 서명 키 관리 / 주입 / 교체

##### 9주차 — 서비스 간 인증 (SPIFFE/SPIRE)
